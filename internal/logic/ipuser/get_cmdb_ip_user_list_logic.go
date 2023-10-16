package ipuser

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbipuser"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/predicate"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdbIpUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdbIpUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdbIpUserListLogic {
	return &GetCmdbIpUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdbIpUserListLogic) GetCmdbIpUserList(in *cmdb.CmdbIpUserListReq) (*cmdb.CmdbIpUserListResp, error) {
	var predicates []predicate.CmdbIpUser
	if in.User != nil {
		predicates = append(predicates, cmdbipuser.UserContains(*in.User))
	}
	if in.Department != nil {
		predicates = append(predicates, cmdbipuser.DepartmentContains(*in.Department))
	}
	if in.Mobile != nil {
		predicates = append(predicates, cmdbipuser.MobileContains(*in.Mobile))
	}
	result, err := l.svcCtx.DB.CmdbIpUser.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &cmdb.CmdbIpUserListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &cmdb.CmdbIpUserInfo{
			Id:         &v.ID,
			CreatedAt:  pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:  pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:     pointy.GetPointer(uint32(v.Status)),
			User:       &v.User,
			Department: &v.Department,
			Mobile:     &v.Mobile,
			Remark:     &v.Remark,
		})
	}

	return resp, nil
}
