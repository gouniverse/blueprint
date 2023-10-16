package models

type UserQueryOptions struct {
	ID        string
	IDIn      string
	Status    string
	StatusIn  []string
	Offset    int
	Limit     int
	SortOrder string
	OrderBy   string
	CountOnly bool
}

type userRepositoryInterface interface {
	UserCreate(user *User) error
	UserDelete(user *User) error
	UserDeleteByID(userID string) error
	UserFindByEmail(email string) (*User, error)
	UserFindByID(userID string) (*User, error)
	UserList(options UserQueryOptions) ([]User, error)
	UserUpdate(user *User) error
}
