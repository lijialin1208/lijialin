package main

import (
	"douyin-publish/dal"
	"douyin-publish/pb"
	service "douyin-publish/server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	dal.Init()
	listen, err := net.Listen("tcp", "127.0.0.1:82")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	server := grpc.NewServer()
	//TODO
	pb.RegisterPublishServer(server, service.NewPublishServer())
	pb.RegisterFeedServer(server, service.NewFeedServer())
	pb.RegisterIssueServer(server, service.NewIssueServer())
	pb.RegisterLikeServer(server, service.NewLikeServer())
	pb.RegisterFavoriteListServer(server, service.NewFavoriteList())
	//TODO
	reflection.Register(server)

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
