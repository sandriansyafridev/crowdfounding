package entity

import "time"

type CampaignImage struct {
	ID                uint64
	CampaignID        uint64    `gorm:"foreignKey"`
	PathCampaignImage string    `gorm:"not null;type:varchar(255); default:default_campaign_image.jpg"`
	IsPrimary         bool      `gorm:"not null;type:boolean;default:false"`
	CreatedAt         time.Time `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
}
