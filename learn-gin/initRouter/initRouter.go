package initRouter

import (
	"learn-gin/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// template
	router.LoadHTMLGlob("/Users/butong/learn-go/learn-gin/templates/*.tmpl")
	// static
	router.Static("/statics", "/Users/butong/learn-go/learn-gin/statics")
	// favicon
	router.StaticFile("/favicon.ico", "/Users/butong/learn-go/learn-gin/favicon.ico")

	index := router.Group("/")
	{
		index.Any("/", handler.Index)
	}

	// user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}

	return router
}
