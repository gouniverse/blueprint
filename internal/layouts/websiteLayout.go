package layouts

// import (
// 	"net/http"

// 	"github.com/gouniverse/hb"
// )

// func NewWebsiteLayout(options Options) *websiteLayout {
// 	layout := &websiteLayout{}
// 	layout.title = options.Title // + " | " + config.AppName
// 	layout.content = options.Content
// 	layout.scriptURLs = options.ScriptURLs
// 	layout.scripts = options.Scripts
// 	layout.styleURLs = options.StyleURLs
// 	layout.styles = options.Styles
// 	return layout
// }

// type websiteLayout struct {
// 	request        *http.Request
// 	websiteSection string // i.e. Blog, Website
// 	title          string
// 	content        *hb.Tag
// 	scriptURLs     []string
// 	scripts        []string
// 	styleURLs      []string
// 	styles         []string
// }

// // NewWebsiteLayout creates a new WebsiteLayout object.
// //
// // Parameters:
// // - none
// //
// // Returns:
// // - *WebsiteLayout: a pointer to the WebsiteLayout object.
// func NewWebsiteOldLayout() *websiteLayout {
// 	return &websiteLayout{}
// }

// func (l *websiteLayout) fnLayout(r *http.Request, content string) string {
// 	return content
// }
