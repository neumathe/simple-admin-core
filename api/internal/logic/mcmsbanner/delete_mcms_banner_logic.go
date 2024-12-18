package mcmsbanner

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"
	"github.com/neumathe/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMcmsBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMcmsBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMcmsBannerLogic {
	return &DeleteMcmsBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteMcmsBannerLogic) DeleteMcmsBanner(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.DeleteEmailLog(l.ctx, &mcms.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil

}
