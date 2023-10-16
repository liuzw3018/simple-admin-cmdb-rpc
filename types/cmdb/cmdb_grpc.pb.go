// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: cmdb.proto

package cmdb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Cmdb_InitDatabase_FullMethodName      = "/cmdb.Cmdb/initDatabase"
	Cmdb_CreateCmdbAssets_FullMethodName  = "/cmdb.Cmdb/createCmdbAssets"
	Cmdb_UpdateCmdbAssets_FullMethodName  = "/cmdb.Cmdb/updateCmdbAssets"
	Cmdb_GetCmdbAssetsList_FullMethodName = "/cmdb.Cmdb/getCmdbAssetsList"
	Cmdb_GetCmdbAssetsById_FullMethodName = "/cmdb.Cmdb/getCmdbAssetsById"
	Cmdb_DeleteCmdbAssets_FullMethodName  = "/cmdb.Cmdb/deleteCmdbAssets"
	Cmdb_CreateCmdbIp_FullMethodName      = "/cmdb.Cmdb/createCmdbIp"
	Cmdb_UpdateCmdbIp_FullMethodName      = "/cmdb.Cmdb/updateCmdbIp"
	Cmdb_GetCmdbIpList_FullMethodName     = "/cmdb.Cmdb/getCmdbIpList"
	Cmdb_GetCmdbIpById_FullMethodName     = "/cmdb.Cmdb/getCmdbIpById"
	Cmdb_DeleteCmdbIp_FullMethodName      = "/cmdb.Cmdb/deleteCmdbIp"
	Cmdb_CreateCmdbIpUser_FullMethodName  = "/cmdb.Cmdb/createCmdbIpUser"
	Cmdb_UpdateCmdbIpUser_FullMethodName  = "/cmdb.Cmdb/updateCmdbIpUser"
	Cmdb_GetCmdbIpUserList_FullMethodName = "/cmdb.Cmdb/getCmdbIpUserList"
	Cmdb_GetCmdbIpUserById_FullMethodName = "/cmdb.Cmdb/getCmdbIpUserById"
	Cmdb_DeleteCmdbIpUser_FullMethodName  = "/cmdb.Cmdb/deleteCmdbIpUser"
)

