package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore-oauth-api/domain/accesstoken"
	"github.com/yesseneon/bookstore-utils/errors"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
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

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if restErr := h.service.Create(at); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, at)
}
