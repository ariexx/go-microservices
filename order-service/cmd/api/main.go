package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
	"net/http"
	"order-service/cmd/gapi"
	"order-service/data"
	"order-service/data/seeds"
	"order-service/pb"
	"os"
	"time"
)

type Config struct {
	db *gorm.DB
}

const grpcPort = ":9090"
const httpPort = ":80"

func main() {
	db := openDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error while connecting to database %s", err)
		}

		sqlDB.Close()
	}()

	//call data model
	orderRepo := data.NewOrderConfig(db)
	_ = orderRepo.AutoMigrate()

	err := seeds.Run(db, seeds.All())
	if err != nil {
		log.Fatalf("Error while running seeds %s", err)
	}

	go runGatewayServer(db)
	runGRPCServer(db)

}

func runGatewayServer(db *gorm.DB) {
	jsonOptions := runtime.WithMarshalerOption("application/json+pretty", &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			Indent:    "  ",
			Multiline: true, // Optional, implied by presence of "Indent".
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	newServerGapi := gapi.NewServer(db)
	grpcMux := runtime.NewServeMux(jsonOptions)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := pb.RegisterOrderServiceHandlerServer(ctx, grpcMux, newServerGapi)
	if err != nil {
		log.Fatalf("Error while registering gateway server %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", httpPort)
	if err != nil {
		log.Fatalf("Error while listening gateway to port %s", err)
	}

	log.Println("Starting HTTP server at port " + httpPort)

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("Error while serving gateway HTTP %s", err)
	}
}

func runGRPCServer(db *gorm.DB) {
	newServerGapi := gapi.NewServer(db)
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Error while listening GRPC to port %s", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, newServerGapi)
	reflection.Register(grpcServer)
	log.Println("Starting GRPC server at port " + grpcPort)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error while serving GRPC %s", err)
	}
}

func openDB() *gorm.DB {

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
	}

	//config database
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
