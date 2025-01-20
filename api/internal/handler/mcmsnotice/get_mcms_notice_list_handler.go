package mcmsnotice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/neumathe/simple-admin-core/api/internal/logic/mcmsnotice"
	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_notice/list mcmsnotice GetMcmsNoticeList
//
// Get mcms notice list | 获取McmsNotice信息列表
//
// Get mcms notice list | 获取McmsNotice信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: McmsNoticeListReq
//
// Responses:
//  200: McmsNoticeListResp

func GetMcmsNoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.McmsNoticeListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsnotice.NewGetMcmsNoticeListLogic(r.Context(), svcCtx)
		resp, err := l.GetMcmsNoticeList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
