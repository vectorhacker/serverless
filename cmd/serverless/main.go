package main

import (
	"log"

	pb "github.com/Azure/azure-functions-go-worker/internal/rpc"
)

func main() {
	msg := &pb.StreamingMessage{
		RequestId: "asdfasdf",
	}

	log.Printf("StreamingMessage: %#v\n", msg)
}
