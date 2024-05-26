package service

import (
	"github.com/fatmuh/go-gin-gorm/data/request"
	"github.com/fatmuh/go-gin-gorm/data/response"
	"github.com/fatmuh/go-gin-gorm/helper"
	"github.com/fatmuh/go-gin-gorm/model"
	"github.com/fatmuh/go-gin-gorm/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func (t TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t TagsServiceImpl) Update(tags request.UpdateTagsRequest) error {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	if err != nil {
		return err
	}

	tagData.Name = tags.Name
	err = t.TagsRepository.Update(tagData)
	if err != nil {
		return err
	}
	return nil
}

func (t TagsServiceImpl) Delete(tagsId int) error {
	return t.TagsRepository.Delete(tagsId)
}

func (t TagsServiceImpl) FindById(tagsId int) (response.TagsResponse, error) {
	tagData, err := t.TagsRepository.FindById(tagsId)
	if err != nil {
		return response.TagsResponse{}, err
	}

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse, nil
}

func (t TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{TagsRepository: tagsRepository, Validate: validate}
}
