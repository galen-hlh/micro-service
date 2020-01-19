package tests

import (
	"context"
	"fmt"
	"github.com/galen-hlh/micro-sdk/go/proto/helper"
	"mirco-service/modules/service/snowflake"
	"testing"
)

func TestGetDistributeId(t *testing.T) {
	var ctx context.Context
	snow := snowflake.SnowFlake{}
	request := &helper.IdRequest{}
	response := &helper.IdResponse{}
	err := snow.GetDistributeId(ctx, request, response)
	if err != nil {
		t.Errorf("error:%s", err.Error())
	}
	fmt.Println(response.Result)
}
