package tibrv

import (
	"os"
	"testing"
)

func TestRvQueue(t *testing.T) {
	var queue RvQueue

	status := queue.Create(Label("TEST_QUEUE"))
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
	count, status := queue.GetCount()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	} else if count != 0 {
		t.Fatalf("Expected 0, got %v", count)
	}
	var in uint32 = 2
	status = queue.SetPriority(in)
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
	out, status := queue.GetPriority()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	} else if out != in {
		t.Fatalf("Expected %d, got %d", out, in)
	}
	status = queue.Poll()
	if status == nil {
		t.Fatalf("Expected %d, got nil", TibrvTimeout)
	} else if status.(*RvError).Code != TibrvTimeout {
		t.Fatalf("Expected %d, got %v", TibrvTimeout, status)
	}
	status = queue.TimedDispatch(NoWait)
	if status == nil {
		t.Fatalf("Expected %d, got nil", TibrvTimeout)
	} else if status.(*RvError).Code != TibrvTimeout {
		t.Fatalf("Expected %d, got %v", TibrvTimeout, status)
	}

	status = queue.Destroy()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
}

func TestRvQueueGroup(t *testing.T) {
	var queue RvQueue
	status := queue.Create()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
	var group RvQueueGroup
	status = group.Create()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}

	status = group.Add(queue)
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
	status = group.Poll()
	if status == nil {
		t.Fatalf("Expected %d, got nil", TibrvTimeout)
	} else if status.(*RvError).Code != TibrvTimeout {
		t.Fatalf("Expected %d, got %v", TibrvTimeout, status)
	}
	status = group.TimedDispatch(NoWait)
	if status == nil {
		t.Fatalf("Expected %d, got nil", TibrvTimeout)
	} else if status.(*RvError).Code != TibrvTimeout {
		t.Fatalf("Expected %d, got %v", TibrvTimeout, status)
	}

	var transport RvNetTransport
	err := transport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", TibrvTimeout)
	}
	subject := "TEST.QGROUP"
	var listener RvListener
	err = listener.Create(
		queue,
		func(m *RvMessage) {},
		transport,
		subject,
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer listener.Destroy()

	var m RvMessage
	m.Create()
	status = m.SetSendSubject(subject)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	status = transport.Send(m)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	status = group.Dispatch()
	if status != nil {
		t.Fatalf("Expected nil, got %v", TibrvTimeout)
	}
	status = group.Remove(queue)
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
	status = queue.Destroy()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
	status = group.Destroy()
	if status != nil {
		t.Fatalf("Expected nil, got %v", status)
	}
}
