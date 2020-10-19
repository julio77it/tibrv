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

// Create initialize a distribuited queue transport. Options: name, workerWeight, workerTasks, schedulerWeight, schedulerHeartbeat, schedulerActivation
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
		C.free(unsafe.Pointer(name))
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

// dqTransportConfig Distribuited Queue configuratiom
type dqTransportConfig struct {
	Name                                    string
	WorkerWeight, WorkerTasks               uint32
	SchedulerWeight                         uint16
	SchedulerHeartbeat, SchedulerActivation float64
}

// DqTransportOption func type for alter default distribuited queue transport options
type DqTransportOption = func(t dqTransportConfig) dqTransportConfig

// Name Distribuited queue name option
func Name(name string) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.Name = name
		return t
	}
}

// WorkerWeight Distribuited queue workerWeight option
func WorkerWeight(workerWeight uint32) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.WorkerWeight = workerWeight
		return t
	}
}

// WorkerTasks Distribuited queue workerTasks option
func WorkerTasks(workerTasks uint32) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.WorkerTasks = workerTasks
		return t
	}
}

// SchedulerWeight Distribuited queue schedulerWeight option
func SchedulerWeight(schedulerWeight uint16) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.SchedulerWeight = schedulerWeight
		return t
	}
}

// SchedulerHeartbeat Distribuited queue schedulerHeartbeat option
func SchedulerHeartbeat(schedulerHeartbeat float64) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.SchedulerHeartbeat = schedulerHeartbeat
		return t
	}
}

// SchedulerActivation Distribuited queue schedulerActivation option
func SchedulerActivation(schedulerActivation float64) DqTransportOption {
	return func(t dqTransportConfig) dqTransportConfig {
		t.SchedulerActivation = schedulerActivation
		return t
	}
}

func (t RvDqTransport) getInternal() uint {
	return uint(t.internal)
}
