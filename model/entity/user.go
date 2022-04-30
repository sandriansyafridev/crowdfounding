package entity

import "time"

type User struct {
	ID               uint64
	Name             string `gorm:"type:varchar(50); not null"`
	Email            string `gorm:"type:varchar(255); not null"`
	Password         string `gorm:"type:varchar(255); not null"`
	Token            string `gorm:"-"`
	PathProfileImage string `gorm:"type:varchar(255);default:default_profile_image.jpg; null;"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
