package links

import "strings"

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

func (l *websiteLinks) PaymentCanceled(paymentKey string) string {
	params := map[string]string{}
	params["payment_key"] = paymentKey
	return URL(PAYMENT_CANCELED, params)
}

func (l *websiteLinks) PaymentSuccess(paymentKey string) string {
	params := map[string]string{}
	params["payment_key"] = paymentKey
	return URL(PAYMENT_SUCCESS, params)
}

func (l *websiteLinks) Resource(resourcePath string, params map[string]string) string {
	if resourcePath == "" {
		return ""
	}
	if resourcePath[0] != '/' {
		resourcePath = "/" + resourcePath
	}

	return URL(RESOURCES+resourcePath, params)
}

func (l *websiteLinks) Theme(params map[string]string) string {
	return URL(THEME, params)
}

func (l *websiteLinks) Thumbnail(extension, width, height, quality, path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		url := strings.ReplaceAll(path, "https://", "https/")
		path = strings.ReplaceAll(url, "http://", "http/")
	}
	return RootURL() + "/th/" + extension + "/" + width + "x" + height + "/" + quality + "/" + path
}

func (l *websiteLinks) Widget(alias string, params map[string]string) string {
	params["alias"] = alias
	return URL(WIDGET, params)
}
