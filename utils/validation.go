package utils

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/stoewer/go-strcase"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func ValidateStruct(s interface{}) map[string][]string {
	en := en.New()

	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()

	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(s)

	fields := TranslateError(err, trans)

	return fields
}

func TranslateError(err error, trans ut.Translator) map[string][]string {
	if err == nil {
		return nil
	}

	fields := map[string][]string{}

	for _, e := range err.(validator.ValidationErrors) {
		translatedErr := e.Translate(trans)

		splittedErr := strings.Split(translatedErr, " ")
		arrFiledName := strings.Split(strcase.SnakeCase(splittedErr[0]), "_")
		fieldName := strings.Title(strings.Join(arrFiledName[:], " "))
		splittedErr[0] = fieldName

		translatedErr = strings.Join(splittedErr[:], " ")

		errMessages := fields[e.Field()]
		errMessages = append(errMessages, translatedErr)

		fields[strcase.SnakeCase(e.Field())] = errMessages
	}

	return fields
}
