package router

import (
	"github.com/fatmuh/go-gin-gorm/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(tagsController *controller.TagsController, authenticationController *controller.AuthenticationController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome home"})
	})

	baseRouter := router.Group("/api")
	tagRouter := baseRouter.Group("/tags")
	tagRouter.GET("", tagsController.FindAll)
	tagRouter.GET("/:tagId", tagsController.FindById)
	tagRouter.POST("", tagsController.Create)
	tagRouter.PATCH("/:tagId", tagsController.Update)
	tagRouter.DELETE("/:tagId", tagsController.Delete)

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/register", authenticationController.Register)
	authRouter.POST("/login", authenticationController.Login)

	return router
}
