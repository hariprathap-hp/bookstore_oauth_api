package httpHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hari/bookstore_oauth_api/src/domain/token"
	"github.com/hari/bookstore_oauth_api/src/services"
	"github.com/hari/bookstore_users_api/utils/errors"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service services.UserService
}

func NewAccessTokenHandler(s services.UserService) AccessTokenHandler {
	return &accessTokenHandler{
		service: s,
	}
}

func (a *accessTokenHandler) GetById(c *gin.Context) {
	at, err := a.service.GetToken(c.Param("access_token_id"))
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusOK, at)
}

func (a *accessTokenHandler) Create(c *gin.Context) {
	var request token.AccessTokenRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewHTTPBadRequestError("invalid json body"))
		return
	}
	at, createErr := a.service.CreateToken(request)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}
	c.JSON(http.StatusOK, at)
}
