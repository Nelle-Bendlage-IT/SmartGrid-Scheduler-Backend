// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: api/protobuf/gsi_prediction.proto

package gsi_prediction

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

// LocalPricePredictionServiceClient is the client API for LocalPricePredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalPricePredictionServiceClient interface {
	GetGSIPrediction(ctx context.Context, in *GSIPrediction, opts ...grpc.CallOption) (*GetGSIPredictionResponse, error)
}

type localPricePredictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalPricePredictionServiceClient(cc grpc.ClientConnInterface) LocalPricePredictionServiceClient {
	return &localPricePredictionServiceClient{cc}
}

func (c *localPricePredictionServiceClient) GetGSIPrediction(ctx context.Context, in *GSIPrediction, opts ...grpc.CallOption) (*GetGSIPredictionResponse, error) {
	out := new(GetGSIPredictionResponse)
	err := c.cc.Invoke(ctx, "/gsi_predicition.LocalPricePredictionService/GetGSIPrediction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalPricePredictionServiceServer is the server API for LocalPricePredictionService service.
// All implementations must embed UnimplementedLocalPricePredictionServiceServer
// for forward compatibility
type LocalPricePredictionServiceServer interface {
	GetGSIPrediction(context.Context, *GSIPrediction) (*GetGSIPredictionResponse, error)
	mustEmbedUnimplementedLocalPricePredictionServiceServer()
}

// UnimplementedLocalPricePredictionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocalPricePredictionServiceServer struct {
}

func (UnimplementedLocalPricePredictionServiceServer) GetGSIPrediction(context.Context, *GSIPrediction) (*GetGSIPredictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGSIPrediction not implemented")
}
func (UnimplementedLocalPricePredictionServiceServer) mustEmbedUnimplementedLocalPricePredictionServiceServer() {
}

// UnsafeLocalPricePredictionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalPricePredictionServiceServer will
// result in compilation errors.
type UnsafeLocalPricePredictionServiceServer interface {
	mustEmbedUnimplementedLocalPricePredictionServiceServer()
}

func RegisterLocalPricePredictionServiceServer(s grpc.ServiceRegistrar, srv LocalPricePredictionServiceServer) {
	s.RegisterService(&LocalPricePredictionService_ServiceDesc, srv)
}

func _LocalPricePredictionService_GetGSIPrediction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSIPrediction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalPricePredictionServiceServer).GetGSIPrediction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsi_predicition.LocalPricePredictionService/GetGSIPrediction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalPricePredictionServiceServer).GetGSIPrediction(ctx, req.(*GSIPrediction))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalPricePredictionService_ServiceDesc is the grpc.ServiceDesc for LocalPricePredictionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalPricePredictionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsi_predicition.LocalPricePredictionService",
	HandlerType: (*LocalPricePredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGSIPrediction",
			Handler:    _LocalPricePredictionService_GetGSIPrediction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf/gsi_prediction.proto",
}