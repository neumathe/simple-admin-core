package mcmsnotice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/mcmsnotice"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_notice mcmsnotice GetMcmsNoticeById
//
// Get mcms notice by ID | 通过ID获取McmsNotice信息
//
// Get mcms notice by ID | 通过ID获取McmsNotice信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDReq
//
// Responses:
//  200: McmsNoticeInfoResp

func GetMcmsNoticeByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsnotice.NewGetMcmsNoticeByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetMcmsNoticeById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
