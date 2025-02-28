package pdftasklog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/neumathe/simple-admin-core/api/internal/logic/pdftasklog"
	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"
)

// swagger:route post /pdflog pdftasklog GetPdfTaskInfo
//
// Create task information | 创建定时任务
//
// Create task information | 创建定时任务
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDReq
//
// Responses:
//  200: BaseMsgResp

func GetPdfTaskInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pdftasklog.NewGetPdfTaskInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetPdfTaskInfo(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
