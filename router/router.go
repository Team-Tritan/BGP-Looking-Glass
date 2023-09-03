package router

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"tritan.dev/bgp-tool/commands"
	"tritan.dev/bgp-tool/regex"
)

func SetupRoutes(app *fiber.App) {
	endpoints := []string{
		"/show-route?subnet=<subnet>",
		"/bgp-routes?asn=<asn>",
		"/ping?ip=<ip>",
		"/traceroute?ip=<ip>",
		"/mtr?ip=<ip>",
	}

	app.Get("/", func(c *fiber.Ctx) error {
		endpointList := fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\nEndpoints:\n%s", strings.Join(endpoints, "\n"))
		return c.SendString(endpointList)
	})

	app.Get("/show-route", func(c *fiber.Ctx) error {
		subnet := c.Query("subnet")
		if !regex.IsValidSubnet(subnet) {
			return c.Status(fiber.StatusBadRequest).SendString("~as393577 looking glass (ง'̀-'́)ง♡~\n\nInvalid subnet.")
		}
		response, err := commands.ExecuteBirdCommand(subnet)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
	})

	app.Get("/bgp-routes", func(c *fiber.Ctx) error {
		asn := c.Query("asn")
		if !regex.IsValidASN(asn) {
			return c.Status(fiber.StatusBadRequest).SendString("~as393577 looking glass (ง'̀-'́)ง♡~\n\nInvalid ASN.")
		}
		response, err := commands.ExecuteBirdCommand(fmt.Sprintf("where bgp_path ~ [= * %s * =] all", asn))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !regex.IsValidIP(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("~as393577 looking glass (ง'̀-'́)ง♡~\n\nInvalid IP.")
		}
		response, err := commands.ExecutePing(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
	})

	app.Get("/traceroute", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !regex.IsValidIP(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("~as393577 looking glass (ง'̀-'́)ง♡~\n\nInvalid IP.")
		}
		response, err := commands.ExecuteTraceroute(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
	})

	app.Get("/mtr", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !regex.IsValidIP(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("~as393577 looking glass (ง'̀-'́)ง♡~\n\nInvalid IP.")
		}
		response, err := commands.ExecuteMTR(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
	})
}
