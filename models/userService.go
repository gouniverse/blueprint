package models

import "project/config"

type userService struct {
	repository userRepositoryInterface
}

func NewUserService() *userService {
	var theRepository userRepositoryInterface
	if config.UserStore != nil {
		theRepository = NewUserEntityRepository()
	} else {
		panic("User store not setup")
	}
	return &userService{repository: theRepository}
}

func (service *userService) UserCreate(user *User) error {
	return service.repository.UserCreate(user)
}

func (service *userService) UserFindByID(id string) (*User, error) {
	return service.repository.UserFindByID(id)
}

func (service *userService) UserList(options UserQueryOptions) ([]User, error) {
	return service.repository.UserList(options)
}

func (service *userService) UserDelete(user *User) error {
	return service.repository.UserDelete(user)
}

func (service *userService) UserUpdate(user *User) error {
	return service.repository.UserUpdate(user)
}
