package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/slog"
)

type LoggerConfig struct {
	DefaultLevel     slog.Level
	ClientErrorLevel slog.Level
	ServerErrorLevel slog.Level

	WithRequestID bool
}

// New returns a fiber.Handler (middleware) that logs requests using slog.
//
// Requests with errors are logged using slog.Error().
// Requests without errors are logged using slog.Info().
func NewFiberLogger(logger *slog.Logger) fiber.Handler {
	return NewFiberLoggerWithConfig(logger, LoggerConfig{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		WithRequestID: true,
	})
}

// NewWithConfig returns a fiber.Handler (middleware) that logs requests using slog.
func NewFiberLoggerWithConfig(logger *slog.Logger, config LoggerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Path()
		path := c.Path()

		requestID := uuid.New().String()
		if config.WithRequestID {
			c.Context().SetUserValue("request-id", requestID)
			c.Set("X-Request-ID", requestID)
		}

		err := c.Next()

		attributes := []slog.Attr{
			slog.Int("status", c.Response().StatusCode()),
			slog.String("method", string(c.Context().Method())),
			slog.String("path", path),
			slog.String("ip", c.Context().RemoteIP().String()),
		}

		if config.WithRequestID {
			attributes = append(attributes, slog.String("request-id", requestID))
		}

		switch {
		case c.Response().StatusCode() >= fasthttp.StatusBadRequest && c.Response().StatusCode() < fasthttp.StatusInternalServerError:
			logger.LogAttrs(context.Background(), config.ClientErrorLevel, err.Error(), attributes...)
		case c.Response().StatusCode() >= fasthttp.StatusInternalServerError:
			logger.LogAttrs(context.Background(), config.ServerErrorLevel, err.Error(), attributes...)
		default:
			logger.LogAttrs(context.Background(), config.DefaultLevel, "Incoming request", attributes...)
		}

		return err
	}
}

// GetRequestID returns the request identifier
func GetRequestID(c *fiber.Ctx) string {
	requestID, ok := c.Context().UserValue("request-id").(string)
	if !ok {
		return ""
	}

	return requestID
}
