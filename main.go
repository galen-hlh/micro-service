package main

import (
	"fmt"
	"github.com/galen-hlh/micro-sdk/go/proto/helper"
	"github.com/micro/go-micro"
	"mirco-service/modules/service/snowflake"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("helper"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	_ = helper.RegisterHelperHandler(service.Server(), new(snowflake.SnowFlake))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
