package main

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect to server: %s", err.Error())

	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow",
		Text:   "Note created",
		Author: "Sam",
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("note index: ", res.Index)

}
