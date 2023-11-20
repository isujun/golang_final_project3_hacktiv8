package service

import (
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/pkg/helpers"
	user_repository "final_project_3/repository/user_repository"
)

type userService struct {
	UserRepo user_repository.Repository
}

type UserService interface {
	CreateUser(userPayload *dto.NewUserRequest) (*dto.NewUserResponse, errs.Error)
	Login(loginPayload *dto.NewLoginRequest) (*dto.NewLoginResponse, errs.Error)
	UpdateUser(userPayload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.Error)
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{UserRepo: userRepo}
}

func (us *userService) CreateUser(userPayload *dto.NewUserRequest) (*dto.NewUserResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(userPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	user := entity.User{
		FullName: userPayload.FullName,
		Email:    userPayload.Email,
		Password: userPayload.Password,
	}
	createdUser, err := us.UserRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	response := dto.NewUserResponse{
		ID:        createdUser.ID,
		FullName:  createdUser.FullName,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}
	return &response, nil
}

func (us *userService) Login(loginPayload *dto.NewLoginRequest) (*dto.NewLoginResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(loginPayload)
	if validateErr != nil {
		return nil, validateErr
	}

	user, err := us.UserRepo.FindOneUserByEmail(loginPayload.Email)
	if err != nil {
		return nil, err
	}
	var response dto.NewLoginResponse
	passwordMatch := helpers.ComparePass([]byte(user.Password), []byte(loginPayload.Password))
	if !passwordMatch {
		return nil, errs.NewUnauthenticatedError("wrong password")
	} else {
		token, generateErr := helpers.GenerateToken(int(user.ID), user.Email, user.Role)
		if generateErr != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		response.Token = token
	}
	return &response, nil
}

func (us *userService) UpdateUser(userPayload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(userPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	var user = entity.User{
		ID: userPayload.ID,
		FullName: userPayload.FullName,
		Email: userPayload.Email,
	}
	updatedUser, err := us.UserRepo.UpdateUser(&user)
	if err != nil {
		return nil, err
	}
	
	response := dto.UpdateUserResponse{
		ID: updatedUser.ID,
		FullName: updatedUser.FullName,
		Email: updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt,
	}
	return &response, nil
} 