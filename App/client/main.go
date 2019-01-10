package main

import (
	"VotingApp/App/protocolBuffers/api"
	"context"
	"log"
	"io"
	"google.golang.org/grpc"
)

const (
	address = ":50051"
)

func CreateSurvey(client api.SurveyClient, survey *api.SurveyMessage){
	resp, err := client.CreateSurvey(context.Background(), survey)
	if err != nil {
		log.Fatalf("Cannot create Survey: %v\n", err)
	}

	if resp.Success{
		log.Printf("New survey with id: %v , was created. \n", resp.Id)
	}
}


func GetSurveys(client api.SurveyClient, label *api.Topic){

	stream, err := client.GetSurveys(context.Background(), label)
	if err != nil {
		log.Fatalf("Error in getting survey: %v\n",err)
	}

	log.Println("********** List of surveys **********")
	index := 1
	for {
		survey, err := stream.Recv()


		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v,GetCustomers(_) = _,%v\n", client, err)
		}

		log.Printf("%v. %v\n",index,survey.Name)
		log.Println("With options:")
		for _, t := range survey.Options{
			log.Printf("- %v\n",t)
		}
		log.Println("---------------------------------------")
		index ++

	}
}

func main() {

	conn, err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	defer conn.Close()
	client := api.NewSurveyClient(conn)

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
	CreateSurvey(client, &survey1)

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
	CreateSurvey(client, &survey2)

	t := api.Topic{Label:""}
	GetSurveys(client, &t)
}
