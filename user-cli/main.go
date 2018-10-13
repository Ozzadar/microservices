package main

import (
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/ozzadar/microservices/user-service/proto/auth"
	"golang.org/x/net/context"
)

func main() {
	cmd.Init()

	client := pb.NewAuthClient("shippy.auth", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := "Paul"
			email := "things@stuff"
			password := "mypassword"
			company := "mauVILLE Technologies"

			log.Printf("%v", c)
			// Call our user service

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Password: password,
				Email:    email,
				Company:  company,
			})

			log.Printf("%s, %s, %s, %s", name, email, password, company)
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Create: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})

			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}

			for _, v := range getAll.Users {
				log.Println(v)
			}

			authResponse, err := client.Auth(context.TODO(), &pb.User{
				Email:    email,
				Password: password,
			})

			if err != nil {
				log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
			}

			log.Printf("Your access Token is: %s \n", authResponse.Token)

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
