package main

import (
	"encoding/json"
	"log"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	pb "github.com/ozzadar/microservices/user-service/proto/user"
)

const topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	// Get the broker instance using our env variables
	pubsub := srv.Server().Options().Broker

	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	_, err := pubsub.Subscribe(topic, func(p broker.Publication) error {
		var user *pb.User

		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}

		log.Println(user)

		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	//Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("sending email to: ", user.Name)
	return nil
}
