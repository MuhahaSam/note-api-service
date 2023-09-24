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
	contextBackground := context.Background()

	client := desc.NewNoteServiceClient(con)

	var createResponse *desc.CreateNoteResponse
	var getResponse *desc.GetNoteResponse
	var resErr error

	createResponse, resErr = client.CreateNote(contextBackground, &desc.CreateNoteRequest{
		Title:  "Wow",
		Text:   "Note created",
		Author: "Sam",
	})

	uuid := &desc.UUID{Value: createResponse.Uuid.Value}

	log.Println("note Id: ", createResponse.Uuid)

	getResponse, resErr = client.GetNote(contextBackground, &desc.GetNoteRequest{
		Uuid: uuid,
	})

	log.Println("read note: ", getResponse)

	type StringValue = wrapperspb.StringValue

	_, resErr = client.UpdateNote(contextBackground, &desc.UpdateNoteRequest{
		Uuid: uuid,
		UpdateBody: &desc.UpdateNoteBody{
			Author: &StringValue{Value: "kim"},
			Title:  &StringValue{Value: "Kim's story"},
			Text:   &StringValue{Value: "This is my first crud on go"},
		},
	})

	getResponse, _ = client.GetNote(contextBackground, &desc.GetNoteRequest{
		Uuid: uuid,
	})

	log.Println("read note after update: ", getResponse)

	_, resErr = client.DeleteNote(contextBackground, &desc.DeleteNoteRequest{
		Uuid: uuid,
	})

	getResponse, _ = client.GetNote(contextBackground, &desc.GetNoteRequest{
		Uuid: uuid,
	})

	log.Println("read note after delete: ", getResponse)

	if resErr != nil {
		log.Println(resErr.Error())
	}

}
