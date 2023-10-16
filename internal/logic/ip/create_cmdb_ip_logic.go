package ip

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCmdbIpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCmdbIpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCmdbIpLogic {
	return &CreateCmdbIpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCmdbIpLogic) CreateCmdbIp(in *cmdb.CmdbIpInfo) (*cmdb.BaseIDResp, error) {
	query := l.svcCtx.DB.CmdbIp.Create().
		SetNotNilUser(in.User).
		SetNotNilDepartment(in.Department).
		SetNotNilMobile(in.Mobile).
		SetNotNilRemark(in.Remark).
		SetNotNilIP(in.Ip).
		SetNotNilMask(in.Mask).
		SetNotNilGateway(in.Gateway).
		SetNotNilOnlineTime(pointy.GetTimePointer(in.OnlineTime, 0)).
		SetNotNilOfflineTime(pointy.GetTimePointer(in.OfflineTime, 0)).
		SetNotNilDeviceType(in.DeviceType)

	if in.IsLeisure != nil {
		query.SetNotNilIsLeisure(pointy.GetPointer(uint(*in.IsLeisure)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseIDResp{Id: result.ID, Msg: errormsg.CreateSuccess}, nil
}
