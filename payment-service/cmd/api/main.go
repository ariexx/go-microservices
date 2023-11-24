package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"payment-service/cmd/gapi"
	"payment-service/config"
	"payment-service/pb"
)

const grpcPort = ":9090"

func main() {
	db := config.InitDatabase()

	runGRPCServer(db)
}

func runGRPCServer(db *gorm.DB) {
	newServerApi := gapi.NewServer(db)
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Error while listening to payment service grpc port  %s", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, newServerApi)
	reflection.Register(grpcServer)
	log.Println("Starting GRPC server payment service at port " + grpcPort)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error while serving GRPC Payment Service %s", err)
	}

}
