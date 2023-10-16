package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/MuhahaSam/golangPractice/config"
	"github.com/MuhahaSam/golangPractice/internal/app/api/note_v1"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatalf("error during reading env file Err: %s", err)
	}

	go func() {
		defer wg.Done()
		startGRPC()
	}()
	go func() {
		defer wg.Done()
		startHTTP()
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", config.GetConfig().GrpcHost)
	if err != nil {
		log.Fatalf("failed mapping port: %s ", err.Error())
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	note := note_v1.NewNote()
	err = note.Init()
	if err != nil {
		log.Fatalf("failed init project: %s ", err.Error())
	}

	desc.RegisterNoteServiceServer(server, note)

	fmt.Printf("server is running on port: %s \n", config.GetConfig().GrpcHost)

	defer note.Destructor()
	if err = server.Serve(list); err != nil {
		log.Fatalf("failed serve: %s ", err.Error())
		return err
	}
	return nil
}

func startHTTP() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, config.GetConfig().GrpcHost, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.GetConfig().HttpHost, mux)
}
