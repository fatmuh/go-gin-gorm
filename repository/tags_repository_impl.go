package repository

import (
	"errors"
	"github.com/fatmuh/go-gin-gorm/data/request"
	"github.com/fatmuh/go-gin-gorm/helper"
	"github.com/fatmuh/go-gin-gorm/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func (t TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) Update(tags model.Tags) error {
	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}

	result := t.Db.Model(&tags).Updates(updateTag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t TagsRepositoryImpl) Delete(tagsId int) error {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Tag is not found")
	}
	return nil
}

func (t TagsRepositoryImpl) FindById(tagsId int) (model.Tags, error) {
	var tag model.Tags
	result := t.Db.First(&tag, tagsId) // Use First instead of Find to ensure single record is fetched
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tag, errors.New("Tag is not found")
		}
		return tag, result.Error
	}
	return tag, nil
}

func (t TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}
