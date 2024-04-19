package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"project/internal/testutils"
	"project/pkg/userstore"
	"strings"

	"github.com/gouniverse/auth"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type authenticationController struct {
}

func NewAuthenticationController() *authenticationController {
	return &authenticationController{}
}

func (c *authenticationController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	once := strings.TrimSpace(utils.Req(r, "once", ""))

	if once == "" {
		return helpers.ToFlashError(w, r, "System Error. Once is required field", links.NewWebsiteLinks().Home(), 5)
	}

	response, err := c.callAuthKnight(once)
	if err != nil {
		return helpers.ToFlashError(w, r, "System Error. No response from authentication provider", links.NewWebsiteLinks().Home(), 5)
	}

	status := lo.ValueOr(response, "status", "")
	message := lo.ValueOr(response, "message", "")
	data := lo.ValueOr(response, "data", "")

	if status == "" {
		return helpers.ToFlashError(w, r, "System Error. No status found", links.NewWebsiteLinks().Home(), 5)
	}

	if message == "" {
		return helpers.ToFlashError(w, r, "System Error. No message found", links.NewWebsiteLinks().Home(), 5)
	}

	if data == "" {
		return helpers.ToFlashError(w, r, "System Error. No data found", links.NewWebsiteLinks().Home(), 5)
	}

	if status != "success" {
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > Response Status: ", message.(string))
		return helpers.ToFlashError(w, r, "System Error. Invalid authentication response status", links.NewWebsiteLinks().Home(), 5)
	}

	mapData := data.(map[string]any)

	email := lo.ValueOr(mapData, "email", "")

	if email == "" {
		return helpers.ToFlashError(w, r, "System Error. No email", links.NewWebsiteLinks().Home(), 5)
	}

	user, errUser := findOrCreateUser(email.(string))

	if errUser != nil {
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > User Create Error: ", errUser.Error())
		return helpers.ToFlashError(w, r, "Error finding user", links.NewWebsiteLinks().Home(), 5)

	}
	
	if user == nil {
		return helpers.ToFlashError(w, r, "User account not found", links.NewWebsiteLinks().Home(), 5)
	}

	if !user.IsActive() {
		return helpers.ToFlashError(w, r, "User account not active. Please contact support", links.NewWebsiteLinks().Home(), 5)
	}

	sessionKey := utils.StrRandomFromGamma(64, "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz")
	errSession := config.SessionStore.Set(sessionKey, user.ID(), 2*60*60, sessionstore.SessionOptions{
		UserID:    user.ID(),
		UserAgent: r.UserAgent(),
		IPAddress: utils.IP(r),
	})

	if errSession != nil {
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > Session Store Error: ", errSession.Error())
		return helpers.ToFlashError(w, r, "Error creating session", links.NewWebsiteLinks().Home(), 5)
	}

	auth.AuthCookieSet(w, r, sessionKey)

	redirectUrl := c.calculateRedirectURL(*user)
	
	return helpers.ToFlashSuccess(w, r, "Login was successful", redirectUrl, 5)
}

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
func (c *authenticationController) calculateRedirectURL(user userstore.User) string {
	// 1. By default all users redirect to home
	redirectUrl := links.NewUserLinks().Home()

	// 2. If user is manager or admin, redirect to admin panel
	if user.IsManager() || user.IsAdministrator() || user.IsSuperuser() {
		redirectUrl = links.NewAdminLinks().Home()
	}
	
	return redirectUrl
}

func findOrCreateUser(email string) (*userstore.User, error) {
	existingUser, errUser := config.UserStore.UserFindByEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	if existingUser != nil {
		return existingUser, nil
	}

	newUser := userstore.NewUser().
		SetEmail(email).
		SetStatus(userstore.USER_STATUS_ACTIVE)

	errCreate := config.UserStore.UserCreate(newUser)

	if errCreate != nil {
		return nil, errCreate
	}

	return newUser, nil
}
