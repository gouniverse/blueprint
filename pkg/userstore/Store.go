package userstore

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
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

	if store.db == nil {
		return errors.New("userstore: database is nil")
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

func (store *Store) UserCount(options UserQueryOptions) (int64, error) {
	options.CountOnly = true
	q := store.userQuery(options)

	sqlStr, params, errSql := q.Prepared(true).
		Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)
	mapped, err := db.SelectToMapString(sqlStr, params...)
	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
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

	if store.db == nil {
		return errors.New("userstore: database is nil")
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

	if store.db == nil {
		return []User{}, errors.New("userstore: database is nil")
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)

	if db == nil {
		return []User{}, errors.New("userstore: database is nil")
	}

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
		return errors.New("user is nil")
	}

	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := user.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.userTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(user.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	_, err := store.db.Exec(sqlStr, params...)

	user.MarkAsNotDirty()

	return err
}

func (store *Store) userQuery(options UserQueryOptions) *goqu.SelectDataset {
	q := goqu.Dialect(store.dbDriverName).From(store.userTableName)

	if options.ID != "" {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID))
	}

	if len(options.IDIn) > 0 {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn))
	}

	if options.Status != "" {
		q = q.Where(goqu.C(COLUMN_STATUS).Eq(options.Status))
	}

	if len(options.StatusIn) > 0 {
		q = q.Where(goqu.C(COLUMN_STATUS).In(options.StatusIn))
	}

	if options.Email != "" {
		q = q.Where(goqu.C(COLUMN_EMAIL).Eq(options.Email))
	}

	if options.CreatedAtGte != "" && options.CreatedAtLte != "" {
		q = q.Where(
			goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte),
			goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte),
		)
	} else if options.CreatedAtGte != "" {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte))
	} else if options.CreatedAtLte != "" {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte))
	}

	if !options.CountOnly {
		if options.Limit > 0 {
			q = q.Limit(uint(options.Limit))
		}

		if options.Offset > 0 {
			q = q.Offset(uint(options.Offset))
		}
	}

	sortOrder := sb.DESC
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
		q = q.Where(goqu.C(COLUMN_DELETED_AT).Eq(sb.NULL_DATETIME))
	}

	return q
}

type UserQueryOptions struct {
	ID           string
	IDIn         []string
	Status       string
	StatusIn     []string
	Email        string
	CreatedAtGte string
	CreatedAtLte string
	Offset       int
	Limit        int
	SortOrder    string
	OrderBy      string
	CountOnly    bool
	WithDeleted  bool
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
