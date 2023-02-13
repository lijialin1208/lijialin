package main

import (
	"douyin-comment/dal"
	"douyin-comment/pb"
	service "douyin-comment/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	dal.Init()
	listen, err := net.Listen("tcp", "127.0.0.1:83")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCommentServer(server, service.NewCommentServer())
	pb.RegisterCommentListServer(server, service.NewCommentListServer())
	reflection.Register(server)

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
