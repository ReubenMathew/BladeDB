package raft

import (
	"fmt"
	"sync"
)

type Log interface {
	Append()
	GetEntry(int) entry
	GetEntries() []entry
}

type RaftLog struct {
	mutex     *sync.RWMutex
	entries   []entry
	commitIdx int
}

type entry struct {
	term    int    `json:"term"`
	command string `json:"command"`
}

func NewLog() *RaftLog {
	return &RaftLog{
		mutex:     new(sync.RWMutex),
		entries:   []entry{},
		commitIdx: 0,
	}
}

func (e entry) GetCommand() string {
	return e.command
}

func (r RaftLog) Append(term int, command string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	e := entry{
		term:    term,
		command: command,
	}
	r.entries = append(r.entries, e)
}

func (r RaftLog) GetEntries() []entry {
	return r.entries
}

func (r RaftLog) GetEntry(idx int) entry {
	return r.entries[idx]
}

func (r RaftLog) ClearEntries() {
	r.entries = nil
}

func PrintLog(l Log) {
	entries := l.GetEntries()

	for _, v := range entries {
		fmt.Println(v.command)
	}
}
