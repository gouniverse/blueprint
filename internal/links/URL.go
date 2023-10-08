package links

import "strings"

func URL(path string, params map[string]string) string {
	return RootURL() + "/" + strings.TrimPrefix(path, "/") + query(params)
}
