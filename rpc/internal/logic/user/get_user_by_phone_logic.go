package user

import (
	"context"
	"github.com/neumathe/simple-admin-common/utils/pointy"
	"github.com/neumathe/simple-admin-core/rpc/ent/user"
	"github.com/neumathe/simple-admin-core/rpc/internal/utils/dberrorhandler"

	"github.com/neumathe/simple-admin-core/rpc/internal/svc"
	"github.com/neumathe/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByPhoneLogic {
	return &GetUserByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByPhoneLogic) GetUserByPhone(in *core.PhoneReq) (*core.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.MobileEQ(in.Phone)).WithRoles().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.UserInfo{
		Nickname:     &result.Nickname,
		Avatar:       &result.Avatar,
		Password:     &result.Password,
		RoleIds:      GetRoleIds(result.Edges.Roles),
		RoleCodes:    GetRoleCodes(result.Edges.Roles),
		Mobile:       &result.Mobile,
		Email:        &result.Email,
		Status:       pointy.GetPointer(uint32(result.Status)),
		Id:           pointy.GetPointer(result.ID.String()),
		Salt:         &result.Salt,
		HomePath:     &result.HomePath,
		Description:  &result.Description,
		DepartmentId: &result.DepartmentID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
	}, nil
}
