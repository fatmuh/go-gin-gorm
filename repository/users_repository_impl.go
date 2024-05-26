package repository

import (
	"errors"
	"github.com/fatmuh/go-gin-gorm/data/request"
	"github.com/fatmuh/go-gin-gorm/helper"
	"github.com/fatmuh/go-gin-gorm/model"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func (u UsersRepositoryImpl) Save(users model.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

func (u UsersRepositoryImpl) Update(users model.Users) error {
	var updateTag = request.UpdateUsersRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	result := u.Db.Model(&users).Updates(updateTag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UsersRepositoryImpl) Delete(userId int) error {
	var users model.Users
	result := u.Db.Where("id = ?", userId).Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("User is not found")
	}
	return nil
}

func (u UsersRepositoryImpl) FindById(userId int) (model.Users, error) {
	var user model.Users
	result := u.Db.First(&user, userId) // Use First instead of Find to ensure single record is fetched
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("User is not found")
		}
		return user, result.Error
	}
	return user, nil
}

func (u UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

func (u UsersRepositoryImpl) FindByUsername(username string) (model.Users, error) {
	var user model.Users
	result := u.Db.First(&user, "username = ?", username) // Use First instead of Find to ensure single record is fetched
	if result.Error != nil {
		return user, errors.New("Invalid username or password")
	}
	return user, nil
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}
