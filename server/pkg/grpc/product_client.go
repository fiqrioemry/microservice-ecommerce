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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
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

func (p *ProductGRPCClient) CheckAvailability(productID string) (*productpb.CheckAvailabilityResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := p.client.CheckProductAvailability(ctx, &productpb.CheckAvailabilityRequest{
		ProductId: productID,
	})
	if err != nil {
		log.Printf("Error calling CheckProductAvailability: %v", err)
		return nil, err
	}
	return resp, nil
}

func (p *ProductGRPCClient) GetMultipleSnapshots(productIDs []string) (*productpb.MultipleProductSnapshotResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	items := []*productpb.ProductItemInput{}
	for _, id := range productIDs {
		items = append(items, &productpb.ProductItemInput{
			ProductId: id,
		})
	}

	resp, err := p.client.GetMultipleProductSnapshots(ctx, &productpb.GetMultipleProductRequest{
		Items: items,
	})
	if err != nil {
		log.Printf("Error calling GetMultipleProductSnapshots: %v", err)
		return nil, err
	}
	return resp, nil
}

func (p *ProductGRPCClient) ReduceStock(items []*productpb.StockUpdateItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := p.client.UpdateProductStock(ctx, &productpb.UpdateStockRequest{
		Items: items,
	})
	return err
}
