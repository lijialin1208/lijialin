package message

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

func MessageAction(ctx context.Context, requestContext *app.RequestContext) {
	log.Println("MessageAction")
	token := requestContext.Query("token")
	to_user_id := requestContext.Query("to_user_id")
	content := requestContext.Query("content")
	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "fail",
		})
		return
	}
	dial, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	defer dial.Close()
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "fail",
		})
		return
	}
	client := user.NewMessageActionClient(dial)
	_, err = client.MessageAction(context.TODO(), &user.MessageActionRequest{
		Uid:     to_user_id,
		Mid:     strconv.Itoa(userClaim.ID),
		Content: content,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "fail",
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "OK",
	})
}
