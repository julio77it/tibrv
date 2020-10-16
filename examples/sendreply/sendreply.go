package main

import (
	"github.com/julio77it/tibrv"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	subject := "UNIT.TEST"

	var queue tibrv.RvQueue
	err := queue.Create()
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	defer queue.Destroy()

	var transport tibrv.RvNetTransport
	err = transport.Create(tibrv.Description("REPLIER"))
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	defer transport.Destroy()

	var callback tibrv.RvCallback = func(t *tibrv.RvNetTransport) func(msg *tibrv.RvMessage) {
		return func(imsg *tibrv.RvMessage) {
			var omsg tibrv.RvMessage
			omsg.Create()
			defer omsg.Destroy()
			e := omsg.SetInt32("Integer32bit", -25)
			if e != nil {
				log.Printf("Expected nil, got %v\n", e)
			}

			sendSubject, _ := imsg.GetSendSubject()
			replySubject, _ := imsg.GetReplySubject()
			log.Printf("|%s|%s|\n", sendSubject, replySubject)
			log.Printf("|%s|%s|\n", imsg.String(), omsg.String())

			e = t.SendReply(omsg, *imsg)
			if e != nil {
				log.Printf("Expected nil, got %v\n", e)
			}
			log.Println("---------------------------------------------------")
		}
	}(&transport)

	var listener tibrv.RvListener
	err = listener.Create(
		queue,
		callback,
		transport,
		subject,
	)
	if err != nil {
		log.Printf("Expected nil, got %v\n", err)
	}
	defer listener.Destroy()

	for {
		queue.Dispatch()
	}
}
