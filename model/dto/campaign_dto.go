package dto

type CampaignCreateDTO struct {
	UserID     uint64
	Name       string `json:"name" binding:"required"`
	ShortDesc  string `json:"short_desc" binding:"required"`
	LongDesc   string `json:"long_desc" binding:"required"`
	Perk       string `json:"perk" binding:"required"`
	GoalAmount uint64 `json:"goal_amount" binding:"required"`
}

type CampaignUpdateDTO struct {
	ID         uint64
	UserID     uint64
	Name       string `json:"name" binding:"required"`
	ShortDesc  string `json:"short_desc" binding:"required"`
	LongDesc   string `json:"long_desc" binding:"required"`
	Perk       string `json:"perk" binding:"required"`
	GoalAmount uint64 `json:"goal_amount" binding:"required"`
}
