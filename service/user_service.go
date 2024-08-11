package service

import (
	"context"

	"boilerplate/store/dtos"
	"boilerplate/store/entity"
	"boilerplate/store/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(ctx context.Context, input *dtos.UserDTO) error
	Get(ctx context.Context, id string) (dtos.UserDTO, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

// TODO - change implementation to return pointer to service and make receivers pointer based for lite container

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (u UserService) Create(ctx context.Context, input *dtos.UserDTO) error {
	var payload entity.User
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return err
	}

	payload.Name = input.Name
	payload.Email = input.Email
	payload.Password = string(passwordHashBytes)

	id, insertErr := u.userRepository.Create(ctx, payload)
	if insertErr != nil {
		return insertErr
	}

	input.ID = id

	return nil
}

func (u UserService) Get(ctx context.Context, id string) (dtos.UserDTO, error) {
	var result dtos.UserDTO
	userEntity, err := u.userRepository.Get(ctx, id)
	if err != nil {
		return result, err
	}

	result.ID = userEntity.ID
	result.Name = userEntity.Name
	result.Email = userEntity.Email
	result.CreatedAt = userEntity.CreatedAt
	result.UpdatedAt = userEntity.UpdatedAt

	return result, nil
}
