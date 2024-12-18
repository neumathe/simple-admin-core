package mcmsnotice

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMcmsNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMcmsNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMcmsNoticeLogic {
	return &DeleteMcmsNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteMcmsNoticeLogic) DeleteMcmsNotice(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.DeleteMcmsNotice(l.ctx, &mcms.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
