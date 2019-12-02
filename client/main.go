package main
import (
	pb "aws.com/rpc/smartcycle"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewSmartcycleServiceClient(conn)

	response, err := c.GetSensorData (context.Background(), &pb.SensorRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetSensorData: %s", err)
	}
	log.Printf("Temperature from the server: %d", response.Temperature )
}