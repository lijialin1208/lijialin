package comment

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"testing"
)

func TestComment(t *testing.T) {
	//评论操作
	t.Run("评论", func(t *testing.T) {
		router := route.NewEngine(config.NewOptions([]config.Option{}))
		router.POST("/douyin/comment/action/", Comment)
		w := ut.PerformRequest(router, "POST", "/douyin/comment/action?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlck5hbWUiOiJsaWppYWxpbiIsIk5pY2tOYW1lIjoiZHk5NjUwMzQ3NyIsImlzcyI6InRlc3QiLCJzdWIiOiJzb21lYm9keSIsImF1ZCI6WyJzb21lYm9keV9lbHNlIl0sImV4cCI6MTY3NzYzNDYxOCwibmJmIjoxNjc2MzM4NjE4LCJpYXQiOjE2NzYzMzg2MTgsImp0aSI6IjEifQ.hKHiiIzLnHsUIvpAGvs4m1K9JM_kGkVOTxAq8Cacz_I&video_id=1&action_type=1&comment_text='哈哈哈'", &ut.Body{}, ut.Header{})
		resp := w.Result()
		assert.DeepEqual(t, 200, resp.StatusCode())
	})
	//删除评论
	t.Run("删除评论", func(t *testing.T) {
		target := "/douyin/comment/action/" +
			"?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlck5hbWUiOiJsaWppYWxpbiIsIk5pY2tOYW1lIjoiZHk5NjUwMzQ3NyIsImlzcyI6InRlc3QiLCJzdWIiOiJzb21lYm9keSIsImF1ZCI6WyJzb21lYm9keV9lbHNlIl0sImV4cCI6MTY3NzYzNDYxOCwibmJmIjoxNjc2MzM4NjE4LCJpYXQiOjE2NzYzMzg2MTgsImp0aSI6IjEifQ.hKHiiIzLnHsUIvpAGvs4m1K9JM_kGkVOTxAq8Cacz_I&video_id=1&action_type=2&comment_text=&comment_id=2"
		h := server.Default()
		h.POST("/douyin/comment/action/", Comment)
		w := ut.PerformRequest(h.Engine, "GET", target, &ut.Body{}, ut.Header{})
		resp := w.Result()
		assert.DeepEqual(t, 200, resp.StatusCode())
	})
}

func TestCommentList(t *testing.T) {

}

func TestCom(t *testing.T) {

}
