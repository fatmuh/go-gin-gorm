package service

import (
	"github.com/fatmuh/go-gin-gorm/data/request"
	"github.com/fatmuh/go-gin-gorm/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest) error
	Delete(tagsId int) error
	FindById(tagsId int) (response.TagsResponse, error)
	FindAll() []response.TagsResponse
}
