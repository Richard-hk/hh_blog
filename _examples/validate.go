package main

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	ID   int64  `json:"id" validate:"gt=0" label:"id"`
	Name string `json:"name" validate:"required,max=2" label:"姓名"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()
	validateVar()
	validateStruct()
	validateStructWithZh()
}

func validateStruct() {
	user := User{
		-1,
		"",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		// Key: 'User.ID' Error:Field validation for 'ID' failed on the 'gt' tag
		// Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
	}
	user = User{
		1,
		"ha",
	}
	fmt.Println()
	err = validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
	// ok
}

func validateVar() {
	email := "123c"
	err := validate.Var(email, "required,email")
	if err != nil {
		fmt.Println(err)
		// Key: '' Error:Field validation for '' failed on the 'email' tag
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		// Key: '' Error:Field validation for '' failed on the 'email' tag
	}

	email = "123@qq.com"
	err = validate.Var(email, "required,email")
	if err != nil {
		fmt.Println(err)
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
	}
	// ok
}

func validateStructWithZh() {
	user := User{
		-1,
		"222",
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	err := zh_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		fmt.Println(err)
		return
	}

	errs := validate.Struct(user)
	fmt.Println()
	fmt.Println(errs)
	fmt.Println()
	if errs != nil {
		a := errs.(validator.ValidationErrors)
		fmt.Println()
		for _, err := range a {
			fmt.Println(err.Translate(trans))
		}
	}
}
