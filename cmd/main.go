package main

import (
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal"
	"github.com/cbotte21/judicial-go/internal/schema"
	pb "github.com/cbotte21/judicial-go/pb"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//Verify enviroment variables exist
	enviroment.VerifyEnvVariable("port")
	enviroment.VerifyEnvVariable("hive-internal_addr")

	//Get port
	port := enviroment.GetEnvVariable("port")

	//Setup tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", port)
	}
	grpcServer := grpc.NewServer()

	//Register handlers to attach
	hiveClient := hive.NewHiveServiceClient(getHiveConn(enviroment.GetEnvVariable("hive_port")))
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

func getHiveConn(port string) *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(enviroment.GetEnvVariable("hive-internal_addr"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}
