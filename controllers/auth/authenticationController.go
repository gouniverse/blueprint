package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"project/internal/testutils"
	"strings"

	"github.com/gouniverse/auth"
	"github.com/gouniverse/blindindexstore"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/strutils"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

const msgAccountNotFound = `Your account may have been deactivated or deleted. Please contact our support team for assistance.`
const msgAccountNotActive = `Your account is not active. Please contact our support team for assistance.`
const msgUserNotFound = `An unexpected error has occurred trying to find your account. The support team has been notified.`

// == CONTROLLER ==============================================================

// authenticationController handles the authentication of the user,
// once the user has logged in successfully via the AuthKnight service.
type authenticationController struct{}

// == CONSTRUCTOR =============================================================

// NewAuthenticationController creates a new instance of the authenticationController struct.
//
// Parameters:
// - none
//
// Returns:
// - *authenticationController: a pointer to the authenticationController struct.
func NewAuthenticationController() *authenticationController {
	return &authenticationController{}
}

// == PUBLIC METHODS ==========================================================

// AnyIndex handles the authentication.
//
// 1. Checks if there is a once parameter in the request from the AuthKnight service.
// 2. Calls the AuthKnight service with the once parameter.
// 3. Verifies the response from the AuthKnight service.
// 4. Based on the email, it will find or create a user in the database.
// 5. Creates a new session for the user.
// 6. Checks if theuser has completed their profile.
// 7. If not, it will redirect the user to the profile page.
// 8. If yes, it will redirect the user to the home page, or the admin panel.
//
// Parameters:
// - w: http.ResponseWriter: the response writer.
// - r: *http.Request: the incoming request.
//
// Return:
// - string: the result of the authentication request.
func (c *authenticationController) Handler(w http.ResponseWriter, r *http.Request) string {
	homeURL := links.NewWebsiteLinks().Home()
	email, errorMessage := c.emailFromAuthKnightRequest(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, "Authentication Provider Error. "+errorMessage, homeURL, 5)
	}

	user, errUser := c.userFindByEmailOrCreate(email, userstore.USER_STATUS_ACTIVE)

	if errUser != nil {
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > User Create Error: ", errUser.Error())
		return helpers.ToFlashError(w, r, msgUserNotFound, homeURL, 5)
	}

	if user == nil {
		return helpers.ToFlashError(w, r, msgAccountNotFound, homeURL, 5)
	}

	if !user.IsActive() {
		return helpers.ToFlashError(w, r, msgAccountNotActive, homeURL, 5)
	}

	sessionKey := strutils.RandomFromGamma(64, "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz")
	errSession := config.SessionStore.Set(sessionKey, user.ID(), 2*60*60, sessionstore.SessionOptions{
		UserID:    user.ID(),
		UserAgent: r.UserAgent(),
		IPAddress: utils.IP(r),
	})

	if errSession != nil {
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > Session Store Error: ", errSession.Error())
		return helpers.ToFlashError(w, r, "Error creating session", homeURL, 5)
	}

	auth.AuthCookieSet(w, r, sessionKey)

	redirectUrl := c.calculateRedirectURL(user)

	return helpers.ToFlashSuccess(w, r, "Login was successful", redirectUrl, 5)
}

// == PRIVATE METHODS =========================================================

func (c *authenticationController) userFindByEmailOrCreate(email string, status string) (userstore.UserInterface, error) {
	userID, err := c.findEmailInBlindIndex(email)

	if err != nil {
		return nil, err
	}

	if userID == "" {
		return c.userCreate(email, status)
	}

	user, err := config.UserStore.UserFindByID(userID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		config.Logger.Error("At Auth Controller > userFindByEmailOrCreate",
			slog.String("error", "User not found, even though email was found in the blind index, and user ID returned successfully"),
			"user", userID)
		return nil, nil
	}

	return user, nil
}

