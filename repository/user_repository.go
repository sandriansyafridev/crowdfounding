package repository

import (
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() (users []entity.User, err error)
	FindByID(UserID uint64) (user entity.User, err error)
	FindByEmail(email string) (user entity.User, err error)
	Delete(user entity.User) (err error)
}

type UserRepositoryImpl struct {
	gormDB *gorm.DB
}

func NewUserRepositoryImpl(gormDB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		gormDB: gormDB,
	}
}

func (userRepository *UserRepositoryImpl) FindAll() (users []entity.User, err error) {

	if err = userRepository.gormDB.Find(&users).Error; err != nil {
		return users, err
	} else {
		return users, nil
	}

}
func (userRepository *UserRepositoryImpl) FindByID(UserID uint64) (user entity.User, err error) {
	if err = userRepository.gormDB.First(&user, UserID).Error; err != nil {
		return user, err
	} else if user.ID == 0 {
		return user, err
	} else {
		return user, nil
	}
}

func (userRepository *UserRepositoryImpl) FindByEmail(email string) (user entity.User, err error) {
	if err = userRepository.gormDB.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	} else if user.ID == 0 {
		return user, err
	} else {
		return user, nil
	}
}

func (userRepository *UserRepositoryImpl) Delete(user entity.User) (err error) {

	if err = userRepository.gormDB.Delete(&user).Error; err != nil {
		return err
	} else {
		return nil
	}

}
