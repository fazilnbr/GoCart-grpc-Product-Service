package http

import (
	"fmt"
	"log"
	"net"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/api/services"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewGRPCServer(ProductService *services.ProductService, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, ProductService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServerHTTP(userHandler *services.ProductService) *ServerHTTP {
	engine := gin.New()
	go NewGRPCServer(userHandler, "50082")
	// Use logger from Gin
	engine.Use(gin.Logger())

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8001")
}
