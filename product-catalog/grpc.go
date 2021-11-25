package main

import (
	"context"
	"log"
	"net"
	pb "northwind-on-dapr/product-catalog/main/northwind.proto"

	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

type server struct {
	pb.UnimplementedProductApiServer
}

func (s *server) Initialize(user, password, host, dbname string) {
	connectDb(user, password, host, dbname)
}

func (s *server) Ping(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	log.Printf("Received: %v", "GetProducts")

	products := Products{}
	err := products.getProducts()
	if err != nil {
		log.Fatal(err)
	}

	var res = pb.GetProductsResponse{}

	for _, v := range products {
		res.Products = append(res.Products, &pb.Product{
			Id:          v.ID,
			ProductName: v.ProductName,
		})
	}

	return &res, nil
}

func (s *server) RunGrpcApp(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductApiServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
