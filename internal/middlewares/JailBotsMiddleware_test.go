package middlewares

import (
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"project/config"
	"testing"

	"github.com/gouniverse/auth/tests"
	"github.com/gouniverse/responses"
	"github.com/spf13/cast"
)

func TestJailBotsMiddlewareName(t *testing.T) {

	config.TestsConfigureAndInitialize()

	// Act
	m := NewJailBotsMiddleware()

	// Assert
	if m.Name != "Jail Bots Middleware" {
		t.Fatal("JailBotsMiddleware.Name must be Jail Bots Middleware. Got ", m.Name)
	}
}

func TestJailBotsMiddlewareAllowedResponse(t *testing.T) {
	config.TestsConfigureAndInitialize()

	allowedUris := []string{
		"/robots.txt",
		"/sitemap.xml",
		"/favicon.ico",
		"/",
		"/auth/login",
	}

	// Act

	for _, allowedUri := range allowedUris {
		m := NewJailBotsMiddleware()
		req, err := tests.NewRequest("GET", allowedUri, tests.NewRequestOptions{})

		if err != nil {
			t.Fatal(err)
		}

		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// DEBUG: t.Log("Passes as expected")
			w.WriteHeader(http.StatusOK)
			responses.HTMLResponse(w, r, "gone through")
		})

		rw := httptest.NewRecorder()
		handler := m.Handler(testHandler)
		handler.ServeHTTP(rw, req)

		// Assert
		body := rw.Body.String()

		if body != "gone through" {
			t.Fatal("Response SHOULD NOT BE 'gone through' but found: ", rw.Body.String())
		}
	}
}

func TestJailBotsMiddlewareJailedResponse(t *testing.T) {
	config.TestsConfigureAndInitialize()

	allowedUris := []string{
		"/.env",
		"/backup",
		"/db",
		"/wp-admin",
	}

	// Act

	for _, allowedUri := range allowedUris {
		m := NewJailBotsMiddleware()
		randInt := rand.IntN(1000)
		req, err := tests.NewRequest("GET", allowedUri, tests.NewRequestOptions{
			Headers: map[string]string{
				"User-Agent":      "test-agent",
				"Referer":         "test-referer",
				"X-Forwarded-For": "127.0.0." + cast.ToString(randInt),
			},
		})

		if err != nil {
			t.Fatal(err)
		}

		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// DEBUG: t.Log("Passes as expected")
			w.WriteHeader(http.StatusOK)
			responses.HTMLResponse(w, r, "gone through")
		})

		rw := httptest.NewRecorder()
		handler := m.Handler(testHandler)
		handler.ServeHTTP(rw, req)

		// Assert
		body := rw.Body.String()

		if body != "malicious access not allowed (jb)" {
			t.Fatal("Response SHOULD NOT BE 'malicious access not allowed (jb)' but found: ", rw.Body.String())
		}
	}
}

func TestJailBotsMiddlewareHandler(t *testing.T) {
	config.TestsConfigureAndInitialize()

	// Act
	m := NewJailBotsMiddleware()

	// Assert
	if m.Handler == nil {
		t.Error("JailBotsMiddleware.Handler is nil")
	}
}

func TestJailBotsMiddlewareIsJailable(t *testing.T) {
	config.TestsConfigureAndInitialize()

	data := []struct {
		uri      string
		jailable bool
	}{
		{uri: "/", jailable: false},
		{uri: "/robots.txt", jailable: false},
		{uri: "/sitemap.xml", jailable: false},
		{uri: "/favicon.ico", jailable: false},
		{uri: "/.env", jailable: true},
		{uri: "/.well-known/ALFA_DATA", jailable: true},
		{uri: "/.well-known/alfacgiapi", jailable: true},
		{uri: "/.well-known/cgialfa", jailable: true},
		{uri: "/api/search?folderIds=0", jailable: true},
		{uri: "/aws/credentials", jailable: true},
		{uri: "/backup", jailable: true},
		{uri: "/backup/license.txt", jailable: true},
		{uri: "/bc", jailable: true},
		{uri: "/bk", jailable: true},
		{uri: "/blog/license.txt", jailable: true},
		{uri: "/bin/", jailable: true},
		{uri: "/cgialfa", jailable: true},
		{uri: "/content/sitetree", jailable: true},
		{uri: "/config.json", jailable: true},
		{uri: "/cgi-bin", jailable: true},
		{uri: "/credentials", jailable: true},
		{uri: "/db", jailable: true},
		{uri: "/db/license.txt", jailable: true},
		{uri: "/wp-config.php", jailable: true},
		{uri: "/wp-content/plugins", jailable: true},
		{uri: "/wp-content/themes", jailable: true},
		{uri: "/wp-includes", jailable: true},
		{uri: "/wp-includes/css", jailable: true},
		{uri: "/wp-includes/js", jailable: true},
		{uri: "/wp-includes/images", jailable: true},
		{uri: "/wp-includes/javascript", jailable: true},
	}

	// Act
	m := jailBotsMiddleware{}

	// Assert
	for i := 0; i < len(data); i++ {
		if m.isJailable(data[i].uri) != data[i].jailable {
			t.Fatal("JailBotsMiddleware.isJailable(", data[i].uri, ") must be ", data[i].jailable)
		}
	}
}
