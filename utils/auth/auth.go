// (c) Jisin0
//
// utils/auth.go contains Authorization methods.

package auth

import "github.com/Jisin0/TGMessageStore/config"

// CheckUser checks if a user is allowed to use the bot.
// It returns true if ALLOW_PUBLIC is set to true or checks the Admins list.
// An empty admin list would also result in a success.
func CheckUser(userID int64) bool {
	if !config.AllowPublic {
		return false
	}

	for _, i := range config.Admins {
		if i == userID {
			return true
		}
	}

	return false
}
