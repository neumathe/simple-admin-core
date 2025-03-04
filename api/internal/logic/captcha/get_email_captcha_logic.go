package captcha

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/neumathe/neumathe-message-center/types/mcms"
	"github.com/neumathe/simple-admin-common/config"
	"github.com/neumathe/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCaptchaLogic {
	return &GetEmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetEmailCaptchaLogic) GetEmailCaptcha(req *types.EmailCaptchaReq) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeInvalidArgumentError("captcha.mcmsNotEnabled")
	}

	captcha := random.RandInt(10000, 99999)
	_, err = l.svcCtx.McmsRpc.SendEmail(l.ctx, &mcms.EmailInfo{
		Target:  []string{req.Email},
		Subject: l.svcCtx.Trans.Trans(l.ctx, "mcms.email.subject"),
		Content: fmt.Sprintf("%s%d", l.svcCtx.Trans.Trans(l.ctx, "mcms.email.content"), captcha),
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Redis.Set(l.ctx, config.RedisCaptchaPrefix+req.Email, strconv.Itoa(captcha), time.Duration(l.svcCtx.Config.ProjectConf.EmailCaptchaExpiredTime)*time.Second).Err()
	if err != nil {
		logx.Errorw("failed to write email captcha to redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
