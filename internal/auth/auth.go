package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API key from the
// headers of an HTTP request
// Example:
// Authorization: ApiKey {apikey}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("auth info could not be found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("first part of auth header is not 'ApiKey'")
	}

	return vals[1], nil
}
