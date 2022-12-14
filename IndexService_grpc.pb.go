// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: IndexService.proto

package indexserver_grpc_go

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

// IndexServiceClient is the client API for IndexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexServiceClient interface {
	AddEventLog(ctx context.Context, in *AddEventLogRequest, opts ...grpc.CallOption) (*OperResponse, error)
	SaveOrUpdateUserOnLineAndLocation(ctx context.Context, in *SaveOrUpdateUserOnLineAndLocationRequest, opts ...grpc.CallOption) (*OperResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*OperResponse, error)
	AddUserFragment(ctx context.Context, in *AddUserFragmentRequest, opts ...grpc.CallOption) (*OperResponse, error)
	DeleteFragment(ctx context.Context, in *DeleteFragmentRequest, opts ...grpc.CallOption) (*OperResponse, error)
	DeleteFragmentByUserId(ctx context.Context, in *DeleteFragmentByUserIdRequest, opts ...grpc.CallOption) (*OperResponse, error)
	QueryNearbyPerson(ctx context.Context, in *QueryNearbyPersonRequest, opts ...grpc.CallOption) (*SearchUserOnLineAndLocationResponse, error)
	QueryNearbyFragment(ctx context.Context, in *QueryNearbyFragmentRequest, opts ...grpc.CallOption) (*SearchUserFragmentResponse, error)
	QueryEventLog(ctx context.Context, in *QueryEventLogRequest, opts ...grpc.CallOption) (*SearchEventLogResponse, error)
}

type indexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexServiceClient(cc grpc.ClientConnInterface) IndexServiceClient {
	return &indexServiceClient{cc}
}

