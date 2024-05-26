package repository

import "github.com/fatmuh/go-gin-gorm/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags) error
	Delete(tagsId int) error
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
