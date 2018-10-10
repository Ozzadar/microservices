package main

import (
	"fmt"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/ozzadar/microservices/user-service/proto/user"
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

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}