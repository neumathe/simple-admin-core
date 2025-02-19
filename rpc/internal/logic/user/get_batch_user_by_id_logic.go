package user

import (
	"context"
	"encoding/json"
	"github.com/neumathe/simple-admin-common/utils/pointy"
	"github.com/neumathe/simple-admin-common/utils/uuidx"
	"github.com/neumathe/simple-admin-core/rpc/ent/user"
	"github.com/neumathe/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"time"

	"github.com/neumathe/simple-admin-core/rpc/internal/svc"
	"github.com/neumathe/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBatchUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBatchUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBatchUserByIdLogic {
	return &GetBatchUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBatchUserByIdLogic) GetBatchUserById(in *core.UUIDsReq) (*core.UserListResp, error) {
	resp := &core.UserListResp{}
	var idsToQuery []string

	for _, id := range in.Ids {
		userInfo := &core.UserInfo{}
		userBytes, err := l.svcCtx.Redis.Get(l.ctx, "userInfo_"+id).Result()
		if err == nil { // 缓存命中
			err := json.Unmarshal([]byte(userBytes), userInfo)
			if err != nil {
				l.Logger.Errorf("unmarshal user info error: %v", err)
				idsToQuery = append(idsToQuery, id)
			} else {
				resp.Data = append(resp.Data, userInfo)
			}
		} else { // 缓存未命中
			idsToQuery = append(idsToQuery, id)
		}
	}

	if len(idsToQuery) > 0 {
		users, err := l.svcCtx.DB.User.Query().Where(user.IDIn(uuidx.ParseUUIDSlice(idsToQuery)...)).WithRoles().All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		for _, v := range users {
			userInfo := &core.UserInfo{
				Id:          pointy.GetPointer(v.ID.String()),
				Avatar:      &v.Avatar,
				RoleIds:     GetRoleIds(v.Edges.Roles),
				RoleCodes:   GetRoleCodes(v.Edges.Roles),
				Nickname:    &v.Nickname,
				PositionIds: GetPositionIds(v.Edges.Positions),
				Description: &v.Description,
			}
			resp.Data = append(resp.Data, userInfo)
			userBytes, err := json.Marshal(userInfo)
			if err == nil {
				err := l.svcCtx.Redis.SetEx(l.ctx, "userInfo_"+v.ID.String(), userBytes, time.Hour*48).Err()
				if err != nil {
					l.Logger.Errorf("set redis error: %v", err)
				}
			} else {
				l.Logger.Errorf("marshal user info error: %v", err)
			}
		}
	}

	return resp, nil
}
