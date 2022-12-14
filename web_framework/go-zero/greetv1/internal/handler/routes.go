// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"greetv1/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: Greetv1Handler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/another/:name",
				Handler: AnotherHandler(serverCtx),
			},
		},
	)
}
