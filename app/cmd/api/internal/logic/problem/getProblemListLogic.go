package problem

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"time"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemListLogic {
	return &GetProblemListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemListLogic) GetProblemList(req *types.ProblemPageReq) (resp *types.ProblemListResp, err error) {
	problemDao := l.svcCtx.Repository.Model.Problem
	req.Page = (req.Page - 1) * req.Size
	titleObscure := "%" + req.KeyWord + "%"
	contentObscure := "%" + req.KeyWord + "%"

	problemDo := problemDao.WithContext(l.ctx).
		Where(problemDao.Title.Like(titleObscure)).
		Or(problemDao.Content.Like(contentObscure)).
		Offset(req.Page).
		Limit(req.Size)

	// 1.根据分类唯一标识获取category的id
	var categoryID int32
	categoryDao := l.svcCtx.Repository.Model.Category
	err = categoryDao.WithContext(l.ctx).
		Where(categoryDao.Identity.Eq(req.CategoryIdentity)).
		Select(categoryDao.ID).
		Scan(&categoryID)
	if err != nil {
		l.Logger.Info("查询种类id出错", err.Error())
		return nil, errors.New(consts.ErrUnKnow)
	}

	// 2.查询中间表 根据
	problems, err := problemDo.Find()
	if err != nil {
		return nil, errors.New(consts.ErrUnKnow)
	}

	problemCategoryDao := l.svcCtx.Repository.Model.ProblemCategory
	problemList := make([]types.Problem, 0)

	for _, problem := range problems {
		// 向中间表查询当前遍历的problem所关联的categoryID
		cids := make([]int32, 0)
		problemCategoryDo := problemCategoryDao.WithContext(l.ctx).Select(problemCategoryDao.CategoryID)
		// 判断category_identity是否是为空，不为空则查询当前遍历的问题的所有关联的category种类信息
		// 否者则只会查询当前category的信息
		if req.CategoryIdentity == "" {
			problemCategoryDo = problemCategoryDo.Where(problemCategoryDao.ProblemID.Eq(problem.ID))
		} else {
			problemCategoryDo = problemCategoryDo.Where(problemCategoryDao.ProblemID.Eq(problem.ID), problemCategoryDao.CategoryID.Eq(categoryID))
		}
		if err := problemCategoryDo.Scan(&cids); err != nil {
			return nil, errors.New(consts.ErrUnKnow)
		}

		// 根据id去查询关联问题的category分类的相关信息
		categories, err := categoryDao.WithContext(l.ctx).Where(categoryDao.ID.In(cids...)).Find()
		if err != nil {
			return nil, errors.New(consts.ErrUnKnow)
		}

		category := buildCategory(categories)

		// 设置删除时间

		problemList = append(problemList, types.Problem{
			Identity:   problem.Identity,
			Title:      problem.Title,
			MaxRuntime: problem.MaxRuntime,
			MaxMem:     problem.MaxMem,
			Content:    problem.Content,
			Category:   category,
			CreatedAt:  problem.CreatedAt.Format(time.DateTime),
			UpdatedAt:  problem.UpdatedAt.Format(time.DateTime),
			DeletedAt:  utils.ParseDeleteTime(problem.DeletedAt),
		})
	}
	return &types.ProblemListResp{
		Count:       int64(len(problemList)),
		ProblemList: problemList,
	}, nil
}

// buildCategory 将model.Category转换成types.Category
func buildCategory(categoryList []*model.Category) []types.Category {
	resp := make([]types.Category, 0)
	for _, c := range categoryList {
		category := types.Category{
			Identity:  c.Identity,
			Name:      c.Name,
			ParentId:  c.ParentID,
			CreatedAt: c.CreatedAt.Format(time.DateTime),
			UpdatedAt: c.UpdatedAt.Format(time.DateTime),
			DeletedAt: utils.ParseDeleteTime(c.DeletedAt),
		}
		resp = append(resp, category)
	}
	return resp
}
