package user

import (
	"context"
	"douyin-api/pb/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func GetMessage(ctx context.Context, requestContext *app.RequestContext) {

	user_id, err := strconv.Atoi(requestContext.Query("user_id"))
	token := requestContext.Query("token")
	dial, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if err != nil {
		log.Println("连接 gPRC 服务失败,", err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 0,
			"status_msg":  "fail",
			"user":        nil,
		})
		return
	}
	defer dial.Close()
	client := user.NewUserInfoClient(dial)
	response, err := client.UserInfo(context.TODO(), &user.UserInfoRequest{
		UserId: int64(user_id),
		Token:  token,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user":        nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "OK",
		"user": utils.H{
			"id":             response.ID,
			"name":           response.Nick,
			"follow_count":   response.FollowCount,
			"follower_count": response.FollowerCount,
			"is_follow":      response.IsFollow,
		},
	})
}
