// (c) Jisin0
//
// config/numerals.go creates user configured numbers.

package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	stringTrue = "true"
)

var (
	DBChannel           = int64Environ("DB_CHANNEL")                                                  // id of database channel
	Admins              = int64ListEnviron("ADMINS")                                                  // admins list
	FsubChannels        = int64ListEnviron("FSUB")                                                    // list of force subscribe channels
	BatchSizeLimit      = int64Environ("BATCH_SIZE_LIMIT", 50)                                        // maximum messages allowed in a batch
	DisableNotification = strings.ToLower(os.Getenv("DISABLE_NOTIFICATION")) == stringTrue            // messages will be sent without a notification
	ProtectContent      = strings.ToLower(os.Getenv("PUBLIC_CONTENT")) == stringTrue                  // disable forwarding content from bot
	AllowPublic         = strings.ToLower(os.Getenv("ALLOW_PUBLIC")) == stringTrue || len(Admins) < 1 // indicates wether anyone can use the bot
)

// int64Environ gets a environment variable as an int64.
func int64Environ(envVar string, defaultVal ...int64) int64 {
	s := os.Getenv(envVar)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}

		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("config.int64environ: failed to parse %s: %v\n", envVar, err)
	}

	return i
}

// int64ListEnviron returns a slice of int64 for an environment variable.
func int64ListEnviron(envVar string) (result []int64) {
	s := os.Getenv(envVar)
	if s == "" {
		return result
	}

	for _, str := range strings.Split(s, " ") {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Printf("config.int64listenviron: failed to parse %s: %v\n", envVar, err)
		}

		result = append(result, i)
	}

	return result
}
