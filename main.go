package main

import (
	"github.com/fatmuh/go-gin-gorm/config"
	"github.com/fatmuh/go-gin-gorm/controller"
	"github.com/fatmuh/go-gin-gorm/helper"
	"github.com/fatmuh/go-gin-gorm/model"
	"github.com/fatmuh/go-gin-gorm/repository"
	"github.com/fatmuh/go-gin-gorm/router"
	"github.com/fatmuh/go-gin-gorm/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("Starting server...")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	routes.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
