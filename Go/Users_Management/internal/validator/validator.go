package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	vi_translations "github.com/go-playground/validator/v10/translations/vi"
)

var Validate = validator.New()
var Uni *ut.UniversalTranslator
var TransEN ut.Translator
var TransVI ut.Translator

func init() {
	enLocale := en.New()
	viLocale := vi.New()
	Uni = ut.New(enLocale, enLocale, viLocale)

	TransEN, _ = Uni.GetTranslator("en")
	TransVI, _ = Uni.GetTranslator("vi")

	_ = en_translations.RegisterDefaultTranslations(Validate, TransEN)
	_ = vi_translations.RegisterDefaultTranslations(Validate, TransVI)
}

// TranslateError translates validator.ValidationErrors to a human-readable string based on language.
func TranslateError(err error, lang string) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	var trans ut.Translator
	if lang == "vi" {
		trans = TransVI
	} else {
		trans = TransEN
	}

	var errors []string
	for _, e := range errs {
		errors = append(errors, e.Translate(trans))
	}

	if len(errors) > 0 {
		return errors[0] // Return the first error message for clean UI presentation
	}
	return err.Error()
}