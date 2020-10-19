package tibrv

import (
	"os"
	"testing"
	"time"
)

func TestDqListenerPublishSubscribe(t *testing.T) {
	subject := "UNIT.TEST.DQ"

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

	var transport RvDqTransport
	err = transport.Create(
		&ntransport,
		Name("TestDqListenerPublishSubscribe"),
		WorkerWeight(DefaultWorkerWeight),
		WorkerTasks(DefaultWorkerTasks),
		SchedulerWeight(DefaultSchedulerWeight),
		SchedulerHeartbeat(DefaultSchedulerHB),
		SchedulerActivation(DefaultSchedulerActive),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()
	time.Sleep(time.Second * 2)

	var output string
	var callback RvCallback = func(s *string) func(msg *RvMessage) {
		return func(msg *RvMessage) {
			*s = msg.String()
		}
	}(&output)

	var listener RvDqListener
	err = listener.Create(
		queue,
		callback,
		transport,
		subject,
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	time.Sleep(time.Second * 2)

	var msg RvMessage
	if err := msg.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := msg.SetInt32("Integer32bit", -25); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := msg.SetSendSubject(subject); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()

	input := msg.String()

	err = ntransport.Send(msg)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := queue.Dispatch(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
	}
	if err := listener.Destroy(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}
