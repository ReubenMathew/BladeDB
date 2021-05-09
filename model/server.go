package model

import (
	"math/rand"
)

const (
	minElectionTimeout = 500
	maxElectionTimeout = 1200
)

type Server struct {
	id                 int
	follower           bool
	leader             bool
	candidate          bool
	electionTimeout    int
	appendEntryTimeout int
	currentTerm        int // latest term server has seen
	votedFor           int // ID of candidate this server has voted for
	leaderID           int // ID of leader server
}

func generateElectionTimeout(min int, max int) int {
	return int(rand.Intn(100)*(max-min+1) + min)
}

/* Constructor */
func NewServer(id int) *Server {
	server := Server{
		id:                 id,
		follower:           true,
		leader:             false,
		candidate:          false,
		appendEntryTimeout: 200,
		currentTerm:        0,
		votedFor:           -1, // -1 means did not vote for anyone
		leaderID:           -1, // -1 means no leader selected
	}
	server.electionTimeout = generateElectionTimeout(minElectionTimeout, maxElectionTimeout)

	return &server
}
