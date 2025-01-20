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

type GetMcmsBannerByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMcmsBannerByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMcmsBannerByIdLogic {
	return &GetMcmsBannerByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMcmsBannerByIdLogic) GetMcmsBannerById(req *types.UUIDReq) (resp *types.McmsBannerInfoResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetMcmsBannerById(l.ctx, &mcms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.McmsBannerInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.McmsBannerInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:      data.Status,
			Image:       data.Image,
			Type:        data.Type,
			Name:        data.Name,
			Description: data.Description,
			Reference:   data.Reference,
		},
	}, nil
}
