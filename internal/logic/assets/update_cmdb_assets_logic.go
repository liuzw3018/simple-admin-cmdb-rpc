package assets

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCmdbAssetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCmdbAssetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCmdbAssetsLogic {
	return &UpdateCmdbAssetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCmdbAssetsLogic) UpdateCmdbAssets(in *cmdb.CmdbAssetsInfo) (*cmdb.BaseResp, error) {
	query := l.svcCtx.DB.CmdbAssets.UpdateOneID(*in.Id).
		SetNotNilUser(in.User).
		SetNotNilDepartment(in.Department).
		SetNotNilMobile(in.Mobile).
		SetNotNilRemark(in.Remark).
		SetNotNilIP(in.Ip).
		SetNotNilMask(in.Mask).
		SetNotNilGateway(in.Gateway).
		SetNotNilOnlineTime(pointy.GetTimePointer(in.OnlineTime, 0)).
		SetNotNilOfflineTime(pointy.GetTimePointer(in.OfflineTime, 0)).
		SetNotNilServerHostname(in.ServerHostname).
		SetNotNilServerOs(in.ServerOs).
		SetNotNilServerOsVersion(in.ServerOsVersion).
		SetNotNilServerOsArch(in.ServerOsArch).
		SetNotNilCpu(in.Cpu).
		SetNotNilMemory(in.Memory).
		SetNotNilDisk(in.Disk).
		SetNotNilNetworkSpeed(in.NetworkSpeed).
		SetNotNilDeviceAddress(in.DeviceAddress)

	if in.PowerStatus != nil {
		query.SetNotNilPowerStatus(pointy.GetPointer(uint(*in.PowerStatus)))
	}
	if in.IsServer != nil {
		query.SetNotNilIsServer(pointy.GetPointer(uint(*in.IsServer)))
	}
	if in.ServerType != nil {
		query.SetNotNilServerType(pointy.GetPointer(uint(*in.ServerType)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseResp{Msg: errormsg.UpdateSuccess}, nil
}
