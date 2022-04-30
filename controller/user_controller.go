package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/format"
	"github.com/sandriansyafridev/crowdfounding/model/response"
	"github.com/sandriansyafridev/crowdfounding/service"
)

type UserController interface {
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	Delete(c *gin.Context)
	UploadProfileImage(c *gin.Context)
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

func (userController *UserControllerImpl) UploadProfileImage(c *gin.Context) {
	request := dto.UserProfileImageDTO{}
	request.UserID = 3 //hardcore temporary :)

	if file, err := c.FormFile("file_profile_image"); err != nil {
		responseFail := response.ResponseFail("no file image upload", err)
		c.JSON(http.StatusNotFound, responseFail)
	} else {
		dst := "public/assets/images/users/"
		fileName := fmt.Sprintf("profile-user-image-%d-%s", request.UserID, file.Filename)
		pathProfileImage := dst + fileName

		if err := c.SaveUploadedFile(file, pathProfileImage); err != nil {
			responseFail := response.ResponseFail("fail to upload profile image", err)
			c.JSON(http.StatusNotFound, responseFail)
		} else {
			request.PathProfileImage = pathProfileImage
			if user, err := userController.UserService.UploadProfileImage(request); err != nil {
				responseFail := response.ResponseFail("fail update profile image user", err)
				c.JSON(http.StatusNotFound, responseFail)
			} else {
				responseSuccess := response.ResponseSuccess("succes update profile image user", gin.H{
					"is_upload": true,
					"user":      user,
				})
				c.JSON(http.StatusNotFound, responseSuccess)
			}
		}

	}

}
