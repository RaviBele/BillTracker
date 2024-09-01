package controllers

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func PhoneValidation(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	log.Printf("Validating phone: %s", value)
	re, err := regexp.Compile(`^\+\d{1,3}\d{9}$`)
	if err != nil {
		log.Println("Error compiling regular expression")
		return false
	}

	return re.MatchString(value)
}

func RegisterValidations() {
	err := validate.RegisterValidation("phone", PhoneValidation)
	if err != nil {
		log.Fatal("Error registering custom validation :", err.Error())
	}
}
