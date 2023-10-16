package assets

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdbAssetsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdbAssetsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdbAssetsByIdLogic {
	return &GetCmdbAssetsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdbAssetsByIdLogic) GetCmdbAssetsById(in *cmdb.IDReq) (*cmdb.CmdbAssetsInfo, error) {
	result, err := l.svcCtx.DB.CmdbAssets.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.CmdbAssetsInfo{
		Id:              &result.ID,
		CreatedAt:       pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:       pointy.GetPointer(result.UpdatedAt.Unix()),
		User:            &result.User,
		Department:      &result.Department,
		Mobile:          &result.Mobile,
		Remark:          &result.Remark,
		Ip:              &result.IP,
		Mask:            &result.Mask,
		Gateway:         &result.Gateway,
		OnlineTime:      pointy.GetPointer(result.OnlineTime.UnixMilli()),
		OfflineTime:     pointy.GetPointer(result.OfflineTime.UnixMilli()),
		PowerStatus:     pointy.GetPointer(uint64(result.PowerStatus)),
		IsServer:        pointy.GetPointer(uint64(result.IsServer)),
		ServerType:      pointy.GetPointer(uint64(result.ServerType)),
		ServerHostname:  &result.ServerHostname,
		ServerOs:        &result.ServerOs,
		ServerOsVersion: &result.ServerOsVersion,
		ServerOsArch:    &result.ServerOsArch,
		Cpu:             &result.Cpu,
		Memory:          &result.Memory,
		Disk:            &result.Disk,
		NetworkSpeed:    &result.NetworkSpeed,
		DeviceAddress:   &result.DeviceAddress,
	}, nil
}
