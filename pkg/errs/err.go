package errs

import "net/http"

type Error interface {
	Message() string
	Status() int
	Error() string
}

type ErrorData struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func (e *ErrorData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewBadRequest(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus: http.StatusBadRequest,
		ErrError: "BAD_REQUEST",
	}
}

func NewUnprocessableEntityResponse(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: "INVALID_REQUEST_BODY",
	}
}

func NewNotFoundError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus: http.StatusNotFound,
		ErrError: "NOT_FOUND",
	}
}

func NewUnauthenticatedError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus: http.StatusUnauthorized,
		ErrError: "NOT_AUTHORIZED",
	}
}
