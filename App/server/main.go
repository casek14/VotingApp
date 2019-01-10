package main

import (
	"VotingApp/App/protocolBuffers/api"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
)

const (
	port = ":50051"
)



func main() {
	lis, err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterSurveyServer(grpcServer, &surveyServer{})
	fmt.Printf("Listening on port *%v \n",port)

	grpcServer.Serve(lis)


	}
