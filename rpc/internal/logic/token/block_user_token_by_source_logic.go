package token

import (
	"context"
	"github.com/neumathe/simple-admin-common/config"
	"github.com/neumathe/simple-admin-common/enum/common"
	"github.com/neumathe/simple-admin-common/i18n"
	"github.com/neumathe/simple-admin-common/msg/logmsg"
	"github.com/neumathe/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-core/rpc/ent/token"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUserTokenBySourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUserTokenBySourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUserTokenBySourceLogic {
	return &BlockUserTokenBySourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUserTokenBySourceLogic) BlockUserTokenBySource(in *core.BlockUserTokenBySourceReq) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Token.Update().Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).Where(token.SourceContains(in.Source)).SetStatus(common.StatusBanned).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	tokenData, err := l.svcCtx.DB.Token.Query().
		Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).
		Where(token.StatusEQ(common.StatusBanned)).
		Where(token.ExpiredAtGT(time.Now())).
		All(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	for _, v := range tokenData {
		expiredTime := v.ExpiredAt.Sub(time.Now())
		if expiredTime > 0 {
			err = l.svcCtx.Redis.Set(l.ctx, config.RedisTokenPrefix+v.Token, "1", expiredTime).Err()
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return nil, errorx.NewInternalError(i18n.RedisError)
			}
		}
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
