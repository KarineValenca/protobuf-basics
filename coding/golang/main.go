package main

import (
	"fmt"

	example_simple "github.com/KarineValenca/protobuf-basics/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 6, 8},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm.GetId())
}
