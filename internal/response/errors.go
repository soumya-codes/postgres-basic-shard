package response

import (
	"fmt"
	"net/http"
	"time"
)

// Define error codes using iota
type ErrorCode int

const (
	ErrCodeUnknown ErrorCode = iota
	ErrCodeNotFound
	ErrCodeInvalidInput
	ErrCodeInternalServerError
)

// Error messages corresponding to the error codes
var errorMessages = map[ErrorCode]string{
	ErrCodeUnknown:             "An unknown error occurred",
	ErrCodeNotFound:            "Resource not found",
	ErrCodeInvalidInput:        "Invalid input parameters",
	ErrCodeInternalServerError: "Internal server error",
}

// CustomResponse represents the structure of an error response
type CustomResponse struct {
	RespCode      int `json:"resp_code"`
	ResponseError `json:"custom_error"`
}

func (e CustomResponse) Error() string {
	return fmt.Sprintf("respcode: %d, code: %d, message: %s, details: %v", e.RespCode, e.Code, e.Message, e.Details)
}

type ResponseError struct {
	Code      string        `json:"code"`
	Message   string        `json:"message"`
	Details   []ErrorDetail `json:"details,omitempty"`
	Timestamp string        `json:"timestamp"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, details: %v", e.Code, e.Message, e.Details)
}

// ErrorDetail represents the structure of detailed error information
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// MapErrorToErrorResponse maps an error to our standard ResponseError
func MapErrorToErrorResponse(errCode ErrorCode, errDetail ErrorDetail) CustomResponse {
	errorResponse := CustomResponse{
		ResponseError: ResponseError{
			Details:   []ErrorDetail{errDetail},
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
	switch errCode {
	case ErrCodeNotFound:
		errorResponse.RespCode = http.StatusNotFound
		errorResponse.Code = "not_found"
		errorResponse.Message = errorMessages[ErrCodeNotFound]
	case ErrCodeInvalidInput:
		errorResponse.RespCode = http.StatusBadRequest
		errorResponse.Code = "invalid_input"
		errorResponse.Message = errorMessages[ErrCodeInvalidInput]
	case ErrCodeInternalServerError:
		errorResponse.RespCode = http.StatusInternalServerError
		errorResponse.Code = "internal_server_error"
		errorResponse.Message = errorMessages[ErrCodeInternalServerError]
	default:
		errorResponse.RespCode = http.StatusInternalServerError
		errorResponse.Code = "unknown_error"
		errorResponse.Message = errorMessages[ErrCodeUnknown]
	}

	return errorResponse
}