func (c *indexServiceClient) AddEventLog(ctx context.Context, in *AddEventLogRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/AddEventLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) SaveOrUpdateUserOnLineAndLocation(ctx context.Context, in *SaveOrUpdateUserOnLineAndLocationRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/SaveOrUpdateUserOnLineAndLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) AddUserFragment(ctx context.Context, in *AddUserFragmentRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/AddUserFragment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) DeleteFragment(ctx context.Context, in *DeleteFragmentRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/DeleteFragment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) DeleteFragmentByUserId(ctx context.Context, in *DeleteFragmentByUserIdRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/DeleteFragmentByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) QueryNearbyPerson(ctx context.Context, in *QueryNearbyPersonRequest, opts ...grpc.CallOption) (*SearchUserOnLineAndLocationResponse, error) {
	out := new(SearchUserOnLineAndLocationResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/QueryNearbyPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) QueryNearbyFragment(ctx context.Context, in *QueryNearbyFragmentRequest, opts ...grpc.CallOption) (*SearchUserFragmentResponse, error) {
	out := new(SearchUserFragmentResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/QueryNearbyFragment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) QueryEventLog(ctx context.Context, in *QueryEventLogRequest, opts ...grpc.CallOption) (*SearchEventLogResponse, error) {
	out := new(SearchEventLogResponse)
	err := c.cc.Invoke(ctx, "/indexserver_grpc.IndexService/QueryEventLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServiceServer is the server API for IndexService service.
// All implementations must embed UnimplementedIndexServiceServer
// for forward compatibility
type IndexServiceServer interface {
	AddEventLog(context.Context, *AddEventLogRequest) (*OperResponse, error)
	SaveOrUpdateUserOnLineAndLocation(context.Context, *SaveOrUpdateUserOnLineAndLocationRequest) (*OperResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*OperResponse, error)
	AddUserFragment(context.Context, *AddUserFragmentRequest) (*OperResponse, error)
	DeleteFragment(context.Context, *DeleteFragmentRequest) (*OperResponse, error)
	DeleteFragmentByUserId(context.Context, *DeleteFragmentByUserIdRequest) (*OperResponse, error)
	QueryNearbyPerson(context.Context, *QueryNearbyPersonRequest) (*SearchUserOnLineAndLocationResponse, error)
	QueryNearbyFragment(context.Context, *QueryNearbyFragmentRequest) (*SearchUserFragmentResponse, error)
	QueryEventLog(context.Context, *QueryEventLogRequest) (*SearchEventLogResponse, error)
	mustEmbedUnimplementedIndexServiceServer()
}

// UnimplementedIndexServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIndexServiceServer struct {
}

func (UnimplementedIndexServiceServer) AddEventLog(context.Context, *AddEventLogRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEventLog not implemented")
}
func (UnimplementedIndexServiceServer) SaveOrUpdateUserOnLineAndLocation(context.Context, *SaveOrUpdateUserOnLineAndLocationRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveOrUpdateUserOnLineAndLocation not implemented")
}
func (UnimplementedIndexServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedIndexServiceServer) AddUserFragment(context.Context, *AddUserFragmentRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserFragment not implemented")
}
func (UnimplementedIndexServiceServer) DeleteFragment(context.Context, *DeleteFragmentRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFragment not implemented")
}
func (UnimplementedIndexServiceServer) DeleteFragmentByUserId(context.Context, *DeleteFragmentByUserIdRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFragmentByUserId not implemented")
}
func (UnimplementedIndexServiceServer) QueryNearbyPerson(context.Context, *QueryNearbyPersonRequest) (*SearchUserOnLineAndLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNearbyPerson not implemented")
}
func (UnimplementedIndexServiceServer) QueryNearbyFragment(context.Context, *QueryNearbyFragmentRequest) (*SearchUserFragmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNearbyFragment not implemented")
}
func (UnimplementedIndexServiceServer) QueryEventLog(context.Context, *QueryEventLogRequest) (*SearchEventLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEventLog not implemented")
}
func (UnimplementedIndexServiceServer) mustEmbedUnimplementedIndexServiceServer() {}

// UnsafeIndexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexServiceServer will
// result in compilation errors.
type UnsafeIndexServiceServer interface {
	mustEmbedUnimplementedIndexServiceServer()
}

func RegisterIndexServiceServer(s grpc.ServiceRegistrar, srv IndexServiceServer) {
	s.RegisterService(&IndexService_ServiceDesc, srv)
}

func _IndexService_AddEventLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).AddEventLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/AddEventLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).AddEventLog(ctx, req.(*AddEventLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_SaveOrUpdateUserOnLineAndLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveOrUpdateUserOnLineAndLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).SaveOrUpdateUserOnLineAndLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/SaveOrUpdateUserOnLineAndLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).SaveOrUpdateUserOnLineAndLocation(ctx, req.(*SaveOrUpdateUserOnLineAndLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_AddUserFragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserFragmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).AddUserFragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/AddUserFragment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).AddUserFragment(ctx, req.(*AddUserFragmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_DeleteFragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFragmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DeleteFragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/DeleteFragment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DeleteFragment(ctx, req.(*DeleteFragmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_DeleteFragmentByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFragmentByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DeleteFragmentByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/DeleteFragmentByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DeleteFragmentByUserId(ctx, req.(*DeleteFragmentByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_QueryNearbyPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNearbyPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).QueryNearbyPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/QueryNearbyPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).QueryNearbyPerson(ctx, req.(*QueryNearbyPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_QueryNearbyFragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNearbyFragmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).QueryNearbyFragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/QueryNearbyFragment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).QueryNearbyFragment(ctx, req.(*QueryNearbyFragmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_QueryEventLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEventLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).QueryEventLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexserver_grpc.IndexService/QueryEventLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).QueryEventLog(ctx, req.(*QueryEventLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndexService_ServiceDesc is the grpc.ServiceDesc for IndexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indexserver_grpc.IndexService",
	HandlerType: (*IndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEventLog",
			Handler:    _IndexService_AddEventLog_Handler,
		},
		{
			MethodName: "SaveOrUpdateUserOnLineAndLocation",
			Handler:    _IndexService_SaveOrUpdateUserOnLineAndLocation_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _IndexService_DeleteUser_Handler,
		},
		{
			MethodName: "AddUserFragment",
			Handler:    _IndexService_AddUserFragment_Handler,
		},
		{
			MethodName: "DeleteFragment",
			Handler:    _IndexService_DeleteFragment_Handler,
		},
		{
			MethodName: "DeleteFragmentByUserId",
			Handler:    _IndexService_DeleteFragmentByUserId_Handler,
		},
		{
			MethodName: "QueryNearbyPerson",
			Handler:    _IndexService_QueryNearbyPerson_Handler,
		},
		{
			MethodName: "QueryNearbyFragment",
			Handler:    _IndexService_QueryNearbyFragment_Handler,
		},
		{
			MethodName: "QueryEventLog",
			Handler:    _IndexService_QueryEventLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "IndexService.proto",
}
