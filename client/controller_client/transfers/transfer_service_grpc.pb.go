// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: transfer_service.proto

package transfers

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
	Transfers_Create_Transfer_FullMethodName     = "/transfer.Transfers/Create_Transfer"
	Transfers_Create_Deposit_FullMethodName      = "/transfer.Transfers/Create_Deposit"
	Transfers_Create_Withdraw_FullMethodName     = "/transfer.Transfers/Create_Withdraw"
	Transfers_GetTransfer_ById_FullMethodName    = "/transfer.Transfers/GetTransfer_ById"
	Transfers_GetTransfer_ByOwner_FullMethodName = "/transfer.Transfers/GetTransfer_ByOwner"
	Transfers_ListStatement_FullMethodName       = "/transfer.Transfers/ListStatement"
)

// TransfersClient is the client API for Transfers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransfersClient interface {
	Create_Transfer(ctx context.Context, in *Create_TransferRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Create_Deposit(ctx context.Context, in *Create_DepositRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Create_Withdraw(ctx context.Context, in *Create_WithdrawRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetTransfer_ById(ctx context.Context, in *GetTransfer_ByIdRequest, opts ...grpc.CallOption) (*GetTransfer_Response, error)
	GetTransfer_ByOwner(ctx context.Context, in *GetTransfer_ByOwnerRequest, opts ...grpc.CallOption) (*GetTransfer_Response, error)
	ListStatement(ctx context.Context, in *ListStatementRequest, opts ...grpc.CallOption) (*ListStatementResponse, error)
}

type transfersClient struct {
	cc grpc.ClientConnInterface
}

func NewTransfersClient(cc grpc.ClientConnInterface) TransfersClient {
	return &transfersClient{cc}
}

func (c *transfersClient) Create_Transfer(ctx context.Context, in *Create_TransferRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Transfers_Create_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) Create_Deposit(ctx context.Context, in *Create_DepositRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Transfers_Create_Deposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) Create_Withdraw(ctx context.Context, in *Create_WithdrawRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Transfers_Create_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) GetTransfer_ById(ctx context.Context, in *GetTransfer_ByIdRequest, opts ...grpc.CallOption) (*GetTransfer_Response, error) {
	out := new(GetTransfer_Response)
	err := c.cc.Invoke(ctx, Transfers_GetTransfer_ById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) GetTransfer_ByOwner(ctx context.Context, in *GetTransfer_ByOwnerRequest, opts ...grpc.CallOption) (*GetTransfer_Response, error) {
	out := new(GetTransfer_Response)
	err := c.cc.Invoke(ctx, Transfers_GetTransfer_ByOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) ListStatement(ctx context.Context, in *ListStatementRequest, opts ...grpc.CallOption) (*ListStatementResponse, error) {
	out := new(ListStatementResponse)
	err := c.cc.Invoke(ctx, Transfers_ListStatement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfersServer is the server API for Transfers service.
// All implementations must embed UnimplementedTransfersServer
// for forward compatibility
type TransfersServer interface {
	Create_Transfer(context.Context, *Create_TransferRequest) (*StatusResponse, error)
	Create_Deposit(context.Context, *Create_DepositRequest) (*StatusResponse, error)
	Create_Withdraw(context.Context, *Create_WithdrawRequest) (*StatusResponse, error)
	GetTransfer_ById(context.Context, *GetTransfer_ByIdRequest) (*GetTransfer_Response, error)
	GetTransfer_ByOwner(context.Context, *GetTransfer_ByOwnerRequest) (*GetTransfer_Response, error)
	ListStatement(context.Context, *ListStatementRequest) (*ListStatementResponse, error)
	mustEmbedUnimplementedTransfersServer()
}

// UnimplementedTransfersServer must be embedded to have forward compatible implementations.
type UnimplementedTransfersServer struct {
}

func (UnimplementedTransfersServer) Create_Transfer(context.Context, *Create_TransferRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create_Transfer not implemented")
}
func (UnimplementedTransfersServer) Create_Deposit(context.Context, *Create_DepositRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create_Deposit not implemented")
}
func (UnimplementedTransfersServer) Create_Withdraw(context.Context, *Create_WithdrawRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create_Withdraw not implemented")
}
func (UnimplementedTransfersServer) GetTransfer_ById(context.Context, *GetTransfer_ByIdRequest) (*GetTransfer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfer_ById not implemented")
}
func (UnimplementedTransfersServer) GetTransfer_ByOwner(context.Context, *GetTransfer_ByOwnerRequest) (*GetTransfer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfer_ByOwner not implemented")
}
func (UnimplementedTransfersServer) ListStatement(context.Context, *ListStatementRequest) (*ListStatementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatement not implemented")
}
func (UnimplementedTransfersServer) mustEmbedUnimplementedTransfersServer() {}

// UnsafeTransfersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransfersServer will
// result in compilation errors.
type UnsafeTransfersServer interface {
	mustEmbedUnimplementedTransfersServer()
}

func RegisterTransfersServer(s grpc.ServiceRegistrar, srv TransfersServer) {
	s.RegisterService(&Transfers_ServiceDesc, srv)
}

func _Transfers_Create_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Create_TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).Create_Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfers_Create_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).Create_Transfer(ctx, req.(*Create_TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_Create_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Create_DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).Create_Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfers_Create_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).Create_Deposit(ctx, req.(*Create_DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_Create_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Create_WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).Create_Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfers_Create_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).Create_Withdraw(ctx, req.(*Create_WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_GetTransfer_ById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransfer_ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).GetTransfer_ById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfers_GetTransfer_ById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).GetTransfer_ById(ctx, req.(*GetTransfer_ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_GetTransfer_ByOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransfer_ByOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).GetTransfer_ByOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfers_GetTransfer_ByOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).GetTransfer_ByOwner(ctx, req.(*GetTransfer_ByOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_ListStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).ListStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfers_ListStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).ListStatement(ctx, req.(*ListStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transfers_ServiceDesc is the grpc.ServiceDesc for Transfers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transfers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transfer.Transfers",
	HandlerType: (*TransfersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create_Transfer",
			Handler:    _Transfers_Create_Transfer_Handler,
		},
		{
			MethodName: "Create_Deposit",
			Handler:    _Transfers_Create_Deposit_Handler,
		},
		{
			MethodName: "Create_Withdraw",
			Handler:    _Transfers_Create_Withdraw_Handler,
		},
		{
			MethodName: "GetTransfer_ById",
			Handler:    _Transfers_GetTransfer_ById_Handler,
		},
		{
			MethodName: "GetTransfer_ByOwner",
			Handler:    _Transfers_GetTransfer_ByOwner_Handler,
		},
		{
			MethodName: "ListStatement",
			Handler:    _Transfers_ListStatement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer_service.proto",
}