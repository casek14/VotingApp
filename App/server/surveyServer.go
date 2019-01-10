package main

import (
	"VotingApp/App/protocolBuffers/api"
	"fmt"
	"strings"
	"context"
)

type surveyServer struct {
	savedSurveys []*api.SurveyMessage
}

func (s *surveyServer)CreateSurvey(ctx context.Context, survey *api.SurveyMessage)(*api.SurveyResponse, error){
	s.savedSurveys = append(s.savedSurveys, survey)
	fmt.Printf("Request: Create Survey: %v-%v\n",survey.Id,survey.Name)
	return &api.SurveyResponse{Id:survey.Id,Success:true}, nil
}

func (s *surveyServer)GetSurveys(topic *api.Topic, stream api.Survey_GetSurveysServer) error {
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







