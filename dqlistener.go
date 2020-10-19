package tibrv

/*
#include <tibrv/tibrv.h>
#include <tibrv/cm.h>
#include <stdlib.h>

void dqcallbackProxy(tibrvcmEvent event,
				     tibrvMsg     message,
					 void*        closure);
*/
import "C"
import (
	"github.com/mattn/go-pointer"
	"unsafe"
)

// RvDqListener event for listening on subject
type RvDqListener struct {
	internal C.tibrvcmEvent
}

//export dqcallbackProxy
func dqcallbackProxy(cEvent C.tibrvcmEvent, cMessage C.tibrvMsg, closure unsafe.Pointer) {
	var msg RvMessage

	msg.create(cMessage)

	callback := *pointer.Restore(closure).(*RvCallback)

	callback(&msg)

	msg.Destroy()
}

// Create initialize listener and start to collect message in queue
func (l *RvDqListener) Create(queue RvQueue, callback RvCallback, transport RvDqTransport, subject string) error {
	cstr := C.CString(subject)
	defer C.free(unsafe.Pointer(cstr))

	status := C.tibrvcmEvent_CreateListener(
		&l.internal,
		queue.internal,
		(*[0]byte)(C.dqcallbackProxy),
		C.uint(transport.getInternal()),
		cstr,
		pointer.Save(&callback))

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the listener in an invalid state, cleaning memory
func (l *RvDqListener) Destroy() error {
	status := C.tibrvcmEvent_DestroyEx(l.internal, C.TIBRVCM_CANCEL, nil)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}
