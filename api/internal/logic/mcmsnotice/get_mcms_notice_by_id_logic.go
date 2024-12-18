package mcmsnotice

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"
	"github.com/neumathe/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMcmsNoticeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMcmsNoticeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMcmsNoticeByIdLogic {
	return &GetMcmsNoticeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMcmsNoticeByIdLogic) GetMcmsNoticeById(req *types.UUIDReq) (resp *types.McmsNoticeInfoResp, err error) {
	data, err := l.svcCtx.McmsRpc.GetMcmsNoticeById(l.ctx, &mcms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.McmsNoticeInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.McmsNoticeInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:    data.Status,
			Title:     data.Title,
			Content:   data.Content,
			Type:      data.Type,
			Duration:  data.Duration,
			Once:      data.Once,
			Reference: data.Reference,
		},
	}, nil
}
