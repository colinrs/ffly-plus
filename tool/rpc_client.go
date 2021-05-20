package main

import (
	"log"

	pb "github.com/colinrs/ffly-plus/internal/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.

	r, err := c.SelectUser(context.Background(), &pb.SelectUserRequest{UserId: 1, UserName: "user_name_10"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s,%s", r.Message, r.Data)
}
