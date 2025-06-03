// utils/json.go
package utils

import (
	"encoding/json"
	"errors"
	"strings"
)

// SafeJSONPtr validates the string is a valid JSON object or array, returns a pointer to it.
func SafeJSONPtr(s string) (*string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	var js json.RawMessage
	if err := json.Unmarshal([]byte(s), &js); err != nil {
		return nil, errors.New("invalid JSON format")
	}

	return &s, nil
}

// Int64PtrIfNonZero returns a pointer to the int64 if it's not zero, otherwise nil.
func Int64PtrIfNonZero(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}

func NilIfEmpty(s string) *string {
	if len(s) == 0 || len(trimWhitespace(s)) == 0 {
		return nil
	}
	return &s
}

// internal helper to trim all whitespace
func trimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

func ValidateRequiredJSON(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("cannot be empty")
	}
	var js json.RawMessage
	if err := json.Unmarshal([]byte(s), &js); err != nil {
		return "", errors.New("invalid JSON format")
	}
	return s, nil
}
