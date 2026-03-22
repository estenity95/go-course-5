package task09specificerror

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func GetEmail(users map[string]string, username string) (string, error) {
	email, ok := users[username]
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrUserNotFound, username)
	}

	return email, nil
}
