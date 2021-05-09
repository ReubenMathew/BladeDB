package model

import (
	"math/rand"
)

type Server struct {
	ID                 int
	Follower           bool
	Leader             bool
	Candidate          bool
	ElectionTimeout    int
	AppendEntryTimeout int
	CurrentTerm        int // latest term server has seen
	VotedFor           int // ID of candidate this server has voted for
}

func generateElectionTimeout(min int, max int) int {
	return int(rand.Intn(100)*(max-min+1) + min)
}

/* Constructor */
func NewServer(id int) *Server {
	server := Server{
		ID:                 id,
		Follower:           true,
		Leader:             false,
		Candidate:          false,
		AppendEntryTimeout: 200,
		CurrentTerm:        0,
		VotedFor:           -1, // -1 means did not vote for anyone
	}
	server.ElectionTimeout = generateElectionTimeout(500, 1200)

	return &server
}
