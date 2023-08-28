package env

import (
	"os"
	"strconv"
)

// API client configuration
var ApiRequestTimeout = GetEnvVar("", "ETZ_API_REQUEST_TIMEOUT")

// API server authentication
var ApiAuthMethod = GetEnvVar("", "ETZ_API_AUTH_METHOD")
var ApiToken = GetEnvVar("", "ETZ_API_TOKEN")

func GetEnvVar(defaultValue, key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func GetEnvInt(defaultValue int, key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil || value == 0 {
		return defaultValue
	}
	return value
}
