package service

import (
	"errors"
	config2 "github.com/fatmuh/go-gin-gorm/config"
	"github.com/fatmuh/go-gin-gorm/data/request"
	"github.com/fatmuh/go-gin-gorm/helper"
	"github.com/fatmuh/go-gin-gorm/model"
	"github.com/fatmuh/go-gin-gorm/repository"
	"github.com/fatmuh/go-gin-gorm/utils"
	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func (a AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	new_user, user_err := a.UsersRepository.FindByUsername(users.Username)
	if user_err != nil {
		return "", errors.New("Invalid username or password")
	}

	config, _ := config2.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("Invalid username or password")
	}

	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}

func (a AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	if err != nil {
		panic(err)
	}

	newuser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	a.UsersRepository.Save(newuser)
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}
