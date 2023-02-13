package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
)

type FriendListServer struct {
}

func NewFriendListServer() *FriendListServer {
	return &FriendListServer{}
}
func (f *FriendListServer) FriendList(ctx context.Context, in *pb.FriendListRequest) (*pb.FriendListResponse, error) {
	userId := in.UserId
	usersID := make([]int, 0)
	db.DB.Table("user_friend").Where("user_id = ?", userId).Pluck("friend_id", &usersID)
	user_list := make([]*pb.Author, 0)
	users := make([]*pojo.User, 0)
	db.DB.Table("users").Where("id in (?)", usersID).Scan(&users)
	for index := 0; index < len(users); index++ {
		user := users[index]
		user_list = append(user_list, &pb.Author{
			Id:            int64(user.ID),
			Name:          user.NickName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		})
	}
	return &pb.FriendListResponse{UserList: user_list}, nil
}
