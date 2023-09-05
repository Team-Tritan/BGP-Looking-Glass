package handlers

import "github.com/gofiber/fiber/v2"

func SendErrorResponse(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).SendString("~as393577 looking glass (ง'̀-'́)ง♡~\n\n" + message)
}
