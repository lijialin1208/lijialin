package server

import (
	"context"
	"douyin-user/dal/db"
	"douyin-user/pb"
	"douyin-user/pojo"
	"strconv"
	"time"
)

type MessageChatServer struct {
}

func NewMessageChatServer() *MessageChatServer {
	return &MessageChatServer{}
}

func (m *MessageChatServer) MessageChat(ctx context.Context, in *pb.MessageChatRequest) (*pb.MessageChatResponse, error) {
	mid, err := strconv.Atoi(in.Mid)
	uid, err := strconv.Atoi(in.Uid)
	pre_msg_time, err := strconv.Atoi(in.PreMsgTime)
	if err != nil {
		return &pb.MessageChatResponse{}, err
	}
	messages := make([]*pb.Message, 0)
	db_messages := make([]*pojo.Message, 0)

	pm_time := time.UnixMilli(int64(pre_msg_time))

	db.DB.Model(&pojo.Message{}).Where("(from_user_id = ? AND to_user_id = ? || from_user_id = ? AND to_user_id = ?) AND created_at > ?", mid, uid, uid, mid, pm_time).Select("id", "from_user_id", "to_user_id", "content", "created_at").Find(&db_messages)
	for i := 0; i < len(db_messages); i++ {
		message := db_messages[i]
		messages = append(messages, &pb.Message{
			ID:         strconv.Itoa(int(message.ID)),
			FromUserId: int64(message.FromUserId),
			ToUserId:   int64(message.ToUserId),
			Content:    message.Content,
			//CreateTime: message.CreatedAt.Format("03:04 PM"),
			CreateTime: strconv.FormatInt(message.CreatedAt.UnixMilli(), 10),
		})
	}
	times := strconv.FormatInt(time.Now().UnixMilli(), 10)
	return &pb.MessageChatResponse{MessageList: messages, PreMsgTime: times}, nil
}
