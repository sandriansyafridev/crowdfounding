package format

import (
	"strings"
	"time"

	"github.com/sandriansyafridev/crowdfounding/model/entity"
)

type User struct {
	ID               uint64 `json:"id"`
	Name             string `json:"name"`
	PathProfileImage string `json:"path_profile_image"`
}

type CampaignImage struct {
	PathCampaignImage string `json:"path_campaign_image"`
	IsPrimary         bool   `json:"is_primary"`
}

type CampaignDetailResponse struct {
	ID uint64 `json:"id"`

	Name                     string          `json:"name"`
	ShortDesc                string          `json:"short_desc"`
	LongDesc                 string          `json:"long_desc"`
	Slug                     string          `json:"slug"`
	PathPrimaryCampaignImage string          `json:"path_primary_campaign_image"`
	Perks                    []string        `json:"perks"`
	GoalAmount               uint64          `json:"goal_amount"`
	CurrentAmount            uint64          `json:"current_amount"`
	User                     User            `json:"user"`
	CampaignImage            []CampaignImage `json:"campaign_images"`
	CreatedAt                time.Time       `json:"created_at"`
	UpdatedAt                time.Time       `json:"updated_at"`
}

func ToCampaignImageResponse(campaignImages *[]entity.CampaignImage) (campaignImagesResponse []CampaignImage) {

	for _, item := range *campaignImages {
		campaignImage := CampaignImage{}
		campaignImage.PathCampaignImage = item.PathCampaignImage
		campaignImage.IsPrimary = item.IsPrimary
		campaignImagesResponse = append(campaignImagesResponse, campaignImage)
	}

	return campaignImagesResponse
}

func ToCampaignDetailResponse(campaign entity.Campaign) (campaignDetailResponse CampaignDetailResponse) {

	campaignDetailResponse.ID = campaign.ID
	campaignDetailResponse.Name = campaign.Name
	campaignDetailResponse.ShortDesc = campaign.ShortDesc
	campaignDetailResponse.LongDesc = campaign.LongDesc
	campaignDetailResponse.Slug = campaign.Slug
	campaignDetailResponse.GoalAmount = campaign.GoalAmount
	campaignDetailResponse.CurrentAmount = campaign.CurrentAmount
	campaignDetailResponse.CreatedAt = campaign.CreatedAt
	campaignDetailResponse.UpdatedAt = campaign.UpdatedAt
	campaignDetailResponse.PathPrimaryCampaignImage = FindIsPrimary(&campaign.CampaignImage)

	campaignDetailResponse.User.ID = campaign.User.ID
	campaignDetailResponse.User.Name = campaign.User.Name
	campaignDetailResponse.User.PathProfileImage = campaign.User.PathProfileImage

	campaignDetailResponse.CampaignImage = ToCampaignImageResponse(&campaign.CampaignImage)

	splitPerks := strings.Split(campaign.Perk, ",")
	campaignDetailResponse.Perks = splitPerks

	return

}
