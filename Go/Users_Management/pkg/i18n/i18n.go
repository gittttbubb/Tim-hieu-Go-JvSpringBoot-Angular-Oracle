package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed locales/*.json
var localesFS embed.FS

var translations = make(map[string]map[string]string)

// Init loads the embedded JSON translations into memory.
func Init() error {
	// Load en.json
	enBytes, err := localesFS.ReadFile("locales/en.json")
	if err != nil {
		return fmt.Errorf("failed to read locales/en.json: %w", err)
	}
	var enMap map[string]string
	if err := json.Unmarshal(enBytes, &enMap); err != nil {
		return fmt.Errorf("failed to parse locales/en.json: %w", err)
	}
	translations["en"] = enMap

	// Load vi.json
	viBytes, err := localesFS.ReadFile("locales/vi.json")
	if err != nil {
		return fmt.Errorf("failed to read locales/vi.json: %w", err)
	}
	var viMap map[string]string
	if err := json.Unmarshal(viBytes, &viMap); err != nil {
		return fmt.Errorf("failed to parse locales/vi.json: %w", err)
	}
	translations["vi"] = viMap

	return nil
}

// Get returns the translation for the given language and key.
// If the key doesn't exist, it falls back to English, and finally the key itself.
func Get(lang string, key string) string {
	lang = strings.ToLower(strings.TrimSpace(lang))
	if lang == "" {
		lang = "en"
	}
	if len(lang) > 2 && lang[2] == '-' {
		lang = lang[:2]
	}

	if dict, ok := translations[lang]; ok {
		if val, ok := dict[key]; ok {
			return val
		}
	}

	// Fallback to "en"
	if lang != "en" {
		if dict, ok := translations["en"]; ok {
			if val, ok := dict[key]; ok {
				return val
			}
		}
	}

	return key
}

// GetCatalog returns the dictionary for a specific language.
func GetCatalog(lang string) map[string]string {
	lang = strings.ToLower(strings.TrimSpace(lang))
	if lang == "" {
		lang = "en"
	}
	if len(lang) > 2 && lang[2] == '-' {
		lang = lang[:2]
	}

	if dict, ok := translations[lang]; ok {
		return dict
	}
	return translations["en"]
}
