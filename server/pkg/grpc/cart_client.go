package grpc

import (
	"context"
	"time"

	cartpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CartGRPCClient struct {
	client cartpb.CartServiceClient
}

func NewCartGRPCClient(address string) (*CartGRPCClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &CartGRPCClient{client: cartpb.NewCartServiceClient(conn)}, nil
}

func (c *CartGRPCClient) GetCart(userID string) ([]*cartpb.CartItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.client.GetCartForCheckout(ctx, &cartpb.GetCartRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}

func (c *CartGRPCClient) ClearCart(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := c.client.ClearCart(ctx, &cartpb.ClearCartRequest{
		UserId: userID,
	})
	return err
}
