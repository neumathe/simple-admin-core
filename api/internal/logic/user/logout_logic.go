package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.BaseMsgResp, err error) {
	if req.Source != "" {
		result, err := l.svcCtx.CoreRpc.BlockUserTokenBySource(l.ctx, &core.BlockUserTokenBySourceReq{Id: l.ctx.Value("userId").(string), Source: req.Source})
		if err != nil {
			return nil, err
		}
		return &types.BaseMsgResp{Msg: result.Msg}, nil
	} else {
		result, err := l.svcCtx.CoreRpc.BlockUserAllToken(l.ctx, &core.UUIDReq{Id: l.ctx.Value("userId").(string)})
		if err != nil {
			return nil, err
		}
		return &types.BaseMsgResp{Msg: result.Msg}, nil
	}
}
