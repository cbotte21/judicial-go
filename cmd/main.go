package main

import (
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal"
	"github.com/cbotte21/judicial-go/internal/schema"
	"github.com/cbotte21/judicial-go/pb"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/environment"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//Verify environment variables exist
	environment.VerifyEnvVariable("port")
	environment.VerifyEnvVariable("hive_internal_addr")

	//Get port
	port := environment.GetEnvVariable("port")

	//Setup tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", port)
	}
	grpcServer := grpc.NewServer()

	//Register handlers to attach
	hiveClient := hive.NewHiveServiceClient(getHiveConn())
	mongoBanClient := datastore.MongoClient[schema.Ban]{}
	err = mongoBanClient.Init()
	if err != nil {
		panic(err)
	}
	mongoUnbanClient := datastore.MongoClient[schema.Unban]{}
	err = mongoUnbanClient.Init()
	if err != nil {
		panic(err)
	}
	//Initialize judicial
	jury := internal.NewJudicial(&hiveClient, &mongoBanClient, &mongoUnbanClient)

	pb.RegisterJudicialServiceServer(grpcServer, &jury)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

func getHiveConn() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(environment.GetEnvVariable("hive_internal_addr"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}
