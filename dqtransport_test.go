package tibrv

import "testing"

func TestRvDqTransportSend(t *testing.T) {
	var transport RvNetTransport
	err := transport.Create(
		Service("7500"),
		Network(""),
		Daemon("7500"),
		Description("TestDescription"),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	var dqueue RvDqTransport
	err = dqueue.Create(
		&transport,
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

	var msg RvMessage
	if err := msg.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
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
	defer msg.Destroy()

	if err := transport.Send(msg); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := dqueue.Destroy(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := transport.Destroy(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}
