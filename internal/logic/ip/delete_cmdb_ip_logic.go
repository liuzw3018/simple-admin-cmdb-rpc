package ip

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbip"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCmdbIpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCmdbIpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCmdbIpLogic {
	return &DeleteCmdbIpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCmdbIpLogic) DeleteCmdbIp(in *cmdb.IDsReq) (*cmdb.BaseResp, error) {
	_, err := l.svcCtx.DB.CmdbIp.Delete().Where(cmdbip.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseResp{Msg: errormsg.DeleteSuccess}, nil
}
