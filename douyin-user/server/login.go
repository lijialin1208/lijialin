package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
	"douyin-user/tool"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type LoginServer struct {
}

func NewLoginServer() *LoginServer {
	return &LoginServer{}
}
func (l *LoginServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := request.Username
	password := request.Password
	user := pojo.User{}
	var count int64
	err := db.DB.Where("user_name = ?", username).Find(&user).Count(&count).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	token, err := tool.GetToken(user.ID, user.UserName, user.NickName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.LoginResponse{
		ID:    int64(user.ID),
		Token: token,
	}, nil
}
