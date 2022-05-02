package repository

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() (campaigns []entity.Campaign, err error)
	FindByUserID(UserID uint64) (campaigns []entity.Campaign, err error)
	FindByID(CampaignID uint64) (campaign entity.Campaign, err error)
	Save(campaign entity.Campaign) (entity.Campaign, error)
	CreateSlug(request string) string
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
			return campaigns, errors.New("user no create campaign")
		} else {
			return campaigns, nil
		}
	}

}

func (campaignRepository *CampaignRepositoryImpl) FindByID(CampaignID uint64) (campaign entity.Campaign, err error) {

	if err = campaignRepository.gormDB.Preload("User").Preload("CampaignImage").First(&campaign, CampaignID).Error; err != nil {
		return campaign, err
	} else {
		return campaign, nil
	}

}

func (campaignRepository *CampaignRepositoryImpl) Save(campaign entity.Campaign) (entity.Campaign, error) {
	if err := campaignRepository.gormDB.Save(&campaign).Error; err != nil {
		return campaign, err
	} else {
		return campaign, nil
	}
}

func (campaignRepository *CampaignRepositoryImpl) CreateSlug(request string) string {

	var count int64
	var slug string = slug.Make(request)
	if err := campaignRepository.gormDB.Model(&entity.Campaign{}).Where("name = ?", request).Count(&count).Error; err != nil {
		return ""
	} else {
		switch {
		case count == 0:
			return slug
		default:
			return fmt.Sprintf("%s%d", slug, 1-(count+1))
		}

	}

}
