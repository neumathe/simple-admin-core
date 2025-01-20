package publicuser

import (
	"context"
	"github.com/neumathe/simple-admin-core/rpc/coreclient"
	"regexp"
	"strings"
	"time"

	"github.com/neumathe/simple-admin-common/config"
	"github.com/neumathe/simple-admin-common/enum/common"

	"github.com/neumathe/simple-admin-common/utils/encrypt"
	"github.com/neumathe/simple-admin-common/utils/jwt"
	"github.com/neumathe/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"
	"github.com/neumathe/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" && l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	if ok := l.svcCtx.Captcha.Verify(config.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true); ok {
		emailRegex := regexp.MustCompile(`^[\w.!#$%&'*+/=?^` + "`" + `{|}~-]+@[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$`)
		phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
		var user *coreclient.UserInfo
		if emailRegex.MatchString(req.EmailOrPhone) {
			user, err = l.svcCtx.CoreRpc.GetUserByEmail(l.ctx, &core.EmailReq{
				Email: req.EmailOrPhone,
			})
			if err != nil {
				return nil, err
			}
		} else if phoneRegex.MatchString(req.EmailOrPhone) {
			user, err = l.svcCtx.CoreRpc.GetUserByPhone(l.ctx, &core.PhoneReq{
				Phone: req.EmailOrPhone,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errorx.NewInvalidArgumentError("login.emailOrMobileFormatError")
		}
		if !encrypt.MD5Check(req.Password, *user.Salt, *user.Password) {
			return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
		}

		if user.Status != nil && *user.Status != uint32(common.StatusNormal) {
			return nil, errorx.NewCodeInvalidArgumentError("login.userBanned")
		}

		token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", user.Id), jwt.WithOption("roleId",
				strings.Join(user.RoleCodes, ",")), jwt.WithOption("deptId", user.DepartmentId))
		if err != nil {
			return nil, err
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
		_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
			Uuid:      user.Id,
			Token:     pointy.GetPointer(token),
			Source:    pointy.GetPointer("web"),
			Status:    pointy.GetPointer(uint32(common.StatusNormal)),
			Username:  user.Nickname,
			ExpiredAt: pointy.GetPointer(expiredAt),
		})

		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.CaptchaId).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
			Data: types.LoginInfo{
				UserId: *user.Id,
				Token:  token,
				Expire: uint64(expiredAt),
			},
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongCaptcha")
	}
}
