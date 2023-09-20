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
	createRes, createErr := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow",
		Text:   "Note created",
		Author: "Sam",
	})

	if createErr != nil {
		log.Println(createErr.Error())
	}

	log.Println("note index: ", createRes.Index)

	GetRes, GetErr := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Index: createRes.Index,
	})

	if GetErr != nil {
		log.Println(GetErr.Error())
	}

	log.Println("red note: ", GetRes)

	_, UpdateErr := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Index: GetRes.GetIndex(),
		UpdateBody: &desc.UpdateNoteBody{
			Author: "kim",
			Title:  "Kim's story",
			Text:   "This is my first crud on go",
		},
	})

	if UpdateErr != nil {
		log.Println(UpdateErr.Error())
	}

	GetResAfterUpdate, GetResAfterUpdateErr := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Index: GetRes.GetIndex(),
	})

	if GetResAfterUpdateErr != nil {
		log.Println(GetResAfterUpdateErr.Error())
	}

	log.Println("note after update: ", GetResAfterUpdate)

	_, DeleteErr := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Index: GetRes.GetIndex(),
	})

	if DeleteErr != nil {
		log.Println(DeleteErr.Error())
	}

	GetResDelete, GetResAfterDeleteErr := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Index: GetRes.GetIndex(),
	})

	if GetResAfterDeleteErr != nil {
		log.Println(GetResAfterDeleteErr.Error())
	}

	log.Println("note after delete: ", GetResDelete)

}
