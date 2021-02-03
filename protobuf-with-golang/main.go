package main

import (
	"Protocol-Buffers/protobuf-with-golang/src/simple"
	"fmt"
)

func main() {
	sm := simple.SingleMessage{
		Id:         1,
		Name:       "Petros",
		IsSimple:   true,
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(sm)

	sm.Name = "Petros Trak"
	fmt.Println(sm)

	fmt.Println(sm.GetId())
}
