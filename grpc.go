package main

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/slimaneakalie/go-plugin-poc/pb/proto"
	"google.golang.org/grpc"
)

type DestinationPlugin interface {
	BatchUpsertUsers(request *UpsertUsersRequest) UpsertUsersResponse
}

type DestGRPCPlugin struct {
	// GRPCPlugin must still implement the DestGRPCPlugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl DestinationPlugin
}

func (d *DestGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterDestinationPluginServer(s, &GRPCServer{Impl: d.Impl})
	return nil
}

func (d *DestGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewDestinationPluginClient(c)}, nil
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

// GRPCClient is an implementation of DestinationPlugin that talks over RPC.
type GRPCClient struct {
	client proto.DestinationPluginClient
}

func (c *GRPCClient) BatchUpsertUsers(request *UpsertUsersRequest) UpsertUsersResponse {
	var users []*proto.User
	for _, u := range request.Users {
		users = append(users, &proto.User{
			Id:        u.Id,
			FullName:  u.FullName,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
		})
	}

	protoResp, _ := c.client.BatchUpsertUsers(context.Background(), &proto.BatchUpsertUsersRequest{
		Users:    users,
		Mappings: request.Mappings,
		Settings: request.Settings,
	})

	var errors []*ResponseError
	for _, err := range protoResp.Errors {
		errors = append(errors, &ResponseError{
			Code:    err.Code,
			UserId:  err.UserId,
			Message: err.Message,
		})
	}

	return UpsertUsersResponse{
		Status: protoResp.Status,
		Errors: errors,
	}
}
