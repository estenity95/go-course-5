package task06mock

import "fmt"

type User struct {
	Name  string
	Email string
}

type EmailSender interface {
	Send(to, subject, body string) error
}

func SendEmail(sender EmailSender, user User) error {
	if user.Email == "" {
		return fmt.Errorf("email is empty")
	}

	subject := fmt.Sprintf("Welcome, %s!", user.Name)
	body := "Your account has been created."

	return sender.Send(user.Email, subject, body)
}
