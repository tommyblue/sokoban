package utils

import "os"

// IsDebugEnv identifies debugging env. Run with `DEBUG=1 ./command`
func IsDebugEnv() bool {
	return os.Getenv("DEBUG") == "1"
}
