package server

import (
	"context"
	"douyin-publish/dal/db"
	"douyin-publish/pb"
	"douyin-publish/pojo"
	"douyin-publish/tool"
	"errors"
	"gorm.io/gorm"
)

type LikeServer struct {
}

func NewLikeServer() *LikeServer {
	return &LikeServer{}
}
func (l *LikeServer) Like(c context.Context, in *pb.LikeRequest) (*pb.LikeResponse, error) {
	t := in.Type
	token := in.Token
	videoId := in.VideoId
	userClaim, err := tool.ParseToken(token)
	userId := userClaim.ID
	if err != nil {
		return &pb.LikeResponse{}, err
	}
	if t == 1 {
		var count int64
		err := db.DB.Table("user_publish").Where("publish_id = ? AND user_id = ?", videoId, userId).Count(&count).Error
		if err != nil {
			return &pb.LikeResponse{}, err
		}
		if count == 1 {
			return &pb.LikeResponse{}, nil
		} else if count == 0 {
			m := map[string]interface{}{
				"publish_id": videoId,
				"user_id":    userId,
			}
			db.DB.Model(&pojo.Publish{}).Where("id = ?", videoId).UpdateColumn("number_likes", gorm.Expr("number_likes + ?", 1))
			db.DB.Table("user_publish").Create(&m)
			return &pb.LikeResponse{}, nil
		} else {
			return &pb.LikeResponse{}, errors.New("like is fail")
		}
	} else if t == 2 {
		var count int64
		err := db.DB.Table("user_publish").Where("publish_id = ? AND user_id = ?", videoId, userId).Count(&count).Error
		if err != nil {
			return &pb.LikeResponse{}, err
		}
		if count == 0 {
			return &pb.LikeResponse{}, nil
		} else if count == 1 {
			m := map[string]interface{}{
				"publish_id": videoId,
				"user_id":    userId,
			}
			db.DB.Model(&pojo.Publish{}).Where("id = ?", videoId).UpdateColumn("number_likes", gorm.Expr("number_likes - ?", 1))
			db.DB.Table("user_publish").Where("publish_id = ? AND user_id = ?", videoId, userId).Delete(&m)
			return &pb.LikeResponse{}, nil
		} else {
			return &pb.LikeResponse{}, errors.New("like is fail")
		}
	} else {
		return &pb.LikeResponse{}, nil
	}
}
