package tibrv

/*
#include <tibrv/tibrv.h>
#include <tibrv/msg.h>
#include <stdlib.h>
#include <malloc.h>
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"unsafe"
)

// FieldID field identifier type
type FieldID = uint16

// RvMessage internal implementation
type RvMessage struct {
	internal C.tibrvMsg
}

// Create initialize message memory structs
func (m *RvMessage) Create() error {
	status := C.tibrvMsg_Create(&m.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

func (m *RvMessage) create(src C.tibrvMsg) error {
	status := C.tibrvMsg_CreateCopy(src, &m.internal)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}

	var sendSubject, replySubject *C.char
	status = C.tibrvMsg_GetSendSubject(src, &sendSubject)
	if status == C.TIBRV_OK {
		C.tibrvMsg_SetSendSubject(m.internal, sendSubject)
	}
	status = C.tibrvMsg_GetReplySubject(src, &replySubject)
	if status == C.TIBRV_OK {
		C.tibrvMsg_SetReplySubject(m.internal, replySubject)
	}
	return nil
}

// Destroy put the message in a invalid state, cleaning memory
func (m *RvMessage) Destroy() error {
	status := C.tibrvMsg_Destroy(m.internal)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// String return a string representation of the message
func (m RvMessage) String() string {
	var buffer *C.char
	pt := &buffer
	status := C.tibrvMsg_ConvertToString(m.internal, pt)
	if status != C.TIBRV_OK {
		return ""
	}
	return C.GoString(buffer)
}

// SetSendSubject set the publish subject. Used for publish/subscribe and request/reply (request subject)
func (m *RvMessage) SetSendSubject(subject string) error {
	cstr := C.CString(subject)
	defer C.free(unsafe.Pointer(cstr)) //#nosec G103 -- unsafe needed by CGO
	status := C.tibrvMsg_SetSendSubject(m.internal, cstr)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetSendSubject get the publish subject. Used for publish/subscribe and request/reply (request subject)
func (m *RvMessage) GetSendSubject() (string, error) {
	var cstr *C.char

	status := C.tibrvMsg_GetSendSubject(m.internal, &cstr)
	if status != C.TIBRV_OK {
		return "", NewRvError(status)
	}
	return C.GoString(cstr), nil
}

// SetReplySubject set the reply subject. Used for request/reply (reply subject)
func (m *RvMessage) SetReplySubject(subject string) error {
	cstr := C.CString(subject)
	defer C.free(unsafe.Pointer(cstr)) //#nosec G103 -- unsafe needed by CGO
	status := C.tibrvMsg_SetReplySubject(m.internal, cstr)

	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetReplySubject get the reply subject. Used for request/reply (reply subject)
func (m *RvMessage) GetReplySubject() (string, error) {
	var cstr *C.char

	status := C.tibrvMsg_GetReplySubject(m.internal, &cstr)
	if status != C.TIBRV_OK {
		return "", NewRvError(status)
	}
	return C.GoString(cstr), nil
}

// GetFields returns a map with field names as keys and field types values
func (m RvMessage) GetFields() (map[string]uint8, error) {
	fields := make(map[string]uint8)

	n, err := m.GetNumFields()

	if err != nil {
		return nil, err
	}
	for i := uint(0); i < n; i++ {
		var field C.tibrvMsgField

		status := C.tibrvMsg_GetFieldByIndex(
			m.internal,
			&field,
			C.uint(i),
		)
		if status != C.TIBRV_OK {
			return nil, NewRvError(status)
		}
		fields[C.GoString(field.name)] = uint8(field._type)
	}
	return fields, nil
}

// GetNumFields returns the number of fields of the message
func (m RvMessage) GetNumFields() (uint, error) {
	var n C.uint

	status := C.tibrvMsg_GetNumFields(
		m.internal,
		&n,
	)
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint(n), nil
}

// GetBool read a 8bit integer field
func (m RvMessage) GetBool(name string) (bool, error) {
	return m.getBool(name, 0)
}
func (m RvMessage) getBool(name string, fieldID FieldID) (bool, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.tibrv_bool

	status := C.tibrvMsg_GetBoolEx(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return false, NewRvError(status)
	}
	return cv == C.TIBRV_TRUE, nil
}

// GetInt8 read a 8bit integer field
func (m RvMessage) GetInt8(name string) (int8, error) {
	return m.getInt8(name, 0)
}
func (m RvMessage) getInt8(name string, fieldID FieldID) (int8, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.schar

	status := C.tibrvMsg_GetI8Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return int8(cv), nil
}

// GetInt16 read a 16bit integer field
func (m RvMessage) GetInt16(name string) (int16, error) {
	return m.getInt16(name, 0)
}
func (m RvMessage) getInt16(name string, fieldID FieldID) (int16, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.short

	status := C.tibrvMsg_GetI16Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return int16(cv), nil
}

// GetInt32 read a 32bit integer field
func (m RvMessage) GetInt32(name string) (int32, error) {
	return m.getInt32(name, 0)
}
func (m RvMessage) getInt32(name string, fieldID FieldID) (int32, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.int

	status := C.tibrvMsg_GetI32Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return int32(cv), nil
}

// GetInt64 read a 64bit integer field
func (m RvMessage) GetInt64(name string) (int64, error) {
	return m.getInt64(name, 0)
}
func (m RvMessage) getInt64(name string, fieldID FieldID) (int64, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.longlong

	status := C.tibrvMsg_GetI64Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return int64(cv), nil
}

// GetUInt8 read a 8bit unsigned integer field
func (m RvMessage) GetUInt8(name string) (uint8, error) {
	return m.getUInt8(name, 0)
}
func (m RvMessage) getUInt8(name string, fieldID FieldID) (uint8, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.uchar

	status := C.tibrvMsg_GetU8Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint8(cv), nil
}

// GetUInt16 read a 16bit unsigned integer field
func (m RvMessage) GetUInt16(name string) (uint16, error) {
	return m.getUInt16(name, 0)
}
func (m RvMessage) getUInt16(name string, fieldID FieldID) (uint16, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.ushort

	status := C.tibrvMsg_GetU16Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint16(cv), nil
}

// GetUInt32 read a 32bit unsigned integer field
func (m RvMessage) GetUInt32(name string) (uint32, error) {
	return m.getUInt32(name, 0)
}
func (m RvMessage) getUInt32(name string, fieldID FieldID) (uint32, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.uint

	status := C.tibrvMsg_GetU32Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint32(cv), nil
}

// GetUInt64 read a 64bit unsigned integer field
func (m RvMessage) GetUInt64(name string) (uint64, error) {
	return m.getUInt64(name, 0)
}
func (m RvMessage) getUInt64(name string, fieldID FieldID) (uint64, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.ulonglong

	status := C.tibrvMsg_GetU64Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint64(cv), nil
}

// SetBool add a 8bit integer field
func (m *RvMessage) SetBool(name string, value bool) error {
	return m.setBool(name, 0, value)
}
func (m *RvMessage) setBool(name string, fieldID FieldID, value bool) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	var v C.tibrv_bool = C.TIBRV_FALSE

	if value {
		v = C.TIBRV_TRUE
	}
	status := C.tibrvMsg_UpdateBoolEx(m.internal, cn, v, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetInt8 add a 8bit integer field
func (m *RvMessage) SetInt8(name string, value int8) error {
	return m.setInt8(name, 0, value)
}
func (m *RvMessage) setInt8(name string, fieldID FieldID, value int8) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateI8Ex(m.internal, cn, C.schar(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetInt16 add a 16bit integer field
func (m *RvMessage) SetInt16(name string, value int16) error {
	return m.setInt16(name, 0, value)
}
func (m *RvMessage) setInt16(name string, fieldID FieldID, value int16) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateI16Ex(m.internal, cn, C.short(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetInt32 add a 32bit integer field
func (m *RvMessage) SetInt32(name string, value int32) error {
	return m.setInt32(name, 0, value)
}
func (m *RvMessage) setInt32(name string, fieldID FieldID, value int32) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateI32Ex(m.internal, cn, C.int(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetInt64 add a 64bit integer field
func (m *RvMessage) SetInt64(name string, value int64) error {
	return m.setInt64(name, 0, value)
}
func (m *RvMessage) setInt64(name string, fieldID FieldID, value int64) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateI64Ex(m.internal, cn, C.longlong(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetUInt8 add a 8bit unsigned integer field
func (m *RvMessage) SetUInt8(name string, value uint8) error {
	return m.setUInt8(name, 0, value)
}
func (m *RvMessage) setUInt8(name string, fieldID FieldID, value uint8) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateU8Ex(m.internal, cn, C.uchar(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetUInt16 add a 16bit unsigned integer field
func (m *RvMessage) SetUInt16(name string, value uint16) error {
	return m.setUInt16(name, 0, value)
}
func (m *RvMessage) setUInt16(name string, fieldID FieldID, value uint16) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateU16Ex(m.internal, cn, C.ushort(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetUInt32 add a 32bit unsigned integer field
func (m *RvMessage) SetUInt32(name string, value uint32) error {
	return m.setUInt32(name, 0, value)
}
func (m *RvMessage) setUInt32(name string, fieldID FieldID, value uint32) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateU32Ex(m.internal, cn, C.uint(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// SetUInt64 add a 64bit unsigned integer field
func (m *RvMessage) SetUInt64(name string, value uint64) error {
	return m.setUInt64(name, 0, value)
}
func (m *RvMessage) setUInt64(name string, fieldID FieldID, value uint64) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateU64Ex(m.internal, cn, C.ulonglong(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetFloat32 read a 32bit float field
func (m RvMessage) GetFloat32(name string) (float32, error) {
	return m.getFloat32(name, 0)
}
func (m RvMessage) getFloat32(name string, fieldID FieldID) (float32, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.float

	status := C.tibrvMsg_GetF32Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return float32(cv), nil
}

// SetFloat32 add a 32bit float field
func (m *RvMessage) SetFloat32(name string, value float32) error {
	return m.setFloat32(name, 0, value)
}
func (m *RvMessage) setFloat32(name string, fieldID FieldID, value float32) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateF32Ex(m.internal, cn, C.float(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetFloat64 read a 64bit float field
func (m RvMessage) GetFloat64(name string) (float64, error) {
	return m.getFloat64(name, 0)
}
func (m RvMessage) getFloat64(name string, fieldID FieldID) (float64, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.double

	status := C.tibrvMsg_GetF64Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return float64(cv), nil
}

// SetFloat64 add a 64bit float field
func (m *RvMessage) SetFloat64(name string, value float64) error {
	return m.setFloat64(name, 0, value)
}
func (m *RvMessage) setFloat64(name string, fieldID FieldID, value float64) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateF64Ex(m.internal, cn, C.double(value), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetString read a string field
func (m RvMessage) GetString(name string) (string, error) {
	return m.getString(name, 0)
}
func (m RvMessage) getString(name string, fieldID FieldID) (string, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv *C.char

	status := C.tibrvMsg_GetStringEx(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return "", NewRvError(status)
	}
	return C.GoString(cv), nil
}

// SetString add a string field
func (m *RvMessage) SetString(name string, value string) error {
	return m.setString(name, 0, value)
}
func (m *RvMessage) setString(name string, fieldID FieldID, value string) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	cv := C.CString(value)
	defer C.free(unsafe.Pointer(cv)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateStringEx(m.internal, cn, cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetStringArray read a string array field
func (m RvMessage) GetStringArray(name string) ([]string, error) {
	return m.getStringArray(name, 0)
}
func (m RvMessage) getStringArray(name string, fieldID FieldID) ([]string, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues **C.char
	var arrayLen C.uint

	status := C.tibrvMsg_GetStringArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]string, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = C.GoString(*(**C.char)(itemPointer))
	}
	return result, nil
}

// SetStringArray add a string array field
func (m *RvMessage) SetStringArray(name string, value []string) error {
	return m.setStringArray(name, 0, value)
}
func (m *RvMessage) setStringArray(name string, fieldID FieldID, value []string) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var sizer *C.char
	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(sizer)))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                              //#nosec G103 -- unsafe needed by CGO

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(sizer)) //#nosec G103 -- unsafe needed by CGO

		cstr := C.CString(value[i]) // free at the end of the function
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(**C.char)(unsafe.Pointer(itemPointer)) = cstr //#nosec G103 -- unsafe needed by CGO
	}
	status := C.tibrvMsg_UpdateStringArrayEx(
		m.internal,
		arrayName,
		(**C.char)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(sizer)) //#nosec G103 -- unsafe needed by CGO

		// free
		C.free(unsafe.Pointer(*(**C.char)(itemPointer))) //#nosec G103 -- unsafe needed by CGO
	}
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetRvMessage read a nested message
func (m RvMessage) GetRvMessage(name string) (RvMessage, error) {
	return m.getRvMessage(name, 0)
}
func (m RvMessage) getRvMessage(name string, fieldID FieldID) (RvMessage, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO
	var cv C.tibrvMsg
	var result RvMessage

	status := C.tibrvMsg_Create(&cv)
	if status != C.TIBRV_OK {
		return result, NewRvError(status)
	}

	status = C.tibrvMsg_GetMsgEx(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return result, NewRvError(status)
	}
	err := result.create(cv)
	if err != nil {
		return result, err
	}
	return result, nil
}

// SetRvMessage add a nested message
func (m *RvMessage) SetRvMessage(name string, value RvMessage) error {
	return m.setRvMessage(name, 0, value)
}
func (m *RvMessage) setRvMessage(name string, fieldID FieldID, value RvMessage) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn)) //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_UpdateMsgEx(m.internal, cn, C.tibrvMsg(value.internal), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// arrayItemPositionPointer pointer arithmetics using bytes, index and typesize
func arrayItemPositionPointer(base uintptr, index uintptr, itemSize uintptr) unsafe.Pointer {
	return unsafe.Pointer(base + index*itemSize) //#nosec G103 -- unsafe needed by CGO
}

// GetInt8Array read a 8bit integer array field
func (m RvMessage) GetInt8Array(name string) ([]int8, error) {
	return m.getInt8Array(name, 0)
}
func (m RvMessage) getInt8Array(name string, fieldID FieldID) ([]int8, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.schar
	var arrayLen C.uint

	status := C.tibrvMsg_GetI8ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]int8, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = int8(*(*C.schar)(itemPointer))
	}
	return result, nil
}

// SetInt8Array add a 8bit integer array field
func (m *RvMessage) SetInt8Array(name string, value []int8) error {
	return m.setInt8Array(name, 0, value)
}
func (m *RvMessage) setInt8Array(name string, fieldID FieldID, value []int8) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.schar)(unsafe.Pointer(itemPointer)) = C.schar(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateI8ArrayEx(
		m.internal,
		arrayName,
		(*C.schar)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetInt16Array read a 16bit integer array field
func (m RvMessage) GetInt16Array(name string) ([]int16, error) {
	return m.getInt16Array(name, 0)
}
func (m RvMessage) getInt16Array(name string, fieldID FieldID) ([]int16, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.short
	var arrayLen C.uint

	status := C.tibrvMsg_GetI16ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]int16, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = int16(*(*C.short)(itemPointer))
	}
	return result, nil
}

// SetInt16Array add a 16bit integer field
func (m *RvMessage) SetInt16Array(name string, value []int16) error {
	return m.setInt16Array(name, 0, value)
}
func (m *RvMessage) setInt16Array(name string, fieldID FieldID, value []int16) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.short)(unsafe.Pointer(itemPointer)) = C.short(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateI16ArrayEx(
		m.internal,
		arrayName,
		(*C.short)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetInt32Array read a 32bit integer array field
func (m RvMessage) GetInt32Array(name string) ([]int32, error) {
	return m.getInt32Array(name, 0)
}
func (m RvMessage) getInt32Array(name string, fieldID FieldID) ([]int32, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.int
	var arrayLen C.uint

	status := C.tibrvMsg_GetI32ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]int32, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = int32(*(*C.int)(itemPointer))
	}
	return result, nil
}

// SetInt32Array add a 32bit integer field
func (m *RvMessage) SetInt32Array(name string, value []int32) error {
	return m.setInt32Array(name, 0, value)
}
func (m *RvMessage) setInt32Array(name string, fieldID FieldID, value []int32) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.int)(unsafe.Pointer(itemPointer)) = C.int(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateI32ArrayEx(
		m.internal,
		arrayName,
		(*C.int)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetInt64Array read a 64bit integer array field
func (m RvMessage) GetInt64Array(name string) ([]int64, error) {
	return m.getInt64Array(name, 0)
}
func (m RvMessage) getInt64Array(name string, fieldID FieldID) ([]int64, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.longlong
	var arrayLen C.uint

	status := C.tibrvMsg_GetI64ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]int64, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = int64(*(*C.longlong)(itemPointer))
	}
	return result, nil
}

// SetInt64Array add a 64bit integer field
func (m *RvMessage) SetInt64Array(name string, value []int64) error {
	return m.setInt64Array(name, 0, value)
}
func (m *RvMessage) setInt64Array(name string, fieldID FieldID, value []int64) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.longlong)(unsafe.Pointer(itemPointer)) = C.longlong(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateI64ArrayEx(
		m.internal,
		arrayName,
		(*C.longlong)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetUInt8Array read a 8bit integer array field
func (m RvMessage) GetUInt8Array(name string) ([]uint8, error) {
	return m.getUInt8Array(name, 0)
}
func (m RvMessage) getUInt8Array(name string, fieldID FieldID) ([]uint8, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.uchar
	var arrayLen C.uint

	status := C.tibrvMsg_GetU8ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]uint8, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = uint8(*(*C.schar)(itemPointer))
	}
	return result, nil
}

// SetUInt8Array add a 8bit integer field
func (m *RvMessage) SetUInt8Array(name string, value []uint8) error {
	return m.setUInt8Array(name, 0, value)
}
func (m *RvMessage) setUInt8Array(name string, fieldID FieldID, value []uint8) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		*(*C.uchar)(unsafe.Pointer(itemPointer)) = C.uchar(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateU8ArrayEx(
		m.internal,
		arrayName,
		(*C.uchar)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetUInt16Array read a 16bit integer array field
func (m RvMessage) GetUInt16Array(name string) ([]uint16, error) {
	return m.getUInt16Array(name, 0)
}
func (m RvMessage) getUInt16Array(name string, fieldID FieldID) ([]uint16, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.ushort
	var arrayLen C.uint

	status := C.tibrvMsg_GetU16ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]uint16, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = uint16(*(*C.ushort)(itemPointer))
	}
	return result, nil
}

// SetUInt16Array add a 16bit integer field
func (m *RvMessage) SetUInt16Array(name string, value []uint16) error {
	return m.setUInt16Array(name, 0, value)
}
func (m *RvMessage) setUInt16Array(name string, fieldID FieldID, value []uint16) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		*(*C.ushort)(unsafe.Pointer(itemPointer)) = C.ushort(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateU16ArrayEx(
		m.internal,
		arrayName,
		(*C.ushort)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetUInt32Array read a 32bit integer array field
func (m RvMessage) GetUInt32Array(name string) ([]uint32, error) {
	return m.getUInt32Array(name, 0)
}
func (m RvMessage) getUInt32Array(name string, fieldID FieldID) ([]uint32, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.uint
	var arrayLen C.uint

	status := C.tibrvMsg_GetU32ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]uint32, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = uint32(*(*C.uint)(itemPointer))
	}
	return result, nil
}

// SetUInt32Array add a 32bit integer field
func (m *RvMessage) SetUInt32Array(name string, value []uint32) error {
	return m.setUInt32Array(name, 0, value)
}
func (m *RvMessage) setUInt32Array(name string, fieldID FieldID, value []uint32) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		*(*C.uint)(unsafe.Pointer(itemPointer)) = C.uint(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateU32ArrayEx(
		m.internal,
		arrayName,
		(*C.uint)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetUInt64Array read a 64bit integer array field
func (m RvMessage) GetUInt64Array(name string) ([]uint64, error) {
	return m.getUInt64Array(name, 0)
}
func (m RvMessage) getUInt64Array(name string, fieldID FieldID) ([]uint64, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.ulonglong
	var arrayLen C.uint

	status := C.tibrvMsg_GetU64ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]uint64, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = uint64(*(*C.ulonglong)(itemPointer))
	}
	return result, nil
}

// SetUInt64Array add a 64bit integer field
func (m *RvMessage) SetUInt64Array(name string, value []uint64) error {
	return m.setUInt64Array(name, 0, value)
}
func (m *RvMessage) setUInt64Array(name string, fieldID FieldID, value []uint64) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		*(*C.ulonglong)(unsafe.Pointer(itemPointer)) = C.ulonglong(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateU64ArrayEx(
		m.internal,
		arrayName,
		(*C.ulonglong)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetFloat32Array read a 32bit float array field
func (m RvMessage) GetFloat32Array(name string) ([]float32, error) {
	return m.getFloat32Array(name, 0)
}
func (m RvMessage) getFloat32Array(name string, fieldID FieldID) ([]float32, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.float
	var arrayLen C.uint

	status := C.tibrvMsg_GetF32ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]float32, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = float32(*(*C.float)(itemPointer))
	}
	return result, nil
}

// SetFloat32Array add a 32bit float field
func (m *RvMessage) SetFloat32Array(name string, value []float32) error {
	return m.setFloat32Array(name, 0, value)
}
func (m *RvMessage) setFloat32Array(name string, fieldID FieldID, value []float32) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		*(*C.float)(unsafe.Pointer(itemPointer)) = C.float(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateF32ArrayEx(
		m.internal,
		arrayName,
		(*C.float)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// GetFloat64Array read a 64bit float array field
func (m RvMessage) GetFloat64Array(name string) ([]float64, error) {
	return m.getFloat64Array(name, 0)
}
func (m RvMessage) getFloat64Array(name string, fieldID FieldID) ([]float64, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	var arrayValues *C.double
	var arrayLen C.uint

	status := C.tibrvMsg_GetF64ArrayEx(m.internal, arrayName, &arrayValues, &arrayLen, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	// convert to slice
	result := make([]float64, uint(arrayLen))

	for i, len := uintptr(0), uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues)) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from bytes to slice position
		result[i] = float64(*(*C.double)(itemPointer))
	}
	return result, nil
}

// SetFloat64Array add a 64bit float field
func (m *RvMessage) SetFloat64Array(name string, value []float64) error {
	return m.setFloat64Array(name, 0, value)
}
func (m *RvMessage) setFloat64Array(name string, fieldID FieldID, value []float64) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName)) //#nosec G103 -- unsafe needed by CGO

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0])))) //#nosec G103 -- unsafe needed by CGO
	defer C.free(unsafe.Pointer(arrayValues))                                 //#nosec G103 -- unsafe needed by CGO

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0])) //#nosec G103 -- unsafe needed by CGO
		// cast & conversion from slice position to bytes
		*(*C.double)(unsafe.Pointer(itemPointer)) = C.double(value[j]) //#nosec G103 -- unsafe needed by CGO
		j++
	}
	status := C.tibrvMsg_UpdateF64ArrayEx(
		m.internal,
		arrayName,
		(*C.double)(arrayValues),
		C.uint(arrayLen),
		C.ushort(fieldID),
	)
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// RemoveField remove field by name
func (m *RvMessage) RemoveField(name string) error {
	return m.removeField(name, 0)
}
func (m *RvMessage) removeField(name string, fieldID FieldID) error {
	fieldName := C.CString(name)
	defer C.free(unsafe.Pointer(fieldName)) //#nosec G103 -- unsafe needed by CGO

	if status := C.tibrvMsg_RemoveFieldEx(m.internal, fieldName, C.ushort(fieldID)); status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// Return the size of a message (in bytes)
func (m RvMessage) GetByteSize() (uint32, error) {
	var size C.tibrv_u32

	status := C.tibrvMsg_GetByteSize(m.internal, &size)
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint32(size), nil
}

// Extract the data from a message as a byte sequence
func (m RvMessage) GetAsBytes() ([]byte, error) {
	var ptr unsafe.Pointer //#nosec G103 -- unsafe needed by CGO

	status := C.tibrvMsg_GetAsBytes(m.internal, &ptr)
	if status != C.TIBRV_OK {
		return nil, NewRvError(status)
	}
	size, err := m.GetByteSize()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, size)

	for i := uint32(0); i < size; i++ {
		bytes[i] = *(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))) //#nosec G103 -- unsafe needed by CGO
	}
	return bytes, nil
}

// Create a new message, and populate it with data
func (m *RvMessage) CreateFromBytes(bytes []byte) error {
	ptr := C.malloc(C.size_t(len(bytes)))
	defer C.free(ptr)

	for i := 0; i < len(bytes); i++ {
		*(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))) = bytes[i] //#nosec G103 -- unsafe needed by CGO
	}
	if status := C.tibrvMsg_CreateFromBytes(&m.internal, ptr); status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// JSON returns a json string representation of the message
func (m RvMessage) JSON() (string, error) {

	fieldList, err := m.GetFields()
	if err != nil {
		return "", err
	}
	var buffer string
	result := bytes.NewBufferString(buffer)
	fmt.Fprint(result, "{")

	i := 0
	for fieldName, fieldType := range fieldList {
		if i > 0 {
			fmt.Fprint(result, ", ")
		}
		if FieldTypeMsg == fieldType {
			fieldValue, err := m.GetRvMessage(fieldName)
			if err != nil {
				return "", err
			}
			defer fieldValue.Destroy()

			fieldValueText, err := fieldValue.JSON()
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %s", fieldName, fieldValueText)
		} else if FieldTypeString == fieldType {
			fieldValue, err := m.GetString(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": \"%s\"", fieldName, fieldValue)
		} else if FieldTypeBool == fieldType {
			fieldValue, err := m.GetBool(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %v", fieldName, fieldValue)
		} else if FieldTypeInt8 == fieldType {
			fieldValue, err := m.GetInt8(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeUInt8 == fieldType {
			fieldValue, err := m.GetUInt8(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeInt16 == fieldType {
			fieldValue, err := m.GetInt16(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeUInt16 == fieldType {
			fieldValue, err := m.GetUInt16(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeInt32 == fieldType {
			fieldValue, err := m.GetInt32(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeUInt32 == fieldType {
			fieldValue, err := m.GetUInt32(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeInt64 == fieldType {
			fieldValue, err := m.GetInt64(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeUInt64 == fieldType {
			fieldValue, err := m.GetUInt64(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %d", fieldName, fieldValue)
		} else if FieldTypeFloat32 == fieldType {
			fieldValue, err := m.GetFloat32(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %f", fieldName, fieldValue)
		} else if FieldTypeFloat64 == fieldType {
			fieldValue, err := m.GetFloat64(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": %f", fieldName, fieldValue)
		} else if FieldTypeStringArray == fieldType {
			fieldValue, err := m.GetStringArray(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "\"%s\"", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeInt8Array == fieldType {
			fieldValue, err := m.GetInt8Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeUInt8Array == fieldType {
			fieldValue, err := m.GetUInt8Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeInt16Array == fieldType {
			fieldValue, err := m.GetInt16Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeUInt16Array == fieldType {
			fieldValue, err := m.GetUInt16Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeInt32Array == fieldType {
			fieldValue, err := m.GetInt32Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeUInt32Array == fieldType {
			fieldValue, err := m.GetUInt32Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeInt64Array == fieldType {
			fieldValue, err := m.GetInt64Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeUInt64Array == fieldType {
			fieldValue, err := m.GetUInt64Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%d", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeFloat32Array == fieldType {
			fieldValue, err := m.GetFloat32Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%f", v)
			}
			fmt.Fprint(result, "]")
		} else if FieldTypeFloat64Array == fieldType {
			fieldValue, err := m.GetFloat64Array(fieldName)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(result, "\"%s\": [", fieldName)
			for j, v := range fieldValue {
				if j > 0 {
					fmt.Fprintf(result, ", ")
				}
				fmt.Fprintf(result, "%f", v)
			}
			fmt.Fprint(result, "]")
		}
		i++
	}
	fmt.Fprint(result, "}")
	return result.String(), nil
}

type jdoc = map[string]interface{}

func jdoc2RvMessage(doc jdoc) (*RvMessage, error) {
	var msg RvMessage
	if err := msg.Create(); err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(doc))
	for k := range doc {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := doc[k]
		switch v.(type) {
		case bool:
			err := msg.SetBool(k, v.(bool))
			if err != nil {
				return nil, err
			}
		case string:
			err := msg.SetString(k, v.(string))
			if err != nil {
				return nil, err
			}
		case float64:
			err := msg.SetFloat64(k, v.(float64))
			if err != nil {
				return nil, err
			}
		case []interface{}:
			l := len(v.([]interface{}))
			if l == 0 {
				continue
			}
			switch v.([]interface{})[0].(type) {
			case string:
				slice := make([]string, l, l)
				for i, vv := range v.([]interface{}) {
					slice[i] = vv.(string)
				}
				err := msg.SetStringArray(k, slice)
				if err != nil {
					return nil, err
				}
			case float64:
				slice := make([]float64, l, l)
				for i, vv := range v.([]interface{}) {
					slice[i] = vv.(float64)
				}
				err := msg.SetFloat64Array(k, slice)
				if err != nil {
					return nil, err
				}
			}
		case jdoc:
			m, err := jdoc2RvMessage(v.(jdoc))
			if err != nil {
				return nil, err
			}
			err = msg.SetRvMessage(k, *m)
			if err != nil {
				return nil, err
			}
		default:
			fmt.Println("BOH")
		}
	}
	return &msg, nil
}

// JSON convert a json string to RvMessage
func JSON(doc string) (*RvMessage, error) {
	res := make(jdoc)

	if err := json.Unmarshal([]byte(doc), &res); err != nil {
		return nil, err
	}
	return jdoc2RvMessage(res)
}
