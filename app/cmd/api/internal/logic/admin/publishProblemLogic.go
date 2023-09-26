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

type PublishProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishProblemLogic {
	return &PublishProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishProblemLogic) PublishProblem(req *types.PublishProblemReq) error {
	problemDao := l.svcCtx.Repository.Model.Problem
	categoryDao := l.svcCtx.Repository.Model.Category
	// 1.查询数据库标题是否重复
	serachProblem, err := problemDao.
		WithContext(l.ctx).
		Where(problemDao.Title.Eq(req.Title)).First()
	if utils.GormFirstError(err) != nil {
		return errors.New(consts.ErrUnKnow)
	}
	if serachProblem != nil {
		return errors.New(consts.ErrExistProblemTitle)
	}

	// 2.判断种类是否存在
	// 2.1 根据唯一标识查询id
	var categoryIds []int32
	if err = categoryDao.WithContext(l.ctx).
		Where(categoryDao.Identity.In(req.CategoryIdentity...)).
		Select(categoryDao.ID).
		Scan(&categoryIds); utils.GormFirstError(err) != nil {
		return errors.New(consts.ErrUnKnow)
	}

	// todo 增加检测 具体当前哪个种类是否存在
	if len(categoryIds) != len(req.CategoryIdentity) {
		return errors.New(consts.ErrNotExistCategory)
	}

	// 3.创建问题 以及 创建问题种类关联
	q := query.Use(repo.GlobalDB)
	err = q.Transaction(func(tx *query.Query) error {
		uuid := utils.NewUUID()
		// 3.1 创建问题
		if err := tx.Problem.WithContext(l.ctx).Create(&model.Problem{
			Identity:   uuid,
			Title:      req.Title,
			MaxRuntime: req.MaxRuntime,
			MaxMem:     req.MaxMem,
			Path:       req.Path,
			Content:    req.Content,
		}); err != nil {
			return err
		}

		// 3.2 创建问题种类
		// 3.2.1 查询创建问题的id
		searchProblem, err := tx.Problem.WithContext(l.ctx).Where(tx.Problem.Identity.Eq(uuid)).First()
		if err := utils.GormFirstError(err); err != nil {
			return err
		}

		// 3.2.2 创建问题和种类分类
		categories := make([]*model.ProblemCategory, 0)
		for i := 0; i < len(req.CategoryIdentity); i++ {
			categories = append(categories, &model.ProblemCategory{
				ProblemID:  searchProblem.ID,
				CategoryID: categoryIds[i],
			})
		}
		if err := tx.ProblemCategory.WithContext(l.ctx).Create(categories...); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return errors.New(consts.ErrCreateProblem)
	}

	return nil
}
