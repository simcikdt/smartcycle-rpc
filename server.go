package main


//https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2

import (
	pb "aws.com/rpc/smartcycle"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

//server is used to implement smartcycle.SmartcycleServiceServer
type SmartcycleServiceServer struct{
	pb.UnimplementedSmartcycleServiceServer
}

//unused right now
const (
	port = ":50051"
)

func (sc *SmartcycleServiceServer) GetSensorData(ctx context.Context, u *pb.SensorRequest) (*pb.SensorData, error) {
return &pb.SensorData{
	Speed:                15,
	Heartrate:            115,
	Cadence:              40,
	Temperature:          75,
	XXX_NoUnkeyedLiteral: struct{}{},
	XXX_unrecognized:     nil,
	XXX_sizecache:        0,
} , nil
}

func (sc *SmartcycleServiceServer) SwitchHeadlight(ctx context.Context, s *pb.HeadlightRequest) (*pb.HeadlightData, error){
	return &pb.HeadlightData{
		On:                   true,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	} , nil
}

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := SmartcycleServiceServer{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Smartcycle service to the server
	pb.RegisterSmartcycleServiceServer(grpcServer, &s)

	log.Println("Starting gRPC server for Smartcycle....")

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}