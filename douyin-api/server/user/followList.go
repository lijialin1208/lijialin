package user

import (
	"context"
	"douyin-api/pb/user"
	"douyin-api/tool"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func FollowList(ctx context.Context, requestContext *app.RequestContext) {
	token := requestContext.Query("token")
	user_id := requestContext.Query("user_id")
	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "string",
			"user_list":   nil,
		})
		return
	}
	conn, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "string",
			"user_list":   nil,
		})
		return
	}
	client := user.NewFollowListClient(conn)
	response, err := client.FollowList(context.TODO(), &user.FollowListRequest{
		UserId: user_id,
		Mid:    strconv.Itoa(userClaim.ID),
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "string",
			"status_msg":  "string",
			"user_list":   nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": "0",
		"status_msg":  "ok",
		"user_list":   response.UserList,
	})
}
