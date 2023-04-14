package models

type Comment struct {
	BaseModel
	UserID  uint   `json:"user_id" gorm:"not null"`
	PhotoID uint   `json:"photo_id" gorm:"not null"`
	Message string `json:"comment_message" form:"comment_message" gorm:"not null" valid:"required~Comment message is required"`
}

func (c *Comment) TableName() string {
	return "tb_comments"
}
