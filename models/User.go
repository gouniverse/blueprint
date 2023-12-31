package models

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

const USER_ROLE_SUPERUSER = "superuser"
const USER_ROLE_ADMINISTRATOR = "administrator"
const USER_ROLE_MANAGER = "manager"
const USER_ROLE_USER = "user"

const USER_STATUS_ACTIVE = "active"
const USER_STATUS_INACTIVE = "inactive"
const USER_STATUS_DELETED = "deleted"
const USER_STATUS_UNVERIFIED = "unverified"

type User struct {
	dataobject.DataObject
}

func NewUser() *User {
	o := &User{}
	o.SetID(uid.HumanUid()).
		SetFirstName("").
		SetMiddleNames("").
		SetLastName("").
		SetEmail("").
		SetProfileImageUrl("").
		SetRole(USER_ROLE_USER).
		SetBusinessName("").
		SetPhone("").
		SetTimezone("").
		SetCountry("").
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetDeletedAt(sb.NULL_DATETIME)

	return o
}

func NewUserFromExistingData(data map[string]string) *User {
	o := &User{}
	o.Hydrate(data)
	return o
}

// func (o *User) Update() error {
// 	return NewUserRepository().UserUpdate(o)
// }

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
	return carbon.NewCarbon().Parse(o.Get("created_at"), carbon.UTC)
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

func (o *User) LastName() string {
	return o.Get("last_name")
}

func (o *User) MiddleNames() string {
	return o.Get("middle_names")
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

func (o *User) SetMiddleNames(middleNames string) *User {
	o.Set("middle_names", middleNames)
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

func UserNoImageUrl() string {
	return "/user/default.png"
}
