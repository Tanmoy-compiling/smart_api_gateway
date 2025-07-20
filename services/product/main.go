package product

import (
	"context"
	"log"
	"net"

	pb "github.com/your-username/smart_api_gateway/proto/product"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"smart_api_gateway/pkg/common/database"
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	// Mock product data
	product := &pb.Product{
		Id:          req.ProductId,
		Name:        "Sample Product",
		Description: "This is a sample product",
		Price:       99.99,
	}

	return &pb.GetProductResponse{
		Product: product,
	}, nil
}

package main

func main() {
	// Initialize Gin
	r := gin.Default()

	// Initialize DB
	if err := database.InitDB(); err != nil {
		panic(err)
	}

	// Setup routes
	SetupRoutes(r)

	// Start server
	r.Run(":5002")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})

	log.Printf("Product service server starting on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
