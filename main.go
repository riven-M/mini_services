package main

import (
	"mini_services/proto"
	"mini_services/services"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	//启动grpc服务
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8992"),
	)

	proto.RegisterGreeterHandler(service.Server(), new(services.GreeterSer))

	if err := service.Run(); err != nil {
		log.Infof("启动异常:%s", err.Error())
	}
}
