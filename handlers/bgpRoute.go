package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tritan.dev/bgp-tool/commands"
	"tritan.dev/bgp-tool/regex"
)

func BgpRoutes(c *fiber.Ctx) error {
	asn := c.Query("asn")
	if !regex.IsValidASN(asn) {
		return SendErrorResponse(c, "Invalid ASN.", fiber.StatusBadRequest)
	}
	response, err := commands.ExecuteBirdCommand(fmt.Sprintf("where bgp_path ~ [= * %s * =] all", asn))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("~as393577 looking glass (ง'̀-'́)ง♡~\n\n%s", response))
}
