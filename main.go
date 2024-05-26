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

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	// Database
	db := config.DatabaseConnection(&loadConfig)
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("users").AutoMigrate(&model.Users{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)
	usersRepository := repository.NewUsersRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	authenticationService := service.NewAuthenticationServiceImpl(usersRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)
	authenticationController := controller.NewAuthenticationController(authenticationService)

	// Router
	routes := router.NewRouter(tagsController, authenticationController)

	routes.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
