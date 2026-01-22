package repository

import (
	"errors"
	"fmt"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
)

var (
	UserNotFound = errors.New("user was not found")
)

type inMemoryUserRepository struct {
	UserSlice []model.User
}

func NewInMemoryUserRepository() Repository {
	return &inMemoryUserRepository{
		UserSlice: make([]model.User, 0),
	}
}

func (ur *inMemoryUserRepository) GetById(id int) (model.User, error) {
	for _, u := range ur.UserSlice {
		if u.Id == id {
			return u, nil
		}
	}

	return model.User{}, fmt.Errorf("get user by id %d: %v", id, UserNotFound)
}

func (ur *inMemoryUserRepository) GetByEmail(email string) (model.User, error) {
	for _, u := range ur.UserSlice {
		if u.Email == email {
			return u, nil
		}
	}

	return model.User{}, fmt.Errorf("get user by email %s: %v", email, UserNotFound)
}

func (ur *inMemoryUserRepository) GetAll() []model.User {
	return ur.UserSlice
}
