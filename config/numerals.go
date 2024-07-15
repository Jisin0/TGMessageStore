// (c) Jisin0
//
// File config/numerals.go creates user configured numbers.

package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	DBChannel int64 = int64Environ("DB_CHANNEL")
)

// int64Environ gets a environment variable and attempts to cast it into an int64.
func int64Environ(envVar string) int64 {
	s := os.Getenv(envVar)
	if s == "" {
		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("config.int64environ: failed to parse %s: %v\n", envVar, err)
	}

	return i
}
