package admin

import (
	"context"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

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

	return nil
}
