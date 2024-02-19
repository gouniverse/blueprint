package widgets

import "net/http"

type Widget interface {
	Alias() string
	Render(req *http.Request, content string, data map[string]string) string
}
