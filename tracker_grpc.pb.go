// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TrackerClient is the client API for Tracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type trackerClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackerClient(cc grpc.ClientConnInterface) TrackerClient {
	return &trackerClient{cc}
}

var trackerRegisterStreamDesc = &grpc.StreamDesc{
	StreamName: "Register",
}

func (c *trackerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/Tracker/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerService is the service API for Tracker service.
// Fields should be assigned to their respective handler implementations only before
// RegisterTrackerService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type TrackerService struct {
	Register func(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

func (s *TrackerService) register(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/Tracker/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterTrackerService registers a service implementation with a gRPC server.
func RegisterTrackerService(s grpc.ServiceRegistrar, srv *TrackerService) {
	srvCopy := *srv
	if srvCopy.Register == nil {
		srvCopy.Register = func(context.Context, *RegisterRequest) (*RegisterResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "Tracker",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Register",
				Handler:    srvCopy.register,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "tracker.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewTrackerService creates a new TrackerService containing the
// implemented methods of the Tracker service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewTrackerService(s interface{}) *TrackerService {
	ns := &TrackerService{}
	if h, ok := s.(interface {
		Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	}); ok {
		ns.Register = h.Register
	}
	return ns
}

// UnstableTrackerService is the service API for Tracker service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableTrackerService interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}