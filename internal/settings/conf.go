package settings

/*
 * Config of the app like setting environ variables.
 */

import (
	"os"
	"strconv"
)

// getEnv gets an environment or a fallback.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

// DebugLog debug log or not
var DebugLog, _ = strconv.ParseBool(getEnv("DEBUG_LOG", "false"))
