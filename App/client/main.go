package main

import (
	"VotingApp/App/protocolBuffers/api"
	"log"
	"google.golang.org/grpc"
	"math/rand"
	"time"
	"context"
)

const (
	address = ":50051"
)



func main() {

	conn, err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	defer conn.Close()
	surveyClient := api.NewSurveyClient(conn)

	// create surveys
	survey1 := api.SurveyMessage{
		Id: 101,
		Name: "Presidential election 2018",
		Description: "Presidential election in 2018 for Czech Republic",
		Topic: []*api.Topic{
			&api.Topic{Label:"president"},
			&api.Topic{Label:"general election"},
		},
		Options:[]string{"Zeman","Horacek", "Fikus"},
		Action:api.SurveyMessage_CREATE,

	}
	CreateSurvey(surveyClient, &survey1)

	survey2 := api.SurveyMessage{
		Id: 102,
		Name: "FAVOURITE COLOR",
		Description: "Vote for the best color of the world !!!",
		Topic: []*api.Topic{
			&api.Topic{Label:"color"},
		},
		Options:[]string{"Green","Blue", "Yellow", "Pink", "Red", "White", "Black"},
		Action:api.SurveyMessage_CREATE,
	}
	CreateSurvey(surveyClient, &survey2)

	t := api.Topic{Label:""}
	surveys := GetSurveys(surveyClient, &t)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	//send votes
	voteClient := api.NewVoteClient(conn)
	for _,s := range surveys{

		index := r1.Intn(len(s.Options) - 1)
		option := s.Options[index]
		v := &api.VoteMessage{Vote: option, SurveyId:s.Id}
		Vote(context.Background(),voteClient, v)
	}
}
