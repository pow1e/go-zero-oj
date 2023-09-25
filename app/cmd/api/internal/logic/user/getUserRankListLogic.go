package user

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"time"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRankListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRankListLogic {
	return &GetUserRankListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRankListLogic) GetUserRankList(req *types.UserRankListReq) (resp []types.UserRankListResp, err error) {
	req.Page = (req.Page - 1) * req.Size
	userDao := l.svcCtx.Repository.Model.User
	userDo := userDao.WithContext(l.ctx).
		Offset(req.Page).
		Limit(req.Size).
		Where(userDao.ID.Neq(consts.AdminID)).
		Order(userDao.FinishProblemNum.Desc(), userDao.SubmitNum.Desc())
	if req.Name != "" {
		userDo.Where(userDao.Name.Eq(req.Name))
	}
	serachUsers, err := userDo.Find()
	if err != nil {
		return nil, errors.New(consts.ErrSearch)
	}
	return l.buildUserRankListResp(serachUsers), nil
}

// buildUserRankListResp 将model.User进行转换
func (l *GetUserRankListLogic) buildUserRankListResp(user []*model.User) []types.UserRankListResp {
	resp := make([]types.UserRankListResp, 0)
	for _, u := range user {
		resp = append(resp, types.UserRankListResp{
			Identity:         u.Identity,
			Name:             u.Name,
			FinishProblemNum: u.FinishProblemNum,
			SubmitNum:        u.SubmitNum,
			CreatedAt:        u.CreatedAt.Format(time.DateTime),
		})
	}
	return resp
}
