package ipuser

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/utils/dberrorhandler"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdbIpUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdbIpUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdbIpUserByIdLogic {
	return &GetCmdbIpUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdbIpUserByIdLogic) GetCmdbIpUserById(in *cmdb.IDReq) (*cmdb.CmdbIpUserInfo, error) {
	result, err := l.svcCtx.DB.CmdbIpUser.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.CmdbIpUserInfo{
		Id:         &result.ID,
		CreatedAt:  pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:  pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:     pointy.GetPointer(uint32(result.Status)),
		User:       &result.User,
		Department: &result.Department,
		Mobile:     &result.Mobile,
		Remark:     &result.Remark,
	}, nil
}
