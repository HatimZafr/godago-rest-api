package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Error  string       `json:"error"`
	Errors []FieldError `json:"errors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AppError struct {
	StatusCode int
	ErrorType  string
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewDatabaseError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  "DATABASE_ERROR",
		Message:    message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		ErrorType:  "NOT_FOUND",
		Message:    message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		ErrorType:  "BAD_REQUEST",
		Message:    message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  "INTERNAL_SERVER_ERROR",
		Message:    message,
	}
}

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, ErrorResponse{
			Error:   appErr.ErrorType,
			Message: appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error:   "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
	})
}

func HandleValidationError(c *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var fieldErrors []FieldError
		for _, e := range validationErrors {
			fieldErrors = append(fieldErrors, FieldError{
				Field:   e.Field(),
				Message: getValidationMessage(e),
			})
		}

		c.JSON(http.StatusBadRequest, ValidationErrorResponse{
			Error:  "VALIDATION_ERROR",
			Errors: fieldErrors,
		})
		return
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   "BAD_REQUEST",
		Message: err.Error(),
	})
}

func getValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must not exceed " + e.Param() + " characters"
	default:
		return "Validation failed for " + e.Field()
	}
}
