package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"project/config"
	"project/internal/testutils"
	"testing"

	"github.com/gouniverse/responses"
)

func TestAuthControllerOnceIsRequired(t *testing.T) {
	testutils.Setup()

	req, err := testutils.NewRequest(http.MethodPost, "/", testutils.NewRequestOptions{
		PostValues: url.Values{},
	})

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	(http.Handler(responses.JSONHandler(NewAuthenticationController().Handler))).ServeHTTP(recorder, req)

	code := recorder.Code

	if code != http.StatusSeeOther {
		t.Fatal(`Response MUST be 303`, code)
	}

	flashMessage, err := testutils.FlashMessageFindFromResponse(recorder.Result())

	if err != nil {
		t.Fatal(err)
	}

	if flashMessage == nil {
		t.Fatal(`Response MUST contain 'flash message'`)
	}

	if flashMessage.Type != "error" {
		t.Fatal(`Response be of type 'success', but got: `, flashMessage.Type, flashMessage.Message)
	}

	if flashMessage.Message != "Authentication Provider Error. Once is required field" {
		t.Fatal(`Response MUST contain 'Authentication Provider Error. Once is required field', but got: `, flashMessage.Message)
	}
}

func TestAuthControllerOnceMustBeValid(t *testing.T) {
	testutils.Setup()

	req, err := testutils.NewRequest(http.MethodPost, "/", testutils.NewRequestOptions{
		PostValues: url.Values{
			"once": {"test"},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	(http.Handler(responses.JSONHandler(NewAuthenticationController().Handler))).ServeHTTP(recorder, req)
	// response := recorder.Body.String()

	code := recorder.Code

	if code != http.StatusSeeOther {
		t.Fatal(`Response MUST be 303`, code)
	}

	flashMessage, err := testutils.FlashMessageFindFromResponse(recorder.Result())

	if err != nil {
		t.Fatal(err)
	}

	if flashMessage == nil {
		t.Fatal(`Response MUST contain 'flash message'`)
	}

	if flashMessage.Type != "error" {
		t.Fatal(`Response be of type 'success', but got: `, flashMessage.Type, flashMessage.Message)
	}

	if flashMessage.Message != "Authentication Provider Error. Invalid authentication response status" {
		t.Fatal(`Response MUST contain 'Authentication Provider Error. Invalid authentication response status', but got: `, flashMessage.Message, flashMessage.Message)
	}
}

func TestAuthControllerOnceSuccessWithNewUser(t *testing.T) {
	testutils.Setup()

	req, err := testutils.NewRequest(http.MethodPost, "/", testutils.NewRequestOptions{
		PostValues: url.Values{
			"once": {testutils.TestKey()},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	(http.Handler(responses.JSONHandler(NewAuthenticationController().Handler))).ServeHTTP(recorder, req)
	// response := recorder.Body.String()
	code := recorder.Code

	if code != http.StatusSeeOther {
		t.Fatal(`Response MUST be 303`, code)
	}

	flashMessage, err := testutils.FlashMessageFindFromResponse(recorder.Result())

	if err != nil {
		t.Fatal(err)
	}

	if flashMessage == nil {
		t.Fatal(`Response MUST contain 'flash message'`)
	}

	if flashMessage.Type != "success" {
		t.Fatal(`Response be of type 'success', but got: `, flashMessage.Type, flashMessage.Message)
	}

	expected := "Login was successful"
	if flashMessage.Message != "Login was successful" {
		t.Fatal(`Response MUST contain '`+expected+`', but got: `, flashMessage.Message)
	}
}

func TestAuthControllerOnceSuccessWithExistingUser(t *testing.T) {
	testutils.Setup()
	user, err := testutils.SeedUser(testutils.USER_01)

	if err != nil {
		t.Fatal(err)
	}

	user.SetEmail("test@test.com")

	err = config.UserStore.UserUpdate(context.Background(), user)

	if err != nil {
		t.Fatal(err)
	}

	req, err := testutils.NewRequest(http.MethodPost, "/", testutils.NewRequestOptions{
		PostValues: url.Values{
			"once": {testutils.TestKey()},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	(http.Handler(responses.JSONHandler(NewAuthenticationController().Handler))).ServeHTTP(recorder, req)
	// response := recorder.Body.String()
	code := recorder.Code

	if code != http.StatusSeeOther {
		t.Fatal(`Response MUST be 303`, code)
	}

	flashMessage, err := testutils.FlashMessageFindFromResponse(recorder.Result())

	if err != nil {
		t.Fatal(err)
	}

	if flashMessage == nil {
		t.Fatal(`Response MUST contain 'flash message'`)
	}

	if flashMessage.Type != "success" {
		t.Fatal(`Response be of type 'success', but got: `, flashMessage.Type, flashMessage.Message)
	}

	expected := "Login was successful"
	if flashMessage.Message != "Login was successful" {
		t.Fatal(`Response MUST contain '`+expected+`', but got: `, flashMessage.Message)
	}
}
