package task09specificerror

import (
	"errors"
	"testing"
)

func TestGetEmailReturnsErrUserNotFound(t *testing.T) {
	users := map[string]string{
		"sasha": "sasha@example.com",
	}

	_, err := GetEmail(users, "jack")
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("GetEmail error = %v, want error matching %v", err, ErrUserNotFound)
	}
}

func TestGetEmailSuccess(t *testing.T) {
	users := map[string]string{
		"sasha": "sasha@example.com",
	}

	if got, err := GetEmail(users, "sasha"); err != nil || got != "sasha@example.com" {
		t.Fatalf("GetEmail() = (%q, %v), want (%q, nil)", got, err, "sasha@example.com")
	}
}
