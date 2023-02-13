package test

import (
	"context"
	"douyin-publish/pb"
	"douyin-publish/server"
	"fmt"
	"testing"
)

func TestIssue(t *testing.T) {
	issueServer := server.NewIssueServer()
	response, _ := issueServer.Issue(context.TODO(), &pb.IssueRequest{
		Uid: 2,
		Mid: 1,
	})
	fmt.Println(response.Pulishs)
}
