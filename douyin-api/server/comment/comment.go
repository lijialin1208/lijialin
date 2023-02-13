package comment

import (
	"context"
	"douyin-api/pb/comment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
)

func Comment(ctx context.Context, requestContext *app.RequestContext) {
	token := requestContext.Query("token")               //用户鉴权token
	video_id := requestContext.Query("video_id")         //视频id
	action_type := requestContext.Query("action_type")   //1-发布评论，2-删除评论
	comment_id := requestContext.Query("comment_id")     //要删除的评论id，在action_type=2的时候使用
	comment_text := requestContext.Query("comment_text") //用户填写的评论内容，在action_type=1的时候使用
	conn, err := grpc.Dial("127.0.0.1:83", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"comment":     nil,
		})
		return
	}
	client := comment.NewCommentClient(conn)
	response, err := client.Comment(context.TODO(), &comment.CommentRequest{
		Token:       token,
		VideoId:     video_id,
		ActionType:  action_type,
		CommentText: comment_text,
		CommentId:   comment_id,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"comment":     nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "ok",
		"comment": utils.H{
			"id": response.Id,
			"user": utils.H{
				"id":             response.User.Id,
				"name":           response.User.Name,
				"follow_count":   response.User.FollowCount,
				"follower_count": response.User.FollowerCount,
				"is_follow":      response.User.IsFollow,
			},
			"content":     response.Content,
			"create_date": response.CreateDate,
		},
	})
}
