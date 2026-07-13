package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// I18nMiddleware determines the request language and stores it in context.
func I18nMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Check query parameter `lang` (e.g. ?lang=vi)
		lang := c.Query("lang")

		// 2. Check X-Lang header
		if lang == "" {
			lang = c.Get("X-Lang")
		}

		// 3. Check Accept-Language header
		if lang == "" {
			acceptLang := c.Get("Accept-Language")
			if acceptLang != "" {
				// Simple check: if "vi" exists in Accept-Language, we prefer Vietnamese
				if strings.Contains(strings.ToLower(acceptLang), "vi") {
					lang = "vi"
				} else {
					lang = "en"
				}
			}
		}

		// Normalize to "en" or "vi", default to "en"
		lang = strings.ToLower(strings.TrimSpace(lang))
		if lang == "" {
			lang = "en"
		}
		if strings.HasPrefix(lang, "vi") {
			lang = "vi"
		} else {
			lang = "en"
		}

		// Store in Fiber Locals
		c.Locals("lang", lang)

		return c.Next()
	}
}
