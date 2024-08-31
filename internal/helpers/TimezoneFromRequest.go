package helpers

import "net/http"

// TimezoneFromRequest returns the timezone from the authenticated user
// or the default timezone (UTC) if the user is not authenticated.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - string: the timezone
func TimezoneFromRequest(r *http.Request) string {
	defaultTimezone := "UTC"
	user := GetAuthUser(r)

	if user == nil {
		return defaultTimezone
	}

	userTimezone := user.Timezone()

	if userTimezone != "" {
		return userTimezone
	}

	return defaultTimezone
}
