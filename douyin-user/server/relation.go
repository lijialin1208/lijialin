package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"gorm.io/gorm"
	"strconv"
)

type RelationServer struct {
}

func NewRelationServer() *RelationServer {
	return &RelationServer{}
}
func (r *RelationServer) Relation(ctx context.Context, in *pb.RelationRequest) (*pb.RelationResponse, error) {
	mid, err := strconv.Atoi(in.Mid)
	toid, err := strconv.Atoi(in.ToId)
	if err != nil {
		return &pb.RelationResponse{}, err
	}
	actionType := in.ActionType
	if actionType == "1" {
		db.DB.Table("follow_fans").Create(map[string]interface{}{
			"follow_id": toid,
			"fans_id":   mid,
		})
		db.DB.Table("users").Where("id = ?", mid).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		db.DB.Table("users").Where("id = ?", toid).Update("follower_count", gorm.Expr("follower_count + ?", 1))

		var count int64
		db.DB.Table("follow_fans").Where("follow_id = ? AND fans_id = ?", mid, toid).Count(&count)
		if count == 1 {
			db.DB.Table("user_friend").Create(map[string]interface{}{
				"user_id":   toid,
				"friend_id": mid,
			})
			db.DB.Table("user_friend").Create(map[string]interface{}{
				"user_id":   mid,
				"friend_id": toid,
			})
		}
	} else {
		db.DB.Table("users").Where("id = ?", mid).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		db.DB.Table("users").Where("id = ?", toid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
		db.DB.Table("follow_fans").Where("follow_id = ? AND fans_id = ?", toid, mid).Delete(map[string]interface{}{
			"follow_id": toid,
			"fans_id":   mid,
		})
		db.DB.Table("user_friend").Where("user_id = ? AND friend_id = ?", toid, mid).Delete(map[string]interface{}{
			"user_id":   toid,
			"friend_id": mid,
		})
		db.DB.Table("user_friend").Where("user_id = ? AND friend_id = ?", mid, toid).Delete(map[string]interface{}{
			"user_id":   mid,
			"friend_id": toid,
		})
	}
	return &pb.RelationResponse{}, nil
}
