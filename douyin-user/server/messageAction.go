package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
	"strconv"
)

type MessageActionServer struct {
}

func NewMessageActionServer() *MessageActionServer {
	return &MessageActionServer{}
}
func (m *MessageActionServer) MessageAction(ctx context.Context, in *pb.MessageActionRequest) (*pb.MessageActionResponse, error) {
	uid, err := strconv.Atoi(in.Uid)
	mid, err := strconv.Atoi(in.Mid)
	if err != nil {
		return &pb.MessageActionResponse{}, err
	}
	content := in.Content
	message := &pojo.Message{
		FromUserId: mid,
		ToUserId:   uid,
		Content:    content,
	}
	db.DB.Create(message)
	return &pb.MessageActionResponse{}, nil
}
