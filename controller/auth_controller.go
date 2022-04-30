package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/format"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	service.AuthService
	service.JWTService
}

func NewAuthControllerImpl(authService service.AuthService, jwtService service.JWTService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
		JWTService:  jwtService,
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
			if token, err := authController.JWTService.GenerateToken(user.ID); err != nil {
				response := response.ResponseFail("fail generate token", err)
				c.JSON(http.StatusBadRequest, response)
			} else {
				user.Token = token
				response := response.ResponseSuccess("login success", format.ToUserResponse(user))
				c.JSON(http.StatusOK, response)
			}

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
			if token, err := authController.JWTService.GenerateToken(userCreated.ID); err != nil {
				response := response.ResponseFail("fail generate token", err)
				c.JSON(http.StatusBadRequest, response)
			} else {
				userCreated.Token = token
				response := response.ResponseSuccess("register success", format.ToUserResponse(userCreated))
				c.JSON(http.StatusOK, response)
			}

		}
	}
}
