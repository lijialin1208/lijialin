package pojo

import (
	"gorm.io/gorm"
)

type Publish struct {
	gorm.Model
	Title          string
	NumberLikes    int
	NumberComments int
	PlayUrl        string
	CoverUrl       string
	UserID         uint
	Comments       []Comment `gorm:"foreignKey:PublishID;references:ID"`
	LikeUser       []User    `gorm:"many2many:user_publish;"`
}
