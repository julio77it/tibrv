package tibrv

/*
#include <tibrv/tibrv.h>
#include <stdlib.h>

void callbackProxy(tibrvEvent event,
				   tibrvMsg   message,
				   void*      closure);

void vectcallbackProxy(tibrvMsg  messages[],
					   tibrv_u32 numMessages);

*/
import "C"
import (
	"github.com/mattn/go-pointer"

	"unsafe"
)

// RvListener event for listening on subject
type RvListener struct {
	internal C.tibrvEvent
}

//export callbackProxy
func callbackProxy(cEvent C.tibrvEvent, cMessage C.tibrvMsg, closure unsafe.Pointer) {
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
func (l *RvListener) Create(queue RvQueue, callback RvCallback, transport RvNetTransport, subject string) error {
	cstr := C.CString(subject)
	defer C.free(unsafe.Pointer(cstr)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvEvent_CreateListener(
		&l.internal,
		queue.internal,
		(*[0]byte)(C.callbackProxy),
		C.uint(transport.getInternal()),
		cstr,
		pointer.Save(&callback))

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the listener in an invalid state, cleaning memory
func (l *RvListener) Destroy() error {
	status := C.tibrvEvent_DestroyEx(l.internal, nil)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// RvVectListener event for listening on subject
type RvVectListener struct {
	internal C.tibrvEvent
}

//export vectcallbackProxy
func vectcallbackProxy(cMessage *C.tibrvMsg, numMessages C.tibrv_u32) {
	var i uintptr
	var len uintptr = uintptr(numMessages)

	for i = 0; i < len; i++ {
		// pointer arithmetics is discouraged, but need with C bindings
		//                                                                  |    offset in the array    |
		//                               |         array index ZERO         |
		// nested cast for get work done |
		p := (*C.tibrvMsg)(unsafe.Pointer(uintptr(unsafe.Pointer(cMessage)) + i*unsafe.Sizeof(*cMessage))) //#nosec G103 -- unsafe needed by CGO

		var closure unsafe.Pointer
		C.tibrvMsg_GetClosure(*p, &closure)

		callback := *pointer.Restore(closure).(*RvCallback)

		var msg RvMessage
		if err := msg.create(*p); err != nil {
			panic(err)
		}
		callback(&msg)
		if err := msg.Destroy(); err != nil {
			panic(err)
		}
	}
}

// Create initialize listener and start to collect message in queue
func (l *RvVectListener) Create(queue RvQueue, callback RvCallback, transport RvNetTransport, subject string) error {
	cstr := C.CString(subject)
	defer C.free(unsafe.Pointer(cstr)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvEvent_CreateVectorListener(
		&l.internal,
		queue.internal,
		(*[0]byte)(C.vectcallbackProxy),
		C.uint(transport.getInternal()),
		cstr,
		pointer.Save(&callback))

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the listener in an invalid state, cleaning memory
func (l *RvVectListener) Destroy() error {
	status := C.tibrvEvent_DestroyEx(l.internal, nil)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}
