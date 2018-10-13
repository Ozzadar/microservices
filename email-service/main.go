package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	pb "github.com/ozzadar/microservices/user-service/proto/auth"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up new message")
	log.Println("Sending email to: ", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	//Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("sending email to: ", user.Name)
	return nil
}
