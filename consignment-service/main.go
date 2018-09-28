package main

import (

	// Import gnereated protobuf code
	"fmt"
	"log"

	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/ozzadar/microservices/consignment-service/proto/consignment"
	vesselProto "github.com/ozzadar/microservices/vessel-service/proto/vessel"
)

const (
	defaulthost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaulthost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	// Init will parse the command line flags
	srv.Init()

	// Register Handler

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
