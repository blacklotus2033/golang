package logic

import (
	"context"

	"greetv1/internal/svc"
	"greetv1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnotherV2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnotherV2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *AnotherV2Logic {
	return &AnotherV2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnotherV2Logic) AnotherV2(req *types.AnotherReq) (resp *types.AnotherRes, err error) {
	// todo: add your logic here and delete this line

	return
}