// CmdbClient is the client API for Cmdb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmdbClient interface {
	// group: base
	InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
	// CmdbAssets management
	// group: assets
	CreateCmdbAssets(ctx context.Context, in *CmdbAssetsInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
	// group: assets
	UpdateCmdbAssets(ctx context.Context, in *CmdbAssetsInfo, opts ...grpc.CallOption) (*BaseResp, error)
	// group: assets
	GetCmdbAssetsList(ctx context.Context, in *CmdbAssetsListReq, opts ...grpc.CallOption) (*CmdbAssetsListResp, error)
	// group: assets
	GetCmdbAssetsById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CmdbAssetsInfo, error)
	// group: assets
	DeleteCmdbAssets(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	// CmdbIp management
	// group: ip
	CreateCmdbIp(ctx context.Context, in *CmdbIpInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
	// group: ip
	UpdateCmdbIp(ctx context.Context, in *CmdbIpInfo, opts ...grpc.CallOption) (*BaseResp, error)
	// group: ip
	GetCmdbIpList(ctx context.Context, in *CmdbIpListReq, opts ...grpc.CallOption) (*CmdbIpListResp, error)
	// group: ip
	GetCmdbIpById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CmdbIpInfo, error)
	// group: ip
	DeleteCmdbIp(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	// CmdbIpUser management
	// group: ipuser
	CreateCmdbIpUser(ctx context.Context, in *CmdbIpUserInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
	// group: ipuser
	UpdateCmdbIpUser(ctx context.Context, in *CmdbIpUserInfo, opts ...grpc.CallOption) (*BaseResp, error)
	// group: ipuser
	GetCmdbIpUserList(ctx context.Context, in *CmdbIpUserListReq, opts ...grpc.CallOption) (*CmdbIpUserListResp, error)
	// group: ipuser
	GetCmdbIpUserById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CmdbIpUserInfo, error)
	// group: ipuser
	DeleteCmdbIpUser(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
}

type cmdbClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdbClient(cc grpc.ClientConnInterface) CmdbClient {
	return &cmdbClient{cc}
}

func (c *cmdbClient) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_InitDatabase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) CreateCmdbAssets(ctx context.Context, in *CmdbAssetsInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	out := new(BaseIDResp)
	err := c.cc.Invoke(ctx, Cmdb_CreateCmdbAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) UpdateCmdbAssets(ctx context.Context, in *CmdbAssetsInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_UpdateCmdbAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) GetCmdbAssetsList(ctx context.Context, in *CmdbAssetsListReq, opts ...grpc.CallOption) (*CmdbAssetsListResp, error) {
	out := new(CmdbAssetsListResp)
	err := c.cc.Invoke(ctx, Cmdb_GetCmdbAssetsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) GetCmdbAssetsById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CmdbAssetsInfo, error) {
	out := new(CmdbAssetsInfo)
	err := c.cc.Invoke(ctx, Cmdb_GetCmdbAssetsById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) DeleteCmdbAssets(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_DeleteCmdbAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) CreateCmdbIp(ctx context.Context, in *CmdbIpInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	out := new(BaseIDResp)
	err := c.cc.Invoke(ctx, Cmdb_CreateCmdbIp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) UpdateCmdbIp(ctx context.Context, in *CmdbIpInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_UpdateCmdbIp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) GetCmdbIpList(ctx context.Context, in *CmdbIpListReq, opts ...grpc.CallOption) (*CmdbIpListResp, error) {
	out := new(CmdbIpListResp)
	err := c.cc.Invoke(ctx, Cmdb_GetCmdbIpList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) GetCmdbIpById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CmdbIpInfo, error) {
	out := new(CmdbIpInfo)
	err := c.cc.Invoke(ctx, Cmdb_GetCmdbIpById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) DeleteCmdbIp(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_DeleteCmdbIp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) CreateCmdbIpUser(ctx context.Context, in *CmdbIpUserInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	out := new(BaseIDResp)
	err := c.cc.Invoke(ctx, Cmdb_CreateCmdbIpUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) UpdateCmdbIpUser(ctx context.Context, in *CmdbIpUserInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_UpdateCmdbIpUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) GetCmdbIpUserList(ctx context.Context, in *CmdbIpUserListReq, opts ...grpc.CallOption) (*CmdbIpUserListResp, error) {
	out := new(CmdbIpUserListResp)
	err := c.cc.Invoke(ctx, Cmdb_GetCmdbIpUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) GetCmdbIpUserById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CmdbIpUserInfo, error) {
	out := new(CmdbIpUserInfo)
	err := c.cc.Invoke(ctx, Cmdb_GetCmdbIpUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) DeleteCmdbIpUser(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, Cmdb_DeleteCmdbIpUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdbServer is the server API for Cmdb service.
// All implementations must embed UnimplementedCmdbServer
// for forward compatibility
type CmdbServer interface {
	// group: base
	InitDatabase(context.Context, *Empty) (*BaseResp, error)
	// CmdbAssets management
	// group: assets
	CreateCmdbAssets(context.Context, *CmdbAssetsInfo) (*BaseIDResp, error)
	// group: assets
	UpdateCmdbAssets(context.Context, *CmdbAssetsInfo) (*BaseResp, error)
	// group: assets
	GetCmdbAssetsList(context.Context, *CmdbAssetsListReq) (*CmdbAssetsListResp, error)
	// group: assets
	GetCmdbAssetsById(context.Context, *IDReq) (*CmdbAssetsInfo, error)
	// group: assets
	DeleteCmdbAssets(context.Context, *IDsReq) (*BaseResp, error)
	// CmdbIp management
	// group: ip
	CreateCmdbIp(context.Context, *CmdbIpInfo) (*BaseIDResp, error)
	// group: ip
	UpdateCmdbIp(context.Context, *CmdbIpInfo) (*BaseResp, error)
	// group: ip
	GetCmdbIpList(context.Context, *CmdbIpListReq) (*CmdbIpListResp, error)
	// group: ip
	GetCmdbIpById(context.Context, *IDReq) (*CmdbIpInfo, error)
	// group: ip
	DeleteCmdbIp(context.Context, *IDsReq) (*BaseResp, error)
	// CmdbIpUser management
	// group: ipuser
	CreateCmdbIpUser(context.Context, *CmdbIpUserInfo) (*BaseIDResp, error)
	// group: ipuser
	UpdateCmdbIpUser(context.Context, *CmdbIpUserInfo) (*BaseResp, error)
	// group: ipuser
	GetCmdbIpUserList(context.Context, *CmdbIpUserListReq) (*CmdbIpUserListResp, error)
	// group: ipuser
	GetCmdbIpUserById(context.Context, *IDReq) (*CmdbIpUserInfo, error)
	// group: ipuser
	DeleteCmdbIpUser(context.Context, *IDsReq) (*BaseResp, error)
	mustEmbedUnimplementedCmdbServer()
}

// UnimplementedCmdbServer must be embedded to have forward compatible implementations.
type UnimplementedCmdbServer struct {
}

func (UnimplementedCmdbServer) InitDatabase(context.Context, *Empty) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDatabase not implemented")
}
func (UnimplementedCmdbServer) CreateCmdbAssets(context.Context, *CmdbAssetsInfo) (*BaseIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCmdbAssets not implemented")
}
func (UnimplementedCmdbServer) UpdateCmdbAssets(context.Context, *CmdbAssetsInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCmdbAssets not implemented")
}
func (UnimplementedCmdbServer) GetCmdbAssetsList(context.Context, *CmdbAssetsListReq) (*CmdbAssetsListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmdbAssetsList not implemented")
}
func (UnimplementedCmdbServer) GetCmdbAssetsById(context.Context, *IDReq) (*CmdbAssetsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmdbAssetsById not implemented")
}
func (UnimplementedCmdbServer) DeleteCmdbAssets(context.Context, *IDsReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCmdbAssets not implemented")
}
func (UnimplementedCmdbServer) CreateCmdbIp(context.Context, *CmdbIpInfo) (*BaseIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCmdbIp not implemented")
}
func (UnimplementedCmdbServer) UpdateCmdbIp(context.Context, *CmdbIpInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCmdbIp not implemented")
}
func (UnimplementedCmdbServer) GetCmdbIpList(context.Context, *CmdbIpListReq) (*CmdbIpListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmdbIpList not implemented")
}
func (UnimplementedCmdbServer) GetCmdbIpById(context.Context, *IDReq) (*CmdbIpInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmdbIpById not implemented")
}
func (UnimplementedCmdbServer) DeleteCmdbIp(context.Context, *IDsReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCmdbIp not implemented")
}
func (UnimplementedCmdbServer) CreateCmdbIpUser(context.Context, *CmdbIpUserInfo) (*BaseIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCmdbIpUser not implemented")
}
func (UnimplementedCmdbServer) UpdateCmdbIpUser(context.Context, *CmdbIpUserInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCmdbIpUser not implemented")
}
func (UnimplementedCmdbServer) GetCmdbIpUserList(context.Context, *CmdbIpUserListReq) (*CmdbIpUserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmdbIpUserList not implemented")
}
func (UnimplementedCmdbServer) GetCmdbIpUserById(context.Context, *IDReq) (*CmdbIpUserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmdbIpUserById not implemented")
}
func (UnimplementedCmdbServer) DeleteCmdbIpUser(context.Context, *IDsReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCmdbIpUser not implemented")
}
func (UnimplementedCmdbServer) mustEmbedUnimplementedCmdbServer() {}

// UnsafeCmdbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmdbServer will
// result in compilation errors.
type UnsafeCmdbServer interface {
	mustEmbedUnimplementedCmdbServer()
}

func RegisterCmdbServer(s grpc.ServiceRegistrar, srv CmdbServer) {
	s.RegisterService(&Cmdb_ServiceDesc, srv)
}

func _Cmdb_InitDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).InitDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_InitDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).InitDatabase(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_CreateCmdbAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbAssetsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).CreateCmdbAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_CreateCmdbAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).CreateCmdbAssets(ctx, req.(*CmdbAssetsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_UpdateCmdbAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbAssetsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).UpdateCmdbAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_UpdateCmdbAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).UpdateCmdbAssets(ctx, req.(*CmdbAssetsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_GetCmdbAssetsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbAssetsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).GetCmdbAssetsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_GetCmdbAssetsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).GetCmdbAssetsList(ctx, req.(*CmdbAssetsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_GetCmdbAssetsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).GetCmdbAssetsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_GetCmdbAssetsById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).GetCmdbAssetsById(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_DeleteCmdbAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).DeleteCmdbAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_DeleteCmdbAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).DeleteCmdbAssets(ctx, req.(*IDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_CreateCmdbIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbIpInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).CreateCmdbIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_CreateCmdbIp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).CreateCmdbIp(ctx, req.(*CmdbIpInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_UpdateCmdbIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbIpInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).UpdateCmdbIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_UpdateCmdbIp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).UpdateCmdbIp(ctx, req.(*CmdbIpInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_GetCmdbIpList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbIpListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).GetCmdbIpList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_GetCmdbIpList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).GetCmdbIpList(ctx, req.(*CmdbIpListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_GetCmdbIpById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).GetCmdbIpById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_GetCmdbIpById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).GetCmdbIpById(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_DeleteCmdbIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).DeleteCmdbIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_DeleteCmdbIp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).DeleteCmdbIp(ctx, req.(*IDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_CreateCmdbIpUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbIpUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).CreateCmdbIpUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_CreateCmdbIpUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).CreateCmdbIpUser(ctx, req.(*CmdbIpUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_UpdateCmdbIpUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbIpUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).UpdateCmdbIpUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_UpdateCmdbIpUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).UpdateCmdbIpUser(ctx, req.(*CmdbIpUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_GetCmdbIpUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdbIpUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).GetCmdbIpUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_GetCmdbIpUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).GetCmdbIpUserList(ctx, req.(*CmdbIpUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_GetCmdbIpUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).GetCmdbIpUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_GetCmdbIpUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).GetCmdbIpUserById(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_DeleteCmdbIpUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).DeleteCmdbIpUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cmdb_DeleteCmdbIpUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).DeleteCmdbIpUser(ctx, req.(*IDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cmdb_ServiceDesc is the grpc.ServiceDesc for Cmdb service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cmdb_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cmdb.Cmdb",
	HandlerType: (*CmdbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "initDatabase",
			Handler:    _Cmdb_InitDatabase_Handler,
		},
		{
			MethodName: "createCmdbAssets",
			Handler:    _Cmdb_CreateCmdbAssets_Handler,
		},
		{
			MethodName: "updateCmdbAssets",
			Handler:    _Cmdb_UpdateCmdbAssets_Handler,
		},
		{
			MethodName: "getCmdbAssetsList",
			Handler:    _Cmdb_GetCmdbAssetsList_Handler,
		},
		{
			MethodName: "getCmdbAssetsById",
			Handler:    _Cmdb_GetCmdbAssetsById_Handler,
		},
		{
			MethodName: "deleteCmdbAssets",
			Handler:    _Cmdb_DeleteCmdbAssets_Handler,
		},
		{
			MethodName: "createCmdbIp",
			Handler:    _Cmdb_CreateCmdbIp_Handler,
		},
		{
			MethodName: "updateCmdbIp",
			Handler:    _Cmdb_UpdateCmdbIp_Handler,
		},
		{
			MethodName: "getCmdbIpList",
			Handler:    _Cmdb_GetCmdbIpList_Handler,
		},
		{
			MethodName: "getCmdbIpById",
			Handler:    _Cmdb_GetCmdbIpById_Handler,
		},
		{
			MethodName: "deleteCmdbIp",
			Handler:    _Cmdb_DeleteCmdbIp_Handler,
		},
		{
			MethodName: "createCmdbIpUser",
			Handler:    _Cmdb_CreateCmdbIpUser_Handler,
		},
		{
			MethodName: "updateCmdbIpUser",
			Handler:    _Cmdb_UpdateCmdbIpUser_Handler,
		},
		{
			MethodName: "getCmdbIpUserList",
			Handler:    _Cmdb_GetCmdbIpUserList_Handler,
		},
		{
			MethodName: "getCmdbIpUserById",
			Handler:    _Cmdb_GetCmdbIpUserById_Handler,
		},
		{
			MethodName: "deleteCmdbIpUser",
			Handler:    _Cmdb_DeleteCmdbIpUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}