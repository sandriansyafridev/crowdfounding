package format

import (
	"time"

	"github.com/sandriansyafridev/crowdfounding/model/entity"
)

type CampaignResponse struct {
	ID                uint64    `json:"id"`
	UserID            uint64    `json:"user_id"`
	Name              string    `json:"name"`
	ShortDesc         string    `json:"short_desc"`
	Slug              string    `json:"slug"`
	PathCampaignImage string    `json:"path_campaign_image"`
	GoalAmount        uint64    `json:"goal_amount"`
	CurrentAmount     uint64    `json:"current_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FindIsPrimary(campaignImages *[]entity.CampaignImage) (path string) {

	if len(*campaignImages) != 0 {
		for _, campaignImage := range *campaignImages {
			if campaignImage.IsPrimary {
				path = campaignImage.PathCampaignImage
				break
			} else {
				path = "default_campaign_image.jpg"
			}
		}
	}

	return path

}

func ToCampaignResponse(campaign entity.Campaign) (campaignResponse CampaignResponse) {

	campaignResponse.ID = campaign.ID
	campaignResponse.UserID = campaign.UserID
	campaignResponse.Name = campaign.Name
	campaignResponse.ShortDesc = campaign.ShortDesc
	campaignResponse.Slug = campaign.Slug
	campaignResponse.PathCampaignImage = FindIsPrimary(&campaign.CampaignImage)
	campaignResponse.GoalAmount = campaign.GoalAmount
	campaignResponse.CurrentAmount = campaign.CurrentAmount
	campaignResponse.CreatedAt = campaign.CreatedAt
	campaignResponse.UpdatedAt = campaign.UpdatedAt

	return campaignResponse
}

func ToCampaignsResponse(campaigns []entity.Campaign) (campaignsResponse []CampaignResponse) {

	for _, campaign := range campaigns {
		campaignsResponse = append(campaignsResponse, ToCampaignResponse(campaign))
	}

	return campaignsResponse

}
