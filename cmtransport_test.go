package tibrv

import "testing"

func TestRvCmTransportSend(t *testing.T) {
	var ntransport RvNetTransport

	err := ntransport.Create(
		Service("7500"),
		Network(""),
		Daemon("7500"),
		Description("TestDescription"),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer ntransport.Destroy()

	var transport RvCmTransport
	err = transport.Create(
		&ntransport,
		Session("TestRvCmTransportSend"),
		Ledger("${TESTDIR}/TestRvCmTransportSend.ledger"),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	err = transport.AddListener(
		"SessionName",
		"Subject",
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	var msg RvMessage
	err = msg.Create()
	if err != nil {
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
	err = transport.Send(msg)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}
