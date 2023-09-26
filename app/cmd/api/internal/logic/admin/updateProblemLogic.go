package admin

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/query"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/repo"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProblemLogic) UpdateProblem(req *types.UpdateProblemReq) error {
	// 1.查询唯一标识是否存在
	problemDao := l.svcCtx.Repository.Model.Problem
	serachProblem, err := problemDao.WithContext(l.ctx).
		Where(problemDao.Identity.Eq(req.ProblemIdentity)).
		First()
	if utils.GormFirstError(err) != nil {
		return errors.New(consts.ErrUnKnow)
	}

	if serachProblem == nil {
		return errors.New(consts.ErrNotExistProblem)
	}

	// 2.修改种类以及问题
	q := query.Use(repo.GlobalDB)
	if err = q.Transaction(func(tx *query.Query) error {
		problemCategoryTx := tx.ProblemCategory
		categoryTx := tx.Category
		// 2.1 根据id修改问题表
		updateProblemModel := &model.Problem{
			Identity:   req.ProblemIdentity,
			Title:      req.Title,
			MaxRuntime: req.MaxRuntime,
			MaxMem:     req.MaxMem,
			Path:       req.Path,
			Content:    req.Content,
		}
		_, err = tx.Problem.WithContext(l.ctx).
			Where(tx.Problem.ID.Eq(serachProblem.ID)).
			Updates(updateProblemModel)
		if err != nil {
			return err
		}

		// 2.2 修改该问题的种类-->先删除后插入
		// 2.2.1 查找当前问题的所有关联种类的id
		// todo 添加输入种类不存在的情况
		var deleteIds []int32
		if err = problemCategoryTx.WithContext(l.ctx).
			Where(problemCategoryTx.ProblemID.Eq(serachProblem.ID)).
			Select(problemCategoryTx.ID).Scan(&deleteIds); err != nil {
			return err
		}

		// 2.2.2 根据id删除所有问题关联的种类
		if _, err = problemCategoryTx.WithContext(l.ctx).
			Where(problemCategoryTx.ID.In(deleteIds...)).
			Delete(); err != nil {
			return err
		}

		// 2.2.3 根据请求中的唯一标识查询种类的id
		var categoryIds []int32
		if err = categoryTx.WithContext(l.ctx).
			Where(categoryTx.Identity.In(req.CategoryIdentity...)).
			Select(categoryTx.ID).
			Scan(&categoryIds); utils.GormFirstError(err) != nil {
			return errors.New(consts.ErrUnKnow)
		}

		if len(categoryIds) == 0 {
			return errors.New(consts.ErrNotExistCategory)
		}

		// 2.2.4 插入当前请求的种类
		createCategoryData := make([]*model.ProblemCategory, 0)
		for i := 0; i < len(categoryIds); i++ {
			createCategoryData = append(createCategoryData, &model.ProblemCategory{
				ProblemID:  serachProblem.ID,
				CategoryID: categoryIds[i],
			})
		}

		if err = problemCategoryTx.
			WithContext(l.ctx).
			Create(createCategoryData...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
