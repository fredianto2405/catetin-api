package errors

import (
	"github.com/fredianto2405/catetin-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})
}

func getErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "tidak boleh kosong"
	case "email":
		return "email tidak valid"
	case "min":
		return "minimal " + e.Param() + " karakter"
	case "max":
		return "maksimal " + e.Param() + " karakter"
	default:
		return "nilai tidak valid"
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors
		if len(errs) > 0 {
			err := errs.Last().Err

			switch e := err.(type) {
			case validator.ValidationErrors:
				validationErrors := make(map[string]string)
				for _, ve := range e {
					field := ve.Field()
					validationErrors[field] = getErrorMsg(ve)
				}
				response.Respond(c, http.StatusBadRequest, false, "errors validasi", validationErrors, nil)
				return

			default:
				response.Respond(c, http.StatusInternalServerError, false, err.Error(), nil, nil)
				return
			}
		}
	}
}
