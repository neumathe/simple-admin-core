package mcmsnotice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/neumathe/simple-admin-core/api/internal/logic/mcmsnotice"
	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_notice/delete mcmsnotice DeleteMcmsNotice
//
// Delete mcms notice information | 删除McmsNotice信息
//
// Delete mcms notice information | 删除McmsNotice信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteMcmsNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsnotice.NewDeleteMcmsNoticeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMcmsNotice(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
