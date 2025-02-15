package ext

import (
	"net/http"

	"github.com/gouniverse/base/req"
)

// IsHtmx checks if the given HTTP request is an HTMX request.
//
// An HTMX request is determined by the presence of the "HX-Request" header.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - bool: true if it's an HTMX request, false otherwise
func IsHtmx(r *http.Request) bool {
	return r.Header.Get("HX-Request") != ""
}

// IsHxBoosted checks if the request has been boosted for performance.
//
// Returns true if the "HX-Boosted" header is present and its value is "true",
// otherwise returns false.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - bool: true if the request has been boosted, false otherwise
func IsHxBoosted(r *http.Request) bool {
	value := req.ValueOr(r, "HX-Boosted", "")
	return value == "true"
}

// IsHxHistoryRestoreRequest checks if the request is related to restoring the browser's history.
//
// Returns true if the "HX-History-Restore-Request" header is present and its
// value is "true", otherwise returns false.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - bool: true if the request is a history restore request, false otherwise
func IsHxHistoryRestoreRequest(r *http.Request) bool {
	value := req.ValueOr(r, "HX-History-Restore-Request", "")
	return value == "true"
}

// IsHxRequest checks if this is an HTMX request.
//
// Returns true if the "HX-Request" header is present and its value is "true",
// otherwise returns false.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - bool: true if it's an HTMX request, false otherwise
func IsHxRequest(r *http.Request) bool {
	value := req.ValueOr(r, "HX-Request", "")
	return value == "true"
}

// IsHxTrigger checks if the request was triggered by an event.
//
// Returns true if the "HX-Trigger" header is present and its value is "true",
// otherwise returns false.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - bool: true if the request was triggered, false otherwise
func IsHxTrigger(r *http.Request) bool {
	value := req.ValueOr(r, "HX-Trigger", "")
	return value == "true"
}

// HxPrompt gets the prompt message for the user.
//
// Returns the value of the "HX-Prompt" header, or an empty string if the
// header is not present.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - string: the prompt message
func HxPrompt(r *http.Request) string {
	return req.ValueOr(r, "HX-Prompt", "")
}

// HxTarget gets the target element for the response.
//
// Returns the value of the "HX-Target" header, or an empty string if the
// header is not present.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - string: the target element
func HxTarget(r *http.Request) string {
	return req.ValueOr(r, "HX-Target", "")
}

// HxTriggerName gets the name of the event or trigger that initiated the request.
//
// Returns the value of the "HX-Trigger-Name" header, or an empty string
// if the header is not present.
//
// Parameters:
//   - r: the http request
//
// Returns:
//   - string: the trigger name
func HxTriggerName(r *http.Request) string {
	return req.ValueOr(r, "HX-Trigger-Name", "")
}
