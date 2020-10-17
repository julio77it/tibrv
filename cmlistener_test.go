package tibrv

import (
	"os"
	"sync"
	"testing"
)

func TestCmRvListenerPublishSubscribe(t *testing.T) {
	subject := "UNIT.TEST.CM"
	session := "TestRvListenerCmPublishSubscribe"
	ledger := "${TESTDIR}/TestRvListenerCmPublishSubscribe.ledger"

	var queue RvQueue
	if err := queue.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer queue.Destroy()

	var ntransport RvNetTransport
	err := ntransport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer ntransport.Destroy()

	var transport RvCmTransport
	err = transport.Create(&ntransport, Session(session), Ledger(ledger))
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	var output string

	var callback RvCallback = func(s *string) func(msg *RvMessage) {
		return func(msg *RvMessage) {
			*s = msg.String()
		}
	}(&output)

	var listener RvCmListener
	err = listener.Create(
		queue,
		callback,
		transport,
		subject,
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer listener.Destroy()

	var msg RvMessage
	if err := msg.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()

	if err := msg.SetInt32("Integer32bit", -25); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := msg.SetSendSubject(subject); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = transport.Send(msg)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	for i := 0; i < 5; i++ {
		queue.TimedDispatch(0.1)
	}
	input := msg.String()

	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
	}
}

func TestCmRvListenerRequestReply(t *testing.T) {
	subject := "UNIT.TEST.CM"
	session := "TestRvListenerCmPublishSubscribe"
	ledger := "${TESTDIR}/TestRvListenerCmPublishSubscribe.ledger"

	var queue RvQueue
	if err := queue.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer queue.Destroy()

	var ntransport RvNetTransport
	err := ntransport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer ntransport.Destroy()

	var transport RvCmTransport
	err = transport.Create(&ntransport, Session(session), Ledger(ledger))
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	var callback RvCallback = func(t *RvCmTransport) func(msg *RvMessage) {
		return func(imsg *RvMessage) {
			var omsg RvMessage
			omsg.create(imsg.internal)
			t.SendReply(omsg, *imsg)
			omsg.Destroy()
		}
	}(&transport)

	var listener RvCmListener
	err = listener.Create(
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
		defer wg.Done()
		for i := 0; i < 5; i++ {
			queue.TimedDispatch(0.1)
		}
	}()
	go func() {
		defer wg.Done()
		if err = transport.SendRequest(request, &reply, WaitForEver); err != nil {
			t.Fatalf("Expected nil, got %v", err)
		}
	}()
	wg.Wait()

	input := request.String()
	output := reply.String()

	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
	}
}
