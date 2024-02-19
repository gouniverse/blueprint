package testutils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gouniverse/responses"
)

func CallHtmlEndpoint(method string, f func(w http.ResponseWriter, r *http.Request) string, options NewRequestOptions) (response string, res *http.Response, err error) {
	req, err := NewRequest(method, "/", options)

	if err != nil {
		return "", nil, err
	}

	recorder := httptest.NewRecorder()
	handler := http.Handler(responses.HTMLHandler(f))
	handler.ServeHTTP(recorder, req)
	response = recorder.Body.String()

	return response, recorder.Result(), nil
}
