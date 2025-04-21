package grpc

import (
	"context"

	userpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/user"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"
	"github.com/google/uuid"
)

type UserGRPCServer struct {
	userpb.UnimplementedUserServiceServer
	Service services.AddressServiceInterface
}

func NewUserGRPCServer(service services.AddressServiceInterface) *UserGRPCServer {
	return &UserGRPCServer{Service: service}
}

func (s *UserGRPCServer) GetAddressByID(ctx context.Context, req *userpb.AddressRequest) (*userpb.AddressResponse, error) {
	addressID, err := uuid.Parse(req.GetAddressId())
	if err != nil {
		return nil, err
	}

	address, err := s.Service.GetAddressById(addressID.String())

	if err != nil {
		return nil, err
	}

	return &userpb.AddressResponse{
		Id:       address.ID.String(),
		Name:     address.Name,
		Address:  address.Address,
		City:     address.City,
		Province: address.Province,
		Zipcode:  address.Zipcode,
		Phone:    address.Phone,
	}, nil
}
