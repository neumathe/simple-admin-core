package user

import (
	"context"

	"github.com/neumathe/simple-admin-common/utils/encrypt"
	"github.com/neumathe/simple-admin-common/utils/pointy"

	"github.com/neumathe/simple-admin-core/rpc/ent/user"
	"github.com/neumathe/simple-admin-core/rpc/internal/svc"
	"github.com/neumathe/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/neumathe/simple-admin-core/rpc/types/core"

	"github.com/neumathe/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
	if in.Mobile != nil {
		checkMobile, err := l.svcCtx.DB.User.Query().Where(user.MobileEQ(*in.Mobile)).Exist(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		if checkMobile {
			return nil, errorx.NewInvalidArgumentError("login.mobileExist")
		}
	}

	if in.Email != nil {
		checkEmail, err := l.svcCtx.DB.User.Query().Where(user.EmailEQ(*in.Email)).Exist(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		if checkEmail {
			return nil, errorx.NewInvalidArgumentError("login.signupUserExist")
		}
	}

	salt := encrypt.CreateSalt(4)
	result, err := l.svcCtx.DB.User.Create().
		SetNotNilPassword(pointy.GetPointer(encrypt.MD5Encrypt(*in.Password, salt))).
		SetSalt(salt).
		SetNotNilNickname(in.Nickname).
		SetNotNilEmail(in.Email).
		SetNotNilMobile(in.Mobile).
		SetNotNilAvatar(in.Avatar).
		AddRoleIDs(in.RoleIds...).
		SetNotNilHomePath(in.HomePath).
		SetNotNilDescription(in.Description).
		SetNotNilDepartmentID(in.DepartmentId).
		AddPositionIDs(in.PositionIds...).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
