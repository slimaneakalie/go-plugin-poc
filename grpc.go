package main

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/slimaneakalie/go-plugin-poc/proto"
	"google.golang.org/grpc"
)

type DestinationPlugin interface {
	BatchUpsertUsers(request *UpsertUsersRequest) UpsertUsersResponse
}

type DestPluginImpl struct {
	// GRPCPlugin must still implement the DestPluginImpl interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl DestinationPlugin
}

func (d *DestPluginImpl) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterDestinationPluginServer(s, &GRPCServer{Impl: d.Impl})
	return nil
}

type GRPCServer struct {
	proto.UnimplementedDestinationPluginServer
	// This is the real implementation
	Impl DestinationPlugin
}

func (s *GRPCServer) BatchUpsertUsers(ctx context.Context, request *proto.BatchUpsertUsersRequest) (*proto.BatchUpsertUsersResponse, error) {
	var users []*User
	for _, u := range request.Users {
		users = append(users, &User{
			Id:        u.Id,
			FullName:  u.FullName,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
		})
	}
	req := &UpsertUsersRequest{
		Users:    users,
		Mappings: request.Mappings,
		Settings: request.Settings,
	}

	resp := s.Impl.BatchUpsertUsers(req)
	var errors []*proto.ResponseError
	for _, err := range resp.Errors {
		errors = append(errors, &proto.ResponseError{
			Code:    err.Code,
			UserId:  err.UserId,
			Message: err.Message,
		})
	}

	return &proto.BatchUpsertUsersResponse{
		Status: resp.Status,
		Errors: errors,
	}, nil
}

func (s *GRPCServer) BatchUpdateUsers(ctx context.Context, request *proto.BatchUpdateUsersRequest) (*proto.BatchUpdateUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) BatchDeleteUsers(ctx context.Context, request *proto.BatchDeleteUsersRequest) (*proto.BatchDeleteUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}
