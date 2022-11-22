package validators

import (
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var validate *validator.Validate

// 验证结构体
func ValidateStructLabel(s interface{}) interface{} {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	errTrans := zh_translations.RegisterDefaultTranslations(validate, trans)
	if errTrans != nil {
		return errTrans
	}
	errS := validate.Struct(s)
	if errS == nil {
		return nil
	}
	errV := errS.(validator.ValidationErrors)
	res := errV[0].Translate(trans)
	for _, s := range errV[1:] {
		res += "," + s.Translate(trans)
	}
	return res
}
