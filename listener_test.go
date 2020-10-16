package tibrv

import (
	"sync"
	"testing"
)

func TestNetRvListenerPublishSubscribe(t *testing.T) {
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

	var output string

	var callback RvCallback = func(s *string) func(msg *RvMessage) {
		return func(msg *RvMessage) {
			*s = msg.String()
		}
	}(&output)

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
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = transport.Send(msg)
		if err != nil {
			t.Fatalf("Expected nil, got %v", err)
		}
	}()
	wg.Wait()
	if err := queue.Dispatch(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	input := msg.String()

	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
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

func TestVectRvListenerPublishSubscribe(t *testing.T) {
	subject := "UNIT.TEST.VECT"

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

	var output string

	var callback RvCallback = func(s *string) func(msg *RvMessage) {
		return func(msg *RvMessage) {
			*s = msg.String()
		}
	}(&output)

	var listener RvVectListener
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
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = transport.Send(msg)
		if err != nil {
			t.Fatalf("Expected nil, got %v", err)
		}
	}()
	wg.Wait()
	if err := queue.Dispatch(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	input := msg.String()

	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
	}
}
