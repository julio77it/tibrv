package main

import (
	"github.com/julio77it/tibrv"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	subject := "UNIT.TEST"

	var transport tibrv.RvNetTransport
	err := transport.Create(tibrv.Description("SENDER"))
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	defer transport.Destroy()

	var request, reply tibrv.RvMessage
	err = request.Create()
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	defer request.Destroy()
	err = request.SetSendSubject(subject)
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	err = request.SetInt32("Integer32bit", -25)
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}

	err = reply.Create()
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	defer reply.Destroy()

	err = transport.SendRequest(request, &reply, tibrv.WaitForEver)
	if err != nil {
		log.Println(err)
		log.Printf("Expected nil, got %v\n", err)
	}
	input := request.String()
	output := reply.String()

	sendSubj, _ := request.GetSendSubject()
	replySubj, _ := request.GetReplySubject()
	log.Printf("INPUT |%s|%s|%s|\n", sendSubj, replySubj, input)
	sendSubj, _ = reply.GetSendSubject()
	replySubj, _ = reply.GetReplySubject()
	log.Printf("OUPUT |%s|%s|%s|\n", sendSubj, replySubj, output)

	if output != input {
		log.Printf("Expected %s, got %s\n", input, output)
	}

	time.Sleep(2 * time.Second)
}
