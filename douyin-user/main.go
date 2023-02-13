package main

import (
	"douyin-user/dal"
	"douyin-user/pb"
	service "douyin-user/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	dal.Init()
	listener, err := net.Listen("tcp", "127.0.0.1:81")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterLoginServer(server, service.NewLoginServer())
	pb.RegisterRegisterServer(server, service.NewRegisterServer())
	pb.RegisterUserInfoServer(server, service.NewGetMessageServer())
	pb.RegisterRelationServer(server, service.NewRelationServer())
	pb.RegisterFollowListServer(server, service.NewFollowListServer())
	pb.RegisterFriendListServer(server, service.NewFriendListServer())
	pb.RegisterFollowerServer(server, service.NewFollowerServer())
	pb.RegisterMessageActionServer(server, service.NewMessageActionServer())
	pb.RegisterMessageChatServer(server, service.NewMessageChatServer())
	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
