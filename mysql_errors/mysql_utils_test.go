package mysqlUtils

import (
	"errors"
	"fmt"
	"testing"

	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// TestParseErrorNoRecords aims to test mysqlUtils.ParseError when there is no returned rows
func TestParseErrorNoRecords(t *testing.T) {
	err := errors.New("no rows in result set")
	expected := restErrors.NewNotFoundError("no record matching given id")
	assert.EqualValues(t, expected, ParseError(err))
}

// TestParseErrorParsingData aims to test mysqlUtils.ParseError when there is an error while parsing database response
func TestParseErrorParsingData(t *testing.T) {
	err := errors.New("something went wrong")
	expected := restErrors.NewInternalServerError("error parsing database response")
	assert.EqualValues(t, expected, ParseError(err))
}

// TestParseErrorBadRequest aims to test mysqlUtils.ParseError when there is an invalid request
func TestParseErrorBadRequest(t *testing.T) {
	mysqlErr := mysql.MySQLError{Number: 1062, Message: fmt.Sprintf("Error %d: %s", 1062, "Duplicate entry 'esequiel@gmail.com' for key 'email_UNIQUE'")}
	expected := restErrors.NewBadRequestError("invalid data")
	assert.EqualValues(t, expected, ParseError(&mysqlErr))
}

// TestParseErrorBadRequest aims to test mysqlUtils.ParseError when there is an error processing request on the database
func TestParseErrorProcessingRequest(t *testing.T) {
	mysqlErr := mysql.MySQLError{Number: 1060, Message: fmt.Sprintf("Error %d: %s", 1060, "Duplicate entry 'esequiel@gmail.com' for key 'email_UNIQUE'")}
	expected := restErrors.NewInternalServerError("error processing request on the database")
	assert.EqualValues(t, expected, ParseError(&mysqlErr))
}
