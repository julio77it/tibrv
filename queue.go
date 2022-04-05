package tibrv

/*
#include <stdlib.h>
#include <tibrv/tibrv.h>
*/
import "C"
import (
	"unsafe"
)

// RvQueue event queue
type RvQueue struct {
	internal C.tibrvQueue
}

// Create initialize the event queue. Options: label
func (q *RvQueue) Create(opts ...QueueOption) error {
	conf := queueConfig{}
	for _, opt := range opts {
		conf = opt(conf)
	}
	if status := C.tibrvQueue_Create(&q.internal); status != C.TIBRV_OK {
		return NewRvError(status)
	}
	if len(conf.Label) > 0 {
		var label *C.char
		label = C.CString(conf.Label)
		defer C.free(unsafe.Pointer(label)) //#nosec G103 -- unsafe needed by CGO

		if status := C.tibrvQueue_SetName(q.internal, label); status != C.TIBRV_OK {
			return NewRvError(status)
		}
	}
	return nil
}

// Destroy put the queue in an invalid state, cleaning memory
func (q *RvQueue) Destroy() error {
	status := C.tibrvQueue_DestroyEx(q.internal, nil, nil)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Dispatch ask for next message, if empty, blocks
func (q *RvQueue) Dispatch() error {
	status := C.tibrvQueue_TimedDispatch(q.internal, C.TIBRV_WAIT_FOREVER)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Poll ask for next message, if empty continue
func (q *RvQueue) Poll() error {
	status := C.tibrvQueue_TimedDispatch(q.internal, C.TIBRV_NO_WAIT)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// TimedDispatch ask for next message, if empty wait for
func (q *RvQueue) TimedDispatch(timeout float64) error {
	status := C.tibrvQueue_TimedDispatch(q.internal, C.double(timeout))

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetCount number of pending messages in queue
func (q *RvQueue) GetCount() (uint32, error) {
	var result C.uint
	status := C.tibrvQueue_GetCount(q.internal, &result)

	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint32(result), nil
}

// GetPriority return queue priority number
func (q *RvQueue) GetPriority() (uint32, error) {
	var result C.uint
	status := C.tibrvQueue_GetPriority(q.internal, &result)

	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint32(result), nil
}

// SetPriority set queue priority number
func (q *RvQueue) SetPriority(priority uint32) error {
	var cpriority C.uint = C.uint(priority)
	status := C.tibrvQueue_SetPriority(q.internal, cpriority)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// queueConfig Network Transport configuratio
type queueConfig struct {
	Label string
}

// QueueOption func type for alter default net transport options
type QueueOption = func(t queueConfig) queueConfig

// Label RendezVous bus service option
func Label(label string) QueueOption {
	return func(t queueConfig) queueConfig {
		t.Label = label
		return t
	}
}

// RvQueueGroup event queue
type RvQueueGroup struct {
	internal C.tibrvQueueGroup
}

// Create initialize the event queue
func (g *RvQueueGroup) Create() error {
	if status := C.tibrvQueueGroup_Create(&g.internal); status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the queue in an invalid state, cleaning memory
func (g *RvQueueGroup) Destroy() error {
	if status := C.tibrvQueueGroup_Destroy(g.internal); status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Dispatch ask for next message, if empty, blocks
func (g *RvQueueGroup) Dispatch() error {
	status := C.tibrvQueueGroup_TimedDispatch(g.internal, C.TIBRV_WAIT_FOREVER)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Poll ask for next message, if empty continue
func (g *RvQueueGroup) Poll() error {
	status := C.tibrvQueueGroup_TimedDispatch(g.internal, C.TIBRV_NO_WAIT)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// TimedDispatch ask for next message, if empty wait for
func (g *RvQueueGroup) TimedDispatch(timeout float64) error {
	status := C.tibrvQueueGroup_TimedDispatch(g.internal, C.double(timeout))

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Add add a queue to the group
func (g *RvQueueGroup) Add(q RvQueue) error {
	status := C.tibrvQueueGroup_Add(g.internal, q.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Remove remove a queue from the group
func (g *RvQueueGroup) Remove(q RvQueue) error {
	status := C.tibrvQueueGroup_Remove(g.internal, q.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}
