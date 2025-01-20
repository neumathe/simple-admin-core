package mcmsnotice

import (
	"context"
	"github.com/neumathe/neumathe-message-center/types/mcms"
	"github.com/neumathe/simple-admin-common/i18n"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMcmsNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMcmsNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMcmsNoticeListLogic {
	return &GetMcmsNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMcmsNoticeListLogic) GetMcmsNoticeList(req *types.McmsNoticeListReq) (resp *types.McmsNoticeListResp, err error) {
	data, err := l.svcCtx.McmsRpc.GetMcmsNoticeList(l.ctx,
		&mcms.McmsNoticeListReq{
			Page:      req.Page,
			PageSize:  req.PageSize,
			Title:     req.Title,
			Content:   req.Content,
			Reference: req.Reference,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.McmsNoticeListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.McmsNoticeInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:    v.Status,
				Title:     v.Title,
				Content:   v.Content,
				Type:      v.Type,
				Duration:  v.Duration,
				Once:      v.Once,
				Reference: v.Reference,
			})
	}
	return resp, nil
}
