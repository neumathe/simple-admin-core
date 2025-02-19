// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"github.com/neumathe/simple-admin-core/rpc/internal/logic/api"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/authority"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/base"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/configuration"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/department"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/dictionary"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/dictionarydetail"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/menu"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/oauthprovider"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/position"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/role"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/token"
	"github.com/neumathe/simple-admin-core/rpc/internal/logic/user"
	"github.com/neumathe/simple-admin-core/rpc/internal/svc"
	"github.com/neumathe/simple-admin-core/rpc/types/core"
)

type CoreServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedCoreServer
}

func NewCoreServer(svcCtx *svc.ServiceContext) *CoreServer {
	return &CoreServer{
		svcCtx: svcCtx,
	}
}

// API management
func (s *CoreServer) CreateApi(ctx context.Context, in *core.ApiInfo) (*core.BaseIDResp, error) {
	l := api.NewCreateApiLogic(ctx, s.svcCtx)
	return l.CreateApi(in)
}

func (s *CoreServer) UpdateApi(ctx context.Context, in *core.ApiInfo) (*core.BaseResp, error) {
	l := api.NewUpdateApiLogic(ctx, s.svcCtx)
	return l.UpdateApi(in)
}

func (s *CoreServer) GetApiList(ctx context.Context, in *core.ApiListReq) (*core.ApiListResp, error) {
	l := api.NewGetApiListLogic(ctx, s.svcCtx)
	return l.GetApiList(in)
}

func (s *CoreServer) GetApiById(ctx context.Context, in *core.IDReq) (*core.ApiInfo, error) {
	l := api.NewGetApiByIdLogic(ctx, s.svcCtx)
	return l.GetApiById(in)
}

func (s *CoreServer) DeleteApi(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := api.NewDeleteApiLogic(ctx, s.svcCtx)
	return l.DeleteApi(in)
}

func (s *CoreServer) GetMenuAuthority(ctx context.Context, in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	l := authority.NewGetMenuAuthorityLogic(ctx, s.svcCtx)
	return l.GetMenuAuthority(in)
}

func (s *CoreServer) CreateOrUpdateMenuAuthority(ctx context.Context, in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	l := authority.NewCreateOrUpdateMenuAuthorityLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenuAuthority(in)
}

