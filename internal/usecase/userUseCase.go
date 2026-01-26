package usecase

import (
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/repository"
	"time"
)

type UserUseCase interface {
	Create(user model.PostUserRequest) (*model.UserResponse, error)
	GetById(id int) (*model.UserResponse, error)
	GetAll() []*model.UserResponse
	Delete(id int) error
	Put(id int, user model.PutUserRequest) (*model.UserResponse, error)
	Patch(id int, user model.PatchUserRequest) (*model.UserResponse, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (uc *userUseCase) Create(user model.PostUserRequest) (*model.UserResponse, error) {
	toSave := &model.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func (uc *userUseCase) Put(id int, user model.PutUserRequest) (*model.UserResponse, error) {
	userFromDb, err := uc.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	userFromDb.Username = user.Username
	userFromDb.Email = user.Email
	userFromDb.Password = user.Password
	userFromDb.UpdatedAt = time.Now()
	if err = uc.userRepo.Put(userFromDb); err != nil {
		return nil, err
	}

	return uc.toResponse(userFromDb), nil
}

func (uc *userUseCase) Patch(id int, user model.PatchUserRequest) (*model.UserResponse, error) {
	userToUpdate, err := uc.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	updated := false

	if user.Username != nil && *user.Username != userToUpdate.Username {
		userToUpdate.Username = *user.Username
		updated = true
	}
	if user.Email != nil && *user.Email != userToUpdate.Email {
		userToUpdate.Email = *user.Email
	}
	if user.Password != nil && *user.Password != userToUpdate.Password {
		userToUpdate.Password = *user.Password
		updated = true
	}

	if !updated {
		return uc.toResponse(userToUpdate), nil
	}

	userToUpdate.UpdatedAt = time.Now()

	err = uc.userRepo.Patch(userToUpdate)
	if err != nil {
		return nil, err
	}

	return uc.toResponse(userToUpdate), nil
}

func (uc *userUseCase) toResponse(user *model.User) *model.UserResponse {
	return &model.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
