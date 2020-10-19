package tibrv

import (
	"os"
	"testing"
)

func TestRvDqTransportSend(t *testing.T) {
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
		Name("TestRvDqTransportSend"),
		WorkerWeight(1),
		WorkerTasks(1),
		SchedulerWeight(1),
		SchedulerHeartbeat(1.0),
		SchedulerActivation(3.5),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := transport.Destroy(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}
