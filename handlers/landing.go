package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Landing(c *fiber.Ctx) error {
	endpoints := []string{
		"/show-route?subnet=<subnet>",
		"/bgp-routes?asn=<asn>",
		"/ping?ip=<ip>",
		"/traceroute?ip=<ip>",
		"/mtr?ip=<ip>",
	}
	endpointList := fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\nEndpoints:\n%s", strings.Join(endpoints, "\n"))
	return c.SendString(endpointList)
}
