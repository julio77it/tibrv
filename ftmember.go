package tibrv

/*
#include <tibrv/tibrv.h>
#include <tibrv/ft.h>
#include <stdlib.h>

void ftcallbackProxy(unsigned member,
					 char* groupName,
					 int action,
				     void* closure);
*/
import "C"
import (
	"unsafe"

	"github.com/mattn/go-pointer"
)

// FtMember event for listening on subject
type FtMember struct {
	internal C.tibrvftMember
}

//export ftcallbackProxy
func ftcallbackProxy(member C.uint, groupName *C.char, action C.int, closure unsafe.Pointer) {
	callback := *pointer.Restore(closure).(*FtCallback)

	callback(C.GoString(groupName), uint(action))
}

// Create initialize fault tollerance member group
func (ft *FtMember) Create(
	queue RvQueue,
	callback FtCallback,
	transport RvNetTransport,
	groupName string,
	weight uint16,
	activeGoal uint16,
	heartbeatInterval float64,
	preparationInterval float64,
	activationInterval float64,
) error {
	cstr := C.CString(groupName)
	defer C.free(unsafe.Pointer(cstr)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvftMember_Create(
		&ft.internal,
		queue.internal,
		(*[0]byte)(C.ftcallbackProxy),
		transport.internal,
		cstr,
		C.ushort(weight),
		C.ushort(activeGoal),
		C.double(heartbeatInterval),
		C.double(preparationInterval),
		C.double(activationInterval),
		pointer.Save(&callback),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the member in an invalid state, cleaning memory
func (ft *FtMember) Destroy() error {
	status := C.tibrvftMember_Destroy(ft.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}
