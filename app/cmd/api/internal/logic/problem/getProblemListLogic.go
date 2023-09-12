package problem

import (
	"context"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
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
	// TODO 增加缓存
	p := l.svcCtx.Repository.Model.Problem
	// page == 1 ==》 offset = 0
	req.Page = (req.Page - 1) * req.Size
	titleObscure := "%" + req.KeyWord + "%"
	contentObscure := "%" + req.KeyWord + "%"
	problemList, err := p.WithContext(l.ctx).
		Where(p.Title.Like(titleObscure)).
		Or(p.Content.Like(contentObscure)).
		Offset(req.Page).
		Limit(req.Size).
		Find()
	if err != nil {
		l.Logger.Info("获取问题列表失败", err.Error())
	}
	return buildProblemList(problemList), nil
}

func buildProblemList(problemList []*model.Problem) *types.ProblemListResp {
	resp := &types.ProblemListResp{}
	resp.ProblemList = make([]types.Problem, 0)
	for _, pb := range problemList {
		var p types.Problem
		p.ID = pb.ID
		p.Identity = pb.Identity
		p.Cid = pb.Cid
		p.Title = pb.Title
		p.MaxRuntime = pb.MaxRuntime
		p.MaxMem = pb.MaxMem
		p.Content = pb.Content
		p.CreatedAt = pb.CreatedAt.Format(time.DateTime)
		p.UpdatedAt = pb.UpdatedAt.Format(time.DateTime)
		if pb.DeletedAt.Valid {
			p.DeletedAt = pb.DeletedAt.Time.Format(time.DateTime)
		}
		resp.ProblemList = append(resp.ProblemList, p)
	}
	resp.Count = int64(len(resp.ProblemList))
	return resp
}
