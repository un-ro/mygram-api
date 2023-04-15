package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	BaseModel
	Name           string `json:"name" form:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" gorm:"not null" valid:"required~Social media url is required, url~Url social media not valid"`
	UserID         uint   `json:"user_id" gorm:"not null"`
}

func (s *SocialMedia) TableName() string {
	return "tb_social_media"
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return
	}

	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return
	}

	return
}
