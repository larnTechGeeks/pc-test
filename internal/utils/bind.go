package utils

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type invalidArguement struct {
	Field string `json:"field"`
	Value any    `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func Bind(c *gin.Context, req any) bool {

	if err := c.ShouldBind(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArguement

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArguement{
					getTagName(reflect.TypeOf(req).Elem(), err.Field()),
					err.Value(),
					err.Tag(),
					err.Param(),
				})
			}

			appErr := errors.New("Submitted invalid form")
			response := gin.H{}
			response["invalid_arguments"] = invalidArgs
			response["error"] = appErr.Error()

			c.JSON(400, response)

			return false
		}

		appErr := errors.New("failed to bind form")

		c.JSON(400, gin.H{"error_message": appErr.Error()})

		return false
	}

	return true
}

func getTagName(t reflect.Type, fieldName string) string {

	field, ok := t.FieldByName(fieldName)
	if !ok {
		return ""
	}

	tag := field.Tag.Get("json")
	if tag != "" {
		return tag
	}

	tag = field.Tag.Get("form")
	if tag != "" {
		return tag
	}

	return field.Tag.Get("xml")
}
