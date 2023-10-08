package links

type websiteLinks struct{}

func NewWebsiteLinks() *websiteLinks {
	return &websiteLinks{}
}

func (l *websiteLinks) Home() string {
	return URL(HOME, map[string]string{})
}

func (l *websiteLinks) Flash(params map[string]string) string {
	return URL(FLASH, params)
}

func (l *websiteLinks) Theme(params map[string]string) string {
	return URL(THEME, params)
}
