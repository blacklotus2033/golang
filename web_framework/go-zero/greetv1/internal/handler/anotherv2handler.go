package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"greetv1/internal/logic"
	"greetv1/internal/svc"
	"greetv1/internal/types"
)

func AnotherV2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnotherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAnotherV2Logic(r.Context(), svcCtx)
		resp, err := l.AnotherV2(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
