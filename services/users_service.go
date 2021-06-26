package services

import (
	"main/domain/users"
	"main/utils/crypto_utils"
	"main/utils/date_utils"
	"main/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUser(string) (users.Users, *errors.RestError)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.NewBadReqError("invalid user id")
	}
	result := &users.User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowString()
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {

	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
