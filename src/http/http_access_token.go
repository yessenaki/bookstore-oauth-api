package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_oauth_api/src/domain/accesstoken"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, restErr := h.service.GetByID(c.Param("access_token_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
