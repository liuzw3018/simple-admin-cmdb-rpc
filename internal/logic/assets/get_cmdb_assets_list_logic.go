package assets

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbassets"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/predicate"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdbAssetsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdbAssetsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdbAssetsListLogic {
	return &GetCmdbAssetsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdbAssetsListLogic) GetCmdbAssetsList(in *cmdb.CmdbAssetsListReq) (*cmdb.CmdbAssetsListResp, error) {
	var predicates []predicate.CmdbAssets
	if in.User != nil {
		predicates = append(predicates, cmdbassets.UserContains(*in.User))
	}
	if in.Department != nil {
		predicates = append(predicates, cmdbassets.DepartmentContains(*in.Department))
	}
	if in.Mobile != nil {
		predicates = append(predicates, cmdbassets.MobileContains(*in.Mobile))
	}
	result, err := l.svcCtx.DB.CmdbAssets.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &cmdb.CmdbAssetsListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &cmdb.CmdbAssetsInfo{
			Id:              &v.ID,
			CreatedAt:       pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:       pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			User:            &v.User,
			Department:      &v.Department,
			Mobile:          &v.Mobile,
			Remark:          &v.Remark,
			Ip:              &v.IP,
			Mask:            &v.Mask,
			Gateway:         &v.Gateway,
			OnlineTime:      pointy.GetPointer(v.OnlineTime.UnixMilli()),
			OfflineTime:     pointy.GetPointer(v.OfflineTime.UnixMilli()),
			PowerStatus:     pointy.GetPointer(uint64(v.PowerStatus)),
			IsServer:        pointy.GetPointer(uint64(v.IsServer)),
			ServerType:      pointy.GetPointer(uint64(v.ServerType)),
			ServerHostname:  &v.ServerHostname,
			ServerOs:        &v.ServerOs,
			ServerOsVersion: &v.ServerOsVersion,
			ServerOsArch:    &v.ServerOsArch,
			Cpu:             &v.Cpu,
			Memory:          &v.Memory,
			Disk:            &v.Disk,
			NetworkSpeed:    &v.NetworkSpeed,
			DeviceAddress:   &v.DeviceAddress,
		})
	}

	return resp, nil
}
