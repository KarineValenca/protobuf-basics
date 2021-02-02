package main

import (
	"fmt"
	"io/ioutil"
	"log"

	example_complex "github.com/KarineValenca/protobuf-basics/complex"
	example_enum "github.com/KarineValenca/protobuf-basics/enum_example"
	example_simple "github.com/KarineValenca/protobuf-basics/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()
	doComplex()
}

func doComplex() {
	cm := example_complex.ComplexMessage{
		OneDummy: &example_complex.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*example_complex.DummyMessage{
			&example_complex.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&example_complex.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := example_enum.EnumExample{
		Id:           42,
		DayOfTheWeek: example_enum.DaysOfTheWeek_THURSDAY,
	}
	fmt.Println(em)
	em.DayOfTheWeek = example_enum.DaysOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println(smAsString)
	sm2 := &example_simple.SimpleMessage{}
	fromJson(smAsString, sm2)
	fmt.Println("ProtoStruct", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return out
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {

	writeToFile("simple.bin", sm)

	sm2 := &example_simple.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content: ", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	log.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return err
	}
	log.Println("Saved bytes into the protocol buffer struct")
	return nil
}

func doSimple() *example_simple.SimpleMessage {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 6, 8},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm.GetId())

	return &sm
}
