package emails

import (
	"project/app/links"
	"project/config"

	"github.com/gouniverse/hb"
)

func NewEmailToAdminOnNewContactFormSubmitted() *emailToAdminOnNewContactFormSubmitted {
	return &emailToAdminOnNewContactFormSubmitted{}
}

type emailToAdminOnNewContactFormSubmitted struct{}

// EmailSendOnRegister sends the email when user registers
func (e *emailToAdminOnNewContactFormSubmitted) Send() error {
	emailSubject := config.AppName + ". New Contact Form Submitted"
	emailContent := e.template()

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

func (e *emailToAdminOnNewContactFormSubmitted) template() string {
	urlHome := hb.Hyperlink().
		HTML(config.AppName).
		Href(links.NewWebsiteLinks().Home()).
		ToHTML()

	h1 := hb.Heading1().
		HTML(`New Contact Form Submitted`).
		Style(STYLE_HEADING)

	p1 := hb.Paragraph().
		HTML(`There is a new contact form request submitted into ` + config.AppName + `.`).
		Style(STYLE_PARAGRAPH)

	p2 := hb.Paragraph().
		HTML(`Please login to admin panel to check the new contact request.`).
		Style(STYLE_PARAGRAPH)

	p6 := hb.Paragraph().
		Children([]hb.TagInterface{
			hb.Text(`Thank you for choosing ` + urlHome + `.`),
			hb.BR(),
			hb.Text(`The new way to learn`),
		}).
		Style(STYLE_PARAGRAPH)

	return hb.Div().Children([]hb.TagInterface{
		h1,
		p1,
		p2,
		hb.BR(),
		hb.BR(),
		p6,
	}).ToHTML()
}
