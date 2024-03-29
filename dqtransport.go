package tibrv

/*
#include <stdlib.h>
#include <tibrv/tibrv.h>
#include <tibrv/cm.h>
*/
import "C"
import "unsafe"

// RvDqTransport the distributed queue transport transport
type RvDqTransport struct {
	internal C.tibrvcmTransport
}

// Create initialize a distributed queue transport. Options: name, workerWeight, workerTasks, schedulerWeight, schedulerHeartbeat, schedulerActivation
func (t *RvDqTransport) Create(transport *RvNetTransport, opts ...DqTransportOption) error {
	conf := dqTransportConfig{
		WorkerWeight:        C.TIBRVCM_DEFAULT_WORKER_WEIGHT,
		WorkerTasks:         C.TIBRVCM_DEFAULT_WORKER_TASKS,
		SchedulerWeight:     C.TIBRVCM_DEFAULT_SCHEDULER_WEIGHT,
		SchedulerHeartbeat:  C.TIBRVCM_DEFAULT_SCHEDULER_HB,
		SchedulerActivation: C.TIBRVCM_DEFAULT_SCHEDULER_ACTIVE,
	}
	for _, opt := range opts {
		conf = opt(conf)
	}
	var name *C.char = nil
	var workerWeight, workerTasks C.uint
	var schedulerWeight C.ushort
	var schedulerHeartbeat, schedulerActivation C.double
	if len(conf.Name) > 0 {
		name = C.CString(conf.Name)
	}
	if conf.WorkerWeight > 0.0 {
		workerWeight = C.uint(conf.WorkerWeight)
	}
	if conf.WorkerTasks > 0.0 {
		workerTasks = C.uint(conf.WorkerTasks)
	}
	if conf.SchedulerWeight > 0.0 {
		schedulerWeight = C.ushort(conf.SchedulerWeight)
	}
	if conf.SchedulerHeartbeat > 0.0 {
		schedulerHeartbeat = C.double(conf.SchedulerHeartbeat)
	}
	if conf.SchedulerActivation > 0.0 {
		schedulerActivation = C.double(conf.SchedulerActivation)
	}
	status := C.tibrvcmTransport_CreateDistributedQueueEx(
		&t.internal,
		transport.internal,
		name,
		workerWeight,
		workerTasks,
		schedulerWeight,
		schedulerHeartbeat,
		schedulerActivation,
	)
	if name != nil {
		C.free(unsafe.Pointer(name)) //#nosec G103 -- unsafe needed by CGO
	}
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Destroy put the transport in an invalid state, cleaning memory
func (t *RvDqTransport) Destroy() error {
	status := C.tibrvcmTransport_Destroy(t.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// dqTransportConfig Distributed Queue configuratiom
type dqTransportConfig struct {
	Name                                    string
	WorkerWeight, WorkerTasks               uint32
	SchedulerWeight                         uint16
	SchedulerHeartbeat, SchedulerActivation float64
}

// DqTransportOption func type for alter default distributed queue transport options
type DqTransportOption = func(t dqTransportConfig) dqTransportConfig

// Name Distributed queue name option
func Name(name string) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.Name = name
		return t
	}
}

// WorkerWeight Distributed queue workerWeight option
func WorkerWeight(workerWeight uint32) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.WorkerWeight = workerWeight
		return t
	}
}

// WorkerTasks Distributed queue workerTasks option
func WorkerTasks(workerTasks uint32) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.WorkerTasks = workerTasks
		return t
	}
}

// SchedulerWeight Distributed queue schedulerWeight option
func SchedulerWeight(schedulerWeight uint16) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.SchedulerWeight = schedulerWeight
		return t
	}
}

// SchedulerHeartbeat Distributed queue schedulerHeartbeat option
func SchedulerHeartbeat(schedulerHeartbeat float64) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.SchedulerHeartbeat = schedulerHeartbeat
		return t
	}
}

// SchedulerActivation Distributed queue schedulerActivation option
func SchedulerActivation(schedulerActivation float64) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.SchedulerActivation = schedulerActivation
		return t
	}
}

func (t RvDqTransport) getInternal() uint {
	return uint(t.internal)
}
