package grpc

import (
	"context"
	"log"
	"time"

	productpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductGRPCClient struct {
	client productpb.ProductServiceClient
}

func NewProductGRPCClient(address string) (*ProductGRPCClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := productpb.NewProductServiceClient(conn)
	return &ProductGRPCClient{client: client}, nil
}

func (p *ProductGRPCClient) GetProductSnapshot(productID, variantID string) (*productpb.ProductSnapshotResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resp, err := p.client.GetProductSnapshot(ctx, &productpb.GetProductRequest{
		ProductId: productID,
		VariantId: variantID,
	})

	if err != nil {
		log.Printf("Error calling GetProductSnapshot: %v", err)
		return nil, err
	}
	return resp, nil
}
