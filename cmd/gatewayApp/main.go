package main

import (
	pkg "github.com/gatewayApp/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pkg.RegisterGatewayServiceServer(s, &pkg.GatewayService{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
