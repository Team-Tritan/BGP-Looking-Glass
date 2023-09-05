package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"tritan.dev/bgp-tool/handlers"
)

func main() {
	app := setupApp()
	err := app.Listen(":4000")
	if err != nil {
		log.Fatal(err)
	}
}

func setupApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	setupRoutes(app)
	return app
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Landing)
	app.Get("/show-route", handlers.ShowRoute)
	app.Get("/bgp-routes", handlers.BgpRoutes)
	app.Get("/ping", handlers.Ping)
	app.Get("/traceroute", handlers.Traceroute)
	app.Get("/mtr", handlers.Mtr)
}

