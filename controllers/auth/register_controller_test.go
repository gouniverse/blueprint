package auth

import (
	"net/http"
	"net/url"
	"project/config"
	"project/internal/links"
	"project/internal/testutils"
	"strings"
	"testing"

	"github.com/gouniverse/auth"
)

func TestRegisterController_RequiresAuthenticatedUser(t *testing.T) {
	testutils.Setup()

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodGet, NewRegisterController().Handler, testutils.NewRequestOptions{
		GetValues: url.Values{},
		Context:   map[any]any{},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusSeeOther {
		t.Fatal(`Response MUST be `, http.StatusSeeOther, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<a href="/flash?message_id=`,
		`">See Other</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}

	flashMessage, err := testutils.FlashMessageFindFromResponse(response)

	if err != nil {
		t.Fatal(err)
	}

	if flashMessage == nil {
		t.Fatal(`Response MUST contain 'flash message'`)
	}

	if flashMessage.Type != "error" {
		t.Fatal(`Response be of type 'success', but got: `, flashMessage.Type, flashMessage.Message)
	}

	expected := "You must be logged in to access this page"
	if flashMessage.Message != expected {
		t.Fatal(`Response MUST contain '`, expected, `', but got: `, flashMessage.Message)
	}
}

func TestRegisterController_ShowsRegisterForm(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	if err != nil {
		t.Fatal(err)
	}

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodGet, NewRegisterController().Handler, testutils.NewRequestOptions{
		GetValues: url.Values{},
		Context: map[any]any{
			auth.AuthenticatedUserID{}:    user.ID(),
			config.AuthenticatedUserKey{}: user,
		},
	})

	if err != nil {
		t.Fatal("Response MUST NOT trigger error, but was:", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`id="FormRegister"`,
		`name="email"`,
		`name="first_name"`,
		`name="last_name"`,
		`name="country"`,
		`name="timezone"`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}
}

func TestRegisterController_RequiresFirstName(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewRegisterController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"email": {user.Email()},
		},
		Context: map[any]any{
			auth.AuthenticatedUserID{}:    user.ID(),
			config.AuthenticatedUserKey{}: user,
		},
	})

	if err != nil {
		t.Fatal("Response MUST NOT trigger error, but was:", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`id="FormRegister"`,
		`name="email"`,
		`name="first_name"`,
		`name="last_name"`,
		`name="country"`,
		`name="timezone"`,
		`First name is required field`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}
}

func TestRegisterController_RequiresLastName(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewRegisterController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"email":      {user.Email()},
			"first_name": {"FirstName"},
		},
		Context: map[any]any{
			auth.AuthenticatedUserID{}:    user.ID(),
			config.AuthenticatedUserKey{}: user,
		},
	})

	if err != nil {
		t.Fatal("Response MUST NOT trigger error, but was:", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`id="FormRegister"`,
		`name="email"`,
		`name="first_name"`,
		`name="last_name"`,
		`name="country"`,
		`name="timezone"`,
		`Last name is required field`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}
}

func TestRegisterController_RequiresCountry(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewRegisterController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"email":      {user.Email()},
			"first_name": {"FirstName"},
			"last_name":  {"LastName"},
		},
		Context: map[any]any{
			auth.AuthenticatedUserID{}:    user.ID(),
			config.AuthenticatedUserKey{}: user,
		},
	})

	if err != nil {
		t.Fatal("Response MUST NOT trigger error, but was:", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`id="FormRegister"`,
		`name="email"`,
		`name="first_name"`,
		`name="last_name"`,
		`name="country"`,
		`name="timezone"`,
		`Country is required field`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}
}

func TestRegisterController_RequiresTimezone(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewRegisterController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"email":      {user.Email()},
			"first_name": {"FirstName"},
			"last_name":  {"LastName"},
			"country":    {"Country"},
		},
		Context: map[any]any{
			auth.AuthenticatedUserID{}:    user.ID(),
			config.AuthenticatedUserKey{}: user,
		},
	})

	if err != nil {
		t.Fatal("Response MUST NOT trigger error, but was:", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`id="FormRegister"`,
		`name="email"`,
		`name="first_name"`,
		`name="last_name"`,
		`name="country"`,
		`name="timezone"`,
		`Timezone is required field`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}
}

func TestRegisterController_Success(t *testing.T) {
	testutils.Setup()

	user, err := testutils.SeedUser(testutils.USER_01)

	responseHTML, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewRegisterController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"email":      {user.Email()},
			"first_name": {"FirstName"},
			"last_name":  {"LastName"},
			"country":    {"Country"},
			"timezone":   {"Timezone"},
		},
		Context: map[any]any{
			auth.AuthenticatedUserID{}:    user.ID(),
			config.AuthenticatedUserKey{}: user,
		},
	})

	if err != nil {
		t.Fatal("Response MUST NOT trigger error, but was:", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`id="FormRegister"`,
		`name="email"`,
		`name="first_name"`,
		`name="last_name"`,
		`name="country"`,
		`name="timezone"`,
		`Your registration completed successfully. You can now continue browsing the website.`,
		`<script>window.location.href = '` + links.NewUserLinks().Home(map[string]string{}) + `'</script>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(responseHTML, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, responseHTML)
		}
	}
}
