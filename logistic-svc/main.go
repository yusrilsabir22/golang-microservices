package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/config"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/db"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/pb"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Logistic Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterLogisticServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
