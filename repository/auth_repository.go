package repository

import (
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	IsEmailAvailable(email string) bool
	FindByEmail(email string) (user entity.User, err error)
	Create(user entity.User) (entity.User, error)
}

type AuthRepositoryImpl struct {
	gormDB *gorm.DB
}

func NewAuthRepositoryImpl(gormDB *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		gormDB: gormDB,
	}
}
func (authRepository *AuthRepositoryImpl) IsEmailAvailable(email string) bool {

	if err := authRepository.gormDB.Where("email = ?", email).Error; err != nil {
		return true
	} else {
		return false
	}
}

func (authRepository *AuthRepositoryImpl) FindByEmail(email string) (user entity.User, err error) {

	if err = authRepository.gormDB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (authRepository *AuthRepositoryImpl) Create(user entity.User) (entity.User, error) {

	//password hash
	if err := authRepository.gormDB.Create(&user).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}

}
