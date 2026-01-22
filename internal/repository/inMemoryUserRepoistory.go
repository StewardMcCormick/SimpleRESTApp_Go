package repository

import (
	"errors"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
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

	return model.User{}, errors.New("")
}

func (ur *inMemoryUserRepository) GetByEmail(email string) (model.User, error) {
	for _, u := range ur.UserSlice {
		if u.Email == email {
			return u, nil
		}
	}

	return model.User{}, errors.New("")
}

func (ur *inMemoryUserRepository) GetAll() []model.User {
	return ur.UserSlice
}
