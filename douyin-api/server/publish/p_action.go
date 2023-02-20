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
	"strings"
)

func Paction(ctx context.Context, requestContext *app.RequestContext) {
	file, err := requestContext.FormFile("data")
	token := requestContext.PostForm("token")
	title := requestContext.PostForm("title")
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 0,
			"status_msg":  "fail",
		})
		return
	}
	play_url := file.Filename
	split := strings.Split(play_url, ".")
	path := "C:/Users/32259/GolandProjects/douyin/douyin-file/public/play_url/" + file.Filename
	err = requestContext.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 0,
			"status_msg":  "fail",
		})
		return
	}
	cover_url, err := tool.GetSnapshot(path, split[0], 1)
	log.Println(cover_url)
	dial, err := grpc.Dial("127.0.0.1:82", grpc.WithInsecure())
	if err != nil {
		log.Println("GPRC连接失败" + err.Error())
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 0,
			"status_msg":  "fail",
		})
		return
	}
	client := pb.NewPublishClient(dial)
	_, err = client.Publish(context.TODO(), &pb.PublishRequest{
		Token:    token,
		PlayUrl:  play_url,
		CoverUrl: cover_url,
		Title:    title,
	})
	if err != nil {
		log.Println(err)
		requestContext.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": 0,
			"status_msg":  "fail",
		})
		return
	}
	requestContext.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "OK",
	})
}
