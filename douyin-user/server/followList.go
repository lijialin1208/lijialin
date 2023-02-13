package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
)

type FollowListServer struct {
}

func NewFollowListServer() *FollowListServer {
	return &FollowListServer{}
}
func (f *FollowListServer) FollowList(ctx context.Context, in *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	userId := in.UserId
	var result []int

	db.DB.Table("follow_fans").Where("fans_id = ?", userId).Pluck("follow_id", &result)
	users := make([]*pb.Author, len(result))
	for i := 0; i < len(result); i++ {
		user := pojo.User{}
		db.DB.Table("users").Where("id = ?", result[i]).Scan(&user)
		users[i] = &pb.Author{
			Id:            int64(user.ID),
			Name:          user.NickName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
	}
	return &pb.FollowListResponse{UserList: users}, nil
}
