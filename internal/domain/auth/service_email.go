package auth

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v3"
)

type EmailService interface {
	SendVerificationEmail(mail string, verificationURL string) error
}

type emailService struct {
	client *resend.Client
}

func (s *emailService) SendVerificationEmail(mail string, verificationURL string) error {
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{mail},
		Subject: "Your verification code from the Group Expenses App.",
		Html: fmt.Sprintf(`
			<div>
				<p>Click the following link to verify your account:</p>
				<a href="%s">Verify account</a>
			</div>
		`, verificationURL),
	}
	_, err := s.client.Emails.Send(params)
	return err
}

func NewEmailService() EmailService {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	return &emailService{
		client: client,
	}
}
