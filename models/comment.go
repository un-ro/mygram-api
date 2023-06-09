package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an Comment
type Comment struct {
	BaseModel
	UserID  uint   `json:"user_id" gorm:"not null"`
	PhotoID uint   `json:"photo_id" gorm:"not null"`
	Message string `json:"comment_message" form:"comment_message" gorm:"not null" valid:"required~Comment message is required"`
}

func (c *Comment) TableName() string {
	return "tb_comments"
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return
	}

	return
}
