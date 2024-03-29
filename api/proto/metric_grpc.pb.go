// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api/proto/metric.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Metrics_Update_FullMethodName = "/metrics.Metrics/Update"
)

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	Update(ctx context.Context, opts ...grpc.CallOption) (Metrics_UpdateClient, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) Update(ctx context.Context, opts ...grpc.CallOption) (Metrics_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[0], Metrics_Update_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsUpdateClient{stream}
	return x, nil
}

type Metrics_UpdateClient interface {
	Send(*Metric) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type metricsUpdateClient struct {
	grpc.ClientStream
}

func (x *metricsUpdateClient) Send(m *Metric) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsUpdateClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	Update(Metrics_UpdateServer) error
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) Update(Metrics_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).Update(&metricsUpdateServer{stream})
}

type Metrics_UpdateServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Metric, error)
	grpc.ServerStream
}

type metricsUpdateServer struct {
	grpc.ServerStream
}

func (x *metricsUpdateServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsUpdateServer) Recv() (*Metric, error) {
	m := new(Metric)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Update",
			Handler:       _Metrics_Update_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/metric.proto",
}
