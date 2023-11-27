package main

import (
	pb "UserManagment/gen/go/protos/v1"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	user := pb.User{
		Uuid:      "1-G-6-9",
		FullName:  "Pavle Djurdjic",
		BirthYear: 1999,
	}

	// Marshal serializes message into binary format
	bytes, err := proto.Marshal(&user)

	if err != nil {
		log.Fatalln("Marshalling error", err)
	}

	// Write bytes to file
	// we have function call and assignment after which we check for errors
	if err := os.WriteFile("user.bin", bytes, 0644); err != nil {
		log.Fatalln("Writing error", err)
	}
}
