package sumit

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"gorm.io/gorm"
	"time"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitListLogic {
	return &SubmitListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitListLogic) SubmitList(req *types.SubmitListReq) (resp *types.SubmitListResp, err error) {
	submitDao := l.svcCtx.Repository.Model.Submit
	req.Page = (req.Page - 1) * req.Size
	submitDo := submitDao.WithContext(l.ctx).Offset(req.Page).Limit(req.Size)
	if req.ProblemIdentity != "" {
		submitDo.Where(submitDao.ProblemIdentity.Eq(req.ProblemIdentity))
	}
	if req.UserIdentity != "" {
		submitDo.Where(submitDao.UserIdentity.Eq(req.UserIdentity))
	}
	// TODO 添加校验是否存在该值
	if req.Status != 0 {
		submitDo.Where(submitDao.Status.Eq(req.Status))
	}
	if req.Language != 0 {
		submitDo.Where(submitDao.Language.Eq(req.Language))
	}
	list, err := submitDo.Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	submitList := l.buildSubmitList(list)
	return &types.SubmitListResp{
		SubmitList: submitList,
		Count:      int64(len(submitList)),
	}, nil
}
func (l *SubmitListLogic) buildSubmitList(submit []*model.Submit) []types.Submit {
	resp := make([]types.Submit, 0, len(submit))
	for _, s := range submit {
		deleteTime := ""
		if s.DeletedAt.Valid {
			deleteTime = s.DeletedAt.Time.Format(time.DateTime)
		}
		resp = append(resp, types.Submit{
			Identity:  s.Identity,
			Status:    s.Status,
			Language:  s.Language,
			RunTime:   s.RunTime,
			RunMem:    s.RunMem,
			CreatedAt: s.CreatedAt.Format(time.DateTime),
			UpdatedAt: s.UpdatedAt.Format(time.DateTime),
			DeletedAt: deleteTime,
		})
	}
	return resp
}
