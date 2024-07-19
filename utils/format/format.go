// (c) Jisin0
//
// utils/format.go contains methods to format strings using a python-like format.

package format

import (
	"fmt"
	"strings"

	"github.com/Jisin0/TGMessageStore/utils/helpers"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// FormatString returns a string formatted with the values from the values.
func FormatString(format string, values map[string]any) string {
	if len(values) < 1 {
		return format
	}

	var result strings.Builder

	length := len(format)

	for i := 0; i < length; {
		if format[i] == '{' {
			end := strings.Index(format[i:], "}")
			if end == -1 {
				result.WriteString(format[i:])
				break
			}

			key := format[i+1 : i+end]
			if value, ok := values[key]; ok {
				result.WriteString(fmt.Sprint(value))
			} else {
				result.WriteString(format[i : i+end+1])
			}

			i += end + 1
		} else {
			result.WriteByte(format[i])
			i++
		}
	}

	return result.String()
}

// BasicFormat formats a string with the mention, name and user_id values.
func BasicFormat(format string, user *gotgbot.User, extraParams ...map[string]any) string {
	if user == nil {
		return format
	}

	userID := fmt.Sprint(user.Id)
	name := FullName(user)

	var mention string
	if user.Username != "" {
		mention = "@" + user.Username
	} else {
		mention = fmt.Sprintf("<a href='tg://user?id=%d'>%s</a>", user.Id, name)
	}

	values := map[string]any{
		"user_id": userID,
		"name":    name,
		"mention": mention,
	}

	if len(extraParams) > 0 {
		helpers.MergeMaps(values, extraParams[0])
	}

	return FormatString(format, values)
}

// FullName returns the full name of a user.
func FullName(user *gotgbot.User) (s string) {
	s = user.FirstName

	if user.LastName != "" {
		s = s + " " + user.LastName
	}

	return s
}

// Mention creates a html string that mentions the user.
func Mention(user *gotgbot.User) (s string) {
	if user.Username != "" {
		s = "@" + user.Username
	} else {
		s = fmt.Sprintf("<a href='tg://user?id=%d'>%s</a>", user.Id, FullName(user))
	}

	return s
}
