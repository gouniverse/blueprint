package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"project/app/links"
	"project/config"
	"project/internal/helpers"
	"project/internal/testutils"
	"strings"
	"testing"

	"github.com/gouniverse/userstore"
)

func TestUserMiddleware_NoUserRedirectsToLogin(t *testing.T) {
	// Arrange
	testutils.Setup()

	// Act
	body, response, err := testutils.CallMiddleware("GET", NewUserMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Should not be called")
		w.WriteHeader(http.StatusOK)
	}, testutils.NewRequestOptions{})

	// Assert

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	if response.StatusCode != http.StatusSeeOther {
		t.Fatalf("Expected status code %d, got %d", http.StatusSeeOther, response.StatusCode)
	}

	if strings.Contains(body, "/flash?message_id=") == false {
		t.Fatalf("Expected response to contain '/flash?message_id=', got %s", body)
	}

	msg, err := testutils.FlashMessageFindFromBody(body)

	if err != nil {
		t.Fatal(err)
	}

	if msg == nil {
		t.Fatal("Flash message should not be nil")
	}

	if msg.Type != helpers.FLASH_ERROR {
		t.Fatalf("Expected message type %s, got %s", helpers.FLASH_ERROR, msg.Type)
	}

	if msg.Message != "Only authenticated users can access this page" {
		t.Fatalf("Expected message %s, got %s", "You are not logged in", msg.Message)
	}

	if !strings.Contains(msg.Url, links.AUTH_LOGIN) {
		t.Fatalf("Expected url %s, got %s", links.AUTH_LOGIN, msg.Url)
	}
}

func TestUserMiddleware_RequiresRegisteredUser(t *testing.T) {
	if config.UserStore == nil {
		t.Fatal("UserStore should not be nil")
	}

	if config.SessionStore == nil {
		t.Fatal("SessionStore should not be nil")
	}

	// Arrange
	testutils.Setup()
	user, session, err := testutils.SeedUserAndSession(testutils.USER_01, httptest.NewRequest("GET", "/", nil), 1)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	body, response, err := testutils.CallMiddleware("GET", NewUserMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Should not be called")
		w.WriteHeader(http.StatusOK)
	}, testutils.NewRequestOptions{
		Context: map[any]any{
			config.AuthenticatedUserContextKey{}:    user,
			config.AuthenticatedSessionContextKey{}: session,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Assert

	// Assert that its a redirect
	if response.StatusCode != http.StatusSeeOther {
		t.Fatalf("Expected status code %d, got %d", http.StatusSeeOther, response.StatusCode)
	}

	if !strings.Contains(body, "/flash?message_id=") {
		t.Fatalf("Expected response to contain '/flash?message_id=', got %s", body)
	}

	msg, err := testutils.FlashMessageFindFromBody(body)

	if err != nil {
		t.Fatal(err)
	}

	if msg == nil {
		t.Fatal("Flash message should not be nil")
	}

	if msg.Type != helpers.FLASH_INFO {
		t.Fatalf("Expected message type %s, got %s", helpers.FLASH_ERROR, msg.Type)
	}

	if msg.Url != links.AUTH_REGISTER {
		t.Fatalf("Expected url %s, got %s", links.AUTH_REGISTER, msg.Url)
	}

	if msg.Message != "Please complete your registration to continue" {
		t.Fatalf("Expected message %s, got %s", "Please complete your registration to continue", msg.Message)
	}
}

func TestUserMiddleware_RequiresActiveUser(t *testing.T) {
	if config.UserStore == nil {
		t.Fatal("UserStore should not be nil")
	}

	if config.SessionStore == nil {
		t.Fatal("SessionStore should not be nil")
	}

	// Arrange
	testutils.Setup()
	user, session, err := testutils.SeedUserAndSession(testutils.USER_01, httptest.NewRequest("GET", "/", nil), 1)

	if err != nil {
		t.Fatal(err)
	}

	user.SetStatus(userstore.USER_STATUS_INACTIVE)
	user.SetFirstName("First Name")
	user.SetLastName("Last Name")
	user.SetCountry("US")
	user.SetTimezone("America/New_York")
	err = config.UserStore.UserUpdate(context.Background(), user)

	if err != nil {
		t.Fatal(err)
	}

	// Act
	body, response, err := testutils.CallMiddleware("GET", NewUserMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Should not be called")
		w.WriteHeader(http.StatusOK)
	}, testutils.NewRequestOptions{
		Context: map[any]any{
			config.AuthenticatedUserContextKey{}:    user,
			config.AuthenticatedSessionContextKey{}: session,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	// Assert

	// Assert that its a redirect
	if response.StatusCode != http.StatusSeeOther {
		t.Fatalf("Expected status code %d, got %d", http.StatusSeeOther, response.StatusCode)
	}

	if !strings.Contains(body, "/flash?message_id=") {
		t.Fatalf("Expected response to contain '/flash?message_id=', got %s", body)
	}

	msg, err := testutils.FlashMessageFindFromBody(body)

	if err != nil {
		t.Fatal(err)
	}

	if msg == nil {
		t.Fatal("Flash message should not be nil")
	}

	if msg.Type != helpers.FLASH_ERROR {
		t.Fatalf("Expected message type %s, got %s", helpers.FLASH_ERROR, msg.Type)
	}

	if msg.Message != "User account not active" {
		t.Fatalf("Expected message %s, got %s", "User account not active", msg.Message)
	}

	if msg.Url != links.HOME {
		t.Fatalf("Expected url %s, got %s", links.HOME, msg.Url)
	}
}

func TestUserMiddleware_Success(t *testing.T) {
	if config.UserStore == nil {
		t.Fatal("UserStore should not be nil")
	}

	if config.SessionStore == nil {
		t.Fatal("SessionStore should not be nil")
	}

	// Arrange
	testutils.Setup()

	user, session, err := testutils.SeedUserAndSession(testutils.USER_01, httptest.NewRequest("GET", "/", nil), 1)

	if err != nil {
		t.Fatal(err)
	}

	// Activate the user
	user.SetStatus(userstore.USER_STATUS_ACTIVE)

	// Register the user
	user.SetFirstName("First Name")
	user.SetLastName("Last Name")
	user.SetCountry("US")
	user.SetTimezone("America/New_York")

	err = config.UserStore.UserUpdate(context.Background(), user)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	body, response, err := testutils.CallMiddleware("GET", NewUserMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success"))
		w.WriteHeader(http.StatusOK)
	}, testutils.NewRequestOptions{
		Context: map[any]any{
			config.AuthenticatedUserContextKey{}:    user,
			config.AuthenticatedSessionContextKey{}: session,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	// Assert

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	if body != "Success" {
		t.Fatalf("Expected body: %s, got: %s", "Success", body)
	}
}
