package models

import (
	"MyGram/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an User
type User struct {
	BaseModel
	Username string `json:"username" form:"username" gorm:"not null" valid:"required~Username is required"`
	Email    string `json:"email" form:"email" gorm:"not null" valid:"required~Email is required, email~Email not valid"`
	Password string `json:"password" form:"password" gorm:"not null" valid:"required~Password is required, stringlength(6|20)~Password must be at least 6 characters"`
	Age      int    `json:"age" form:"age" gorm:"not null" valid:"required~Age is required, range(9|100)~Age must be greater than 8"`

	// Relations
	Photos       []Photo       `json:"photos" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	SocialsMedia []SocialMedia `json:"socials_media" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	Comments     []Comment     `json:"comments" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

func (u *User) TableName() string {
	return "tb_users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}
