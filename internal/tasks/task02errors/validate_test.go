package task02errors

import "testing"

func TestValidateAge(t *testing.T) {
	if err := ValidateAge(25); err != nil {
		t.Fatalf("ValidateAge(25) returned unexpected error: %v", err)
	}

	if err := ValidateAge(-1); err == nil {
		t.Fatal("ValidateAge(-1) returned nil error, expected non-nil error")
	}
}
