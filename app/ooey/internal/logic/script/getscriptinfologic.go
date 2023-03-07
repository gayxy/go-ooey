package script

import (
	"context"
	"github.com/gayxy/go-ooey/app/ooey/internal/svc"
	"github.com/gayxy/go-ooey/app/ooey/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScriptInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScriptInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScriptInfoLogic {
	return &GetScriptInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScriptInfoLogic) GetScriptInfo(req *types.GetScriptInfoRequest) (resp *types.GetScriptInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
