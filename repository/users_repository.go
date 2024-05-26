package repository

import "github.com/fatmuh/go-gin-gorm/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users) error
	Delete(userId int) error
	FindById(userId int) (model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string) (model.Users, error)
}
