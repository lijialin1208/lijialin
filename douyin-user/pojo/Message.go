package pojo

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromUserId int
	ToUserId   int
	Content    string
}
