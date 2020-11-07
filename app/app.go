package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore-oauth-api/domain/accesstoken"
	"github.com/yesseneon/bookstore-oauth-api/http"
	"github.com/yesseneon/bookstore-oauth-api/repository/db"
)

var router = gin.Default()

func StartApp() {
	atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8081")
}
