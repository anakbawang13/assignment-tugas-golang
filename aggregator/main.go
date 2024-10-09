package main

import (
	handler "assignment-tugas-golang/aggregator/handler/grpc"
	pb "assignment-tugas-golang/aggregator/proto/aggregator_service/v1"
	"assignment-tugas-golang/aggregator/router"
	"assignment-tugas-golang/aggregator/service"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Use grpc.NewClient to create a new gRPC connection
	connU, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to create gRPC client:", err)
	}
	defer connU.Close()

	connWT, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to create gRPC client:", err)
	}
	defer connWT.Close()

	userServiceClient := pb.NewUserServiceClient(connU)
	walletServiceClient := pb.NewWalletServiceClient(connWT)
	transactionServiceClient := pb.NewTransactionServiceClient(connWT)

	aggrService := service.NewAggregatorService(userServiceClient, walletServiceClient, transactionServiceClient)

	aggrHandler := handler.NewAggregatorHandler(aggrService)

	r := gin.Default()

	router.SetupRouter(r, aggrHandler)

	log.Fatal(r.Run(":8080"))
}
