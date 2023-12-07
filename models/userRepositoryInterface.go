package models

type UserQueryOptions struct {
	Email       string
	ID          string
	IDIn        []string
	Status      string
	StatusIn    []string
	Offset      int
	Limit       int
	SortOrder   string
	OrderBy     string
	CountOnly   bool
	WithDeleted bool
}

type userRepositoryInterface interface {
	UserCreate(user *User) error
	UserFindByEmail(email string) (*User, error)
	UserFindByID(userID string) (*User, error)
	UserList(options UserQueryOptions) ([]User, error)
	UserSoftDelete(user *User) error
	UserSoftDeleteByID(userID string) error
	UserUpdate(user *User) error
}
