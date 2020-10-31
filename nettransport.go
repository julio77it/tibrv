package tibrv

/*
#include <stdlib.h>
#include <tibrv/tibrv.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// RvNetTransport a regular transport, connected to service-network-daemon RendezVous bus
type RvNetTransport struct {
	internal C.tibrvTransport
}

// Create initialize a regular transport. Options: service, network, daemon, description
func (t *RvNetTransport) Create(opts ...NetTransportOption) error {
	conf := netTransportConfig{}
	for _, opt := range opts {
		conf = opt(conf)
	}
	var service, network, daemon, description *C.char
	if len(conf.Service) > 0 {
		service = C.CString(conf.Service)
	}
	if len(conf.Network) > 0 {
		network = C.CString(conf.Network)
	}
	if len(conf.Daemon) > 0 {
		daemon = C.CString(conf.Daemon)
	}
	status := C.tibrvTransport_Create(&t.internal, service, network, daemon)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}

	if len(conf.Description) > 0 {
		description = C.CString(conf.Description)
		status = C.tibrvTransport_SetDescription(t.internal, description)
		if status != C.TIBRV_OK {
			return NewRvError(status)
		}
	}
	return nil
}

// Destroy put the transport in an invalid state, cleaning memory
func (t *RvNetTransport) Destroy() error {
	status := C.tibrvTransport_Destroy(t.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Send publish/subscribe protocol : publish a message
func (t RvNetTransport) Send(msg RvMessage) error {
	status := C.tibrvTransport_Send(t.internal, msg.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SendRequest request-reply protocol : send a message
func (t RvNetTransport) SendRequest(req RvMessage, res *RvMessage, timeout float64) error {
	status := C.tibrvTransport_SendRequest(t.internal, req.internal, &res.internal, C.double(timeout))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SendReply request-reply protocol : send the response to a message
func (t RvNetTransport) SendReply(res, req RvMessage) error {
	status := C.tibrvTransport_SendReply(t.internal, res.internal, req.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// CreateInbox create a 64byte inbox
func (t RvNetTransport) CreateInbox() (string, error) {
	buffer := C.CString(fmt.Sprintf("%64v", ""))
	defer C.free(unsafe.Pointer(buffer))

	status := C.tibrvTransport_CreateInbox(t.internal, buffer, C.tibrv_u32(64))

	if status != C.TIBRV_OK {
		return "", NewRvError(status)
	}
	return C.GoString(buffer), nil
}

// netTransportConfig Network Transport configuratio
type netTransportConfig struct {
	Service, Network, Daemon, Description string
}

// NetTransportOption func type for alter default net transport options
type NetTransportOption = func(t netTransportConfig) netTransportConfig

// Service RendezVous bus service option
func Service(service string) NetTransportOption {
	return func(t netTransportConfig) netTransportConfig {
		t.Service = service
		return t
	}
}

// Network RendezVous bus service option
func Network(network string) NetTransportOption {
	return func(t netTransportConfig) netTransportConfig {
		t.Network = network
		return t
	}
}

// Daemon RendezVous bus service option
func Daemon(daemon string) NetTransportOption {
	return func(t netTransportConfig) netTransportConfig {
		t.Daemon = daemon
		return t
	}
}

// Description RendezVous bus service option
func Description(description string) NetTransportOption {
	return func(t netTransportConfig) netTransportConfig {
		t.Description = description
		return t
	}
}

func (t RvNetTransport) getInternal() uint {
	return uint(t.internal)
}
