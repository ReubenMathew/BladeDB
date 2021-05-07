package main

import (
	"fmt"

	"github.com/ReubenMathew/BladeDB/raft"
)

func main() {
	fmt.Println("Hello World")
	r := raft.NewRaft(2)
	fmt.Println(r.GetID())
}
