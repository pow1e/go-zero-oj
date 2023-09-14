package global

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"sync"
)

var validateOnce sync.Once
var Validate *validator.Validate
var Translator ut.Translator

func InitValidator() {
	validateOnce.Do(func() {
		val := validator.New()
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		trans, _ := ut.New(zh.New()).GetTranslator("zh")
		if validateErr := zhTranslations.RegisterDefaultTranslations(val, trans); validateErr != nil {
			panic(validateErr)
		}
		Validate = val
		Translator = trans
	})
}
