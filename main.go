package main

import (
	"fmt"
	pb "github.com/Miraddo/protobuf-examples/go-example"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func main()  {


	profile :=&pb.Person{
		Name: "Miraddo",
		Age: 25,
		Eye: pb.Person_Black, // or 1 //just paste the number
		Allow: &pb.Validate{
			Valid: true,
		},
		Address: []*pb.Person_Address{
			{
				City: "Tokyo",
				Address: "All World is Mine but I like Japan more",
			},
			{
				City: "New York",
				Address: "I'd like to see New York too",
			},
		},
	}

	profile.Eye = 2 // just paste the number // it returns brown

	// read data with Get
	fmt.Printf("The name is : '%s' and The age is: '%d' and The color is: '%s' and The Alow is: '%v' and the addresses are: '%v'  \n", profile.GetName(), profile.GetAge(), profile.GetEye(), profile.GetAllow(), profile.GetAddress())

	// make data to binary with proto marshal
	binData, err := proto.Marshal(profile)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("Return Binary : ", binData)

	// read data from binary to main data with proto unmarshal
	profile2 := &pb.Person{}
	err = proto.Unmarshal(binData, profile2)

	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("Return Main Data : ", profile2)

	// make data to binary as json with proto json marshal
	binJSONData, err := protojson.Marshal(profile2)

	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("Return JSON Binary : ", binJSONData)

	// read data from binary to main data as JSON with proto json unmarshal
	profile3 := &pb.Person{}

	err = protojson.Unmarshal(binJSONData, profile3)

	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("Return Main JSON Data : ", profile3)

	// to get json format with proto json format
	fmt.Println("Return JSON Format Data : \n", protojson.Format(profile3))

}
