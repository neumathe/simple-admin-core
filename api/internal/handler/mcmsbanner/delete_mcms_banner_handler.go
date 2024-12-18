package mcmsbanner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/mcmsbanner"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_banner/delete mcmsbanner DeleteMcmsBanner
//
// Delete mcms banner information | 删除McmsBanner信息
//
// Delete mcms banner information | 删除McmsBanner信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteMcmsBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsbanner.NewDeleteMcmsBannerLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMcmsBanner(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
