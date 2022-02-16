package mysqlUtils

import (
	"strings"

	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

// ParseError is responsible to handle mysql errors
func ParseError(err error) restErrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return restErrors.NewNotFoundError("no record matching given id")
		}

		return restErrors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return restErrors.NewBadRequestError("invalid data")
	}

	return restErrors.NewInternalServerError("error processing request on the database")
}
