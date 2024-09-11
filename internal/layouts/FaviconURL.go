package layouts

// FaviconURL returns the URL for the favicon
// can be an URL or a base64 encoded data URI
func FaviconURL() string {
	return "data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>♨️</text></svg>"
}
