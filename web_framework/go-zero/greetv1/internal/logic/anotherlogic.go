package logic

import (
	"context"

	"greetv1/internal/svc"
	"greetv1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnotherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnotherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnotherLogic {
	return &AnotherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnotherLogic) Another(req *types.AnotherReq) (resp *types.AnotherRes, err error) {
	// todo: add your logic here and delete this line

	return
}
