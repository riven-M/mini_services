package services

import (
	"context"
	"fmt"
	"mini_services/proto"

	"github.com/micro/go-micro/v2/util/log"
)

type GreeterSer struct{}

func (this GreeterSer) GetResult(ctx context.Context, Allrequest *proto.Request, out *proto.Response) error {

	fmt.Println("---GetResult----")
	log.Info("requests:page:%d,size:%d", Allrequest.Page, Allrequest.Size)

	out.Code = 200
	out.Content = "it's ok"
	return nil
}
