package pdftasklog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/neumathe/simple-admin-core/api/internal/logic/pdftasklog"
	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"
)

// swagger:route post /pdflog/list pdftasklog GetPdfTaskList
//
// Get task list | 获取定时任务列表
//
// Get task list | 获取定时任务列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PdfTaskListReq
//
// Responses:
//  200: PdfTaskInfoListResp

func GetPdfTaskListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PdfTaskListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pdftasklog.NewGetPdfTaskListLogic(r.Context(), svcCtx)
		resp, err := l.GetPdfTaskList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
