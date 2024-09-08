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
	"github.com/gouniverse/blindindexstore"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/strutils"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

type authenticationController struct {
}

// == CONSTRUCTOR =============================================================

func NewAuthenticationController() *authenticationController {
	return &authenticationController{}
}

// == PUBLIC METHODS ==========================================================

func (c *authenticationController) Handler(w http.ResponseWriter, r *http.Request) string {
	email, errorMessage := c.emailFromAuthKnightRequest(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, "Authentication Provider Error. "+errorMessage, links.NewWebsiteLinks().Home(), 5)
	}

	user, errUser := c.userFindByEmailOrCreate(email, userstore.USER_STATUS_ACTIVE)

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

	sessionKey := strutils.RandomFromGamma(64, "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz")
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

// == PRIVATE METHODS =========================================================

func (c *authenticationController) userFindByEmailOrCreate(email string, status string) (*userstore.User, error) {
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
		return nil, nil
	}

	return user, nil
}

func (c *authenticationController) userCreate(email string, status string) (*userstore.User, error) {
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
		config.LogStore.ErrorWithContext("At Auth Controller > AnyIndex > Call Auth Knight Error: ", err.Error())
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

	// 3. If user does not have any names, redirect to profile
	if !user.IsRegistrationCompleted() {
		redirectUrl = links.NewAuthLinks().Register(map[string]string{})
		redirectUrl = helpers.ToFlashInfoURL("Thank you for logging in. Please complete your data to finish your registration", redirectUrl, 5)
	}

	return redirectUrl
}
