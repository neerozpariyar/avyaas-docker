package utils

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

/*
InitTranslator initializes a validator and a translator for error messages. The validator is set up
with English (en) translations, and the translator is configured accordingly. It is used in
conjunction with the TranslateError function for handling validation errors.
Returns:
  - A pointer to the new validator.Validate instance that can be used to validate struct fields.
  - A ut.Translator instance that contains all available translation functions.
*/
func InitTranslator() (*validator.Validate, ut.Translator) {
	// Create a new validator instance
	validate := validator.New()

	// Initialize English language support for translation
	en := en.New()

	// Create a new universal translator with English as the default language
	uni := ut.New(en, en)

	// Get the translator instance for the English language
	trans, _ := uni.GetTranslator("en")

	// Register default translations for English in the validator
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	// Return the validator and translator instances
	return validate, trans
}

/*
TranslateError translates validation errors using the provided translator and returns a map where
field names are keys and corresponding translated error messages are values. If no validation errors
are provided, an empty map is returned.
Parameters:
  - validationErrs: A validator.ValidationErrors instance containing validation errors.
  - trans: A ut.Translator for translating validation errors into the desired language.

Returns:
  - A map[string]string where keys are field names and values are translated error messages.
*/
func TranslateError(validationErrs validator.ValidationErrors, trans ut.Translator) map[string]string {
	// Create an empty map to store field names as key and their translated error messages as value
	err_map := make(map[string]string)

	// Check if validation errors are provided
	if len(validationErrs) != 0 {
		for _, e := range validationErrs {
			// Store the lowercase field name and its translated error message in the map
			err_map[strings.ToLower(e.Field())] = e.Translate(trans)
		}
	}

	return err_map
}
