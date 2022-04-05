package tibrv

/*
#include <tibrv/tibrv.h>
#include <tibrv/cm.h>
#include <stdlib.h>

void cmcallbackProxy(tibrvcmEvent event,
				     tibrvMsg     message,
					 void*        closure);

*/
import "C"
import (
	"unsafe"

	"github.com/mattn/go-pointer"
)

// RvCmListener event for listening on subject
type RvCmListener struct {
	internal C.tibrvcmEvent
}

//export cmcallbackProxy
func cmcallbackProxy(cEvent C.tibrvcmEvent, cMessage C.tibrvMsg, closure unsafe.Pointer) {
	var msg RvMessage

	if err := msg.create(cMessage); err != nil {
		panic(err)
	}

	callback := *pointer.Restore(closure).(*RvCallback)

	callback(&msg)

	if err := msg.Destroy(); err != nil {
		panic(err)
	}
}

// Create initialize listener and start to collect message in queue
func (l *RvCmListener) Create(queue RvQueue, callback RvCallback, transport RvCmTransport, subject string) error {
	cstr := C.CString(subject)
	defer C.free(unsafe.Pointer(cstr)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvcmEvent_CreateListener(
		&l.internal,
		queue.internal,
		(*[0]byte)(C.cmcallbackProxy),
		C.uint(transport.getInternal()),
		cstr,
		pointer.Save(&callback))

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the listener in an invalid state, cleaning memory
func (l *RvCmListener) Destroy() error {
	status := C.tibrvcmEvent_DestroyEx(l.internal, C.TIBRVCM_PERSIST, nil)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}
