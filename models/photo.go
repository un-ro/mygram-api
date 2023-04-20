package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for an Photo
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

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	return
}
