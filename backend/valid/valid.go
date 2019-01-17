package valid

import (
	"regexp"
)

// Username validates a given username, which:
// must have a length between 5 and 25
// must start with a letter
// must finish with a letter or a digit
// can contain only letters, digits, hyphens and underscores.
func Username(username string) bool {
	if username == "" {
		return false
	}

	return regexp.MustCompile("[a-zA-Z][a-zA-Z0-9-_]{3,23}[a-zA-Z0-9]").FindString(username) == username
}

// Password validates a given password, which:
// must have a length between 8 and 25
// can contain only letters, digits, and the following characters: -_?#!£$%&.
func Password(password string) bool {
	if password == "" {
		return false
	}

	return regexp.MustCompile("[a-zA-Z0-9-_?#!£$%&]{8,24}").FindString(password) == password
}
