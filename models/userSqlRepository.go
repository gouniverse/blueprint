package models

import (
	"errors"
	"log"
	"project/config"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/sb"
	"github.com/mingrammer/cfmt"
	"github.com/samber/lo"
)

const USER_TABLE_NAME = "users_user"

var _ userRepositoryInterface = (*userSqlRepository)(nil) // verify it extends the data object interface

type userSqlRepository struct {
	debug bool
}

func NewUserSqlRepository() *userSqlRepository {
	return &userSqlRepository{}
}

func (repository *userSqlRepository) UserTableCreate() error {
	sql := sb.NewBuilder(sb.DatabaseDriverName(config.Database.DB())).
		Table(USER_TABLE_NAME).
		Column(sb.Column{
			Name:       "id",
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   "status",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "email",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "first_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "middle_names",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "last_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "business_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "role",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "country",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "timezone",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "profile_image_url",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name:   "phone",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name: "memo",
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: "created_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "updated_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "deleted_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	_, err := config.Database.Exec(sql)

	if err != nil {
		cfmt.Errorln("User order table failed to be created:", err.Error())
	}

	return err
}

func (repository *userSqlRepository) UserCreate(user *User) error {
	user.SetCreatedAt(carbon.Now(carbon.UTC).ToDateString(carbon.UTC))
	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateString(carbon.UTC))
	user.SetDeletedAt(sb.NULL_DATETIME)

	data := user.Data()

	sqlStr, params, errSql := goqu.Dialect(config.Database.Type()).
		Insert(USER_TABLE_NAME).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if repository.debug {
		log.Println(sqlStr)
	}

	_, err := config.Database.Exec(sqlStr, params...)

	if err != nil {
		return err
	}

	user.MarkAsNotDirty()

	return nil
}

func (repository *userSqlRepository) UserSoftDelete(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	user.SetDeletedAt(carbon.Now(carbon.UTC).ToDateString(carbon.UTC))

	return repository.UserUpdate(user)
}

func (repository *userSqlRepository) UserSoftDeleteByID(id string) error {
	user, err := repository.UserFindByID(id)

	if err != nil {
		return err
	}

	return repository.UserSoftDelete(user)
}

func (repository *userSqlRepository) UserFindByID(id string) (*User, error) {
	if id == "" {
		return nil, errors.New("user id is empty")
	}

	list, err := repository.UserList(UserQueryOptions{
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

func (repository *userSqlRepository) UserFindByEmail(email string) (*User, error) {
	if email == "" {
		return nil, errors.New("user email is empty")
	}

	list, err := repository.UserList(UserQueryOptions{
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

func (repository *userSqlRepository) UserList(options UserQueryOptions) ([]User, error) {
	q := repository.userQuery(options)

	sqlStr, _, errSql := q.Select().ToSQL()

	if errSql != nil {
		return []User{}, nil
	}

	if repository.debug {
		log.Println(sqlStr)
	}

	modelMaps, err := config.Database.SelectToMapString(sqlStr)
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

func (repository *userSqlRepository) UserUpdate(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	user.SetUpdatedAt(carbon.NewCarbon().Now(carbon.UTC).Format("Y-m-d H:i:s", carbon.UTC))

	dataChanged := user.DataChanged()

	delete(dataChanged, "id") // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(config.Database.Type()).
		Update(USER_TABLE_NAME).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C("id").Eq(user.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if repository.debug {
		log.Println(sqlStr)
	}

	_, err := config.Database.Exec(sqlStr, params...)

	user.MarkAsNotDirty()

	return err
}

func (repository *userSqlRepository) userQuery(options UserQueryOptions) *goqu.SelectDataset {
	q := goqu.Dialect(config.Database.Type()).From(USER_TABLE_NAME)

	if options.ID != "" {
		q = q.Where(goqu.C("id").Eq(options.ID))
	}

	if options.Email != "" {
		q = q.Where(goqu.C("email").Eq(options.Email))
	}

	if options.Status != "" {
		q = q.Where(goqu.C("status").Eq(options.Status))
	}

	if len(options.StatusIn) > 0 {
		q = q.Where(goqu.C("status").In(options.StatusIn))
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
		q = q.Where(goqu.C("deleted").Neq(sb.NULL_DATETIME))
	}

	return q
}
