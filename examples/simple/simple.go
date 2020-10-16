package main

import (
	"fmt"
	"github.com/julio77it/tibrv"
	"time"
)

func main() {
	var m tibrv.RvMessage
	m.Create()
	defer m.Destroy()

	name := "fieldName"
	var inV float64 = 3.145

	var in tibrv.RvMessage
	in.Create()
	in.SetFloat64(name, inV)
	m.SetRvMessage(name, in)

	fmt.Println(m)

	m.SetSendSubject("PROVA.1")

	var transport tibrv.RvNetTransport

	transport.Create(
		//tibrv.Service("7500"),
		//tibrv.Network("127.0.0.1"),
		//tibrv.Daemon("tcp:7500"),
		tibrv.Description("Descrizione"),
	)

	var queue tibrv.RvQueue
	queue.Create()
	defer queue.Destroy()
	go func() {
		for {
			queue.Dispatch()
		}
	}()
	var listener tibrv.RvListener

	cb := func(m *tibrv.RvMessage) {
		fmt.Println("RvMessage:", *m)
	}

	listener.Create(queue, cb, transport, "PROVA.1")
	defer listener.Destroy()

	for i := 0; i < 100; i++ {
		m.SetUInt32("counter", uint32(i))
		transport.Send(m)
	}

	time.Sleep(2 * time.Second)
}
