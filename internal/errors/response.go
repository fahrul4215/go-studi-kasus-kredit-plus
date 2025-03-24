package errors

import (
	"encoding/json"
	"fmt"
	"go-studi-kasus-kredit-plus/internal/utils"
	"net/http"
	"sort"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return ErrorResponse{
		Status:  http.StatusForbidden,
		Message: msg,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InvalidInput creates a new error response representing a data validation error (HTTP 400).
func InvalidInput(errs validation.Errors) ErrorResponse {
	var details []invalidField
	var fields []string
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}

	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "There is some problem with the data you submitted.",
		Details: details,
	}
}

// ParseFieldError parses a validator.FieldError into a more readable string format
func parseFieldError(e validator.FieldError) string {
	// workaround to the fact that the `gt|gtfield=Start` gets passed as an entire tag for some reason
	// https://github.com/go-playground/validator/issues/926
	fieldPrefix := fmt.Sprintf("The field %s", utils.ToSnakeCase(e.Field()))
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "len":
		return fmt.Sprintf("%s must be %s characters long", fieldPrefix, e.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldPrefix)
	case "required":
		return fmt.Sprintf("%s is required", fieldPrefix)
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", fieldPrefix, e.Param())
	case "isdefault":
		return fmt.Sprintf("%s must be %s", fieldPrefix, e.Param())
	case "numeric":
		return fmt.Sprintf("%s must be a number", fieldPrefix)
	case "min":
		return fmt.Sprintf("%s must be at least %s", fieldPrefix, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", fieldPrefix, e.Param())
	default:
		// if it's a tag for which we don't have a good format string yet we'll try using the default english translator
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		}
		return fmt.Errorf("%v", e).Error()
	}
}

// parseFieldErrorID parses a validator.FieldError into a more readable string format
func parseFieldErrorID(e validator.FieldError) string {
	fieldPrefix := fmt.Sprintf("Kolom %s", utils.ToSnakeCase(e.Field()))
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "len":
		return fmt.Sprintf("%s harus %s karater", fieldPrefix, e.Param())
	case "email":
		return fmt.Sprintf("%s harus valid", fieldPrefix)
	case "required":
		return fmt.Sprintf("%s wajib diisi", fieldPrefix)
	case "required_without":
		return fmt.Sprintf("%s wajib diiisi jika %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s harus kurang dari %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s harus lebih dari %s", fieldPrefix, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", fieldPrefix, e.Param())
	case "isdefault":
		return fmt.Sprintf("%s must be %s", fieldPrefix, e.Param())
	case "numeric":
		return fmt.Sprintf("%s harus berupa angka", fieldPrefix)
	case "min":
		return fmt.Sprintf("%s harus minimal %s", fieldPrefix, e.Param())
	case "max":
		return fmt.Sprintf("%s harus maksimal %s", fieldPrefix, e.Param())
	default:
		// if it's a tag for which we don't have a good format string yet we'll try using the default english translator
		// indonesia := id.New()
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("id"); found {
			return e.Translate(translatorInstance)
		}
		return fmt.Errorf("%v", e).Error()
	}
}

// ParseMarshallingError parses an unmarshalling error into a more readable string format
func parseMarshallingError(e json.UnmarshalTypeError) string {
	field := utils.ToSnakeCase(e.Field)
	return fmt.Sprintf("The field %s must be a %s", field, e.Type.String())
}

// parseMarshallingErrorID parses an unmarshalling error into a more readable string format
func parseMarshallingErrorID(e json.UnmarshalTypeError) string {
	field := utils.ToSnakeCase(e.Field)
	return fmt.Sprintf("Kolom %s harus berupa %s", field, e.Type.String())
}

// ParseErrorValidation parses a validator.ValidationErrors into a more readable string format
func ParseErrorValidation(errs ...error) string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
			// loop through each of those and convert them one by one
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}

	return out[0]
}

// customs represents FAHRUL
func Custom(status int, msg string, detail any) ErrorResponse {
	if msg == "" {
		msg = "error acquired, please contact administrator."
	}
	return ErrorResponse{
		Status:  status,
		Message: msg,
		Details: detail,
	}
}
