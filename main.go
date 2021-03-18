package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gugazimmermann/go-protocol-buffers/src/simple/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWriteSimple(sm)
	fmt.Println("")
	jsonDemoSimple(sm)
}

func readAndWriteSimple(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFIle("simple.bin", sm2)
	fmt.Println("PB from simple.bin:", sm2)
}

func jsonDemoSimple(sm proto.Message) {
	smAsJSONString := toJson(sm)
	fmt.Println("PB to JSON:")
	fmt.Println(string(smAsJSONString))
	sm2 := &simplepb.SimpleMessage{}
	fromJson(smAsJSONString, sm2)
	fmt.Println("PB from JSON:", sm2)
}

func doSimple() proto.Message {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Test",
		SampleList: []int32{1, 4, 7, 8},
	}
	return &sm
}

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fName, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written to:", fName)
	return nil
}

func readFromFIle(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Can't unserialize the bytes", err)
		return err
	}

	return nil
}

func toJson(pb proto.Message) []byte {
	marshaler := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	s, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return s
}

func fromJson(in []byte, pb proto.Message) {
	err := protojson.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Can't convert from JSON to PB", err)
	}
}
