package userstore

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
)

// == CLASS ===================================================================

type User struct {
	dataobject.DataObject
}

// == CONSTRUCTORS ============================================================

func NewUser() *User {
	o := &User{}

	o.SetID(uid.HumanUid()).
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("").
		SetMiddleNames("").
		SetLastName("").
		SetEmail("").
		SetProfileImageUrl("").
		SetRole(USER_ROLE_USER).
		SetBusinessName("").
		SetPhone("").
		SetPassword("").
		SetTimezone("").
		SetCountry("").
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetDeletedAt(sb.NULL_DATETIME).
		SetMetas(map[string]string{})

	return o
}

func NewUserFromExistingData(data map[string]string) *User {
	o := &User{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func UserNoImageUrl() string {
	return "/user/default.png"
	//return config.MediaUrl + "/user/default.png"
}

func (o *User) IsActive() bool {
	return o.Status() == USER_STATUS_ACTIVE
}

func (o *User) IsDeleted() bool {
	return o.Status() == USER_STATUS_DELETED
}

func (o *User) IsInactive() bool {
	return o.Status() == USER_STATUS_INACTIVE
}

func (o *User) IsUnverified() bool {
	return o.Status() == USER_STATUS_UNVERIFIED
}

func (o *User) IsAdministrator() bool {
	return o.Role() == USER_ROLE_ADMINISTRATOR
}

func (o *User) IsManager() bool {
	return o.Role() == USER_ROLE_MANAGER
}

func (o *User) IsSuperuser() bool {
	return o.Role() == USER_ROLE_SUPERUSER
}

// IsRegistrationCompleted checks if the user registration is incomplete.
//
// Registration is considered incomplete if the user's first name
// or last name is empty.
//
// Parameters:
// - authUser: a pointer to a userstore.User object representing the authenticated user.
//
// Returns:
// - bool: true if the user registration is incomplete, false otherwise.
func (o *User) IsRegistrationCompleted() bool {
	return o.FirstName() != "" && o.LastName() != ""
}

// == SETTERS AND GETTERS =====================================================

func (o *User) BusinessName() string {
	return o.Get("business_name")
}

func (o *User) SetBusinessName(lastName string) *User {
	o.Set("business_name", lastName)
	return o
}

func (o *User) Country() string {
	return o.Get("country")
}

func (o *User) SetCountry(country string) *User {
	o.Set("country", country)
	return o
}

func (o *User) CreatedAt() string {
	return o.Get("created_at")
}

func (o *User) CreatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *User) SetCreatedAt(createdAt string) *User {
	o.Set("created_at", createdAt)
	return o
}

func (o *User) DeletedAt() string {
	return o.Get("deleted_at")
}

func (o *User) SetDeletedAt(deletedAt string) *User {
	o.Set("deleted_at", deletedAt)
	return o
}

func (o *User) Email() string {
	return o.Get("email")
}

func (o *User) SetEmail(email string) *User {
	o.Set("email", email)
	return o
}

func (o *User) FirstName() string {
	return o.Get("first_name")
}

func (o *User) SetFirstName(firstName string) *User {
	o.Set("first_name", firstName)
	return o
}

func (o *User) ID() string {
	return o.Get("id")
}

func (o *User) SetID(id string) *User {
	o.Set("id", id)
	return o
}

func (o *User) LastName() string {
	return o.Get("last_name")
}

func (o *User) SetLastName(lastName string) *User {
	o.Set("last_name", lastName)
	return o
}

func (o *User) Memo() string {
	return o.Get("memo")
}

func (o *User) SetMemo(memo string) *User {
	o.Set("memo", memo)
	return o
}

func (o *User) MiddleNames() string {
	return o.Get("middle_names")
}

func (o *User) SetMiddleNames(middleNames string) *User {
	o.Set("middle_names", middleNames)
	return o
}

func (o *User) Metas() (map[string]string, error) {
	metasStr := o.Get("metas")

	if metasStr == "" {
		metasStr = "{}"
	}

	metasJson, errJson := utils.FromJSON(metasStr, map[string]string{})
	if errJson != nil {
		return map[string]string{}, errJson
	}

	return maputils.MapStringAnyToMapStringString(metasJson.(map[string]any)), nil
}

func (o *User) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *User) SetMeta(name string, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
	// return config.MetaStore.Set("user", o.ID(), name, value)
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *User) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}
	o.Set("metas", mapString)
	return nil
}

func (o *User) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *User) Password() string {
	return o.Get("password")
}

func (o *User) PasswordCompare(password string) bool {
	hash := o.Get("password")
	return utils.StrToBcryptHashCompare(password, hash)
}

// SetPasswordAndHash hashes the password before saving
func (o *User) SetPasswordAndHash(password string) error {
	hash, err := utils.StrToBcryptHash(password)

	if err != nil {
		return err
	}

	o.SetPassword(hash)

	return nil
}

// SetPassword sets the password as provided, if you want it hashed use SetPasswordAndHash() method
func (o *User) SetPassword(password string) *User {
	o.Set("password", password)
	return o
}

func (o *User) Phone() string {
	return o.Get("phone")
}

func (o *User) SetPhone(phone string) *User {
	o.Set("phone", phone)
	return o
}

func (o *User) ProfileImageUrl() string {
	return o.Get("profile_image_url")
}

func (o *User) ProfileImageOrDefaultUrl() string {
	defaultURL := UserNoImageUrl()

	if o.ProfileImageUrl() != "" {
		return o.ProfileImageUrl()
	}

	return defaultURL
}

func (o *User) SetProfileImageUrl(imageUrl string) *User {
	o.Set("profile_image_url", imageUrl)
	return o
}

func (o *User) Role() string {
	return o.Get("role")
}

func (o *User) SetRole(role string) *User {
	o.Set("role", role)
	return o
}

func (o *User) Status() string {
	return o.Get("status")
}

func (o *User) SetStatus(status string) *User {
	o.Set("status", status)
	return o
}

func (o *User) Timezone() string {
	return o.Get("timezone")
}

func (o *User) SetTimezone(timezone string) *User {
	o.Set("timezone", timezone)
	return o
}

func (o *User) UpdatedAt() string {
	return o.Get("updated_at")
}

func (o *User) UpdatedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.Get("updated_at"), carbon.UTC)
}

func (o *User) SetUpdatedAt(updatedAt string) *User {
	o.Set("updated_at", updatedAt)
	return o
}
