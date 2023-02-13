package server

import (
	"context"
	"douyin-publish/dal/db"
	"douyin-publish/pb"
	"douyin-publish/pojo"
	"douyin-publish/tool"
	"log"
)

type PublishServer struct {
}

func NewPublishServer() *PublishServer {
	return &PublishServer{}
}
func (p *PublishServer) Publish(ctx context.Context, in *pb.PublishRequest) (*pb.PublishResponse, error) {
	pulish := &pojo.Publish{
		Title:          in.Title,
		NumberLikes:    0,
		NumberComments: 0,
		PlayUrl:        in.PlayUrl,
		CoverUrl:       in.CoverUrl,
	}
	token := in.Token
	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	u := pojo.User{}
	err = db.DB.Where("id = ?", userClaim.ID).First(&u).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pulish.UserID = uint(userClaim.ID)
	err = db.DB.Create(&pulish).Error
	db.DB.Table("user_issues").Create(map[string]interface{}{
		"user_id": userClaim.ID, "publish_id": pulish.ID,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.PublishResponse{}, nil
}
