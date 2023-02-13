package server

import (
	"context"
	"douyin-publish/dal/db"
	"douyin-publish/pb"
	"douyin-publish/pojo"
	"fmt"
	"github.com/spf13/viper"
)

type FavoriteList struct {
}

func NewFavoriteList() *FavoriteList {
	return &FavoriteList{}
}
func (f *FavoriteList) FavoriteList(c context.Context, in *pb.FavoriteListRequest) (*pb.FavoriteListResponse, error) {
	uid := in.UserId
	mid := in.MyId
	var publishes []int
	isFollow := false
	var count int64
	db.DB.Table("user_publish").Where("user_id = ?", uid).Pluck("publish_id", &publishes)
	db.DB.Table("follow_fans").Where("follow_id = ? AND fans_id = ?", uid, mid).Count(&count)
	if count == 1 {
		isFollow = !isFollow
	}
	vodies := make([]*pb.Vodie, len(publishes))
	for _, publishID := range publishes {
		publish := pojo.Publish{}
		var authorID int64
		author := pojo.User{}
		isFavorite := false
		db.DB.Where("id = ?", publishID).Find(&publish)
		db.DB.Table("user_issues").Where("publish_id = ?", publishID).Select("user_id").Scan(&authorID)
		db.DB.Table("users").Where("id = ?", authorID).Scan(&author)
		var cnt int64
		db.DB.Table("user_publish").Where("publish_id = ? AND user_id = ?", publishID, mid).Count(&cnt)
		if cnt == 1 {
			isFavorite = !isFavorite
		}
		vodies = append(vodies, &pb.Vodie{
			Id: int64(publish.ID),
			Author: &pb.Author{
				Id:            authorID,
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
		})
	}
	return &pb.FavoriteListResponse{Vodies: vodies}, nil
}
