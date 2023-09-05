package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"tritan.dev/bgp-tool/handlers"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	SetupRoutes(app)

	err := app.Listen(":4000")
	if err != nil {
		log.Fatal(err)
	}
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Landing)
	app.Get("/show-route", handlers.ShowRoute)
	app.Get("/bgp-routes", handlers.BgpRoutes)
	app.Get("/ping", handlers.Ping)
	app.Get("/traceroute", handlers.Traceroute)
	app.Get("/mtr", handlers.Mtr)
}
