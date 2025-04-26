package grpc

import (
	"context"
	"fmt"

	userpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/user"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"
)

type UserGRPCServer struct {
	userpb.UnimplementedUserServiceServer
	Service services.AddressServiceInterface
}

func NewUserGRPCServer(service services.AddressServiceInterface) *UserGRPCServer {
	return &UserGRPCServer{Service: service}
}

func (s *UserGRPCServer) GetMainAddress(ctx context.Context, req *userpb.GetMainAddressRequest) (*userpb.AddressResponse, error) {
	userID := req.GetUserId()

	addresses, err := s.Service.GetAddresses(userID)
	if err != nil || len(addresses) == 0 {
		return nil, fmt.Errorf("no addresses found")
	}

	for _, addr := range addresses {
		if addr.IsMain {
			return &userpb.AddressResponse{
				Id:            addr.ID.String(),
				Name:          addr.Name,
				Address:       addr.Address,
				ProvinceId:    uint32(addr.ProvinceID),
				Province:      addr.Province,
				CityId:        uint32(addr.CityID),
				City:          addr.City,
				DistrictId:    uint32(addr.DistrictID),
				District:      addr.District,
				SubdistrictId: uint32(addr.SubdistrictID),
				Subdistrict:   addr.Subdistrict,
				PostalCodeId:  uint32(addr.PostalCodeID),
				PostalCode:    addr.PostalCode,
				Phone:         addr.Phone,
				IsMain:        addr.IsMain,
			}, nil
		}
	}

	return nil, fmt.Errorf("no main address set")
}
