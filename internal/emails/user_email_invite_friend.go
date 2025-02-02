package emails

import (
	"context"
	"errors"
	"project/config"
	"project/internal/links"

	"github.com/gouniverse/hb"
)

func NewInviteFriendEmail() *inviteFriendEmail {
	return &inviteFriendEmail{}
}

type inviteFriendEmail struct{}

// EmailSendOnRegister sends the email when user registers
func (e *inviteFriendEmail) Send(sendingUserID string, userNote string, recipientEmail string, recipientName string) error {
	if config.UserStore == nil {
		return errors.New("user store not configured")
	}

	user, err := config.UserStore.UserFindByID(context.Background(), sendingUserID)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	userName := user.FirstName()

	if userName == "" {
		userName = user.Email()
	}

	emailSubject := config.AppName + ". Invitation by a Friend"
	emailContent := e.template(userName, userNote, recipientName)

	finalHtml := blankEmailTemplate(emailSubject, emailContent)

	errSend := Send(SendOptions{
		From:     config.MailFromEmailAddress,
		FromName: config.MailFromName,
		To:       []string{recipientEmail},
		Subject:  emailSubject,
		HtmlBody: finalHtml,
	})
	return errSend
}

func (e *inviteFriendEmail) template(userName string, userNote string, recipientName string) string {

	urlHome := hb.Hyperlink().Text("ProvedExpert").
		Href(links.NewWebsiteLinks().Home()).ToHTML()

	urlJoin := hb.Hyperlink().Text("Click to Join Me at ProvedExpert").
		Href(links.NewWebsiteLinks().Home()).ToHTML()

	h1 := hb.Heading1().
		HTML(`You have an awesome friend`).
		Style(STYLE_HEADING)

	p1 := hb.Paragraph().
		HTML(`Hi ` + recipientName + `,`).
		Style(STYLE_PARAGRAPH)

	p2 := hb.Paragraph().
		HTML(`You have been invited by a friend who thinks you will like ` + config.AppName + `.`).
		Style(STYLE_PARAGRAPH)

	p3 := hb.Paragraph().
		HTML(`A note from your friend ` + userName + `:`).
		Style(STYLE_PARAGRAPH)

	p4 := hb.Paragraph().
		HTML(`"` + userNote + `"`).
		Style(STYLE_PARAGRAPH)

	p5 := hb.Paragraph().
		HTML(urlJoin).
		Style(STYLE_PARAGRAPH)

	p6 := hb.Paragraph().
		HTML(``). // Add description
		Style(STYLE_PARAGRAPH)

	p7 := hb.Paragraph().
		Children([]hb.TagInterface{
			hb.Raw(`Thank you for choosing ` + urlHome + `.`),
		}).
		Style(STYLE_PARAGRAPH)

	return hb.Div().Children([]hb.TagInterface{
		h1,
		p1,
		p2,
		p3,
		p4,
		p5,
		p6,
		hb.BR(),
		hb.BR(),
		p7,
	}).ToHTML()
}
