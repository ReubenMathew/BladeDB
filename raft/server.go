package raft

import (
	"math/rand"
	"sync"
	"time"
)

const (
	minElectionTimeout = 500
	maxElectionTimeout = 1200
)

type Server interface {
	ID() int
}

type server struct {
	mutex              sync.RWMutex
	id                 int
	peerIds            []int
	follower           bool
	leader             bool
	candidate          bool
	running            bool
	electionTimeout    time.Duration
	appendEntryTimeout int
	term               int // latest term server has seen
	votedFor           int // ID of candidate this server has voted for
	leaderID           int // ID of leader server
}

func generateElectionTimeout(min int, max int) time.Duration {
	return time.Duration(int(rand.Intn(100)*(max-min+1)+min)) * time.Millisecond
}

/* Constructor */
func NewServer(id int) Server {
	s := &server{
		id:                 id,
		follower:           true,
		leader:             false,
		candidate:          false,
		running:            false,
		appendEntryTimeout: 200,
		term:               0,
		votedFor:           -1, // -1 means did not vote for anyone
		leaderID:           -1, // -1 means no leader selected
	}
	s.electionTimeout = generateElectionTimeout(minElectionTimeout, maxElectionTimeout)

	return s
}

func (s *server) ID() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.id
}
