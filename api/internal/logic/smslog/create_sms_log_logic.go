package smslog

import (
	"context"

	"github.com/neumathe/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/neumathe/neumathe-message-center/types/mcms"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSmsLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSmsLogLogic {
	return &CreateSmsLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSmsLogLogic) CreateSmsLog(req *types.SmsLogInfo) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.CreateSmsLog(l.ctx,
		&mcms.SmsLogInfo{
			PhoneNumber: req.PhoneNumber,
			Content:     req.Content,
			SendStatus:  req.SendStatus,
			Provider:    req.Provider,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
