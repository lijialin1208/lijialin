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

func Follower(ctx context.Context, requestContext *app.RequestContext) {
	user_id := requestContext.Query("user_id")
	token := requestContext.Query("token")
	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_list":   nil,
		})
		return
	}
	dial, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_list":   nil,
		})
		return
	}
	client := user.NewFollowerClient(dial)
	response, err := client.Follower(context.TODO(), &user.FollowerRequest{
		UserId: user_id,
		Mid:    strconv.Itoa(userClaim.ID),
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_list":   nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "OK",
		"user_list":   response.Followers,
	})
}
