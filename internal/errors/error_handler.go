package errors

import (
	"fmt"
)

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func New(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func ErrorResponse(err *CustomError) map[string]interface{} {
	return map[string]interface{}{
		"message": err.Message,
	}
}

func HandleError(err error) (map[string]interface{}, int) {
	// If the error is a CustomError, return it directly
	if customErr, ok := err.(*CustomError); ok {
		return ErrorResponse(customErr), customErr.StatusCode
	}

	fmt.Printf("Internal server error: %v\n", err)

	return ErrorResponse(New("An unexpected error occurred. Please try again later.", 500)), 500
}
