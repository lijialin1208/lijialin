package publish

import (
	"context"
	"douyin-api/pb"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func Like(ctx context.Context, requestContext *app.RequestContext) {
	token := requestContext.Query("token")
	vid := requestContext.Query("video_id")
	video_id, err2 := strconv.Atoi(vid)
	if err2 != nil {
		log.Println(err2)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
		})
		return
	}
	ty := requestContext.Query("action_type")
	action_type, err2 := strconv.Atoi(ty)
	if err2 != nil {
		log.Println(err2)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
		})
		return
	}
	conn, err := grpc.Dial("127.0.0.1:82", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
		})
		return
	}
	defer conn.Close()
	client := pb.NewLikeClient(conn)
	_, err = client.Like(context.TODO(), &pb.LikeRequest{
		VideoId: int64(video_id),
		Type:    int64(action_type),
		Token:   token,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "like ok",
	})
}
