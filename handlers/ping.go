package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tritan.dev/bgp-tool/commands"
	"tritan.dev/bgp-tool/regex"
)

func Ping(c *fiber.Ctx) error {
	ip := c.Query("ip")
	if !regex.IsValidIP(ip) {
		return SendErrorResponse(c, "Invalid IP.", fiber.StatusBadRequest)
	}
	response, err := commands.ExecutePing(ip)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
}
