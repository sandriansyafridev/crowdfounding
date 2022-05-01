package service

import (
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"github.com/sandriansyafridev/crowdfounding/repository"
)

type CampaignService interface {
	GetCampaigns() (campaigns []entity.Campaign, err error)
	GetCampaign(CampaignID uint64) (campaign entity.Campaign, err error)
	GetCampaignsByUserID(UserID uint64) (campaigns []entity.Campaign, err error)
}

type CampaignServiceImp struct {
	repository.CampaignRepository
}

func NewCampaignServiceImpl(campaignRepository repository.CampaignRepository) *CampaignServiceImp {
	return &CampaignServiceImp{
		CampaignRepository: campaignRepository,
	}
}

func (campaignService *CampaignServiceImp) GetCampaigns() (campaigns []entity.Campaign, err error) {
	if campaigns, err = campaignService.CampaignRepository.FindAll(); err != nil {
		return campaigns, err
	} else {
		return campaigns, nil
	}

}

func (campaignService *CampaignServiceImp) GetCampaignsByUserID(UserID uint64) (campaigns []entity.Campaign, err error) {

	if campaigns, err = campaignService.CampaignRepository.FindByUserID(UserID); err != nil {
		return campaigns, err
	} else {
		return campaigns, nil
	}

}

func (campaignService *CampaignServiceImp) GetCampaign(CampaignID uint64) (campaign entity.Campaign, err error) {

	if campaign, err = campaignService.CampaignRepository.FindByID(CampaignID); err != nil {
		return campaign, err
	} else {
		return campaign, nil
	}

}
