package grpcutils

import (
	"log"

	"google.golang.org/grpc"
)

func ClientMain() {
	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcutils.NewInvoicerClient(conn)
	// Call your gRPC methods here
	// Example: response, err := client.Create(context.Background(), &grpcutils.CreateRequest{})
	log.Println(client)
}
