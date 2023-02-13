package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
	"douyin-user/tool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type RegisterServer struct {
}

func NewRegisterServer() *RegisterServer {
	return &RegisterServer{}
}

func (r *RegisterServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	//获取参数
	username := request.Username
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//插入数据
	//随机nick
	rand.Seed(time.Now().Unix())
	nick := "dy" + strconv.Itoa(rand.Intn(90000000)+10000000)
	//插入（不唯一会报错）
	user := pojo.User{UserName: username, PassWord: string(password), NickName: nick}
	tx := db.DB.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//生成token
	token, err := tool.GetToken(user.ID, user.UserName, user.NickName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.RegisterResponse{ID: int64(user.ID), Token: token}, nil
}
