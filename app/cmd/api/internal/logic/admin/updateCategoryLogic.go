package admin

import (
	"context"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) error {
	categoryDao := l.svcCtx.Repository.Model.Category
	updateData := &model.Category{
		Name: req.Name,
	}
	_, err := categoryDao.WithContext(l.ctx).Where(categoryDao.Identity.Eq(req.CategoryIdentity)).Updates(updateData)
	if err != nil {
		return err
	}
	return nil
}
