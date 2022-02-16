package cryptoutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSha256 is a testing.T that verifies GetSha256 crypto output
func TestGetSha256(t *testing.T) {
	assert.EqualValues(t, "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", GetSha256("password"))
	assert.NotEqualValues(t, "test", GetSha256("password"))
}
