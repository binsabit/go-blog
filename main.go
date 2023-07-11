package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.con/binsabi/go-blog/app"
	"github.con/binsabi/go-blog/config"
	"github.con/binsabi/go-blog/db"
	"github.con/binsabi/go-blog/lib/logger/sl"
)

func main() {

	config := config.MustLoad()

	logger := sl.NewLogger(config.LogFile)

	router := fiber.New()
	// router.Use(helmet.New())
	// router.Use(recover.New())
	// router.Use(cors.New())
	// router.Use(compress.New(compress.Config{
	// Level: compress.LevelBestSpeed,
	// }))
	// router.Use(middlewares.NewFiberLogger(logger))
	pool := db.ConnectToDB(config.Storage)
	defer pool.Close()

	application := app.NewApplication(router, pool, logger, config)
	application.RegisterRoutes()
	err := application.Router.Listen(config.HTTPServer.Address)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
