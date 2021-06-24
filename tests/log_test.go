package tests

import (
	"testing"

	"github.com/ReubenMathew/BladeDB/raft"
)

var log *raft.RaftLog

func testSetup_Log() {
	log = raft.NewLog()
}

func TestLog_empty(t *testing.T) {
	testSetup_Log()

	if len(log.GetEntries()) > 0 {
		t.Fatalf("Log should be empty")
	}
}

func TestLog_append(t *testing.T) {
	testSetup_Log()

	log.Append(0, "test log entry")
	if len(log.GetEntries()) == 1 {
		t.Fatal("Entry not added successfully")
	}
	if log.GetEntry(0).GetCommand() != "test log entry" {
		t.Fatal("Incorrect Entry")
	}
}
