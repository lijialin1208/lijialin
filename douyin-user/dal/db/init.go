package db

import (
	"douyin-user/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&pojo.Message{})
	if err != nil {
		panic(err)
	}
}
