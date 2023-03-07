package script

import (
	"github.com/gayxy/go-ooey/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/gayxy/go-ooey/app/ooey/internal/logic/script"
	"github.com/gayxy/go-ooey/app/ooey/internal/svc"
	"github.com/gayxy/go-ooey/app/ooey/internal/types"
)

func GetScriptInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetScriptInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := script.NewGetScriptInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetScriptInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
