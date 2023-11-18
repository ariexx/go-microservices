package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"product-service/cmd/gapi"
	"product-service/config"
	"product-service/pb"
)

const (
	grpcPort = ":9090"
	httpPort = ":80"
)

func main() {
	db := config.InitDatabase()

	runGRPCServer(db)
}

func runGRPCServer(db *gorm.DB) {
	newServerApi := gapi.NewServer(db)
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Error while listening to product service grpc port  %s", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, newServerApi)
	reflection.Register(grpcServer)
	log.Println("Starting GRPC server product service at port " + grpcPort)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error while serving GRPC Product Service %s", err)
	}
}
