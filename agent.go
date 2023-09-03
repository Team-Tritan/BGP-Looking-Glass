package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"tritan.dev/bgp-tool/Regex"
	"tritan.dev/bgp-tool/commands"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("suck yourself ~as393577~")
	})

	app.Get("/route", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !Regex.IsValidSubnet(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid subnet param")
		}
		response, err := commands.ExecuteBirdCommand(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Route Info for IP %s:\n%s", ip, response))
	})

	app.Get("/asn-routes", func(c *fiber.Ctx) error {
		asn := c.Query("asn")
		if !Regex.IsValidASN(asn) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ASN param")
		}
		response, err := commands.ExecuteBirdCommand(fmt.Sprintf("where bgp_path ~ [= * %s * =] all", asn))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("BGP Routes for ASN %s:\n%s", asn, response))
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !Regex.IsValidIP(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid IP param")
		}
		response, err := commands.ExecutePing(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Ping for IP %s:\n%s", ip, response))
	})

	log.Fatal(app.Listen(":8080"))
}
