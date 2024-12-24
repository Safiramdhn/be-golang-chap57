package main

import (
	"log"
	"net"

	pb "be-chap57/notification_service/proto"
	"be-chap57/service"

	"google.golang.org/grpc"
)

func main() {
	// Define the port to listen on.
	const port = ":50051"

	// Create a new gRPC server.
	grpcServer := grpc.NewServer()

	// Register the NotificationService with the gRPC server.
	notificationService := &service.NotificationService{}
	pb.RegisterNotificationServiceServer(grpcServer, notificationService)

	// Create a listener on the specified port.
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("gRPC server listening on %s", port)

	// Start the gRPC server.
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