func (c *authenticationController) userCreate(email string, status string) (userstore.UserInterface, error) {
	user := userstore.NewUser().
		SetStatus(status).
		SetEmail(email)

	err := config.UserStore.UserCreate(user)

	if err != nil {
		return nil, err
	}

	emailToken, err := config.VaultStore.TokenCreate(email, config.VaultKey, 20)

	if err != nil {
		return nil, err
	}

	user.SetEmail(emailToken)

	err = config.UserStore.UserUpdate(user)

	if err != nil {
		return nil, err
	}

	searchValue := blindindexstore.NewSearchValue().
		SetSourceReferenceID(user.ID()).
		SetSearchValue(email)

	err = config.BlindIndexStoreEmail.SearchValueCreate(searchValue)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *authenticationController) findEmailInBlindIndex(email string) (userID string, err error) {
	recordsFound, err := config.BlindIndexStoreEmail.SearchValueList(blindindexstore.SearchValueQueryOptions{
		SearchValue: email,
		SearchType:  blindindexstore.SEARCH_TYPE_EQUALS,
	})

	if err != nil {
		return "", err
	}

	if len(recordsFound) < 1 {
		return "", nil
	}

	return recordsFound[0].SourceReferenceID(), nil
}

func (c *authenticationController) emailFromAuthKnightRequest(r *http.Request) (email string, errorMessage string) {
	once := strings.TrimSpace(utils.Req(r, "once", ""))

	if once == "" {
		return "", "Once is required field"
	}

	response, err := c.callAuthKnight(once)
	if err != nil {
		config.LogStore.ErrorWithContext("At Auth Controller > emailFromAuthKnightRequest > Call Auth Knight Error: ", err.Error())
		return "", "No response from authentication provider"
	}

	status := lo.ValueOr(response, "status", "")
	message := lo.ValueOr(response, "message", "")
	data := lo.ValueOr(response, "data", "")

	if status == "" {
		return "", "No status found"
	}

	if message == "" {
		return "", "No message found"
	}

	if data == "" {
		return "", "No data found"
	}

	if status != "success" {
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > Response Status: ", message.(string))
		return "", "Invalid authentication response status"
	}

	mapData := data.(map[string]any)

	email = strings.TrimSpace(lo.ValueOr(mapData, "email", "").(string))

	if email == "" {
		return "", "No email found"
	}

	return email, ""
}

// callAuthKnight makes a request to the authentication server
// to verify the provided "once" token. The "once" token is provided
// by the AuthKnight service.
//
// Note! If the environment is "testing", it will return a predefined response
// which is used only for testing purposes. In the case of a successful response,
// the email is "test@test.com".
//
// Parameters:
//   - once: The once token to be verified.
//
// Returns:
//   - response: A map containing the response data from the authentication server.
//   - error: An error object if an error occurred during the request.
func (*authenticationController) callAuthKnight(once string) (map[string]interface{}, error) {
	var response map[string]interface{}

	if config.IsEnvTesting() {
		var testResponseJSONString = ""
		if once == testutils.TestKey() {
			testResponseJSONString = `{"status":"success","message":"success","data":{"email":"test@test.com"}}`
		} else {
			testResponseJSONString = `{"status":"error","message":"once data is invalid:test","data":{}}`
		}
		json.NewDecoder(bytes.NewReader([]byte(testResponseJSONString))).Decode(&response)
		return response, nil
	}

	req, err := http.PostForm("https://authknight.com/api/who?once="+once, url.Values{
		"once": {once},
	})

	if err != nil {
		return nil, err
	}

	if req == nil {
		return nil, errors.New("no response")
	}

	defer req.Body.Close()

	json.NewDecoder(req.Body).Decode(&response)

	return response, nil
}

// calculateRedirectURL calculates the redirect URL based on the user's role and profile completeness.
//
// 1. By default all users redirect to home
// 2. If user is manager or admin, redirect to admin panel
// 3. If user does not have any names, redirect to profile
//
// Parameters:
// - user (models.User): The user object.
//
// Returns:
// - string: The redirect URL.
func (c *authenticationController) calculateRedirectURL(user userstore.UserInterface) string {
	// 1. By default all users redirect to home
	redirectUrl := links.NewUserLinks().Home(map[string]string{})

	// 2. If user is manager or admin, redirect to admin panel
	if user.IsManager() || user.IsAdministrator() || user.IsSuperuser() {
		redirectUrl = links.NewAdminLinks().Home(map[string]string{})
	}

	// 3. If user does not have any names, redirect to profile
	if !user.IsRegistrationCompleted() {
		redirectUrl = links.NewAuthLinks().Register(map[string]string{})
		redirectUrl = helpers.ToFlashInfoURL("Thank you for logging in. Please complete your data to finish your registration", redirectUrl, 5)
	}

	return redirectUrl
}
