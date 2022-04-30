package app

import (
	"github.com/sandriansyafridev/crowdfounding/model/entity"
	"gorm.io/gorm"
)

func NewMigration(gormDB *gorm.DB) {
	gormDB.AutoMigrate(
		&entity.User{},
		&entity.Campaign{},
		&entity.CampaignImage{},
	)
}
