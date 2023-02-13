package pojo

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `gorm:"unique;uniqueIndex"`
	PassWord      string
	NickName      string
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	Fans          []User `gorm:"many2many:follow_fans;foreignKey:ID;joinForeignKey:FansID;references:ID;joinReferences:FollowID"`
	Follows       []User `gorm:"many2many:follow_fans;foreignKey:ID;joinForeignKey:FollowID;references:ID;joinReferences:FansID"`
	Friend        []User `gorm:"many2many:user_friend"`
	UserID        uint
	LikePublish   []Publish `gorm:"many2many:user_publish;foreignKey:UserID;joinForeignKey:UserID;references:ID;joinReferences:PublishID"`
}
