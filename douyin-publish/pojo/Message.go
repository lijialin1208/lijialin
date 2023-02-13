package pojo

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID  int
	Content string
}
