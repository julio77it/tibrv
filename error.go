package tibrv

/*
#include <tibrv/tibrv.h>
*/
import "C"
import "fmt"

// RvError Tibco RendezVous error
type RvError struct {
	Code int
	Text string
}

// Error implements the error interface
func (e *RvError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Text)
}

// String convert to message string representation
func (e *RvError) String() string {
	return e.Error()
}

// NewRvError create Tibco RendezVous error
func NewRvError(code C.tibrv_status) *RvError {
	if code == C.TIBRV_OK {
		return nil
	}
	return &RvError{
		Code: int(code),
		Text: C.GoString(C.tibrvStatus_GetText(code)),
	}
}
