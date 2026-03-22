package task02errors

import "fmt"

func ValidateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age cannot be negative: %d", age)
	}

	return nil
}
