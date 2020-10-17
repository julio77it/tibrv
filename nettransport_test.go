package tibrv

import (
	"os"
	"sync"
	"testing"
)

func TestRvNetTransportSend(t *testing.T) {
	var transport RvNetTransport

	err := transport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
		Description("TestDescription"),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	var msg RvMessage
	if err := msg.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()
	err = msg.SetSendSubject("PROVA.TEST.TIBRV")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetString("StringField", "StringValue")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt8("Int8Field", -1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt32("UInt32Field", 70000)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetFloat64("Float64Field", -3.145)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := transport.Send(msg); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}

func TestNetRvListenerRequestReply(t *testing.T) {
	subject := "UNIT.TEST"

	var queue RvQueue
	if err := queue.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer queue.Destroy()

	var transport RvNetTransport
	if err := transport.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	var callback RvCallback = func(t *RvNetTransport) func(msg *RvMessage) {
		return func(imsg *RvMessage) {
			var omsg RvMessage
			omsg.create(imsg.internal)
			t.SendReply(omsg, *imsg)
			omsg.Destroy()
		}
	}(&transport)

	var listener RvListener
	err := listener.Create(
		queue,
		callback,
		transport,
		subject,
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer listener.Destroy()

	var request, reply RvMessage
	if err := request.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer request.Destroy()

	if err := request.SetSendSubject(subject); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := request.SetReplySubject(subject + ".AA"); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := request.SetInt32("Integer32bit", -25); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if err := reply.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer reply.Destroy()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		if err = queue.Dispatch(); err != nil {
			t.Fatalf("Expected nil, got %v", err)
		}
		wg.Done()
	}()
	go func() {
		if err = transport.SendRequest(request, &reply, WaitForEver); err != nil {
			t.Fatalf("Expected nil, got %v", err)
		}
		wg.Done()
	}()
	wg.Wait()
	input := request.String()
	output := reply.String()

	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
	}
}

func BenchmarkRvNetTransportSend(b *testing.B) {
	var transport RvNetTransport

	err := transport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
		Description("BenchmarkDescription"),
	)
	if err != nil {
		b.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	var msg RvMessage
	if err := msg.Create(); err != nil {
		b.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()

	err = msg.SetSendSubject("PROVA.TEST.TIBRV")
	if err != nil {
		b.Fatalf("Expected nil, got %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = msg.SetUInt64("UInt64Field", uint64(n))
		if err != nil {
			b.Fatalf("Expected nil, got %v", err)
		}
		if err := transport.Send(msg); err != nil {
			b.Fatalf("Expected nil, got %v", err)
		}
	}
}
