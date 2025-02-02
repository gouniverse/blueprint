package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	urlpkg "net/url"
)

// NewRequest options for the new request
type NewRequestOptions struct {
	// Body use this to set the request body
	Body string

	// Context allows setting the request context.
	Context map[any]any

	// ContentType allows setting the Content-Type header
	ContentType string

	// Headers allows setting the request headers
	Headers map[string]string

	// GetValues use this to set the GET values, it will be converted to url.Values
	// and set as the request query
	// Deprecated: use QueryValues
	GetValues urlpkg.Values

	// PostValues use this to set the POST values, it will be converted to url.Values
	// and set as the request body, with Content-Type: application/x-www-form-urlencoded
	// Deprecated: use FormValues
	PostValues urlpkg.Values

	// JSONData sets the request body as application/json.
	// If set, Body and FormValues will be ignored.
	JSONData any

	// QueryParams sets the URL query parameters.
	QueryParams urlpkg.Values

	// FormValues sets the request body as application/x-www-form-urlencoded.
	// If set, Body and JSONData will be ignored.
	FormValues urlpkg.Values
}

// NewRequest creates a new Request for testing, but adds RequestURI
// as the default imlemented in GoLang does not add the RequestURI
// and leaves it to the end user to implement
func NewRequest(method string, url string, opts NewRequestOptions) (*http.Request, error) {
	if url == "" {
		url = "/"
	}

	// Ensure headers is properly initialized
	if opts.Headers == nil {
		opts.Headers = map[string]string{}
	}

	// Keep the backwards compatibility until the deprecated fields are removed
	if opts.GetValues != nil {
		opts.QueryParams = opts.GetValues
	}

	// Keep the backwards compatibility until the deprecated fields are removed
	if opts.PostValues != nil {
		opts.FormValues = opts.PostValues
	}

	// Set Content-Type if not explicitly provided
	if opts.ContentType != "" {
		opts.Headers["Content-Type"] = opts.ContentType
	} else if opts.JSONData != nil {
		opts.Headers["Content-Type"] = "application/json"
	} else if opts.FormValues != nil {
		opts.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	var body *bytes.Buffer

	if opts.Body != "" {
		body = bytes.NewBuffer([]byte(opts.Body))
	} else if opts.JSONData != nil {
		jsonData, err := json.Marshal(opts.JSONData)

		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer([]byte(jsonData))
	} else if opts.FormValues != nil {
		body = bytes.NewBuffer([]byte(opts.FormValues.Encode()))
	} else {
		body = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	u, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}

	req.RequestURI = u.Path

	u.RawQuery = opts.QueryParams.Encode()
	req.URL.RawQuery = u.RawQuery

	// Set headers
	for key, value := range opts.Headers {
		req.Header.Set(key, value)
	}

	// Set context values
	for key, value := range opts.Context {
		ctx := context.WithValue(req.Context(), key, value)
		req = req.WithContext(ctx)
	}

	return req, nil
}
