package raft

import "sync"

type raftLog struct {
	mutex     sync.RWMutex
	entries   []entry
	commitIdx int
}

type entry struct {
	Idx     int    `json:"index"`
	Term    int    `json:"term"`
	Command []byte `json:"command"`
}
