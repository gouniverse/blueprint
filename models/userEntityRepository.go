package models

import (
	"errors"
	"project/config"

	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/entitystore"
	"github.com/samber/lo"
)

const USER_ENTITY_TYPE = "user"

var _ userRepositoryInterface = (*userEntityRepository)(nil) // verify it extends the data object interface

type userEntityRepository struct{}

func NewUserEntityRepository() *userEntityRepository {
	return &userEntityRepository{}
}

func (repository *userEntityRepository) UserCreate(user *User) error {
	user.SetCreatedAt(carbon.Now(carbon.UTC).Format(FORMAT_DATETIME, carbon.UTC))
	user.SetUpdatedAt(carbon.Now(carbon.UTC).Format(FORMAT_DATETIME, carbon.UTC))
	data := user.Data()

	entity := config.UserStore.NewEntity(entitystore.NewEntityOptions{
		ID:        user.ID(),
		Type:      USER_ENTITY_TYPE,
		Handle:    user.Email(),
		CreatedAt: user.CreatedAtCarbon().ToStdTime(),
		UpdatedAt: user.UpdatedAtCarbon().ToStdTime(),
	})

	err := config.UserStore.EntityCreate(&entity)

	if err != nil {
		return err
	}

	lo.ForEach(lo.Keys(data), func(key string, index int) {
		err := entity.SetString(key, data[key])
		if err != nil {
			return
		}
	})

	return nil
}

func (repository *userEntityRepository) UserDelete(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	return repository.UserDeleteByID(user.ID())
}

func (repository *userEntityRepository) UserDeleteByID(id string) error {
	_, err := config.UserStore.EntityTrash(id)

	if err != nil {
		return err
	}

	return nil
}

func (repository *userEntityRepository) UserFindByID(id string) (*User, error) {
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

func (repository *userEntityRepository) UserFindByEmail(email string) (*User, error) {
	entity, err := config.UserStore.EntityFindByAttribute(USER_ENTITY_TYPE, "email", email)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, nil
	}

	return repository.entityToUser(*entity)
}

func (repository *userEntityRepository) UserList(options UserQueryOptions) ([]User, error) {
	entityList, errEntityList := config.UserStore.EntityList(entitystore.EntityQueryOptions{
		EntityType: USER_ENTITY_TYPE,
		ID:         options.ID,
		Limit:      uint64(options.Limit),
		Offset:     uint64(options.Offset),
		CountOnly:  options.CountOnly,
	})

	if errEntityList != nil {
		return []User{}, errEntityList
	}

	list := []User{}

	for _, entity := range entityList {
		user, errUser := repository.entityToUser(entity)

		if errUser != nil {
			return []User{}, errUser
		}

		if user == nil {
			return []User{}, errors.New("user " + entity.ID() + " could not be mapped")
		}

		list = append(list, *user)
	}

	return list, nil
}

func (repository *userEntityRepository) UserUpdate(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	user.SetUpdatedAt(carbon.NewCarbon().Now(carbon.UTC).Format("Y-m-d H:i:s", carbon.UTC))
	dataChanged := user.DataChanged()

	delete(dataChanged, "id") // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	entity, err := config.UserStore.EntityFindByID(user.ID())

	if err != nil {
		return err
	}

	lo.ForEach(lo.Keys(dataChanged), func(key string, index int) {
		err := entity.SetString(key, dataChanged[key])
		if err != nil {
			return
		}
	})

	return nil
}

func (repository *userEntityRepository) entityToUser(entity entitystore.Entity) (*User, error) {
	attrs, err := entity.GetAttributes()
	if err != nil {
		return nil, err
	}

	modelMap := map[string]string{
		"id":         entity.ID(),
		"handle":     entity.Handle(),
		"created_at": entity.CreatedAt().String(),
		"updated_at": entity.UpdatedAt().String(),
	}

	lo.ForEach(attrs, func(attr entitystore.Attribute, index int) {
		modelMap[attr.AttributeKey()] = attr.AttributeValue()
	})

	return NewUserFromExistingData(modelMap), nil
}
