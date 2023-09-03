package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func executeCommand(command string) (string, error) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	app := fiber.New()

	app.Get("/route", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if ip == "" {
			return c.Status(fiber.StatusBadRequest).SendString("IP parameter is required")
		}
		cmd := fmt.Sprintf("birdc show route %s", ip)
		response, err := executeCommand(cmd)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Route Info for IP %s:\n%s", ip, response))
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if ip == "" {
			return c.Status(fiber.StatusBadRequest).SendString("IP parameter is required")
		}
		cmd := fmt.Sprintf("ping %s", ip)
		response, err := executeCommand(cmd)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Ping for IP %s:\n%s", ip, response))
	})

	app.Get("/routes/asn", func(c *fiber.Ctx) error {
		asn := c.Query("asn")
		if asn == "" {
			return c.Status(fiber.StatusBadRequest).SendString("ASN parameter is required")
		}
		cmd := fmt.Sprintf("sudo birdc show route where bgp_path ~ [= * %s * =] all", asn)
		response, err := executeCommand(cmd)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("BGP Routes for ASN %s:\n%s", asn, response))
	})

	log.Fatal(app.Listen(":8080"))
}
