package userstore

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
)

// var _ StoreInterface = (*Store)(nil) // verify it extends the interface

type Store struct {
	userTableName      string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debugEnabled       bool
}

// AutoMigrate auto migrate
func (store *Store) AutoMigrate() error {
	sql := store.sqlUserTableCreate()

	if sql == "" {
		return errors.New("user table create sql is empty")
	}

	_, err := store.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

// EnableDebug - enables the debug option
func (st *Store) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

func (store *Store) UserCreate(user *User) error {
	user.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := user.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.userTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	_, err := store.db.Exec(sqlStr, params...)

	if err != nil {
		return err
	}

	user.MarkAsNotDirty()

	return nil
}

func (store *Store) UserDelete(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	return store.UserDeleteByID(user.ID())
}

func (store *Store) UserDeleteByID(id string) error {
	if id == "" {
		return errors.New("user id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.userTableName).
		Prepared(true).
		Where(goqu.C("id").Eq(id)).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	_, err := store.db.Exec(sqlStr, params...)

	return err
}

func (store *Store) UserFindByEmail(email string) (*User, error) {
	if email == "" {
		return nil, errors.New("user email is empty")
	}

	list, err := store.UserList(UserQueryOptions{
		Email: email,
		Limit: 1,
	})

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, nil
}

// UserFindByEmailOrCreate - finds by email or creates a user (with active status)
func (store *Store) UserFindByEmailOrCreate(email string, createStatus string) (*User, error) {
	existingUser, errUser := store.UserFindByEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	if existingUser != nil {
		return existingUser, nil
	}

	newUser := NewUser().
		SetEmail(email).
		SetStatus(createStatus)

	errCreate := store.UserCreate(newUser)

	if errCreate != nil {
		return nil, errCreate
	}

	return newUser, nil
}

func (store *Store) UserFindByID(id string) (*User, error) {
	if id == "" {
		return nil, errors.New("user id is empty")
	}

	list, err := store.UserList(UserQueryOptions{
		ID:    id,
		Limit: 1,
	})

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, nil
}

func (store *Store) UserList(options UserQueryOptions) ([]User, error) {
	q := store.userQuery(options)

	sqlStr, _, errSql := q.Select().ToSQL()

	if errSql != nil {
		return []User{}, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)
	modelMaps, err := db.SelectToMapString(sqlStr)
	if err != nil {
		return []User{}, err
	}

	list := []User{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewUserFromExistingData(modelMap)
		list = append(list, *model)
	})

	return list, nil
}

func (store *Store) UserSoftDelete(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	user.SetDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.UserUpdate(user)
}

func (store *Store) UserSoftDeleteByID(id string) error {
	user, err := store.UserFindByID(id)

	if err != nil {
		return err
	}

	return store.UserSoftDelete(user)
}

func (store *Store) UserUpdate(user *User) error {
	if user == nil {
		return errors.New("order is nil")
	}

	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := user.DataChanged()

	delete(dataChanged, "id") // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.userTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C("id").Eq(user.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	_, err := store.db.Exec(sqlStr, params...)

	user.MarkAsNotDirty()

	return err
}

func (store *Store) userQuery(options UserQueryOptions) *goqu.SelectDataset {
	q := goqu.Dialect(store.dbDriverName).From(store.userTableName)

	if options.ID != "" {
		q = q.Where(goqu.C("id").Eq(options.ID))
	}

	if options.Status != "" {
		q = q.Where(goqu.C("status").Eq(options.Status))
	}

	if len(options.StatusIn) > 0 {
		q = q.Where(goqu.C("status").In(options.StatusIn))
	}

	if options.Email != "" {
		q = q.Where(goqu.C("email").Eq(options.Email))
	}

	if !options.CountOnly {
		if options.Limit > 0 {
			q = q.Limit(uint(options.Limit))
		}

		if options.Offset > 0 {
			q = q.Offset(uint(options.Offset))
		}
	}

	sortOrder := "desc"
	if options.SortOrder != "" {
		sortOrder = options.SortOrder
	}

	if options.OrderBy != "" {
		if strings.EqualFold(sortOrder, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy).Desc())
		}
	}

	if !options.WithDeleted {
		q = q.Where(goqu.C("deleted_at").Eq(sb.NULL_DATETIME))
	}

	return q
}

type UserQueryOptions struct {
	ID          string
	IDIn        []string
	Status      string
	StatusIn    []string
	Email       string
	Offset      int
	Limit       int
	SortOrder   string
	OrderBy     string
	CountOnly   bool
	WithDeleted bool
}

type TimezoneQueryOptions struct {
	ID        string
	Status    string
	StatusIn  []string
	UserCode  string
	Timezone  string
	Offset    int
	Limit     int
	SortOrder string
	SortBy    string
	CountOnly bool
}
