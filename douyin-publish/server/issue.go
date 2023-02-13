package server

import (
	"context"
	"douyin-publish/dal/db"
	"douyin-publish/pb"
	"douyin-publish/pojo"
	"fmt"
	"github.com/spf13/viper"
)

type Issue struct {
}

func NewIssueServer() *Issue {
	return &Issue{}
}

func (i *Issue) Issue(ctx context.Context, in *pb.IssueRequest) (*pb.IssueResponse, error) {
	result := make(map[string]interface{})
	isFollow := false
	var follow int64
	err := db.DB.Table("follow_fans").Where("follow_id = ? AND fans_id = ?", in.Uid, in.Mid).Scan(&result).Count(&follow).Error
	if err != nil {
		return nil, err
	}
	if follow == 1 {
		isFollow = !isFollow
	}
	var publishes []int
	author := pojo.User{}
	err = db.DB.Table("user_issues").Where("user_id = ?", in.Uid).Pluck("publish_id", &publishes).Error
	if err != nil {
		return nil, err
	}
	err = db.DB.Where("id = ?", in.Uid).Find(&author).Limit(1).Error
	if err != nil {
		return nil, err
	}
	isFavorite := false
	vodies := make([]*pb.Vodie, len(publishes))
	for index, publishID := range publishes {
		var count int64
		err := db.DB.Table("user_publish").Where("publish_id = ? AND user_id = ?", publishID, in.Mid).Count(&count).Error
		if err != nil {
			return nil, err
		}
		if count == 1 {
			isFavorite = !isFavorite
		}
		publish := pojo.Publish{}
		db.DB.Where("id = ?", publishID).Find(&publish).Limit(1)
		vodies[index] = &pb.Vodie{
			Id: int64(publishID),
			Author: &pb.Author{
				Id:            int64(author.ID),
				Name:          author.NickName,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       fmt.Sprint("http://", viper.GetString("server.ip"), ":", viper.GetString("server.port"), "/", publish.PlayUrl),
			CoverUrl:      publish.CoverUrl,
			FavoriteCount: int64(publish.NumberLikes),
			CommentCount:  int64(publish.NumberComments),
			IsFavorite:    isFavorite,
			Title:         publish.Title,
		}
	}
	return &pb.IssueResponse{Pulishs: vodies}, nil
}
