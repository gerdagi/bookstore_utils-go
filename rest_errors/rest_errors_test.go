package resterrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("This is the error message", errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, "This is the error message", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])

	errBytes, _ := json.Marshal(err)
	fmt.Println(string(errBytes))
}

func TestNewError(t *testing.T) {
	err := NewError("This is the new error message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "This is the new error message", err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("This is the bad request error message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "This is the bad request error message", err.Message)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("user id not found")

	assert.NotNil(t, err)
	assert.EqualValues(t, "user id not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Error)
}
