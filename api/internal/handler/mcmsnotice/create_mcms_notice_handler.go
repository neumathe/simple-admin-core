package mcmsnotice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/mcmsnotice"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_notice/create mcmsnotice CreateMcmsNotice
//
// Create mcms notice information | 创建McmsNotice信息
//
// Create mcms notice information | 创建McmsNotice信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: McmsNoticeInfo
//
// Responses:
//  200: BaseMsgResp

func CreateMcmsNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.McmsNoticeInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsnotice.NewCreateMcmsNoticeLogic(r.Context(), svcCtx)
		resp, err := l.CreateMcmsNotice(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
