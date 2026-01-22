package usecase

import (
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/repository"
	"time"
)

type UserUseCase interface {
	Create(user model.CreateUserRequest) (*model.UserResponse, error)
	GetById(id int) (*model.UserResponse, error)
	GetAll() []*model.UserResponse
	Delete(id int) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (uc *userUseCase) Create(user model.CreateUserRequest) (*model.UserResponse, error) {
	toSave := &model.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}

	saved, err := uc.userRepo.Save(toSave)
	if err != nil {
		return nil, err
	}

	return uc.toResponse(saved), nil
}

func (uc *userUseCase) GetById(id int) (*model.UserResponse, error) {
	user, err := uc.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return uc.toResponse(user), nil
}

func (uc *userUseCase) GetAll() []*model.UserResponse {
	got := uc.userRepo.GetAll()
	result := make([]*model.UserResponse, 0, len(got))
	for _, u := range got {
		result = append(result, uc.toResponse(u))
	}

	return result
}

func (uc *userUseCase) Delete(id int) error {
	return uc.userRepo.Delete(id)
}

func (uc *userUseCase) toResponse(user *model.User) *model.UserResponse {
	return &model.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
