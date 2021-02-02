package main

import (
	"fmt"
	"io/ioutil"
	"log"

	tutorial "github.com/KarineValenca/golang-exercise/tutorial"
	"google.golang.org/protobuf/proto"
)

func main() {
	pm := tutorial.Person{
		Name:  "Test",
		Id:    123,
		Email: "test@test.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "555-555", Type: tutorial.Person_HOME},
		},
	}
	writeData(pm, "addressbook")
	writeData(pm, "addressbook")
	addressBook := tutorial.AddressBook{}
	readData("addressbook", &addressBook)
	fmt.Println(addressBook)
}

func writeData(pm tutorial.Person, fname string) {
	addressBook := &tutorial.AddressBook{}
	addressBook.People = append(addressBook.People, &pm)
	out, err := proto.Marshal(addressBook)
	if err != nil {
		log.Fatalln("Failed to encode address book", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book", err)
	}
}

func readData(fname string, pb proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Failed to parse address book", err)
	}
}
