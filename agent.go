package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RouteInfoResponse struct {
	IP       string `json:"ip"`
	Response string `json:"response"`
}

type BGPRouteResponse struct {
	ASN      string `json:"asn"`
	Response string `json:"response"`
}

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

	app.Get("/routes/show", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if ip == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "IP parameter is required"})
		}
		cmd := fmt.Sprintf("sudo birdc show route %s", ip)
		response, err := executeCommand(cmd)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(RouteInfoResponse{IP: ip, Response: response})
	})

	app.Get("/routes/bgp", func(c *fiber.Ctx) error {
		asn := c.Query("asn")
		if asn == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ASN parameter is required"})
		}
		cmd := fmt.Sprintf("sudo birdc show route where bgp_path ~ [= * %s * =] all", asn)
		response, err := executeCommand(cmd)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(BGPRouteResponse{ASN: asn, Response: response})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if ip == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "IP parameter is required"})
		}
		cmd := fmt.Sprintf("ping %s", ip)
		response, err := executeCommand(cmd)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(RouteInfoResponse{IP: ip, Response: response})
	})

	log.Fatal(app.Listen(":8080"))
}
