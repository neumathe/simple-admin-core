package pdftasklog

import (
	"context"
	"fmt"
	"github.com/neumathe/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-job/types/job"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/neumathe/simple-admin-core/api/internal/svc"
	"github.com/neumathe/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPdfTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPdfTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPdfTaskListLogic {
	return &GetPdfTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPdfTaskListLogic) GetPdfTaskList(req *types.PdfTaskListReq) (resp *types.PdfTaskInfoListResp, err error) {
	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	userId := l.ctx.Value("userId").(string)
	data, err := l.svcCtx.JobRpc.GetPdfTaskList(l.ctx,
		&job.PdfTaskListReq{
			Page:      req.Page,
			PageSize:  req.PageSize,
			CreatedBy: &userId,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.PdfTaskInfoListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.PdfTaskInfo{
				Id:            v.Id,
				CreatedAt:     v.CreatedAt,
				UpdatedAt:     v.UpdatedAt,
				CreatedBy:     v.CreatedBy,
				Status:        v.Status,
				QuestionCount: v.QuestionCount,
				FileName:      v.FileName,
				Duration:      v.Duration,
				StartedAt:     v.StartedAt,
				DownloadURL:   fmt.Sprintf("%s/%s.pdf", l.svcCtx.Config.ProjectConf.PDFDownloadPrefix, v.FileName),
			})
	}
	return resp, nil
}
