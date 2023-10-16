package ipuser

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCmdbIpUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCmdbIpUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCmdbIpUserLogic {
	return &UpdateCmdbIpUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCmdbIpUserLogic) UpdateCmdbIpUser(in *cmdb.CmdbIpUserInfo) (*cmdb.BaseResp, error) {
	query := l.svcCtx.DB.CmdbIpUser.UpdateOneID(*in.Id).
		SetNotNilUser(in.User).
		SetNotNilDepartment(in.Department).
		SetNotNilMobile(in.Mobile).
		SetNotNilRemark(in.Remark)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseResp{Msg: errormsg.UpdateSuccess}, nil
}
