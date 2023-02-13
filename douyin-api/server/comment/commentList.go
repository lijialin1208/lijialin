package comment

import (
	"context"
	pb "douyin-api/pb/comment"
	"douyin-api/tool"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func CommentList(ctx context.Context, requestContext *app.RequestContext) {
	token := requestContext.Query("token")
	video_id := requestContext.Query("video_id")
	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code":  1,
			"status_msg":   "string",
			"comment_list": nil,
		})
		return
	}
	conn, err := grpc.Dial("127.0.0.1:83", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code":  1,
			"status_msg":   "string",
			"comment_list": nil,
		})
		return
	}
	commentListClient := pb.NewCommentListClient(conn)
	response, err := commentListClient.CommentList(context.TODO(), &pb.CommentListRequest{
		Mid: strconv.Itoa(userClaim.ID),
		Vid: video_id,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code":  1,
			"status_msg":   "string",
			"comment_list": nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code":  0,
		"status_msg":   "ok",
		"comment_list": response.CommentList,
	})
}
