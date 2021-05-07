package raft

type Raft struct {
	ID int
}

func NewRaft(id int) *Raft {
	raft := Raft{}
	raft.ID = id
	return &raft
}

func (raft *Raft) GetID() int {
	return raft.ID
}
