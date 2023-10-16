package ipuser

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbipuser"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCmdbIpUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCmdbIpUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCmdbIpUserLogic {
	return &DeleteCmdbIpUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCmdbIpUserLogic) DeleteCmdbIpUser(in *cmdb.IDsReq) (*cmdb.BaseResp, error) {
	_, err := l.svcCtx.DB.CmdbIpUser.Delete().Where(cmdbipuser.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseResp{Msg: errormsg.DeleteSuccess}, nil
}
