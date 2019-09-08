package main

import (
	"log"
	"root"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := &root.AccountRequest{
		Username: "uid123",
		Password: "password",
	}
	data, err := proto.Marshal(req)

	log.Print(err)
	log.Print(string(data))

	re := &root.AccountRequest{}
	proto.Unmarshal(data, re)
	log.Print(re)
}
