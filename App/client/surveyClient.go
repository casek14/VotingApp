package main

import (
	"log"
	"VotingApp/App/protocolBuffers/api"
	"io"
	"context"
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


func GetSurveys(client api.SurveyClient, label *api.Topic)(surveys []*api.SurveyMessage){

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
		surveys = append(surveys, survey)
		log.Printf("%v. %v\n",index,survey.Name)
		log.Println("With options:")
		for _, t := range survey.Options{
			log.Printf("- %v\n",t)
		}
		log.Println("---------------------------------------")
		index ++

	}
	return surveys
}
