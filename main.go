package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gugazimmermann/go-protocol-buffers/src/complex/complexpb"
	"github.com/gugazimmermann/go-protocol-buffers/src/enum_example/enumpb"
	"github.com/gugazimmermann/go-protocol-buffers/src/simple/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Test",
		SampleList: []int32{1, 4, 7, 8},
	}
	return &sm
}

func doEnum(ID int32) *enumpb.EnumMessage {
	en := enumpb.EnumMessage{
		Id: ID,
	}
	return &en
}

func doDummy(i int32, n string) *complexpb.DummyMessage {
	dm := complexpb.DummyMessage{
		Id:   i,
		Name: n,
	}
	return &dm
}

func doComplex(ID int32) *complexpb.ComplexMessage {
	dm1 := doDummy(23, "Dummy 23")
	cm := complexpb.ComplexMessage{
		Id:       ID,
		OneDummy: dm1,
	}
	return &cm
}

func main() {
	// simple proto
	fmt.Print("\n\n*** Simple ***\n\n")
	sm := doSimple()
	sm2 := &simplepb.SimpleMessage{}
	readAndWrite("simple.bin", sm, sm2)
	fmt.Println("")
	jsonDemo(sm)

	// enum proto
	fmt.Print("\n\n*** Enum ***\n\n")
	en := doEnum(30)
	fmt.Println(en)
	en.DayOfTheWeek = enumpb.DayOfTheWeek_FRIDAY
	fmt.Println(en)
	en2 := &enumpb.EnumMessage{}
	readAndWrite("enum.bin", en, en2)
	fmt.Println("")
	jsonDemo(en)

	// complex proto
	fmt.Print("\n\n*** Complex ***\n\n")
	cm := doComplex(30)
	fmt.Println(cm)
	dmTest := doDummy(66, "Dummy 66")
	dms := []*complexpb.DummyMessage{
		doDummy(50, "Dummy  50"),
		doDummy(51, "Dummy  51"),
		dmTest,
	}
	cm.MultipleDummy = dms
	fmt.Println(cm)
	dmTest.Name = "Renamed Dummy"
	fmt.Println(cm)

	cm2 := &complexpb.ComplexMessage{}
	readAndWrite("enum.bin", cm, cm2)
	fmt.Println("")
	jsonDemo(cm)
}

func readAndWrite(fileName string, pb proto.Message, pb2 proto.Message) {
	writeToFile(fileName, pb)
	readFromFile(fileName, pb2)
	fmt.Println("PB from", fileName)
	fmt.Println(pb2)
}

func jsonDemo(pb proto.Message) {
	pbAsJSONString := toJson(pb)
	fmt.Println("PB to JSON:")
	fmt.Println(string(pbAsJSONString))
	fromJson(pbAsJSONString, pb)
	fmt.Println("PB from JSON:", pb)
}

func writeToFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written to:", fileName)
	return nil
}

func readFromFile(fileName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fileName)
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
