package mcmsbanner

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMcmsBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMcmsBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMcmsBannerLogic {
	return &UpdateMcmsBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateMcmsBannerLogic) UpdateMcmsBanner(req *types.McmsBannerInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.UpdateMcmsBanner(l.ctx,
		&mcms.McmsBannerInfo{
			Id:          req.Id,
			Status:      req.Status,
			Image:       req.Image,
			Type:        req.Type,
			Name:        req.Name,
			Description: req.Description,
			Reference:   req.Reference,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
