package validator

import (
	"github.com/go-playground/locales/en"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"server/utils/result"
)

func Validate(data interface{}) (code int, message *string) {
	uni := unTrans.New(en.New())
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	err := enTrans.RegisterDefaultTranslations(validate, trans)
	err = validate.Struct(data)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			message := fieldError.Translate(trans)
			return result.Error, &message
		}
	}
	return result.Success,nil
}
