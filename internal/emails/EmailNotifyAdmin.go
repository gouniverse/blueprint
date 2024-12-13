package emails

import (
	"project/config"
)

func NewEmailNotifyAdmin() *emailNotifyAdmin {
	return &emailNotifyAdmin{}
}

type emailNotifyAdmin struct{}

// EmailSendOnRegister sends the email when user registers
func (e *emailNotifyAdmin) Send(html string) error {
	emailSubject := config.AppName + ". Admin Notification"
	emailContent := html

	finalHtml := blankEmailTemplate(emailSubject, emailContent)

	recipientEmail := "info@sinevia.com"

	errSend := Send(SendOptions{
		From:     config.MailFromEmailAddress,
		FromName: config.AppName,
		To:       []string{recipientEmail},
		Subject:  emailSubject,
		HtmlBody: finalHtml,
	})
	return errSend
}
