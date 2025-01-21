package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"project/internal/testutils"
	"strings"
	"testing"

	"github.com/gouniverse/userstore"
)

func TestAdminMiddleware_NoUserRedirectsToLogin(t *testing.T) {
	// Arrange
	testutils.Setup()

	// Act
	body, response, err := testutils.CallMiddleware("GET", NewAdminMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Should not be called")
		w.WriteHeader(http.StatusOK)
	}, testutils.NewRequestOptions{})

	// Assert

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

	if msg.Message != "You must be logged in to access this page" {
		t.Fatalf("Expected message %s, got %s", "You must be logged in to access this page", msg.Message)
	}

	if !strings.Contains(msg.Url, links.AUTH_LOGIN) {
		t.Fatalf("Expected url %s, got %s", links.AUTH_LOGIN, msg.Url)
	}
}

func TestAdminMiddleware_RequiresRegisteredUser(t *testing.T) {
	// Arrange

	testutils.Setup()
	user, session, err := testutils.SeedUserAndSession(testutils.USER_01, httptest.NewRequest("GET", "/", nil), 1)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	body, response, err := testutils.CallMiddleware("GET", NewAdminMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
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

func TestAdminMiddleware_RequiresActiveUser(t *testing.T) {
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

	body, response, err := testutils.CallMiddleware("GET", NewAdminMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
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

	if msg.Type != helpers.FLASH_ERROR {
		t.Fatalf("Expected message type %s, got %s", helpers.FLASH_ERROR, msg.Type)
	}

	if msg.Message != "Your account is not active" {
		t.Fatalf("Expected message %s, got %s", "Your account is not active", msg.Message)
	}

	if msg.Url != links.HOME {
		t.Fatalf("Expected url %s, got %s", links.HOME, msg.Url)
	}
}

func TestAdminMiddleware_RequiresAdminUser(t *testing.T) {
	// Arrange

	testutils.Setup()
	user, session, err := testutils.SeedUserAndSession(testutils.USER_01, httptest.NewRequest("GET", "/", nil), 1)

	if err != nil {
		t.Fatal(err)
	}

	user.SetStatus(userstore.USER_STATUS_ACTIVE)
	user.SetFirstName("First Name")
	user.SetLastName("Last Name")
	user.SetCountry("US")
	user.SetTimezone("America/New_York")
	err = config.UserStore.UserUpdate(context.Background(), user)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	body, response, err := testutils.CallMiddleware("GET", NewAdminMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
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

	if msg.Type != helpers.FLASH_ERROR {
		t.Fatalf("Expected message type %s, got %s", helpers.FLASH_ERROR, msg.Type)
	}

	if msg.Message != "You must be an administrator to access this page" {
		t.Fatalf("Expected message %s, got %s", "You must be an administrator to access this page", msg.Message)
	}

	if msg.Url != links.HOME {
		t.Fatalf("Expected url %s, got %s", links.HOME, msg.Url)
	}
}

func TestAdminMiddleware_Success(t *testing.T) {
	// Arrange
	testutils.Setup()

	user, session, err := testutils.SeedUserAndSession(testutils.ADMIN_01, httptest.NewRequest("GET", "/", nil), 1)

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

	body, response, err := testutils.CallMiddleware("GET", NewAdminMiddleware().Handler, func(w http.ResponseWriter, r *http.Request) {
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

	// Assert

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	if body != "Success" {
		t.Fatalf("Expected body: %s, got: %s", "Success", body)
	}
}
