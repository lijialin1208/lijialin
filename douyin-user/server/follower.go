package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
)

type FollowerServer struct {
}

func NewFollowerServer() *FollowerServer {
	return &FollowerServer{}
}

func (f *FollowerServer) Follower(ctx context.Context, in *pb.FollowerRequest) (*pb.FollowerResponse, error) {
	userId := in.UserId
	mid := in.Mid
	fans := make([]int, 0)
	rfans := make([]*pb.Author, 0)
	db.DB.Table("follow_fans").Where("follow_id = ?", userId).Pluck("fans_id", &fans)
	if len(fans) == 0 {
		return &pb.FollowerResponse{Followers: rfans}, nil
	}
	for index := 0; index < len(fans); index++ {
		fanID := fans[index]
		user := pojo.User{}
		isFollow := false
		var cnt int64
		db.DB.Table("follow_fans").Where("fans_id = ? AND follow_id = ?", mid, fanID).Count(&cnt)
		if cnt == 1 {
			isFollow = !isFollow
		}
		db.DB.Table("users").Where("id = ?", fanID).Scan(&user)
		rfans = append(rfans, &pb.Author{
			Id:            int64(user.ID),
			Name:          user.NickName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		})
	}
	return &pb.FollowerResponse{Followers: rfans}, nil
}
