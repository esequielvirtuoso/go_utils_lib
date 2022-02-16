package restErrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

// RestErr is a standard struct to be used while handling errors in REST applications.
type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// NewError returns a new error with an input message
func NewError(msg string) error {
	return errors.New(msg)
}

/* NewBadRequestError returns a standardized struct with the correct status,
*  and error tag for bad request situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

/* NewNotFoundError returns a standardized struct with the correct status,
*  and error tag for not found situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

/* NewInternalServerError returns a standardized struct with the correct status,
*  and error tag for internal error situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}

	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

/* NewUnauthorized returns a standardized struct with the correct status,
*  and error tag for unauthorized access situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewUnauthorized(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}
