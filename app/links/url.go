package links

import "strings"

// URL returns the full URL for a given path with optional query parameters
func URL(path string, params map[string]string) string {
	return RootURL() + "/" + strings.TrimPrefix(path, "/") + query(params)
}
