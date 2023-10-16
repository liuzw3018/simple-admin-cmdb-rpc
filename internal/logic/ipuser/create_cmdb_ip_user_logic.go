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

type CreateCmdbIpUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCmdbIpUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCmdbIpUserLogic {
	return &CreateCmdbIpUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCmdbIpUserLogic) CreateCmdbIpUser(in *cmdb.CmdbIpUserInfo) (*cmdb.BaseIDResp, error) {
	query := l.svcCtx.DB.CmdbIpUser.Create().
		SetNotNilUser(in.User).
		SetNotNilDepartment(in.Department).
		SetNotNilMobile(in.Mobile).
		SetNotNilRemark(in.Remark)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &cmdb.BaseIDResp{Id: result.ID, Msg: errormsg.CreateSuccess}, nil
}
