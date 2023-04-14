package models

type User struct {
	BaseModel
	Username string `json:"username" form:"username" gorm:"not null" valid:"required~Username is required"`
	Email    string `json:"email" form:"email" gorm:"not null" valid:"required~Email is required, email~Email not valid"`
	Password string `json:"password" form:"password" gorm:"not null" valid:"required~Password is required"`
	Age      int    `json:"age" form:"age" gorm:"not null" valid:"required~Age is required, range(9|100)~Age must be greater than 8"`

	// Relations
	Photos       []Photo       `json:"photos" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	SocialsMedia []SocialMedia `json:"socials_media" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	Comments     []Comment     `json:"comments" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

func (u *User) TableName() string {
	return "tb_users"
}
