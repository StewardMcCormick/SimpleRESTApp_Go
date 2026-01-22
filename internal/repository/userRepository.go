package repository

import "github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"

type UserRepository interface {
	GetById(id int) (model.User, error)
	GetAll() []model.User
	Save(user model.User) (model.User, error)
}
