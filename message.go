package tibrv

/*
#include <tibrv/tibrv.h>
#include <stdlib.h>
#include <malloc.h>
*/
import "C"
import "unsafe"

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

	var sendSubject, replySubject *C.char
	C.tibrvMsg_GetSendSubject(src, &sendSubject)
	C.tibrvMsg_SetSendSubject(m.internal, sendSubject)
	C.tibrvMsg_GetReplySubject(src, &replySubject)
	C.tibrvMsg_SetReplySubject(m.internal, replySubject)

	if status != C.TIBRV_OK {
		return NewRvError(status)
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
	defer C.free(unsafe.Pointer(cstr))
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
	defer C.free(unsafe.Pointer(cstr))
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

// GetInt8 read a 8bit integer field
func (m RvMessage) GetInt8(name string) (int8, error) {
	return m.getInt8(name, 0)
}
func (m RvMessage) getInt8(name string, fieldID FieldID) (int8, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
	var cv C.ulonglong

	status := C.tibrvMsg_GetU64Ex(m.internal, cn, &cv, C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return 0, NewRvError(status)
	}
	return uint64(cv), nil
}

// SetInt8 add a 8bit integer field
func (m *RvMessage) SetInt8(name string, value int8) error {
	return m.setInt8(name, 0, value)
}
func (m *RvMessage) setInt8(name string, fieldID FieldID, value int8) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))

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
	defer C.free(unsafe.Pointer(cn))
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
	defer C.free(unsafe.Pointer(cn))
	cv := C.CString(value)
	defer C.free(unsafe.Pointer(cv))

	status := C.tibrvMsg_UpdateStringEx(m.internal, cn, cv, C.ushort(fieldID))
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
	defer C.free(unsafe.Pointer(cn))
	var cv C.tibrvMsg
	var result RvMessage

	status := C.tibrvMsg_GetMsgEx(m.internal, cn, &cv, C.ushort(fieldID))
	if status == C.TIBRV_OK {
		err := result.create(cv)
		if err != nil {
			return result, err
		}
	}
	return result, NewRvError(status)
}

// SetRvMessage add a nested message
func (m *RvMessage) SetRvMessage(name string, value RvMessage) error {
	return m.setRvMessage(name, 0, value)
}
func (m *RvMessage) setRvMessage(name string, fieldID FieldID, value RvMessage) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))

	status := C.tibrvMsg_UpdateMsgEx(m.internal, cn, C.tibrvMsg(value.internal), C.ushort(fieldID))
	if status != C.TIBRV_OK {
		return NewRvError(status)
	}
	return nil
}

// arrayItemPositionPointer pointer arithmetics using bytes, index and typesize
func arrayItemPositionPointer(base uintptr, index uintptr, itemSize uintptr) unsafe.Pointer {
	return unsafe.Pointer(base + index*itemSize)
}

// GetInt8Array read a 8bit integer array field
func (m RvMessage) GetInt8Array(name string) ([]int8, error) {
	return m.getInt8Array(name, 0)
}
func (m RvMessage) getInt8Array(name string, fieldID FieldID) ([]int8, error) {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
		// cast & conversion from bytes to slice position
		result[i] = int8(*(*C.schar)(itemPointer))
	}
	return result, nil
}

// SetInt8Array add a 8bit integer field
func (m *RvMessage) SetInt8Array(name string, value []int8) error {
	return m.setInt8Array(name, 0, value)
}
func (m *RvMessage) setInt8Array(name string, fieldID FieldID, value []int8) error {
	arrayName := C.CString(name)
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.schar)(unsafe.Pointer(itemPointer)) = C.schar(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.short)(unsafe.Pointer(itemPointer)) = C.short(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.int)(unsafe.Pointer(itemPointer)) = C.int(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		//int8(*(*C.schar)(itemPointer))
		*(*C.longlong)(unsafe.Pointer(itemPointer)) = C.longlong(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		*(*C.uchar)(unsafe.Pointer(itemPointer)) = C.uchar(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		*(*C.ushort)(unsafe.Pointer(itemPointer)) = C.ushort(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		*(*C.uint)(unsafe.Pointer(itemPointer)) = C.uint(value[j])
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
	defer C.free(unsafe.Pointer(arrayName))

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
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(*arrayValues))
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
	defer C.free(unsafe.Pointer(arrayName))

	arrayLen := len(value)
	arrayValues := C.malloc(C.ulong(arrayLen * int(unsafe.Sizeof(value[0]))))
	defer C.free(unsafe.Pointer(arrayValues))

	for i, j, len := uintptr(0), 0, uintptr(arrayLen); i < len; i++ {
		// pointer arithmetics inside this function
		itemPointer := arrayItemPositionPointer(uintptr(unsafe.Pointer(arrayValues)), i, unsafe.Sizeof(value[0]))
		// cast & conversion from slice position to bytes
		*(*C.ulonglong)(unsafe.Pointer(itemPointer)) = C.ulonglong(value[j])
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
