package pojo

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content         string
	ParentCommentID uint `gorm:"default:0"`
	PublishID       uint
	UserID          uint
	Comments        []Comment `gorm:"foreignKey:ParentCommentID;references:ID"`
}
