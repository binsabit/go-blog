package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.con/binsabi/go-blog/app"
	"github.con/binsabi/go-blog/config"
	"github.con/binsabi/go-blog/db"
	"github.con/binsabi/go-blog/lib/logger/sl"
	"github.con/binsabi/go-blog/middlewares"
)

func main() {

	config := config.MustLoad()

	logger := sl.NewLogger(config.LogFile)

	router := fiber.New()
	router.Use(cors.New())
	router.Use(middlewares.NewFiberLogger(logger))
	pool := db.ConnectToDB(config.Storage)
	defer pool.Close()

	application := app.NewApplication(router, pool, logger, config)
	application.RegisterRoutes()
	err := application.Router.Listen(config.HTTPServer.Address)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
