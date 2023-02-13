package server

import (
	"context"
	"douyin-comment/dal/db"
	"douyin-comment/pb"
	"douyin-comment/pojo"
)

type CommentListServer struct {
}

func NewCommentListServer() *CommentListServer {
	return &CommentListServer{}
}
func (c *CommentListServer) CommentList(ctx context.Context, in *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	mid := in.Mid
	vid := in.Vid
	var comments []pojo.Comment
	db.DB.Where("publish_id = ?", vid).Find(&comments)
	if len(comments) == 0 {
		return &pb.CommentListResponse{}, nil
	}
	commentList := make([]*pb.CommentResponse, len(comments))
	for index, comment := range comments {
		user := pojo.User{}
		var count int64
		isFollow := false
		db.DB.Table("users").Where("id = ?", comment.UserID).Scan(&user)
		db.DB.Table("follow_fans").Where("follow_id = ? AND fans_id = ?", user.ID, mid).Count(&count)
		if count == 1 {
			isFollow = !isFollow
		}
		commentList[index] = &pb.CommentResponse{
			Id: int64(comment.ID),
			User: &pb.Author{
				Id:            int64(user.ID),
				Name:          user.NickName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      isFollow,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		}
	}
	return &pb.CommentListResponse{CommentList: commentList}, nil
}
