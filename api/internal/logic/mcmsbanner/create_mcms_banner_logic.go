package mcmsbanner

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"
	"github.com/neumathe/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMcmsBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMcmsBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMcmsBannerLogic {
	return &CreateMcmsBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateMcmsBannerLogic) CreateMcmsBanner(req *types.McmsBannerInfo) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.CreateMcmsBanner(l.ctx,
		&mcms.McmsBannerInfo{
			Id:          req.Id,
			CreatedAt:   req.CreatedAt,
			UpdatedAt:   req.UpdatedAt,
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
