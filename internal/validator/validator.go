package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entransations "github.com/go-playground/validator/v10/translations/en"
	log "github.com/sirupsen/logrus"
)

type ErrNamespace string

type Handler interface {
	Validate(input interface{}) *ValidationError
}

type ValidationError struct {
	Namespace    ErrNamespace
	ErrorField   string
	ErrorMessage string
}

type Validator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidator() Handler {
	translator := en.New()

	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")

	if !found {
		log.Errorln("translator not found")
	}

	v := validator.New()

	if err := entransations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Error("error registering the transactions %s", err.Error())
		return nil
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "required field {0} is missing", true)
	},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})

	return &Validator{
		Validator:  v,
		Translator: trans,
	}
}

func (v *Validator) Validate(input interface{}) *ValidationError {
	err := v.Validator.Struct(input)
	if err == nil {
		return nil
	}

	for _, er := range err.(validator.ValidationErrors) {
		vErr := &ValidationError{
			ErrorMessage: er.Translate(v.Translator),
			ErrorField:   er.Field(),
			Namespace:    ErrNamespace(er.Namespace()),
		}

		return vErr
	}

	return nil
}
