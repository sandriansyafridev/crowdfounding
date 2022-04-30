package service

import (
	"errors"

	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"github.com/sandriansyafridev/crowdfounding/repository"
)

type AuthService interface {
	Login(request dto.LoginDTO) (user entity.User, err error)
	Register(request dto.RegisterDTO) (userCreated entity.User, err error)
}

type AuthServiceImpl struct {
	repository.AuthRepository
}

func NewAuthRepositoryImpl(authRepository repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
	}
}
func (authService *AuthServiceImpl) Login(request dto.LoginDTO) (user entity.User, err error) {

	if user, err = authService.AuthRepository.FindByEmail(request.Email); err != nil {
		return user, err
	} else {
		//validation password
		if user.Password != request.Password {
			return user, errors.New("password not match")
		} else {
			return user, nil
		}
	}

}

func (authService *AuthServiceImpl) Register(request dto.RegisterDTO) (userCreated entity.User, err error) {

	user := entity.User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	if userCreated, err = authService.AuthRepository.Create(user); err != nil {
		return user, err
	} else {
		userCreated.Token = "token" //Generate JWT
		return userCreated, nil
	}

}
