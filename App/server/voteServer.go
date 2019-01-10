package main

import (
	"VotingApp/App/protocolBuffers/api"
	"fmt"
	"context"
)

type voteServer struct{
	votes []*api.VoteMessage
}

func (s *voteServer) Vote(ctx context.Context, vote *api.VoteMessage)(*api.VoteResponse,error){
	s.votes = append(s.votes, vote)
	fmt.Printf("New vote registred. SurveyID: %d - vote: %s.\n",vote.SurveyId,vote.Vote)
	return &api.VoteResponse{Success:true}, nil
}


