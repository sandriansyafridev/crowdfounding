package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}
func (authController *AuthControllerImpl) Login(c *gin.Context) {

	request := dto.LoginDTO{}
	if err := c.ShouldBind(&request); err != nil {
		response := response.ResponseFail("bad request", err)
		c.JSON(http.StatusBadRequest, response)
	} else {
		if user, err := authController.AuthService.Login(request); err != nil {
			response := response.ResponseFail("internal error", err)
			c.JSON(http.StatusBadRequest, response)
		} else {
			response := response.ResponseSuccess("login success", user)
			c.JSON(http.StatusOK, response)
		}
	}

}

func (authController *AuthControllerImpl) Register(c *gin.Context) {
	request := dto.RegisterDTO{}
	if err := c.ShouldBind(&request); err != nil {
		response := response.ResponseFail("bad request", err)
		c.JSON(http.StatusBadRequest, response)
	} else {
		if userCreated, err := authController.AuthService.Register(request); err != nil {
			response := response.ResponseFail("internal error", err)
			c.JSON(http.StatusOK, response)
		} else {
			response := response.ResponseSuccess("register success", userCreated)
			c.JSON(http.StatusOK, response)
		}
	}
}
