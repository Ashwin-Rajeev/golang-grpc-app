package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"golang-grpc-app/app"

	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultID int32 = 1
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := app.NewStudentsClient(conn)

	// Contact the server and print out its response.
	id := defaultID
	if len(os.Args) > 1 {
		val, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("did not convert id: %v", err)
		}
		id = int32(val)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetStudents(ctx, &app.StudentRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("******************* Student details *******************")
	log.Printf("\tStudent ID      : %d\n", r.GetId())
	log.Printf("\tStudent Name    : %s\n", r.GetName())
	log.Printf("\tStudent Code    : %s\n", r.GetCode())
	log.Printf("\tStudent Program : %s\n", r.GetProgram())
}
