package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"mirco-service/proto/helper"
)

func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("helper.client"))
	service.Init()

	// 创建新的客户端
	client := helper.NewHelperService("helper", service.Client())

	// 调用greeter
	rsp, err := client.GetDistributeId(context.TODO(), &helper.IdRequest{})
	if err != nil {
		fmt.Println(err)
	}

	// 打印响应请求
	fmt.Println(rsp.Result)
}
