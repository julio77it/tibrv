package tibrv

/*
#include <tibrv/tibrv.h>
*/
import "C"

// Open init Tibco RendezVous Enviroment
func Open() error {
	if code := C.tibrv_Open(); code != C.TIBRV_OK {
		return NewRvError(code)
	}
	return nil
}

// Close destroy Tibco RendezVous Enviroment
func Close() error {
	if code := C.tibrv_Close(); code != C.TIBRV_OK {
		return NewRvError(code)
	}
	return nil
}
