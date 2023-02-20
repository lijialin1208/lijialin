package server

import (
	"context"
	"douyin-publish/dal/db"
	"douyin-publish/pb"
	"douyin-publish/pojo"
	"douyin-publish/tool"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type FeedServer struct {
}

func NewFeedServer() *FeedServer {
	return &FeedServer{}
}
func (f *FeedServer) Feed(ctx context.Context, in *pb.FeedRequest) (*pb.FeedResponse, error) {
	publishes := make([]pojo.Publish, 0)
	if in.LatestTime == "" {
		db.DB.Find(&publishes).Order("created_at DESC").Limit(10)
	} else {
		latestTime, _ := strconv.Atoi(in.LatestTime)

		format := time.UnixMilli(int64(latestTime))
		db.DB.Where("created_at < ?", format).Find(&publishes).Order("created_at DESC").Limit(10)
	}
	if len(publishes) == 0 {
		return &pb.FeedResponse{
			Pulishs:  nil,
			NextTime: 0,
		}, nil
	}
	vodies := make([]*pb.Vodie, len(publishes))
	if in.Token != "" {
		parseToken, err := tool.ParseToken(in.Token)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		UserID := parseToken.ID
		for i, publish := range publishes {
			isFollow := false
			var count1 int64
			isFavorite := false
			var up int64
			db.DB.Table("follow_fans").Where("follow_id = ? AND fans_id = ?", publish.UserID, UserID).Count(&count1)
			if count1 == 1 {
				isFollow = true
			}
			db.DB.Table("user_publish").Where("publish_id = ? AND user_id = ?", publish.ID, UserID).Count(&up)
			if up == 1 {
				isFavorite = true
			}
			u := pojo.User{}
			db.DB.Where("id = ?", publish.UserID).Find(&u)
			u.FollowCount = 10
			vodies[i] = &pb.Vodie{
				Id: int64(publish.ID),
				Author: &pb.Author{
					Id:            int64(u.ID),
					Name:          u.NickName,
					FollowCount:   u.FollowCount,
					FollowerCount: u.FollowerCount,
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
		return &pb.FeedResponse{
			Pulishs:  vodies,
			NextTime: publishes[len(publishes)-1].CreatedAt.Unix(),
		}, nil
	} else {
		for i, publish := range publishes {
			u := pojo.User{
				Model:         gorm.Model{},
				NickName:      "",
				FollowCount:   0,
				FollowerCount: 0,
			}
			db.DB.Where("id = ?", publish.UserID).Find(&u)
			vodies[i] = &pb.Vodie{
				Id: int64(publish.ID),
				Author: &pb.Author{
					Id:            int64(u.ID),
					Name:          u.NickName,
					FollowCount:   u.FollowCount,
					FollowerCount: u.FollowerCount,
					IsFollow:      false,
				},
				PlayUrl:       fmt.Sprint("http://", viper.GetString("server.ip"), ":", viper.GetString("server.port"), "/", publish.PlayUrl),
				CoverUrl:      publish.CoverUrl,
				FavoriteCount: int64(publish.NumberLikes),
				CommentCount:  int64(publish.NumberComments),
				IsFavorite:    false,
				Title:         publish.Title,
			}
		}
		return &pb.FeedResponse{
			Pulishs:  vodies,
			NextTime: publishes[len(publishes)-1].CreatedAt.Unix(),
		}, nil
	}

}
