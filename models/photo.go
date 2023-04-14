package models

type Photo struct {
	BaseModel
	Title    string `json:"title" form:"title" gorm:"not null" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" gorm:"not null" valid:"required~Photo url is required, url~Url photo not valid"`
	UserID   uint   `json:"user_id" gorm:"not null"`

	// Relations
	Comments []Comment `json:"comment_message" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

func (p *Photo) TableName() string {
	return "tb_photo"
}