func (s *CoreServer) InitDatabase(ctx context.Context, in *core.Empty) (*core.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Configuration management
func (s *CoreServer) CreateConfiguration(ctx context.Context, in *core.ConfigurationInfo) (*core.BaseIDResp, error) {
	l := configuration.NewCreateConfigurationLogic(ctx, s.svcCtx)
	return l.CreateConfiguration(in)
}

func (s *CoreServer) UpdateConfiguration(ctx context.Context, in *core.ConfigurationInfo) (*core.BaseResp, error) {
	l := configuration.NewUpdateConfigurationLogic(ctx, s.svcCtx)
	return l.UpdateConfiguration(in)
}

func (s *CoreServer) GetConfigurationList(ctx context.Context, in *core.ConfigurationListReq) (*core.ConfigurationListResp, error) {
	l := configuration.NewGetConfigurationListLogic(ctx, s.svcCtx)
	return l.GetConfigurationList(in)
}

func (s *CoreServer) GetConfigurationById(ctx context.Context, in *core.IDReq) (*core.ConfigurationInfo, error) {
	l := configuration.NewGetConfigurationByIdLogic(ctx, s.svcCtx)
	return l.GetConfigurationById(in)
}

func (s *CoreServer) DeleteConfiguration(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := configuration.NewDeleteConfigurationLogic(ctx, s.svcCtx)
	return l.DeleteConfiguration(in)
}

// Department management
func (s *CoreServer) CreateDepartment(ctx context.Context, in *core.DepartmentInfo) (*core.BaseIDResp, error) {
	l := department.NewCreateDepartmentLogic(ctx, s.svcCtx)
	return l.CreateDepartment(in)
}

func (s *CoreServer) UpdateDepartment(ctx context.Context, in *core.DepartmentInfo) (*core.BaseResp, error) {
	l := department.NewUpdateDepartmentLogic(ctx, s.svcCtx)
	return l.UpdateDepartment(in)
}

func (s *CoreServer) GetDepartmentList(ctx context.Context, in *core.DepartmentListReq) (*core.DepartmentListResp, error) {
	l := department.NewGetDepartmentListLogic(ctx, s.svcCtx)
	return l.GetDepartmentList(in)
}

func (s *CoreServer) GetDepartmentById(ctx context.Context, in *core.IDReq) (*core.DepartmentInfo, error) {
	l := department.NewGetDepartmentByIdLogic(ctx, s.svcCtx)
	return l.GetDepartmentById(in)
}

func (s *CoreServer) DeleteDepartment(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := department.NewDeleteDepartmentLogic(ctx, s.svcCtx)
	return l.DeleteDepartment(in)
}

// Dictionary management
func (s *CoreServer) CreateDictionary(ctx context.Context, in *core.DictionaryInfo) (*core.BaseIDResp, error) {
	l := dictionary.NewCreateDictionaryLogic(ctx, s.svcCtx)
	return l.CreateDictionary(in)
}

func (s *CoreServer) UpdateDictionary(ctx context.Context, in *core.DictionaryInfo) (*core.BaseResp, error) {
	l := dictionary.NewUpdateDictionaryLogic(ctx, s.svcCtx)
	return l.UpdateDictionary(in)
}

func (s *CoreServer) GetDictionaryList(ctx context.Context, in *core.DictionaryListReq) (*core.DictionaryListResp, error) {
	l := dictionary.NewGetDictionaryListLogic(ctx, s.svcCtx)
	return l.GetDictionaryList(in)
}

func (s *CoreServer) GetDictionaryById(ctx context.Context, in *core.IDReq) (*core.DictionaryInfo, error) {
	l := dictionary.NewGetDictionaryByIdLogic(ctx, s.svcCtx)
	return l.GetDictionaryById(in)
}

func (s *CoreServer) DeleteDictionary(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := dictionary.NewDeleteDictionaryLogic(ctx, s.svcCtx)
	return l.DeleteDictionary(in)
}

// DictionaryDetail management
func (s *CoreServer) CreateDictionaryDetail(ctx context.Context, in *core.DictionaryDetailInfo) (*core.BaseIDResp, error) {
	l := dictionarydetail.NewCreateDictionaryDetailLogic(ctx, s.svcCtx)
	return l.CreateDictionaryDetail(in)
}

func (s *CoreServer) UpdateDictionaryDetail(ctx context.Context, in *core.DictionaryDetailInfo) (*core.BaseResp, error) {
	l := dictionarydetail.NewUpdateDictionaryDetailLogic(ctx, s.svcCtx)
	return l.UpdateDictionaryDetail(in)
}

func (s *CoreServer) GetDictionaryDetailList(ctx context.Context, in *core.DictionaryDetailListReq) (*core.DictionaryDetailListResp, error) {
	l := dictionarydetail.NewGetDictionaryDetailListLogic(ctx, s.svcCtx)
	return l.GetDictionaryDetailList(in)
}

func (s *CoreServer) GetDictionaryDetailById(ctx context.Context, in *core.IDReq) (*core.DictionaryDetailInfo, error) {
	l := dictionarydetail.NewGetDictionaryDetailByIdLogic(ctx, s.svcCtx)
	return l.GetDictionaryDetailById(in)
}

func (s *CoreServer) DeleteDictionaryDetail(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := dictionarydetail.NewDeleteDictionaryDetailLogic(ctx, s.svcCtx)
	return l.DeleteDictionaryDetail(in)
}

func (s *CoreServer) GetDictionaryDetailByDictionaryName(ctx context.Context, in *core.BaseMsg) (*core.DictionaryDetailListResp, error) {
	l := dictionarydetail.NewGetDictionaryDetailByDictionaryNameLogic(ctx, s.svcCtx)
	return l.GetDictionaryDetailByDictionaryName(in)
}

func (s *CoreServer) CreateMenu(ctx context.Context, in *core.MenuInfo) (*core.BaseIDResp, error) {
	l := menu.NewCreateMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(in)
}

func (s *CoreServer) UpdateMenu(ctx context.Context, in *core.MenuInfo) (*core.BaseResp, error) {
	l := menu.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *CoreServer) DeleteMenu(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := menu.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *CoreServer) GetMenuListByRole(ctx context.Context, in *core.BaseMsg) (*core.MenuInfoList, error) {
	l := menu.NewGetMenuListByRoleLogic(ctx, s.svcCtx)
	return l.GetMenuListByRole(in)
}

func (s *CoreServer) GetMenuList(ctx context.Context, in *core.PageInfoReq) (*core.MenuInfoList, error) {
	l := menu.NewGetMenuListLogic(ctx, s.svcCtx)
	return l.GetMenuList(in)
}

// OauthProvider management
func (s *CoreServer) CreateOauthProvider(ctx context.Context, in *core.OauthProviderInfo) (*core.BaseIDResp, error) {
	l := oauthprovider.NewCreateOauthProviderLogic(ctx, s.svcCtx)
	return l.CreateOauthProvider(in)
}

func (s *CoreServer) UpdateOauthProvider(ctx context.Context, in *core.OauthProviderInfo) (*core.BaseResp, error) {
	l := oauthprovider.NewUpdateOauthProviderLogic(ctx, s.svcCtx)
	return l.UpdateOauthProvider(in)
}

func (s *CoreServer) GetOauthProviderList(ctx context.Context, in *core.OauthProviderListReq) (*core.OauthProviderListResp, error) {
	l := oauthprovider.NewGetOauthProviderListLogic(ctx, s.svcCtx)
	return l.GetOauthProviderList(in)
}

func (s *CoreServer) GetOauthProviderById(ctx context.Context, in *core.IDReq) (*core.OauthProviderInfo, error) {
	l := oauthprovider.NewGetOauthProviderByIdLogic(ctx, s.svcCtx)
	return l.GetOauthProviderById(in)
}

func (s *CoreServer) DeleteOauthProvider(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := oauthprovider.NewDeleteOauthProviderLogic(ctx, s.svcCtx)
	return l.DeleteOauthProvider(in)
}

func (s *CoreServer) OauthLogin(ctx context.Context, in *core.OauthLoginReq) (*core.OauthRedirectResp, error) {
	l := oauthprovider.NewOauthLoginLogic(ctx, s.svcCtx)
	return l.OauthLogin(in)
}

func (s *CoreServer) OauthCallback(ctx context.Context, in *core.CallbackReq) (*core.UserInfo, error) {
	l := oauthprovider.NewOauthCallbackLogic(ctx, s.svcCtx)
	return l.OauthCallback(in)
}

// Position management
func (s *CoreServer) CreatePosition(ctx context.Context, in *core.PositionInfo) (*core.BaseIDResp, error) {
	l := position.NewCreatePositionLogic(ctx, s.svcCtx)
	return l.CreatePosition(in)
}

func (s *CoreServer) UpdatePosition(ctx context.Context, in *core.PositionInfo) (*core.BaseResp, error) {
	l := position.NewUpdatePositionLogic(ctx, s.svcCtx)
	return l.UpdatePosition(in)
}

func (s *CoreServer) GetPositionList(ctx context.Context, in *core.PositionListReq) (*core.PositionListResp, error) {
	l := position.NewGetPositionListLogic(ctx, s.svcCtx)
	return l.GetPositionList(in)
}

func (s *CoreServer) GetPositionById(ctx context.Context, in *core.IDReq) (*core.PositionInfo, error) {
	l := position.NewGetPositionByIdLogic(ctx, s.svcCtx)
	return l.GetPositionById(in)
}

func (s *CoreServer) DeletePosition(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := position.NewDeletePositionLogic(ctx, s.svcCtx)
	return l.DeletePosition(in)
}

// Role management
func (s *CoreServer) CreateRole(ctx context.Context, in *core.RoleInfo) (*core.BaseIDResp, error) {
	l := role.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *CoreServer) UpdateRole(ctx context.Context, in *core.RoleInfo) (*core.BaseResp, error) {
	l := role.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *CoreServer) GetRoleList(ctx context.Context, in *core.RoleListReq) (*core.RoleListResp, error) {
	l := role.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}

func (s *CoreServer) GetRoleById(ctx context.Context, in *core.IDReq) (*core.RoleInfo, error) {
	l := role.NewGetRoleByIdLogic(ctx, s.svcCtx)
	return l.GetRoleById(in)
}

func (s *CoreServer) DeleteRole(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := role.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

// Token management
func (s *CoreServer) CreateToken(ctx context.Context, in *core.TokenInfo) (*core.BaseUUIDResp, error) {
	l := token.NewCreateTokenLogic(ctx, s.svcCtx)
	return l.CreateToken(in)
}

func (s *CoreServer) DeleteToken(ctx context.Context, in *core.UUIDsReq) (*core.BaseResp, error) {
	l := token.NewDeleteTokenLogic(ctx, s.svcCtx)
	return l.DeleteToken(in)
}

func (s *CoreServer) GetTokenList(ctx context.Context, in *core.TokenListReq) (*core.TokenListResp, error) {
	l := token.NewGetTokenListLogic(ctx, s.svcCtx)
	return l.GetTokenList(in)
}

func (s *CoreServer) GetTokenById(ctx context.Context, in *core.UUIDReq) (*core.TokenInfo, error) {
	l := token.NewGetTokenByIdLogic(ctx, s.svcCtx)
	return l.GetTokenById(in)
}

func (s *CoreServer) BlockUserAllToken(ctx context.Context, in *core.UUIDReq) (*core.BaseResp, error) {
	l := token.NewBlockUserAllTokenLogic(ctx, s.svcCtx)
	return l.BlockUserAllToken(in)
}

func (s *CoreServer) BlockUserTokenBySource(ctx context.Context, in *core.BlockUserTokenBySourceReq) (*core.BaseResp, error) {
	l := token.NewBlockUserTokenBySourceLogic(ctx, s.svcCtx)
	return l.BlockUserTokenBySource(in)
}

func (s *CoreServer) UpdateToken(ctx context.Context, in *core.TokenInfo) (*core.BaseResp, error) {
	l := token.NewUpdateTokenLogic(ctx, s.svcCtx)
	return l.UpdateToken(in)
}

// User management
func (s *CoreServer) CreateUser(ctx context.Context, in *core.UserInfo) (*core.BaseUUIDResp, error) {
	l := user.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *CoreServer) UpdateUser(ctx context.Context, in *core.UserInfo) (*core.BaseResp, error) {
	l := user.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *CoreServer) GetUserList(ctx context.Context, in *core.UserListReq) (*core.UserListResp, error) {
	l := user.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}

func (s *CoreServer) GetUserById(ctx context.Context, in *core.UUIDReq) (*core.UserInfo, error) {
	l := user.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *CoreServer) DeleteUser(ctx context.Context, in *core.UUIDsReq) (*core.BaseResp, error) {
	l := user.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *CoreServer) GetUserByEmail(ctx context.Context, in *core.EmailReq) (*core.UserInfo, error) {
	l := user.NewGetUserByEmailLogic(ctx, s.svcCtx)
	return l.GetUserByEmail(in)
}

func (s *CoreServer) GetUserByPhone(ctx context.Context, in *core.PhoneReq) (*core.UserInfo, error) {
	l := user.NewGetUserByPhoneLogic(ctx, s.svcCtx)
	return l.GetUserByPhone(in)
}

func (s *CoreServer) GetBatchUserById(ctx context.Context, in *core.UUIDsReq) (*core.UserListResp, error) {
	l := user.NewGetBatchUserByIdLogic(ctx, s.svcCtx)
	return l.GetBatchUserById(in)
}
