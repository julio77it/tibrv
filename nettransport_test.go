package tibrv

import (
	"os"
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
