package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/internal/testutils"
	"testing"

	"github.com/gouniverse/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_NoSessionKey(t *testing.T) {
	testutils.Setup()

	// Create a request without a session cookie
	req := httptest.NewRequest("GET", "/", nil)
	responseRecorder := httptest.NewRecorder()

	// Create the middleware handler
	handler := authHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(config.AuthenticatedUserContextKey{}) != nil {
			t.Fatal("User should not be set in context")
		}
		if r.Context().Value(config.AuthenticatedSessionContextKey{}) != nil {
			t.Fatal("Session should not be set in context")
		}
		responseRecorder.WriteHeader(http.StatusOK)
	}))

	// Execute the handler
	handler.ServeHTTP(responseRecorder, req)

	// Assert that the next handler was called
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestAuthHandler_SessionNotFoundError(t *testing.T) {
	testutils.Setup()

	// Create a request with a session cookie
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: "some_session_key"})

	// Create a response recorder
	responseRecorder := httptest.NewRecorder()

	// Create the middleware handler
	handler := authHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(config.AuthenticatedUserContextKey{}) != nil {
			t.Fatal("User should not be set in context")
		}
		if r.Context().Value(config.AuthenticatedSessionContextKey{}) != nil {
			t.Fatal("Session should not be set in context")
		}
		responseRecorder.WriteHeader(http.StatusOK)
	}))

	// Execute the handler
	handler.ServeHTTP(responseRecorder, req)

	// Assert that the next handler was called
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestAuthHandler_SessionExpired(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatal("user should not be nil")
	}

	session, err := testutils.SeedSession(httptest.NewRequest("GET", "/", nil), user, -100)

	if err != nil {
		t.Fatal(err)
	}

	if session == nil {
		t.Fatal("session should not be nil")
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: session.GetKey()})

	responseRecorder := httptest.NewRecorder()

	handler := authHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(config.AuthenticatedUserContextKey{}) != nil {
			t.Fatal("User should not be set in context")
		}
		if r.Context().Value(config.AuthenticatedSessionContextKey{}) != nil {
			t.Fatal("Session should not be set in context")
		}
		responseRecorder.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestAuthHandler_UserNotFound(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatal("user should not be nil")
	}

	session, err := testutils.SeedSession(httptest.NewRequest("GET", "/", nil), user, 1)

	if err != nil {
		t.Fatal(err)
	}

	if session == nil {
		t.Fatal("session should not be nil")
	}

	err = config.UserStore.UserDelete(context.Background(), user)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: session.GetKey()})

	responseRecorder := httptest.NewRecorder()

	handler := authHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(config.AuthenticatedUserContextKey{}) != nil {
			t.Fatal("User should not be set in context")
		}
		if r.Context().Value(config.AuthenticatedSessionContextKey{}) != nil {
			t.Fatal("Session should not be set in context")
		}

		responseRecorder.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestAuthHandler_SessionSuccess(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatal("user should not be nil")
	}

	session, err := testutils.SeedSession(httptest.NewRequest("GET", "/", nil), user, 1)

	if err != nil {
		t.Fatal(err)
	}

	if session == nil {
		t.Fatal("session should not be nil")
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: session.GetKey()})

	responseRecorder := httptest.NewRecorder()

	handler := authHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(config.AuthenticatedUserContextKey{}) == nil {
			t.Fatal("User should be set in context")
		}
		if r.Context().Value(config.AuthenticatedSessionContextKey{}) == nil {
			t.Fatal("Session should be set in context")
		}
		responseRecorder.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}
