package entity

import "time"

type Campaign struct {
	ID            uint64
	UserID        uint64
	User          User      `gorm:"not null;foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name          string    `gorm:"not null;type:varchar(255)"`
	Slug          string    `gorm:"not null;type:varchar(255)"`
	ShortDesc     string    `gorm:"not null;type:varchar(255)"`
	LongDesc      string    `gorm:"not null;type:text"`
	Perk          string    `gorm:"not null;type:text"`
	BackerCount   int       `gorm:"not null; default:0"`
	GoalAmount    uint64    `gorm:"not null; default:0"`
	CurrentAmount uint64    `gorm:"not null; default:0"`
	CreatedAt     time.Time `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
}
