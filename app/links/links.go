package links

import (
	"net/url"
	"project/config"
)

// RootURL returns a URL to the current website
func RootURL() string {
	appURL := config.AppUrl
	if config.IsEnvTesting() {
		appURL = ""
	}
	return appURL
}

func query(queryData map[string]string) string {
	queryString := ""

	if len(queryData) > 0 {
		v := url.Values{}
		for key, value := range queryData {
			v.Set(key, value)
		}
		queryString += "?" + httpBuildQuery(v)
	}

	return queryString
}

func httpBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}
