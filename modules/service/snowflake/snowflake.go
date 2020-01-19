package snowflake

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/galen-hlh/micro-sdk/go/proto/helper"
)

type SnowFlake struct{}

func (g *SnowFlake) GetDistributeId(ctx context.Context, req *helper.IdRequest, rsp *helper.IdResponse) error {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	// Generate a snowflake ID.
	id := node.Generate()

	rsp.Result = int64(id)
	return err
}
