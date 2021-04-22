package logic

import (
	"context"

	"bookstore/rpc/transform/example.com/project/protos/fizz"
	"bookstore/rpc/transform/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *fizz.ExpandReq) (*fizz.ExpandResp, error) {
	// todo: add your logic here and delete this line

	return &fizz.ExpandResp{}, nil
}
