package mcmsbanner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/mcmsbanner"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /mcms_banner/list mcmsbanner GetMcmsBannerList
//
// Get mcms banner list | 获取McmsBanner信息列表
//
// Get mcms banner list | 获取McmsBanner信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: McmsBannerListReq
//
// Responses:
//  200: McmsBannerListResp

func GetMcmsBannerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.McmsBannerListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := mcmsbanner.NewGetMcmsBannerListLogic(r.Context(), svcCtx)
		resp, err := l.GetMcmsBannerList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
