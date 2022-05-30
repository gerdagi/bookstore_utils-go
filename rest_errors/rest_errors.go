package resterrors

import (
	"errors"
	"fmt"
	"net/http"
)

var ()

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message %s - status: %d - error: %s - causes: [%v]", e.message, e.status, e.error, e.causes)
}

func (r restError) Message() string {
	return r.message
}

func (r restError) Status() int {
	return r.status
}

func (r restError) Causes() []interface{} {
	return r.causes
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) RestErr {
	return &restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return &restError{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restError{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

func NewUnauthorizedError(message string) RestErr {
	return &restError{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := &restError{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}

	if err != nil {
		result.causes = append(result.causes, err.Error())
	}

	return result
}
