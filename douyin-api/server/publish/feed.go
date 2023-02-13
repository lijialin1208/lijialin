package publish

import (
	"context"
	"douyin-api/pb"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"log"
)

type FeedResponse struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	NextTime   int         `json:"next_time"`
	VideoList  []*pb.Vodie `json:"video_list"`
}

func Feed(ctx context.Context, requestContext *app.RequestContext) {
	latest_time := requestContext.Query("latest_time")
	//atoi, _ := strconv.Atoi(latest_time)
	//log.Println(time.UnixMicro(int64(atoi)))
	//log.Println(time.UnixMilli(int64(atoi)))
	token := requestContext.Query("token")
	feedRequest := pb.FeedRequest{
		LatestTime: "",
		Token:      "",
	}
	if latest_time != "" {
		feedRequest.LatestTime = latest_time
	}
	if token != "" {
		feedRequest.Token = token
	}
	conn, err := grpc.Dial("127.0.0.1:82", grpc.WithInsecure())
	if err != nil {
		log.Println("Gprc连接出错" + err.Error())
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"next_time":   nil,
			"video_list":  nil,
		})
		return
	}
	client := pb.NewFeedClient(conn)
	response, err := client.Feed(context.TODO(), &feedRequest)
	if err != nil {
		log.Println(err.Error())
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 1,
			"status_msg":  "fail",
			"next_time":   nil,
			"video_list":  nil,
		})
		return
	}
	log.Println(response.Pulishs)
	requestContext.JSON(consts.StatusOK, FeedResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		NextTime:   int(response.NextTime),
		VideoList:  response.Pulishs,
	})
}
