package restErrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewError aims to validate restErrors.NewError return.
func TestNewError(t *testing.T) {
	expectedErr := errors.New("something went wrong")
	assert.EqualValues(t, expectedErr, NewError("something went wrong"))
}

// TestNewBadRequestError aims to validate restErrors.NewBadRequestError return.
func TestNewBadRequestError(t *testing.T) {
	expected := restErr{
		ErrMessage: "invalid request",
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}

	assert.EqualValues(t, expected, NewBadRequestError("invalid request"))
}

// TestNewNotFoundError aims to validate restErrors.NewNotFoundError return.
func TestNewNotFoundError(t *testing.T) {
	expected := restErr{
		ErrMessage: "not found",
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}

	assert.EqualValues(t, expected, NewNotFoundError("not found"))
}

// TestNewInternalServerError aims to validate restErrors.NewInternalServerError return.
func TestNewInternalServerError(t *testing.T) {
	expected := restErr{
		ErrMessage: "internal server error",
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
		ErrCauses:  []interface{}{"database error"},
	}

	actualErr := NewInternalServerError("internal server error", errors.New("database error"))
	assert.EqualValues(t, expected, actualErr)
	assert.NotNil(t, actualErr.Causes())
	assert.EqualValues(t, 1, len(actualErr.Causes()))
	assert.EqualValues(t, "database error", actualErr.Causes()[0])
}

// TestNewInternalServerErrorEmptyMessage aims to validate restErrors.NewInternalServerError return when no message parameter is passed.
func TestNewInternalServerErrorEmptyMessage(t *testing.T) {
	actualErr := NewInternalServerError("", errors.New("database error"))
	assert.Empty(t, actualErr.Message())
}

// TestNewUnauthorized aims to validate restErrors.NewUnauthorized return.
func TestNewUnauthorized(t *testing.T) {
	expected := restErr{
		ErrMessage: "unauthorized error",
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}

	assert.EqualValues(t, expected, NewUnauthorized("unauthorized error"))
}
