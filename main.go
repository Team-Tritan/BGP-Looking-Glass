package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"tritan.dev/bgp-tool/router"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	err := app.Listen(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
