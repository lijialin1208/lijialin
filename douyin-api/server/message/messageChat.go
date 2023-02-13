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

func MessageChat(ctx context.Context, requestContext *app.RequestContext) {
	token := requestContext.Query("token")
	to_user_id := requestContext.Query("to_user_id")
	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code":  "1",
			"message_list": nil,
		})
		return
	}
	dial, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code":  "1",
			"message_list": nil,
		})
		return
	}
	client := user.NewMessageChatClient(dial)
	response, err := client.MessageChat(context.TODO(), &user.MessageChatRequest{
		Uid: to_user_id,
		Mid: strconv.Itoa(userClaim.ID),
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code":  "1",
			"message_list": nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code":  0,
		"message_list": response.MessageList,
	})
}
