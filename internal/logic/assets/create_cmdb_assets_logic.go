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

type CreateCmdbAssetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCmdbAssetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCmdbAssetsLogic {
	return &CreateCmdbAssetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCmdbAssetsLogic) CreateCmdbAssets(in *cmdb.CmdbAssetsInfo) (*cmdb.BaseIDResp, error) {
	query := l.svcCtx.DB.CmdbAssets.Create().
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
		SetNotNilCPU(in.Cpu).
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

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseIDResp{Id: result.ID, Msg: errormsg.CreateSuccess}, nil
}
