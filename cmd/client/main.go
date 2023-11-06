package main

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dIdn't connect to server: %s", err.Error())
	}

	defer con.Close()
	ctx := context.Background()
	client := desc.NewNoteServiceClient(con)

	var createRes *desc.CreateResponse

	createRes, err = client.CreateNote(ctx, &desc.CreateRequest{
		Title:  "Wow",
		Text:   "Note created",
		Author: "Sam",
	})

	uuid := createRes.Uuid

	log.Println("note Id: ", createRes.Uuid)

	getRes, err := client.GetNote(ctx, &desc.GetRequest{
		Uuid: uuid,
	})

	log.Println("read note: ", getRes)

	_, err = client.UpdateNote(ctx, &desc.UpdateRequest{
		Uuid: uuid,
		UpdateBody: &desc.UpdateBody{
			Author: &wrapperspb.StringValue{Value: "kim"},
			Title:  &wrapperspb.StringValue{Value: "Kim's story"},
			Text:   &wrapperspb.StringValue{Value: "Kim's story"},
		},
	})

	getRes, err = client.GetNote(ctx, &desc.GetRequest{
		Uuid: uuid,
	})

	log.Println("read note after update: ", getRes)

	_, err = client.DeleteNote(ctx, &desc.DeleteRequest{Uuid: uuid})
	getRes, err = client.GetNote(ctx, &desc.GetRequest{
		Uuid: uuid,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("read note after delete: ", getRes)
}
