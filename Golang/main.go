package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/proto-buf/Golang/src/addressbook"
	"github.com/proto-buf/Golang/src/complex"
	"github.com/proto-buf/Golang/src/enum_example"
	"github.com/proto-buf/Golang/src/simple"
)

func main() {
	//sm := doSimple()
	// protobuf read and write
	//readAndWrite(sm)

	// Json read and write
	//jsonTest(sm)

	// Enum example
	//doEnum()

	// Complex
	//doComplex()

	// Person example
	doPerson()
}

func doPerson() {
	pm := tutorial.Person{}
	fmt.Println(pm)
}
func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}
	fmt.Println(em)
}

func jsonTest(sm proto.Message) {

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)

}

func toJSON(pb proto.Message) string {
	// didnt define any options
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("couldnt unmarshal JSON into the pb", err)

	}
}

//Read and Write to and from protobuff
func readAndWrite(sm proto.Message) {

	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Reading the content", sm2)
}

//write to protobuff binary
func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written")
	return nil
}

//read protobuff file
func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldnt put the bytes into the protocol buffers struct", err2)
	}

	return nil

}

func doSimple() *simplepb.SimpleMessage {
	//Struct
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "hamed",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	//rename
	sm.Name = "I renamed you"
	fmt.Println(sm)
	// sm.Id is not correct we should use the helper functions like GetId
	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
