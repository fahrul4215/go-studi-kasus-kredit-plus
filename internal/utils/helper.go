package utils

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// Convert UUID string to binary(16)
func UuidToBinary(u uuid.UUID) ([]byte, error) {
	bytes, err := u.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Convert binary(16) back to UUID string
func BinaryToUUID(b []byte) (string, error) {
	u, err := uuid.FromBytes(b)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase convert string to snake case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
