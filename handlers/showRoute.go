package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tritan.dev/bgp-tool/commands"
	"tritan.dev/bgp-tool/regex"
)

func ShowRoute(c *fiber.Ctx) error {
	subnet := c.Query("subnet")
	if !regex.IsValidSubnet(subnet) {
		return SendErrorResponse(c, "Invalid subnet.", fiber.StatusBadRequest)
	}
	response, err := commands.ExecuteBirdCommand(subnet)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
}
