package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"mirco-service/proto/helper"
)

type Helper struct{}

func (g *Helper) GetDistributeId(ctx context.Context, req *helper.IdRequest, rsp *helper.IdResponse) error {
	fmt.Println(1)
	rsp.Result = 12345
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("helper"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	_ = helper.RegisterHelperHandler(service.Server(), new(Helper))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
