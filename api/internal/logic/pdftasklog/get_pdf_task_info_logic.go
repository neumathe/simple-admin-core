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

type GetPdfTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPdfTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPdfTaskInfoLogic {
	return &GetPdfTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPdfTaskInfoLogic) GetPdfTaskInfo(req *types.UUIDReq) (resp *types.PdfTaskInfoResp, err error) {
	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	userId := l.ctx.Value("userId").(string)
	info, err := l.svcCtx.JobRpc.GetPdfTaskInfo(l.ctx, &job.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if info.CreatedBy != userId {
		return nil, errorx.NewApiForbiddenError("无权访问")
	}
	return &types.PdfTaskInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.PdfTaskInfo{
			Id:            info.Id,
			CreatedAt:     info.CreatedAt,
			UpdatedAt:     info.UpdatedAt,
			CreatedBy:     info.CreatedBy,
			Status:        info.Status,
			QuestionCount: info.QuestionCount,
			FileName:      info.FileName,
			Duration:      info.Duration,
			StartedAt:     info.StartedAt,
			DownloadURL:   fmt.Sprintf("%s/%s.pdf", l.svcCtx.Config.ProjectConf.PDFDownloadPrefix, info.FileName),
		},
	}, nil
}
