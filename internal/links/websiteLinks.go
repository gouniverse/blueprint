package links

type websiteLinks struct{}

func NewWebsiteLinks() *websiteLinks {
	return &websiteLinks{}
}

func (l *websiteLinks) Home() string {
	return URL(HOME, map[string]string{})
}

func (l *websiteLinks) Blog(params map[string]string) string {
	return URL(BLOG, params)
}

func (l *websiteLinks) BlogPost(postID string, postSlug string) string {
	uri := BLOG_POST
	uri += "/" + postID
	uri += "/" + postSlug
	return URL(uri, map[string]string{})
}

func (l *websiteLinks) Contact() string {
	return URL(CONTACT, map[string]string{})
}

func (l *websiteLinks) Flash(params map[string]string) string {
	return URL(FLASH, params)
}

func (l *websiteLinks) Theme(params map[string]string) string {
	return URL(THEME, params)
}

func (l *websiteLinks) Widget(alias string, params map[string]string) string {
	params["alias"] = alias
	return URL(WIDGET, params)
}
