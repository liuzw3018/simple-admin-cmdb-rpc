package ip

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdbIpByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdbIpByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdbIpByIdLogic {
	return &GetCmdbIpByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdbIpByIdLogic) GetCmdbIpById(in *cmdb.IDReq) (*cmdb.CmdbIpInfo, error) {
	result, err := l.svcCtx.DB.CmdbIp.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.CmdbIpInfo{
		Id:          &result.ID,
		CreatedAt:   pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:   pointy.GetPointer(result.UpdatedAt.Unix()),
		User:        &result.User,
		Department:  &result.Department,
		Mobile:      &result.Mobile,
		Remark:      &result.Remark,
		Ip:          &result.IP,
		Mask:        &result.Mask,
		Gateway:     &result.Gateway,
		OnlineTime:  pointy.GetPointer(result.OnlineTime.UnixMilli()),
		OfflineTime: pointy.GetPointer(result.OfflineTime.UnixMilli()),
		IsLeisure:   pointy.GetPointer(uint64(result.IsLeisure)),
		DeviceType:  &result.DeviceType,
	}, nil
}
