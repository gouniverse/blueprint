package layouts

import (
	"github.com/gouniverse/hb"
)

// LogoHTML generates the HTML for the logo
func LogoHTML() string {
	primaryColor := "orange"
	secondaryColor := "white"

	left := hb.Span().
		Style("padding: 5px;").
		Style("color: " + secondaryColor + "; font-size: 20px;").
		HTML("BLUE")

	right := hb.Span().
		Style("padding: 5px;").
		Style("background: " + secondaryColor + "; color: " + primaryColor + ";").
		HTML("PRINT")

	frame := hb.Div().
		Style("display: inline-block; justify-content: space-between; align-items: center; width: fit-content;").
		Style("padding: 0px;").
		Style("border: 3px solid " + primaryColor + "; background: " + primaryColor + "; color: " + secondaryColor + ";").
		Style("font-family: sans-serif; font-size: 20px; letter-spacing: 2px;").
		Child(left).
		Child(right)

	return frame.ToHTML()
}
