package main

import (
	"context"
	"errors"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"

	// Import gnereated protobuf code
	"fmt"
	"log"

	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	pb "github.com/ozzadar/microservices/consignment-service/proto/consignment"
	userService "github.com/ozzadar/microservices/user-service/proto/auth"
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
		micro.Name("shippy.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselProto.NewVesselServiceClient("shippy.vessel", srv.Client())
	// Init will parse the command line flags
	srv.Init()

	// Register Handler

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewAuthClient("go.micro.srv.user", client.DefaultClient)

		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})

		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
