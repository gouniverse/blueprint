package layouts

import (
	"net/http"
	"project/internal/links"

	"github.com/gouniverse/hb"
)

func NewAdminCrudLayout(w http.ResponseWriter, r *http.Request, title string, content string, styleURLs []string, style string, jsURLs []string, js string) string {
	return adminCrudLayout(w, r, title, content, styleURLs, style, jsURLs, js)
}

func adminCrudLayout(_ http.ResponseWriter, r *http.Request, title string, content string, styleURLs []string, style string, jsURLs []string, js string) string {
	jsURLs = append([]string{
		"https://code.jquery.com/jquery-3.6.4.min.js",
		"//code.jquery.com/ui/1.11.4/jquery-ui.js",
		links.URL("/resources/blockarea_v0200.js", map[string]string{}),
	}, jsURLs...)
	styleURLs = append([]string{
		// "https://cdn.datatables.net/1.13.4/css/jquery.dataTables.min.css",
		"//code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css",
	}, styleURLs...)
	// cfmt.Infoln(styleURLs)
	dashboard := NewAdminLayout(r, Options{
		Title:      title,
		Content:    hb.NewHTML(content),
		Scripts:    []string{js},
		ScriptURLs: jsURLs,
		StyleURLs:  styleURLs,
		Styles:     []string{style},
	})
	return dashboard.ToHTML()
}
