package tests

import (
	"testing"

	"github.com/ReubenMathew/BladeDB/raft"
)

var log *raft.RaftLog

func testSetup_Log() {
	log = raft.NewLog()
}

func TestLog_write(t *testing.T) {
	testSetup_Log()
}
