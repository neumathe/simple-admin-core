package mcmsbanner

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"
	"github.com/neumathe/simple-admin-common/i18n"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMcmsBannerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMcmsBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMcmsBannerListLogic {
	return &GetMcmsBannerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMcmsBannerListLogic) GetMcmsBannerList(req *types.McmsBannerListReq) (resp *types.McmsBannerListResp, err error) {
	data, err := l.svcCtx.McmsRpc.GetMcmsBannerList(l.ctx,
		&mcms.McmsBannerListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Image:       req.Image,
			Name:        req.Name,
			Description: req.Description,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.McmsBannerListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.McmsBannerInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:      v.Status,
				Image:       v.Image,
				Type:        v.Type,
				Name:        v.Name,
				Description: v.Description,
				Reference:   v.Reference,
			})
	}
	return resp, nil
}
