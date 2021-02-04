package main

import (
	enumpb "Protocol-Buffers/protobuf-with-golang/src/enum_example"
	simplepb "Protocol-Buffers/protobuf-with-golang/src/simple"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := simple()
	// writeToFile("simple.bin", sm)

	// sm2 := &simplepb.SingleMessage{}
	// readFromFile("simple.bin", sm2)
	// fmt.Println(sm2)
	smToString := toJSON(sm)
	fmt.Println(smToString)

	sm2 := &simplepb.SingleMessage{}
	fromJSON(smToString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)

	doEnum()
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_TUESDAY,
	}

	// em.DayOfTheWeek = enumpb.DayOfTheWeek_FRIDAY

	fmt.Println(em)
}

func fromJSON(in string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't conver to JSON", err)
		return ""
	}
	return out
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes.", err)
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data have been written!")
	return nil
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Something went wront when reading the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct.", err)
		return err
	}

	return nil
}

func simple() *simplepb.SingleMessage {
	sm := simplepb.SingleMessage{
		Id:         1,
		Name:       "Petros",
		IsSimple:   true,
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(sm)

	sm.Name = "Petros Trak"
	fmt.Println(sm)

	fmt.Println(sm.GetId())

	return &sm
}
