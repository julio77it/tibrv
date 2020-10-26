package tibrv

import (
	"fmt"
)

func init() {
	err := Open()
	if err != nil {
		fmt.Println(err)
	}
}

// WaitForEver as timeout value in RvTransport.SendRequest, wait until receive a response
const WaitForEver float64 = -1

// NoWait as timeout valie in RvTransport.SendRequest, doesn't wait if no message in queue
const NoWait float64 = 0

// DefaultWorkerWeight When the scheduler receives a task, it assigns the task to the available worker with the greatest worker weight
const DefaultWorkerWeight uint32 = 1

// DefaultWorkerTasks Task capacity is the maximum number of tasks that a worker can accept.
const DefaultWorkerTasks uint32 = 1

// DefaultSchedulerWeight represents the ability of this member to fulfill the role of scheduler, relative to other members with the same name
const DefaultSchedulerWeight uint16 = 1

// DefaultSchedulerHB The scheduler sends heartbeat messages at this interval (in seconds)
const DefaultSchedulerHB float64 = 1.0

// DefaultSchedulerActive When the heartbeat signal from the scheduler has been silent for this interval (in seconds), the member with the greatest scheduler weight takes its place as the new scheduler
const DefaultSchedulerActive float64 = 3.5

// RvTransport a transport object represents a delivery mechanism for messages.
type RvTransport interface {
	Send(msg RvMessage) error
	SendRequest(req RvMessage, res *RvMessage, timeout float64) error
	SendReply(res, req RvMessage) error
	getInternal() uint
}

// RvDispatchable Common interface for queues and queue groups
type RvDispatchable interface {
	Dispatch() error
	Poll() error
	TimedDispatch() error
}

// RvCallback signature for the message handler function
type RvCallback = func(*RvMessage)

// FtPrepareToActivate fault tollerance status : near to activate
const FtPrepareToActivate = 1

// FtActivate fault tollerance status : activating
const FtActivate = 2

// FtDeactivate fault tollerance status : deactivating
const FtDeactivate = 3

// FtCallback signature for the fault tollerance event handler function
type FtCallback = func(string, uint)
