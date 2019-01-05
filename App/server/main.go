package main

import (
	"VotingApp/protocolBuffers/api"
	"strings"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
)

const (
	port = ":50051"
)

type server struct {
	savedSurveys []*api.SurveyMessage
}

func (s *server)CreateSurvey(ctx context.Context, survey *api.SurveyMessage)(*api.SurveyResponse, error){
	s.savedSurveys = append(s.savedSurveys, survey)
	fmt.Printf("Request: Create Survey: %v-%v\n",survey.Id,survey.Name)
	return &api.SurveyResponse{Id:survey.Id,Success:true}, nil
}

func (s *server)GetSurveys(topic *api.Topic, stream api.Survey_GetSurveysServer) error {
	for _, survey := range s.savedSurveys{
		contains := false
		if topic.Label != ""{
			for _, t := range survey.Topic {
				if strings.Contains(t.Label, topic.Label) {
					contains = true
					continue
					}

			}

		}

		if contains == true {
			continue
		}
		err := stream.Send(survey)
		if err != nil {
			return err
		}

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterSurveyServer(grpcServer, &server{})
	fmt.Printf("Listening on port *%v \n",port)

	grpcServer.Serve(lis)


	}
