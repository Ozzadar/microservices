package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/ozzadar/microservices/user-service/proto/auth"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	// Automatically migrates the user struct
	// into databse columns/types etc. This will check for changes
	// and migrate them each time
	// this service is restarted
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("shippy.auth"),
		micro.Version("latest"),
	)

	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())

	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService, publisher})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
