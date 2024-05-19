package links

type websiteLinks struct{}

func NewWebsiteLinks() *websiteLinks {
	return &websiteLinks{}
}

func (l *websiteLinks) Home() string {
	return URL(HOME, map[string]string{})
}

func (l *websiteLinks) Blog() string {
	return URL(BLOG, map[string]string{})
}

func (l *websiteLinks) BlogP(params map[string]string) string {
	return URL(BLOG, params)
}

func (l *websiteLinks) BlogPost(postId string, postTitle string) string {
	return URL(BLOG_POST+"/"+postId+"/"+postTitle, map[string]string{})
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
