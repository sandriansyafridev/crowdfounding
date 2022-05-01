package repository

import (
	"errors"

	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() (campaigns []entity.Campaign, err error)
	FindByUserID(UserID uint64) (campaigns []entity.Campaign, err error)
}

type CampaignRepositoryImpl struct {
	gormDB *gorm.DB
}

func NewCampaignRepositoryImpl(gormDB *gorm.DB) *CampaignRepositoryImpl {
	return &CampaignRepositoryImpl{
		gormDB: gormDB,
	}
}

func (campaignRepository *CampaignRepositoryImpl) FindAll() (campaigns []entity.Campaign, err error) {

	if err = campaignRepository.gormDB.Preload("User").Preload("CampaignImage").Find(&campaigns).Error; err != nil {
		return campaigns, err
	} else {
		return campaigns, nil
	}

}

func (campaignRepository *CampaignRepositoryImpl) FindByUserID(UserID uint64) (campaigns []entity.Campaign, err error) {

	if err = campaignRepository.gormDB.Where("user_id = ?", UserID).Preload("User").Preload("CampaignImage").Find(&campaigns).Error; err != nil {
		return campaigns, err
	} else {
		if len(campaigns) == 0 {
			return campaigns, errors.New("user no create campaign!")
		} else {
			return campaigns, nil
		}
	}

}
