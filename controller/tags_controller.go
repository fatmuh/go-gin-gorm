package controller

import (
	"github.com/fatmuh/go-gin-gorm/data/request"
	"github.com/fatmuh/go-gin-gorm/data/response"
	"github.com/fatmuh/go-gin-gorm/helper"
	"github.com/fatmuh/go-gin-gorm/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(tagsService service.TagsService) *TagsController {
	return &TagsController{
		tagsService: tagsService,
	}
}

// Create Controller
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update Controller
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	err = controller.tagsService.Update(updateTagsRequest)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Controller
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	err = controller.tagsService.Delete(id)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById Controller
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	tagResponse, err := controller.tagsService.FindById(id)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll Controller
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
