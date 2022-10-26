package logic

import (
	"context"

	"greetv1/internal/svc"
	"greetv1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Greetv1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetv1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Greetv1Logic {
	return &Greetv1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Greetv1Logic) Greetv1(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
