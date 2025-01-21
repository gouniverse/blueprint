package testutils

import (
	"net/http"
	"net/http/httptest"
)

func CallMiddleware(method string, middleware func(next http.Handler) http.Handler, next func(w http.ResponseWriter, r *http.Request), options NewRequestOptions) (body string, response *http.Response, err error) {
	req, err := NewRequest(method, "/", options)

	if err != nil {
		return "", nil, err
	}

	recorder := httptest.NewRecorder()
	handler := middleware(http.HandlerFunc(next))
	handler.ServeHTTP(recorder, req)
	body = recorder.Body.String()

	return body, recorder.Result(), nil
}
