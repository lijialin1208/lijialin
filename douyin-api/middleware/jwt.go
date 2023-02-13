package middleware

import (
	"context"
	"douyin-api/tool"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

func JwtAuth() app.HandlerFunc {
	log.Println("jwt")
	return func(ctx context.Context, requestContext *app.RequestContext) {
		token := requestContext.PostForm("token")
		if token == "" {
			token = requestContext.Query("token")
			if token == "" {
				requestContext.Abort()
				return
			}
		}
		_, err := tool.ParseToken(token)
		if err != nil {
			log.Println("2")
			requestContext.Abort()
			return
		}
		requestContext.Next(ctx)
	}
}
