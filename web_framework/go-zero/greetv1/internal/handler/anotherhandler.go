package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"greetv1/internal/logic"
	"greetv1/internal/svc"
	"greetv1/internal/types"
)

func AnotherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnotherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAnotherLogic(r.Context(), svcCtx)
		resp, err := l.Another(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
