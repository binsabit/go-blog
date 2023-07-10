package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.con/binsabi/go-blog/config"
	"golang.org/x/exp/slog"
)

type Application struct {
	Storage *sql.DB
	Router  *fiber.App
	Logger  *slog.Logger
	Config  *config.Config
}

func NewApplication(router *fiber.App, storage *sql.DB, logger *slog.Logger, config *config.Config) *Application {
	return &Application{
		Storage: storage,
		Router:  router,
		Logger:  logger,
		Config:  config,
	}
}

func (app *Application) StartApplication() {

}
