package user

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"gorm.io/gorm"
	"time"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail(req *types.UserDeitalReq) (resp *types.UserDeitalResp, err error) {
	userDao := l.svcCtx.Repository.Model.User
	user, err := userDao.WithContext(l.ctx).
		Where(userDao.Identity.Eq(req.Identity)).
		First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(consts.ErrUnKnow)
	}

	if user == nil {
		return nil, errors.New(consts.ErrUserNotExist)
	}

	var deleteTime string
	if user.DeletedAt.Valid {
		deleteTime = user.DeletedAt.Time.Format(time.DateTime)
	}

	return &types.UserDeitalResp{
		Identity:  user.Identity,
		Name:      user.Name,
		Phone:     user.Phone,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
		DeletedAt: deleteTime,
	}, nil
}
