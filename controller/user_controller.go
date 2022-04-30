package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/model/format"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/service"
)

type UserController interface {
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	Delete(c *gin.Context)
}

type UserControllerImpl struct {
	service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) GetUsers(c *gin.Context) {

	if users, err := userController.UserService.GetUsers(); err != nil {
		responseFail := response.ResponseFail("fail to get users", err)
		c.JSON(http.StatusNotFound, responseFail)
	} else {
		responseSuccess := response.ResponseSuccess("success to get users", format.ToUsersResponse(users))
		c.JSON(http.StatusOK, responseSuccess)
	}

}

func (userController *UserControllerImpl) GetUserByID(c *gin.Context) {

	if UserID, err := strconv.Atoi(c.Param("id")); err != nil {
		responseFail := response.ResponseFail("fail to fetch param user 'id'", err)
		c.JSON(http.StatusNotFound, responseFail)
	} else {
		if user, err := userController.UserService.GetUserByID(uint64(UserID)); err != nil {
			responseFail := response.ResponseFail("fail to get users", err)
			c.JSON(http.StatusNotFound, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("success to get users", format.ToUserResponse(user))
			c.JSON(http.StatusOK, responseSuccess)
		}
	}

}
func (userController *UserControllerImpl) Delete(c *gin.Context) {
	if UserID, err := strconv.Atoi(c.Param("id")); err != nil {
		responseFail := response.ResponseFail("fail to fetch param user 'id'", err)
		c.JSON(http.StatusNotFound, responseFail)
	} else {
		if userDeleted, err := userController.UserService.GetUserByID(uint64(UserID)); err != nil {
			responseFail := response.ResponseFail("fail to fetch param user 'id'", err)
			c.JSON(http.StatusOK, responseFail)
		} else {
			responseSuccess := response.ResponseSuccess("success to delete user", format.ToUserResponse(userDeleted))
			c.JSON(http.StatusOK, responseSuccess)
		}
	}
}
