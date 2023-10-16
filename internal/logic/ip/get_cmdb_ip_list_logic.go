package ip

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbip"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/predicate"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdbIpListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdbIpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdbIpListLogic {
	return &GetCmdbIpListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdbIpListLogic) GetCmdbIpList(in *cmdb.CmdbIpListReq) (*cmdb.CmdbIpListResp, error) {
	var predicates []predicate.CmdbIp
	if in.User != nil {
		predicates = append(predicates, cmdbip.UserContains(*in.User))
	}
	if in.Department != nil {
		predicates = append(predicates, cmdbip.DepartmentContains(*in.Department))
	}
	if in.Mobile != nil {
		predicates = append(predicates, cmdbip.MobileContains(*in.Mobile))
	}
	result, err := l.svcCtx.DB.CmdbIp.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &cmdb.CmdbIpListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &cmdb.CmdbIpInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			User:        &v.User,
			Department:  &v.Department,
			Mobile:      &v.Mobile,
			Remark:      &v.Remark,
			Ip:          &v.IP,
			Mask:        &v.Mask,
			Gateway:     &v.Gateway,
			OnlineTime:  pointy.GetPointer(v.OnlineTime.UnixMilli()),
			OfflineTime: pointy.GetPointer(v.OfflineTime.UnixMilli()),
			IsLeisure:   pointy.GetPointer(uint64(v.IsLeisure)),
			DeviceType:  &v.DeviceType,
		})
	}

	return resp, nil
}
