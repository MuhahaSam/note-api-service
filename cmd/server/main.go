package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MuhahaSam/golangPractice/internal/app/api/note_v1"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed mapping port: %s ", err.Error())
	}

	server := grpc.NewServer()
	desc.RegisterNoteServiceServer(server, note_v1.NewNote())

	fmt.Printf("server is running on port: %s \n", port)

	if err = server.Serve(list); err != nil {
		log.Fatalf("failed serve: %s ", err.Error())
	}
}
