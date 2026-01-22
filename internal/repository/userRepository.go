package repository

import "github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"

type Repository interface {
	GetById(id int) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() []model.User
}
