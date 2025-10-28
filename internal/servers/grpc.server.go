package servers

import (
	grpcpb "go-api-swagger/internal/grpc"
	"go-api-swagger/internal/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(customerService *services.CustomerService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al abrir puerto gRPC: %v", err)
	}

	s := grpc.NewServer()
	grpcpb.RegisterCustomerServiceServer(s, grpcpb.NewCustomerServer(customerService))
	reflection.Register(s)

	log.Println("Servidor gRPC corriendo en :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error en gRPC server: %v", err)
	}
}
