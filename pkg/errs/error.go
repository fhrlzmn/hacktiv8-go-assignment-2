package errs

import "net/http"

type error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

type Error interface {
	GetMessage() string
	GetStatus() int
	GetError() string
}

func (e *error) GetMessage() string {
	return e.Message
}

func (e *error) GetStatus() int {
	return e.Status
}

func (e *error) GetError() string {
	return e.Error
}

func NotFound(message string) Error {
	return &error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "NOT_FOUND",
	}
}

func BadRequest(message string) Error {
	return &error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BAD_REQUEST",
	}
}

func InternalServerError(message string) Error {
	return &error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_ERROR",
	}
}

func UnprocessableEntity(message string) Error {
	return &error{
		Message: message,
		Status:  http.StatusUnprocessableEntity,
		Error:   "UNPROCESSABLE_ENTITY",
	}
}
