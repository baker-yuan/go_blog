package utils

import (
	"fmt"
	"os"
)

var (
	gitHash string
	version string
)

// GetHashAndVersion get the hash and version
func GetHashAndVersion() (string, string) {
	return gitHash, version
}

// PrintVersion print version and git hash to stdout
func PrintVersion() {
	gitHash, version := GetHashAndVersion()
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Version", version)
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "GitHash", gitHash)
}
