package out

import (
	"context"
)

// EmailSender defines the port for sending emails
type EmailSender interface {
	// Send sends an email with the given details
	Send(ctx context.Context, to, subject, body string) error
}
