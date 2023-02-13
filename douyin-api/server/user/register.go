package user

import (
	"context"
	"douyin-api/pb/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
)

func Register(ctx context.Context, requestContext *app.RequestContext) {

	username := requestContext.Query("username")
	password := requestContext.Query("password")
	dial, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if err != nil {
		log.Println("连接 gPRC 服务失败,", err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_id":     0,
			"token":       "",
		})
		return
	}
	defer dial.Close()
	client := user.NewRegisterClient(dial)
	response, err := client.Register(context.TODO(), &user.RegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_id":     0,
			"token":       "",
		})
		return
	}
	if response == nil {
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_id":     0,
			"token":       "",
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "string",
		"user_id":     response.ID,
		"token":       response.Token,
	})
}
