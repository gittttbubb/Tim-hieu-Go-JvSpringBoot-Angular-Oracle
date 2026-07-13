package response

import (
	"go-rbac-system/internal/validator"
	"go-rbac-system/pkg/i18n"

	"github.com/gofiber/fiber/v2"
)

func translateMessage(c *fiber.Ctx, message string) string {
	lang := "en"
	if l := c.Locals("lang"); l != nil {
		lang = l.(string)
	}
	return i18n.Get(lang, message)
}

func Success(c *fiber.Ctx, data interface{}) error {
	if m, ok := data.(fiber.Map); ok {
		if msg, exists := m["message"]; exists {
			if strMsg, isStr := msg.(string); isStr {
				m["message"] = translateMessage(c, strMsg)
			}
		}
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"data":    data,
		},
	)
}

func Created(c *fiber.Ctx, data interface{}) error {
	if m, ok := data.(fiber.Map); ok {
		if msg, exists := m["message"]; exists {
			if strMsg, isStr := msg.(string); isStr {
				m["message"] = translateMessage(c, strMsg)
			}
		}
	}
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"success": true,
			"data":    data,
		},
	)
}

func Error(c *fiber.Ctx, status int, message string) error {
	translatedMsg := translateMessage(c, message)
	return c.Status(status).JSON(
		fiber.Map{
			"success": false,
			"message": translatedMsg,
		},
	)
}

// ValidationError translates and returns a structured validation error response.
func ValidationError(c *fiber.Ctx, err error) error {
	lang := "en"
	if l := c.Locals("lang"); l != nil {
		lang = l.(string)
	}
	msg := validator.TranslateError(err, lang)
	return Error(c, fiber.StatusBadRequest, msg)
}