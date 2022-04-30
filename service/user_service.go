package service

import (
	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"github.com/sandriansyafridev/crowdfounding/repository"
)

type UserService interface {
	GetUsers() (users []entity.User, err error)
	GetUserByID(UserID uint64) (user entity.User, err error)
	DeleteUser(user entity.User) (userDeleted entity.User, err error)
	UploadProfileImage(request dto.UserProfileImageDTO) (user entity.User, err error)
}

type UserServiceImpl struct {
	repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (userService *UserServiceImpl) GetUsers() (users []entity.User, err error) {

	if users, err = userService.UserRepository.FindAll(); err != nil {
		return users, err
	} else {
		return users, nil
	}

}
func (userService *UserServiceImpl) GetUserByID(UserID uint64) (user entity.User, err error) {
	if user, err = userService.UserRepository.FindByID(UserID); err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (userService *UserServiceImpl) DeleteUser(user entity.User) (userDeleted entity.User, err error) {

	if userDeleted, err = userService.FindByID(user.ID); err != nil {
		return userDeleted, err
	} else {
		if err = userService.UserRepository.Delete(userDeleted); err != nil {
			return userDeleted, err
		} else {
			return userDeleted, nil
		}
	}
}

func (userService *UserServiceImpl) UploadProfileImage(request dto.UserProfileImageDTO) (user entity.User, err error) {
	if user, err = userService.UserRepository.FindByID(request.UserID); err != nil {
		return user, err
	} else {
		user.PathProfileImage = request.PathProfileImage
		if userUpdated, err := userService.UserRepository.Update(user); err != nil {
			return userUpdated, err
		} else {
			return userUpdated, nil

		}
	}
}
