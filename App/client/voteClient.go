package main

import (
	"VotingApp/App/protocolBuffers/api"
	"context"
	"log"
)

func Vote(ctx context.Context, client api.VoteClient,message *api.VoteMessage) (*api.VoteResponse,error){
	resp, err := client.Vote(context.Background(),message)
	if err != nil {
		log.Fatalf("Failed to send vote: %v\n",err)
	}

	if ! resp.Success {
		return nil, err
	}

	return resp,nil
}

