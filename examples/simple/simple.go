package main

import (
	"fmt"
	"time"

	"github.com/julio77it/tibrv"
)

func main() {
	var m tibrv.RvMessage
	if err := m.Create(); err != nil {
		panic(err)
	}
	defer m.Destroy()

	name := "fieldName"
	var inV float64 = 3.145

	var in tibrv.RvMessage
	if err := in.Create(); err != nil {
		panic(err)
	}
	if err := in.SetFloat64(name, inV); err != nil {
		panic(err)
	}
	if err := m.SetRvMessage(name, in); err != nil {
		panic(err)
	}

	fmt.Println(m)

	if err := m.SetSendSubject("PROVA.1"); err != nil {
		panic(err)
	}

	var transport tibrv.RvNetTransport

	if err := transport.Create(
		//tibrv.Service("7500"),
		//tibrv.Network("127.0.0.1"),
		//tibrv.Daemon("tcp:7500"),
		tibrv.Description("Descrizione"),
	); err != nil {
		panic(err)
	}

	var queue tibrv.RvQueue
	if err := queue.Create(); err != nil {
		panic(err)
	}
	defer queue.Destroy()
	go func() {
		for {
			if err := queue.Dispatch(); err != nil {
				panic(err)
			}
		}
	}()
	var listener tibrv.RvListener

	cb := func(m *tibrv.RvMessage) {
		fmt.Println("RvMessage:", *m)
	}

	if err := listener.Create(queue, cb, transport, "PROVA.1"); err != nil {
		panic(err)
	}
	defer listener.Destroy()

	for i := 0; i < 100; i++ {
		if err := m.SetUInt32("counter", uint32(i)); err != nil {
			panic(err)
		}
		if err := transport.Send(m); err != nil {
			panic(err)
		}
	}

	time.Sleep(2 * time.Second)
}
