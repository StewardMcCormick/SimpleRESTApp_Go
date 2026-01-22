package repository

import (
	"fmt"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
)

type inMemoryUserRepository struct {
	UserSlice []model.User
}

func NewInMemoryUserRepository() UserRepository {
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

func (ur *inMemoryUserRepository) GetAll() []model.User {
	return ur.UserSlice
}

func (ur *inMemoryUserRepository) Save(user model.User) (model.User, error) {
	ur.UserSlice = append(ur.UserSlice, user)
	return user, nil
}
