package grpcutils

import (
	"context"
	"log"
	"net"

	"github.com/wunicorns/goutils/grpcutils" // Replace with your actual package name
	"google.golang.org/grpc"
)

// Implement the interface myPkgName.InvoicerServer
type myGRPCServer struct {
	grpcutils.UnimplementedInvoicerServer
}

// Implement the Create method
func (m *myGRPCServer) Create(ctx context.Context, request *grpcutils.CreateRequest) (*grpcutils.CreateResponse, error) {
	log.Println("Create called")
	// You can customize the response here
	return &grpcutils.CreateResponse{
		Pdf: []byte("TODO"), // Replace with actual data
	}, nil
}

func ServerMain() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myInvoicerServer := &myGRPCServer{}
	grpcutils.RegisterInvoicerServer(s, myInvoicerServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
