package grpcpb

import (
	"context"
	"go-api-swagger/internal/services"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

// CustomerServer implementa CustomerServiceServer
type CustomerServer struct {
	customerService *services.CustomerService
	UnimplementedCustomerServiceServer
}

// Constructor para inyectar dependencias
func NewCustomerServer(service *services.CustomerService) *CustomerServer {
	return &CustomerServer{customerService: service}
}

// Implementa el método definido en el .proto
func (s *CustomerServer) GetCustomerByID(ctx context.Context, req *GetCustomerRequest) (*GetCustomerResponse, error) {
	customer, err := s.customerService.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &GetCustomerResponse{
		Id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Username:  *customer.Username,
		Phone:     customer.Phone,
	}, nil
}

func (s *CustomerServer) StreamAllCustomers(_ *emptypb.Empty, stream CustomerService_StreamAllCustomersServer) error {
	customers, err := s.customerService.GetAll(context.Background())
	if err != nil {
		return err
	}

	for _, customer := range *customers {
		username := ""
		if customer.Username != nil {
			username = *customer.Username
		}

		resp := &GetCustomerResponse{
			Id:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Username:  username,
			Phone:     customer.Phone,
		}

		// Enviar cada cliente de forma individual
		if err := stream.Send(resp); err != nil {
			log.Printf("Error enviando cliente %d: %v", customer.ID, err)
			return err
		}
	}

	return nil
}
