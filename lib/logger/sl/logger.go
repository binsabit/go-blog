package sl

import (
	"log"
	"os"

	"golang.org/x/exp/slog"
)

func NewLogger(filepath string) *slog.Logger {
	if filepath == "" {
		textHandler := slog.NewTextHandler(os.Stdout, nil)
		logger := slog.New(textHandler)
		return logger
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("could not init log file:%v", err)
	}
	// defer file.Close()

	textHandler := slog.NewTextHandler(file, nil)
	logger := slog.New(textHandler)
	return logger

}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
