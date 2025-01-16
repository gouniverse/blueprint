package testutils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gouniverse/responses"
)

func CallHtmlEndpoint(method string, f func(w http.ResponseWriter, r *http.Request) string, options NewRequestOptions) (body string, response *http.Response, err error) {
	req, err := NewRequest(method, "/", options)

	if err != nil {
		return "", nil, err
	}

	recorder := httptest.NewRecorder()
	handler := http.Handler(responses.HTMLHandler(f))
	handler.ServeHTTP(recorder, req)
	body = recorder.Body.String()

	return body, recorder.Result(), nil
}
