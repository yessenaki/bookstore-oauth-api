package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/yesseneon/bookstore_oauth_api/src/http"
	"github.com/yesseneon/bookstore_oauth_api/src/repository/db"
)

var router = gin.Default()

func StartApp() {
	atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}
