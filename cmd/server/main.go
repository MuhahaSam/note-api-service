package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MuhahaSam/golangPractice/internal/app/api/note_v1"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("error during reading env file Err: %s", err)
	}

	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed mapping port: %s ", err.Error())
	}

	server := grpc.NewServer()
	note := note_v1.NewNote().New()
	desc.RegisterNoteServiceServer(server, note)

	fmt.Printf("server is running on port: %s \n", port)

	defer note.Destructor()
	if err = server.Serve(list); err != nil {
		log.Fatalf("failed serve: %s ", err.Error())
	}
}
