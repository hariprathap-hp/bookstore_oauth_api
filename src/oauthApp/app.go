package oauthApp

import (
	"github.com/gin-gonic/gin"
	"github.com/hari/bookstore_oauth_api/src/httpHandler"
	"github.com/hari/bookstore_oauth_api/src/repo/dbRepo"
	"github.com/hari/bookstore_oauth_api/src/repo/userRepo"
	"github.com/hari/bookstore_oauth_api/src/services"
)

var (
	router = gin.Default()
)

func StartApp() {
	atHandler := httpHandler.NewAccessTokenHandler(
		services.NewService(dbRepo.NewRepository(), userRepo.NewUserRepo()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8081")
}
