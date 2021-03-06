package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	pb "github.com/ozzadar/microservices/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {

	var consignment *pb.Consignment

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)

	return consignment, err
}

func main() {
	cmd.Init()

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response
	file := defaultFilename

	var token string

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecting file and token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create new context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.CreateConsignment(ctx, consignment)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})

	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
