package service

import (
	"errors"
	"strings"
	"time"

	"github.com/sandriansyafridev/crowdfounding/model/dto"
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"github.com/sandriansyafridev/crowdfounding/repository"
)

type CampaignService interface {
	GetCampaigns() (campaigns []entity.Campaign, err error)
	GetCampaign(CampaignID uint64) (campaign entity.Campaign, err error)
	GetCampaignsByUserID(UserID uint64) (campaigns []entity.Campaign, err error)
	CreateCampaign(request dto.CampaignCreateDTO) (entity.Campaign, error)
	UpdateCampaign(request dto.CampaignUpdateDTO) (entity.Campaign, error)
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

func (campaignService *CampaignServiceImp) CreateCampaign(request dto.CampaignCreateDTO) (entity.Campaign, error) {

	campaign := entity.Campaign{}
	campaign.UserID = request.UserID
	campaign.Name = strings.ToLower(request.Name)
	campaign.Slug = campaignService.CampaignRepository.CreateSlug(strings.ToLower(request.Name))
	campaign.ShortDesc = strings.ToLower(request.ShortDesc)
	campaign.LongDesc = request.LongDesc
	campaign.Perk = request.Perk
	campaign.GoalAmount = request.GoalAmount
	if campaignCreated, err := campaignService.CampaignRepository.Save(campaign); err != nil {
		return campaignCreated, err
	} else {
		return campaignCreated, nil
	}

}

func (campaignService *CampaignServiceImp) UpdateCampaign(request dto.CampaignUpdateDTO) (entity.Campaign, error) {

	if campaign, err := campaignService.FindByID(request.ID); err != nil {
		return campaign, err
	} else if ok := request.UserID == campaign.UserID; !ok {
		return campaign, errors.New("not allow to updated campaign")
	} else {
		campaign.Name = strings.ToLower(request.Name)
		campaign.ShortDesc = strings.ToLower(request.ShortDesc)
		campaign.LongDesc = request.LongDesc
		campaign.Perk = request.Perk
		campaign.GoalAmount = request.GoalAmount
		campaign.UpdatedAt = time.Now()
		return campaign, nil
	}

}
