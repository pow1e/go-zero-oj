package problem

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"time"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemDetailLogic {
	return &GetProblemDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemDetailLogic) GetProblemDetail(req *types.ProblemDeitalReq) (resp *types.ProblemDeitalResp, err error) {
	// 1.查询当前表示的详细信息
	problemDao := l.svcCtx.Repository.Model.Problem
	serachProblem, err := problemDao.WithContext(l.ctx).Where(problemDao.Identity.Eq(req.Identity)).First()
	if err != nil {
		return nil, errors.New(consts.ErrUnKnow)
	}

	// 2.查询关联的category种类信息
	var cids []int32
	problemCategoryDao := l.svcCtx.Repository.Model.ProblemCategory
	if err = problemCategoryDao.WithContext(l.ctx).
		Where(problemCategoryDao.ProblemID.Eq(serachProblem.ID)).
		Select(problemCategoryDao.CategoryID).
		Scan(&cids); err != nil {
		return nil, errors.New(consts.ErrUnKnow)
	}

	// 3.根据id去查询category列表
	categoryDao := l.svcCtx.Repository.Model.Category
	categoryList, err := categoryDao.WithContext(l.ctx).Where(categoryDao.ID.In(cids...)).Find()
	if err != nil {
		return nil, errors.New(consts.ErrUnKnow)
	}

	deleteTime := ""
	if serachProblem.DeletedAt.Valid {
		deleteTime = serachProblem.DeletedAt.Time.Format(time.DateTime)
	}
	problem := types.Problem{
		Identity:   serachProblem.Identity,
		Title:      serachProblem.Title,
		MaxRuntime: serachProblem.MaxRuntime,
		MaxMem:     serachProblem.MaxMem,
		Content:    serachProblem.Content,
		Category:   buildCategory(categoryList),
		CreatedAt:  serachProblem.CreatedAt.Format(time.DateTime),
		UpdatedAt:  serachProblem.UpdatedAt.Format(time.DateTime),
		DeletedAt:  deleteTime,
	}

	return &types.ProblemDeitalResp{Problem: problem}, nil
}
