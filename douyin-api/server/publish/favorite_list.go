package publish

import (
	"context"
	"douyin-api/pb"
	"douyin-api/tool"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func FavoriteList(ctx context.Context, requestContext *app.RequestContext) {
	user_id := requestContext.Query("user_id")
	uid, _ := strconv.Atoi(user_id)
	token := requestContext.Query("token")
	userClaim, _ := tool.ParseToken(token)
	conn, err := grpc.Dial("127.0.0.1:82", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "fail",
			"video_list":  nil,
		})
		return
	}
	client := pb.NewFavoriteListClient(conn)
	response, err := client.FavoriteList(context.TODO(), &pb.FavoriteListRequest{UserId: int64(uid), MyId: int64(userClaim.ID)})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusOK, utils.H{
			"status_code": "1",
			"status_msg":  "fail",
			"video_list":  nil,
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": "0",
		"status_msg":  "ok",
		"video_list":  response.Vodies,
	})
}
