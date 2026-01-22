package repository

import (
	"fmt"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
	"slices"
)

type UserRepository interface {
	GetById(id int) (*model.User, error)
	GetAll() []*model.User
	Save(user *model.User) (*model.User, error)
	Delete(id int) error
}

type inMemoryUserRepository struct {
	UserSlice []*model.User
	lastId    int
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		UserSlice: make([]*model.User, 0),
	}
}

func (ur *inMemoryUserRepository) GetById(id int) (*model.User, error) {
	for _, u := range ur.UserSlice {
		if u.Id == id {
			return u, nil
		}
	}

	return nil, fmt.Errorf("get user by id %d: %v", id, UserNotFound)
}

func (ur *inMemoryUserRepository) GetAll() []*model.User {
	return ur.UserSlice
}

func (ur *inMemoryUserRepository) Save(user *model.User) (*model.User, error) {
	user.Id = ur.lastId
	ur.UserSlice = append(ur.UserSlice, user)

	ur.lastId++
	return user, nil
}

func (ur *inMemoryUserRepository) Delete(id int) error {
	for i, u := range ur.UserSlice {
		if u.Id == id {
			ur.UserSlice = slices.Delete(ur.UserSlice, i, i+1)
			return nil
		}
	}

	return UserNotFound
}
