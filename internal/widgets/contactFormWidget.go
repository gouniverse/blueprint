package widgets

import (
	"net/http"

	"github.com/gouniverse/cms"
	"github.com/gouniverse/hb"
)

// NewContactForm creates a new instance of the contactForm widget.
//
// Parameters:
//   - None
//
// Returns:
//   - *contactForm - A pointer to the contactForm shrtcode
func NewContactFormWidget() *contactFormWidget {
	return &contactFormWidget{}
}

var _ cms.ShortcodeInterface = (*contactFormWidget)(nil) // verify it extends the interface

// contactForm is the struct that will be used to render the contactForm shortcode.
//
// This shortcode is used to send a contact message from the website.
type contactFormWidget struct{}

// Alias the shortcode alias to be used in the template.
func (widget *contactFormWidget) Alias() string {
	return "x-contact-form"
}

// Description a user-friendly description of the shortcode.
func (widget *contactFormWidget) Description() string {
	return "Renders the contact form"
}

// Render implements the shortcode interface.
func (widget *contactFormWidget) Render(r *http.Request, content string, params map[string]string) string {
	path := r.URL.Path

	return widget.form(path)
}

func (widget *contactFormWidget) form(path string) string {
	_ = path // not used currently
	link := "https://tiny.vip/swij"
	return hb.NewHyperlink().
		Href(link).
		HTML("Open My Contact Form").
		Target("_blank").
		Class("btn btn-primary").
		ToHTML()
}
