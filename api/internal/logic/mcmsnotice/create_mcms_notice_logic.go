package mcmsnotice

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMcmsNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMcmsNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMcmsNoticeLogic {
	return &CreateMcmsNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateMcmsNoticeLogic) CreateMcmsNotice(req *types.McmsNoticeInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.CreateMcmsNotice(l.ctx,
		&mcms.McmsNoticeInfo{
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
