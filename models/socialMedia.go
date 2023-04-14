package models

type SocialMedia struct {
	BaseModel
	Name           string `json:"name" form:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" gorm:"not null" valid:"required~Social media url is required, url~Url social media not valid"`
	UserID         uint   `json:"user_id" gorm:"not null"`
}

func (s *SocialMedia) TableName() string {
	return "tb_social_media"
}
