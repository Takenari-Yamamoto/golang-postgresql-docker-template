package main

import (
	"fmt"
	"log"
	"protobuf-lesson/pb"
)

func main() {
	employee := &pb.Employee {
		Id: 1,
		Name: "Suzuki",
		Email: "test@example.com",
		Occupation: pb.Occupation_ENGINEER,
		PhoneNumber: []string{"090-1234-5678", "090-1234-1234"},
		Project: map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year: 2000,
			Month: 1,
			Day: 1,
		},
	}

	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Can`t serialize", err)
	// }

	// if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
	// 	log.Fatalln("Can`t serialize", err)
	// }

	// in, err := ioutil.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Can`t serialize", err)
	// }

	// readEmployee := &pb.Employee{}

	// err = proto.Unmarshal(in, readEmployee)

	// if err != nil {
	// 	log.Fatalln("Can`t serialize", err)
	// }

	// fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)

	if err != nil {
		log.Fatalln("Cant marshall to json", err)
	}

	fmt.Println(out)
	// readEmployee := &pb.Employee{}
	// if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
	// 	log.Fatalln("Can`t serialize", err)
	// }

	// fmt.Println(readEmployee)

}
