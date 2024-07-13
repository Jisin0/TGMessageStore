// (c) Jisin0
// Helper methods.

package plugins

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// Returns a html formatted string that mention's the user
func mention(u *gotgbot.User) string {
	name := u.FirstName
	if u.LastName != "" {
		name = name + " " + u.LastName
	}

	return fmt.Sprintf("<a href='tg://user?id=%v'>%v</a>", u.Id, name)
}

// Checks if a string slice Contains an item.
func Contains(l []string, v string) bool {
	for _, i := range l {
		if i == v {
			return true
		}
	}

	return false
}
