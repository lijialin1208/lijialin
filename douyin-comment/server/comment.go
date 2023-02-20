package server

import (
	"context"
	"douyin-comment/dal/db"
	"douyin-comment/pb"
	"douyin-comment/pojo"
	"douyin-comment/tool"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type CommentServer struct {
}

func NewCommentServer() *CommentServer {
	return &CommentServer{}
}
func (c *CommentServer) Comment(ctx context.Context, in *pb.CommentRequest) (*pb.CommentResponse, error) {
	actionType := in.ActionType
	vid, _ := strconv.Atoi(in.VideoId)
	token := in.Token
	userClaim, _ := tool.ParseToken(token)
	id := userClaim.ID
	if actionType == "1" {
		commentText := in.CommentText
		user := pojo.User{}
		p := &pojo.Comment{
			Content:   commentText,
			PublishID: uint(vid),
			UserID:    uint(id),
		}
		db.DB.Create(&p)
		db.DB.Table("publishes").Where("id = ?", vid).UpdateColumn("number_comments", gorm.Expr("number_comments + ?", 1))
		db.DB.Where("id = ?", id).Find(&user)
		createdAt := p.CreatedAt.Format("01-02")
		return &pb.CommentResponse{
			Id: int64(p.ID),
			User: &pb.Author{
				Id:            int64(user.UserID),
				Name:          user.NickName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
			Content:    p.Content,
			CreateDate: createdAt,
		}, nil
	} else if actionType == "2" {
		db.DB.Table("comments").Where("id = ?", in.CommentId).Delete(&pojo.Comment{})
		db.DB.Table("publishes").Where("id = ?", vid).UpdateColumn("number_comments", gorm.Expr("number_comments - ?", 1))
		return &pb.CommentResponse{}, nil
	} else {
		return &pb.CommentResponse{}, errors.New("type fail")
	}
}
