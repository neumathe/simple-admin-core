package mcmsbanner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/mcmsbanner"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_banner/create mcmsbanner CreateMcmsBanner
//
// Create mcms banner information | 创建McmsBanner信息
//
// Create mcms banner information | 创建McmsBanner信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: McmsBannerInfo
//
// Responses:
//  200: BaseMsgResp

func CreateMcmsBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.McmsBannerInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsbanner.NewCreateMcmsBannerLogic(r.Context(), svcCtx)
		resp, err := l.CreateMcmsBanner(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
