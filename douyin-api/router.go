package main

import (
	"douyin-api/middleware"
	"douyin-api/server/comment"
	"douyin-api/server/message"
	"douyin-api/server/publish"
	"douyin-api/server/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter(hertz *server.Hertz) {
	hertz.Static("/static", "./play_url")
	//基础接口
	hertz.POST("/douyin/user/register/", user.Register)
	hertz.POST("/douyin/user/login/", user.Login)
	hertz.GET("/douyin/user/", user.GetMessage)
	hertz.GET("/douyin/feed/", publish.Feed)

	douyin := hertz.Group("/douyin").Use(middleware.JwtAuth())
	douyin.POST("/publish/action/", publish.Paction)
	douyin.GET("/publish/list/", publish.UPublish)
	//互动接口
	douyin.POST("/favorite/action/", publish.Like)
	douyin.POST("/comment/action/", comment.Comment)
	douyin.GET("/favorite/list/", publish.FavoriteList)
	douyin.GET("/comment/list/", comment.CommentList)
	//社交接口
	douyin.POST("/relation/action/", user.Relation)
	douyin.POST("/message/action/", message.MessageAction)
	douyin.GET("/relation/follow/list/", user.FollowList)
	douyin.GET("/relation/follower/list/", user.Follower)
	douyin.GET("/relation/friend/list/", user.Friend)
	douyin.GET("/message/chat/", message.MessageChat)
}
