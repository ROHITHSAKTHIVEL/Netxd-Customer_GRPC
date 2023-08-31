package main

import (
	"Netxd_Project/Netxd_Customer_GRPC/config"
	"Netxd_Project/Netxd_Customer_GRPC/constants"
	"Netxd_Project/Netxd_Customer_GRPC/controller"
	"Netxd_Project/Netxd_DAL/services"
	"context"
	"fmt"
	"net"

	pro "Netxd_Project/Netxd_Customer/customer"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, "DemoBank", "Customer")
	controller.CustomerService = services.InitCustomerService(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controller.RPCserver{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
