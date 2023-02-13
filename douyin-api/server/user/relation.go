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

func Relation(ctx context.Context, requestContext *app.RequestContext) {
	token := requestContext.Query("token")
	action_type := requestContext.Query("action_type")
	to_user_id := requestContext.Query("to_user_id")
	conn, err := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "string",
		})
		return
	}

	userClaim, err := tool.ParseToken(token)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "string",
		})
		return
	}
	client := user.NewRelationClient(conn)
	_, err = client.Relation(context.TODO(), &user.RelationRequest{
		ActionType: action_type,
		ToId:       to_user_id,
		Mid:        strconv.Itoa(userClaim.ID),
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  "string",
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "OK",
	})
}
