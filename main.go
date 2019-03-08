package main

import (
	complexpb "Personal/proto-golang-playground/src/complex"
	enumpb "Personal/proto-golang-playground/src/enum"
	simplepb "Personal/proto-golang-playground/src/simple"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Printf("Hello world!")

	sm := doSimple()
	if err := writeToFile("simple.bin", sm); err != nil {
		fmt.Println(err.Error())
	}

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm3 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm3)

	fmt.Println("Successfully created protobuf struct: ", sm3)

	doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_FRIDAY
	fmt.Println(em)
}

func toJSON(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON!", err)
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the protobuf struct!", err)
	}
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes!", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file!", err)
	}

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you!"

	fmt.Println(sm)
	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
