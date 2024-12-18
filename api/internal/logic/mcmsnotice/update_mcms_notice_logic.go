package mcmsnotice

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMcmsNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMcmsNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMcmsNoticeLogic {
	return &UpdateMcmsNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateMcmsNoticeLogic) UpdateMcmsNotice(req *types.McmsNoticeInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.UpdateMcmsNotice(l.ctx,
		&mcms.McmsNoticeInfo{
			Id:        req.Id,
			Status:    req.Status,
			Title:     req.Title,
			Content:   req.Content,
			Type:      req.Type,
			Duration:  req.Duration,
			Once:      req.Once,
			Reference: req.Reference,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
