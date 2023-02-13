package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
	"douyin-user/tool"
	"log"
)

type GetMessageServer struct {
}

func NewGetMessageServer() *GetMessageServer {
	return &GetMessageServer{}
}

func (g *GetMessageServer) UserInfo(ctx context.Context, request *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	id := request.UserId
	userClaim, err := tool.ParseToken(request.Token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	myid := userClaim.ID

	user := pojo.User{}
	err = db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//根据myid和id判断是否关注
	result := 0
	var isFollow = false
	db.DB.Raw("SELECT COUNT(*) FROM follow_fans WHERE follow_id = ? AND fans_id = ?", id, myid).Scan(&result)
	if result == 1 {
		isFollow = true
	}
	return &pb.UserInfoResponse{
		Nick:          user.NickName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		ID:            int64(user.ID),
		IsFollow:      isFollow,
	}, nil
}
