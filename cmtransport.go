package tibrv

/*
#include <stdlib.h>
#include <tibrv/tibrv.h>
#include <tibrv/cm.h>
*/
import "C"
import (
	"os"
	"unsafe"
)

// RvCmTransport a certified message delivery transport, connected to service-network-daemon RendezVous bus
type RvCmTransport struct {
	internal C.tibrvcmTransport
}

// Create initialize a certified message delivery transport. Options: service, network, daemon, description
func (t *RvCmTransport) Create(transport *RvNetTransport, opts ...CmTransportOption) error {
	conf := cmTransportConfig{}
	for _, opt := range opts {
		conf = opt(conf)
	}
	var session, ledger *C.char
	if len(conf.Session) > 0 {
		session = C.CString(conf.Session)
	}
	if len(conf.Ledger) > 0 {
		ledger = C.CString(conf.Ledger)
	}
	status := C.tibrvcmTransport_Create(&t.internal, transport.internal, session, C.TIBRV_FALSE, ledger, C.TIBRV_TRUE, nil)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the transport in an invalid state, cleaning memory
func (t *RvCmTransport) Destroy() error {
	status := C.tibrvcmTransport_Destroy(t.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Send publish/subscribe protocol : publish a message
func (t RvCmTransport) Send(msg RvMessage) error {
	status := C.tibrvcmTransport_Send(t.internal, msg.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SendRequest request-reply protocol : send a message
func (t RvCmTransport) SendRequest(req RvMessage, res *RvMessage, timeout float64) error {
	status := C.tibrvcmTransport_SendRequest(t.internal, req.internal, &res.internal, C.double(timeout))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SendReply request-reply protocol : send the response to a message
func (t RvCmTransport) SendReply(res, req RvMessage) error {
	status := C.tibrvcmTransport_SendReply(t.internal, res.internal, req.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// AddListener request-reply protocol : send the response to a message
func (t RvCmTransport) AddListener(session, subject string) error {
	csession := C.CString(session)
	defer C.free(unsafe.Pointer(csession)) //#nosec G103 -- unsafe needed by CGO
	csubject := C.CString(subject)
	defer C.free(unsafe.Pointer(csubject)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvcmTransport_AddListener(t.internal, csession, csubject)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// cmTransportConfig Certified Transport configuratiom
type cmTransportConfig struct {
	Session, Ledger string
}

// CmTransportOption func type for alter default net transport options
type CmTransportOption = func(t cmTransportConfig) cmTransportConfig

// Session certified message delivery session name option
func Session(session string) CmTransportOption {
	return func(t cmTransportConfig) cmTransportConfig {
		t.Session = session
		return t
	}
}

// Ledger certified message delivery ledger filename option
func Ledger(ledger string) CmTransportOption {
	return func(t cmTransportConfig) cmTransportConfig {
		t.Ledger = os.ExpandEnv(ledger)
		return t
	}
}

func (t RvCmTransport) getInternal() uint {
	return uint(t.internal)
}
