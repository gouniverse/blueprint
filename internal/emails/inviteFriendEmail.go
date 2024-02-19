package emails

import (
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
	user, err := config.UserStore.UserFindByID(sendingUserID)

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

	urlHome := hb.NewHyperlink().Text("ProvedExpert").
		Href(links.NewWebsiteLinks().Home()).ToHTML()

	urlJoin := hb.NewHyperlink().Text("Click to Join Me at ProvedExpert").
		Href(links.NewWebsiteLinks().Home()).ToHTML()

	h1 := hb.NewHeading1().
		HTML(`You have an awesome friend`).
		Style(STYLE_HEADING)

	p1 := hb.NewParagraph().
		HTML(`Hi ` + recipientName + `,`).
		Style(STYLE_PARAGRAPH)

	p2 := hb.NewParagraph().
		HTML(`You have been invited by a friend who thinks you will like ` + config.AppName + `.`).
		Style(STYLE_PARAGRAPH)

	p3 := hb.NewParagraph().
		HTML(`A note from your friend ` + userName + `:`).
		Style(STYLE_PARAGRAPH)

	p4 := hb.NewParagraph().
		HTML(`"` + userNote + `"`).
		Style(STYLE_PARAGRAPH)

	p5 := hb.NewParagraph().
		HTML(urlJoin).
		Style(STYLE_PARAGRAPH)

	p6 := hb.NewParagraph().
		HTML(``). // Add description
		Style(STYLE_PARAGRAPH)

	p7 := hb.NewParagraph().
		Children([]*hb.Tag{
			hb.NewHTML(`Thank you for choosing ` + urlHome + `.`),
		}).
		Style(STYLE_PARAGRAPH)

	return hb.NewDiv().Children([]*hb.Tag{
		h1,
		p1,
		p2,
		p3,
		p4,
		p5,
		p6,
		hb.NewBR(),
		hb.NewBR(),
		p7,
	}).ToHTML()
}
