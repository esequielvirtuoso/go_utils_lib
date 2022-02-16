// Package cryptoutils handle crypto operations
package cryptoutils

import (
	"crypto/sha256"
	"encoding/hex"
)

// GetSha256 encrypts the input string with GetSha256 algorithm
func GetSha256(input string) string {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
