package layouts

// FaviconURL returns the URL for the favicon
// can be an URL or a base64 encoded data URI
func FaviconURL() string {
	emojiIcon := `♨️`
	svg := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><text y=".9em" font-size="90">` + emojiIcon + `</text></svg>`
	return "data:image/svg+xml," + svg
}
