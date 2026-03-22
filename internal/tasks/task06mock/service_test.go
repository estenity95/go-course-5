package task06mock

import (
	"errors"
	"testing"
)

type emailSenderMock struct {
	called  bool
	to      string
	subject string
	body    string
	err     error
}

func (m *emailSenderMock) Send(to, subject, body string) error {
	m.called = true
	m.to = to
	m.subject = subject
	m.body = body
	return m.err
}

func TestSendEmail(t *testing.T) {
	mock := &emailSenderMock{}
	user := User{Name: "Sasha", Email: "sasha@example.com"}

	if err := SendEmail(mock, user); err != nil {
		t.Fatalf("SendEmail returned unexpected error: %v", err)
	}

	if !mock.called {
		t.Fatal("email sender was not called")
	}

	if mock.to != user.Email {
		t.Fatalf("email sender received to = %q, want %q", mock.to, user.Email)
	}

	if mock.subject != "Welcome, Sasha!" {
		t.Fatalf("email sender received subject = %q, want %q", mock.subject, "Welcome, Alice!")
	}
}

func TestSendEmailPropagatesSenderError(t *testing.T) {
	expectedErr := errors.New("email service unavailable")
	mock := &emailSenderMock{err: expectedErr}

	err := SendEmail(mock, User{Name: "Bob", Email: "bob@example.com"})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("SendEmail error = %v, want %v", err, expectedErr)
	}
}
