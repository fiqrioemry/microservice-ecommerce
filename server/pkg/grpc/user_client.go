package grpc

import (
	"context"
	"time"

	userpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGRPCClient struct {
	client userpb.UserServiceClient
}

func NewUserGRPCClient(address string) (*UserGRPCClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &UserGRPCClient{client: userpb.NewUserServiceClient(conn)}, nil
}

func (u *UserGRPCClient) GetMainAddress(userID string) (*userpb.AddressResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return u.client.GetMainAddress(ctx, &userpb.GetMainAddressRequest{UserId: userID})
}
