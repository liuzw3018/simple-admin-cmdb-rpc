// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package server

import (
	"context"

	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/logic/assets"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/logic/base"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/logic/ip"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/logic/ipuser"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/internal/svc"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/types/cmdb"
)

type CmdbServer struct {
	svcCtx *svc.ServiceContext
	cmdb.UnimplementedCmdbServer
}

func NewCmdbServer(svcCtx *svc.ServiceContext) *CmdbServer {
	return &CmdbServer{
		svcCtx: svcCtx,
	}
}

func (s *CmdbServer) InitDatabase(ctx context.Context, in *cmdb.Empty) (*cmdb.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// CmdbAssets management
func (s *CmdbServer) CreateCmdbAssets(ctx context.Context, in *cmdb.CmdbAssetsInfo) (*cmdb.BaseIDResp, error) {
	l := assets.NewCreateCmdbAssetsLogic(ctx, s.svcCtx)
	return l.CreateCmdbAssets(in)
}

func (s *CmdbServer) UpdateCmdbAssets(ctx context.Context, in *cmdb.CmdbAssetsInfo) (*cmdb.BaseResp, error) {
	l := assets.NewUpdateCmdbAssetsLogic(ctx, s.svcCtx)
	return l.UpdateCmdbAssets(in)
}

func (s *CmdbServer) GetCmdbAssetsList(ctx context.Context, in *cmdb.CmdbAssetsListReq) (*cmdb.CmdbAssetsListResp, error) {
	l := assets.NewGetCmdbAssetsListLogic(ctx, s.svcCtx)
	return l.GetCmdbAssetsList(in)
}

func (s *CmdbServer) GetCmdbAssetsById(ctx context.Context, in *cmdb.IDReq) (*cmdb.CmdbAssetsInfo, error) {
	l := assets.NewGetCmdbAssetsByIdLogic(ctx, s.svcCtx)
	return l.GetCmdbAssetsById(in)
}

func (s *CmdbServer) DeleteCmdbAssets(ctx context.Context, in *cmdb.IDsReq) (*cmdb.BaseResp, error) {
	l := assets.NewDeleteCmdbAssetsLogic(ctx, s.svcCtx)
	return l.DeleteCmdbAssets(in)
}

// CmdbIp management
func (s *CmdbServer) CreateCmdbIp(ctx context.Context, in *cmdb.CmdbIpInfo) (*cmdb.BaseIDResp, error) {
	l := ip.NewCreateCmdbIpLogic(ctx, s.svcCtx)
	return l.CreateCmdbIp(in)
}

func (s *CmdbServer) UpdateCmdbIp(ctx context.Context, in *cmdb.CmdbIpInfo) (*cmdb.BaseResp, error) {
	l := ip.NewUpdateCmdbIpLogic(ctx, s.svcCtx)
	return l.UpdateCmdbIp(in)
}

func (s *CmdbServer) GetCmdbIpList(ctx context.Context, in *cmdb.CmdbIpListReq) (*cmdb.CmdbIpListResp, error) {
	l := ip.NewGetCmdbIpListLogic(ctx, s.svcCtx)
	return l.GetCmdbIpList(in)
}

func (s *CmdbServer) GetCmdbIpById(ctx context.Context, in *cmdb.IDReq) (*cmdb.CmdbIpInfo, error) {
	l := ip.NewGetCmdbIpByIdLogic(ctx, s.svcCtx)
	return l.GetCmdbIpById(in)
}

func (s *CmdbServer) DeleteCmdbIp(ctx context.Context, in *cmdb.IDsReq) (*cmdb.BaseResp, error) {
	l := ip.NewDeleteCmdbIpLogic(ctx, s.svcCtx)
	return l.DeleteCmdbIp(in)
}

// CmdbIpUser management
func (s *CmdbServer) CreateCmdbIpUser(ctx context.Context, in *cmdb.CmdbIpUserInfo) (*cmdb.BaseIDResp, error) {
	l := ipuser.NewCreateCmdbIpUserLogic(ctx, s.svcCtx)
	return l.CreateCmdbIpUser(in)
}

func (s *CmdbServer) UpdateCmdbIpUser(ctx context.Context, in *cmdb.CmdbIpUserInfo) (*cmdb.BaseResp, error) {
	l := ipuser.NewUpdateCmdbIpUserLogic(ctx, s.svcCtx)
	return l.UpdateCmdbIpUser(in)
}

func (s *CmdbServer) GetCmdbIpUserList(ctx context.Context, in *cmdb.CmdbIpUserListReq) (*cmdb.CmdbIpUserListResp, error) {
	l := ipuser.NewGetCmdbIpUserListLogic(ctx, s.svcCtx)
	return l.GetCmdbIpUserList(in)
}

func (s *CmdbServer) GetCmdbIpUserById(ctx context.Context, in *cmdb.IDReq) (*cmdb.CmdbIpUserInfo, error) {
	l := ipuser.NewGetCmdbIpUserByIdLogic(ctx, s.svcCtx)
	return l.GetCmdbIpUserById(in)
}

func (s *CmdbServer) DeleteCmdbIpUser(ctx context.Context, in *cmdb.IDsReq) (*cmdb.BaseResp, error) {
	l := ipuser.NewDeleteCmdbIpUserLogic(ctx, s.svcCtx)
	return l.DeleteCmdbIpUser(in)
}
