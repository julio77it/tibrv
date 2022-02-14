package tibrv

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestRvMessage(t *testing.T) {
	var msg RvMessage

	err := msg.Create()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.Destroy()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}

func TestRvMessageCreate(t *testing.T) {
	var src, dst RvMessage

	sendSubject := "TEST.REQUEST"
	replySubject := "TEST.REPLY"

	src.Create()
	defer src.Destroy()
	src.SetSendSubject(sendSubject)
	src.SetReplySubject(replySubject)

	dst.create(src.internal)
	defer dst.Destroy()

	dstSendSubject, _ := dst.GetSendSubject()
	dstreplySubject, _ := dst.GetReplySubject()

	if sendSubject != dstSendSubject {
		t.Fatalf("Expected %s, got %s", sendSubject, dstSendSubject)
	}
	if replySubject != dstreplySubject {
		t.Fatalf("Expected %s, got %s", replySubject, dstreplySubject)
	}
}

func TestRvMessageSendSubject(t *testing.T) {
	var msg RvMessage

	in := "SUBJECT.PROVA"

	msg.Create()
	defer msg.Destroy()

	err := msg.SetSendSubject(in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetSendSubject()
	if err != nil {
		t.Fatalf("Expected %s, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %s, got %s", in, out)
	}
}

func TestRvMessageReplySubject(t *testing.T) {
	var msg RvMessage

	in := "SUBJECT.PROVA"

	msg.Create()
	defer msg.Destroy()

	err := msg.SetReplySubject(in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetReplySubject()
	if err != nil {
		t.Fatalf("Expected %s, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %s, got %s", in, out)
	}
}

func TestRvMessageToString(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in int8 = -120

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt8(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	expected := "{fieldName=-120}"
	out := msg.String()
	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestRvMessageString(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := "ciao ciao"

	msg.Create()
	defer msg.Destroy()

	err := msg.SetString(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetString(name)
	if err != nil {
		t.Fatalf("Expected %s, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %s, got %s", in, out)
	}
}

func TestRvMessageBoolTrue(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in bool = true

	msg.Create()
	defer msg.Destroy()

	err := msg.SetBool(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetBool(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetBool(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageBoolFalse(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in bool = false

	msg.Create()
	defer msg.Destroy()

	err := msg.SetBool(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetBool(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetBool(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt8(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in int8 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt8(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt8(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetInt8(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt16(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in int16 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt16(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt16(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetInt16(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt32(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in int32 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt32(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt32(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetInt32(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt64(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in int64 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt64(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt64(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	out, err = msg.GetInt64(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt8(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in uint8 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt8(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt8(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetUInt8(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt16(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in uint16 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt16(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt16(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetUInt16(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt32(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in uint32 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt32(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt32(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetUInt32(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt64(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in uint64 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt64(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt64(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	_, err = msg.GetUInt64(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageFloat32(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in float32 = 3.145

	msg.Create()
	defer msg.Destroy()

	err := msg.SetFloat32(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetFloat32(name)
	if err != nil {
		t.Fatalf("Expected %f, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %f, got %f", in, out)
	}
	_, err = msg.GetFloat32(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageFloat64(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in float64 = 3.145

	msg.Create()
	defer msg.Destroy()

	if err := msg.SetFloat64(name, in); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetFloat64(name)
	if err != nil {
		t.Fatalf("Expected %f, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %f, got %f", in, out)
	}
	_, err = msg.GetFloat64(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageRemove(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	var in int8 = 126

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt8(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt8(name)
	if err != nil {
		t.Fatalf("Expected %d, got %v", in, err)
	}
	if out != in {
		t.Fatalf("Expected %d, got %d", in, out)
	}
	err = msg.RemoveField(name)
	if err != nil {
		t.Fatalf("Expected ERR, got nil")
	}
	_, err = msg.GetInt8(name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageMsg(t *testing.T) {
	var msg RvMessage

	msg.Create()
	defer msg.Destroy()

	name := "fieldName"
	var inV float64 = 3.145

	var in RvMessage
	in.Create()
	in.SetFloat64(name, inV)
	msg.SetRvMessage(name, in)

	out, _ := msg.GetRvMessage(name)
	outV, _ := out.GetFloat64(name)

	if outV != inV {
		t.Fatalf("Expected %f, got %f", inV, outV)
	}
}

func TestArrayItemPositionPointer(t *testing.T) {
	if expected, got := uintptr(8), uintptr(arrayItemPositionPointer(0, 8, 1)); expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
	if expected, got := uintptr(16), uintptr(arrayItemPositionPointer(0, 8, 2)); expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
	if expected, got := uintptr(32), uintptr(arrayItemPositionPointer(0, 8, 4)); expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
	if expected, got := uintptr(64), uintptr(arrayItemPositionPointer(0, 8, 8)); expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
}

func TestRvMessageStringArray(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []string{"ABC", "DEF", "GHI"}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetStringArray(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetStringArray(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetStringArray(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt8Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []int8{126, 125, 124}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt8Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt8Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetInt8Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt16Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []int16{126, 125, 124}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt16Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt16Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetInt16Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt32Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := make([]int32, 0, 3)
	in = append(in, 126, 125, 124)

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt32Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt32Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetInt32Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageInt64Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []int64{126, 125, 124}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetInt64Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetInt64Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetInt64Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt8Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []uint8{126, 125, 124}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt8Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt8Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetUInt8Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt16Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []uint16{126, 125, 124}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt16Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt16Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetUInt16Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt32Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := make([]uint32, 0, 3)
	in = append(in, 126, 125, 124)

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt32Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt32Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetUInt32Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageUInt64Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []uint64{126, 125, 124}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetUInt64Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetUInt64Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetUInt64Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageFloat32Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []float32{5.0, 3.145, 1234567.90}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetFloat32Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetFloat32Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetFloat32Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageFloat64Array(t *testing.T) {
	var msg RvMessage

	name := "fieldName"
	in := []float64{5.0, 3.145, 1234567.90}

	msg.Create()
	defer msg.Destroy()

	err := msg.SetFloat64Array(name, in)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	out, err := msg.GetFloat64Array(name)
	if err != nil {
		t.Fatalf("Expected %v, got %v", in, err)
	}
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("Expected %v, got %v", in, out)
	}
	_, err = msg.GetFloat64Array(name + name)
	if err == nil {
		t.Fatalf("Expected ERR, got nil")
	}
}

func TestRvMessageGetFields(t *testing.T) {
	var msg RvMessage

	err := msg.Create()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetString("String", "string")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt8("Int8", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt8("UInt8", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt16("Int16", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt16("UInt16", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt32("Int32", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt32("UInt32", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt64("Int64", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt64("UInt64", 1)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetFloat32("Float32", 1.0)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetFloat64("Float64", 1.0)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetStringArray("StringArray", []string{"string1", "string2"})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt8Array("Int8Array", []int8{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt8Array("UInt8Array", []uint8{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt16Array("Int16Array", []int16{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt16Array("UInt16Array", []uint16{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt32Array("Int32Array", []int32{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt32Array("UInt32Array", []uint32{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetInt64Array("Int64Array", []int64{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetUInt64Array("UInt64Array", []uint64{1, 2})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetFloat32Array("Float32Array", []float32{1.2, 2.3})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = msg.SetFloat64Array("Float64Array", []float64{1.2, 2.3})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	var innerMsg RvMessage
	innerMsg.Create()
	defer innerMsg.Destroy()
	err = msg.SetRvMessage("RvMessage", innerMsg)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	fields, err := msg.GetFields()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if fieldType, ok := fields["String"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeString {
		t.Fatalf("Expected %v, got %v", FieldTypeString, fieldType)
	}
	if fieldType, ok := fields["Int8"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt8 {
		t.Fatalf("Expected %v, got %v", FieldTypeInt8, fieldType)
	}
	if fieldType, ok := fields["UInt8"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt8 {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt8, fieldType)
	}
	if fieldType, ok := fields["Int16"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt16 {
		t.Fatalf("Expected %v, got %v", FieldTypeInt16, fieldType)
	}
	if fieldType, ok := fields["UInt16"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt16 {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt16, fieldType)
	}
	if fieldType, ok := fields["Int32"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt32 {
		t.Fatalf("Expected %v, got %v", FieldTypeInt32, fieldType)
	}
	if fieldType, ok := fields["UInt32"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt32 {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt32, fieldType)
	}
	if fieldType, ok := fields["Int64"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt64 {
		t.Fatalf("Expected %v, got %v", FieldTypeInt64, fieldType)
	}
	if fieldType, ok := fields["UInt64"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt64 {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt64, fieldType)
	}
	if fieldType, ok := fields["Float32"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeFloat32 {
		t.Fatalf("Expected %v, got %v", FieldTypeFloat32, fieldType)
	}
	if fieldType, ok := fields["Float64"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeFloat64 {
		t.Fatalf("Expected %v, got %v", FieldTypeFloat64, fieldType)
	}
	if fieldType, ok := fields["StringArray"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeStringArray {
		t.Fatalf("Expected %v, got %v", FieldTypeStringArray, fieldType)
	}
	if fieldType, ok := fields["Int8Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt8Array {
		t.Fatalf("Expected %v, got %v", FieldTypeInt8Array, fieldType)
	}
	if fieldType, ok := fields["UInt8Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt8Array {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt8Array, fieldType)
	}
	if fieldType, ok := fields["Int16Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt16Array {
		t.Fatalf("Expected %v, got %v", FieldTypeInt16Array, fieldType)
	}
	if fieldType, ok := fields["UInt16Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt16Array {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt16Array, fieldType)
	}
	if fieldType, ok := fields["Int32Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt32Array {
		t.Fatalf("Expected %v, got %v", FieldTypeInt32Array, fieldType)
	}
	if fieldType, ok := fields["UInt32Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt32Array {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt32Array, fieldType)
	}
	if fieldType, ok := fields["Int64Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeInt64Array {
		t.Fatalf("Expected %v, got %v", FieldTypeInt64Array, fieldType)
	}
	if fieldType, ok := fields["UInt64Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeUInt64Array {
		t.Fatalf("Expected %v, got %v", FieldTypeUInt64Array, fieldType)
	}
	if fieldType, ok := fields["Float32Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeFloat32Array {
		t.Fatalf("Expected %v, got %v", FieldTypeFloat32Array, fieldType)
	}
	if fieldType, ok := fields["Float64Array"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeFloat64Array {
		t.Fatalf("Expected %v, got %v", FieldTypeFloat64Array, fieldType)
	}
	if fieldType, ok := fields["RvMessage"]; !ok {
		t.Fatal("Expected ok, got !ok")
	} else if fieldType != FieldTypeMsg {
		t.Fatalf("Expected %v, got %v", FieldTypeMsg, fieldType)
	}
}
func TestRvMessageGetNumFields(t *testing.T) {
	var msg RvMessage

	err := msg.Create()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	n, err := msg.GetNumFields()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if n != 0 {
		t.Fatalf("Expected 0, got %v", n)
	}
	err = msg.SetString("String1", "string1")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	n, err = msg.GetNumFields()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if n != 1 {
		t.Fatalf("Expected 1, got %v", n)
	}
	err = msg.SetString("String2", "string2")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	n, err = msg.GetNumFields()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if n != 2 {
		t.Fatalf("Expected 2, got %v", n)
	}
}

type AutoGenerated struct {
	Int16        int   `json:"Int16"`
	UInt8        int   `json:"UInt8"`
	UInt32Array  []int `json:"UInt32Array"`
	UInt64       int   `json:"UInt64"`
	InnerMessage struct {
		DoveSono string `json:"Dove sono"`
	} `json:"InnerMessage"`
	Float64Array []float64 `json:"Float64Array"`
	StringArray  []string  `json:"StringArray"`
	Int8         int       `json:"Int8"`
	Int16Array   []int     `json:"Int16Array"`
	Int32Array   []int     `json:"Int32Array"`
	UInt16Array  []int     `json:"UInt16Array"`
	String       string    `json:"String"`
	Int64Array   []int     `json:"Int64Array"`
	UInt8Array   []int     `json:"UInt8Array"`
	Float32Array []float64 `json:"Float32Array"`
	Float64      float64   `json:"Float64"`
	UInt64Array  []int     `json:"UInt64Array"`
	Float32      float64   `json:"Float32"`
	Int8Array    []int     `json:"Int8Array"`
	Int32        int       `json:"Int32"`
	Int64        int       `json:"Int64"`
	UInt16       int       `json:"UInt16"`
	UInt32       int       `json:"UInt32"`
}

func TestRvMessageJSON(t *testing.T) {
	var msg RvMessage

	err := msg.Create()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()

	msg.SetString("String", "text1")
	msg.SetStringArray("StringArray", []string{"text1.1", "text1.2"})
	msg.SetInt8("Int8", 1)
	msg.SetInt8Array("Int8Array", []int8{1, 2})
	msg.SetInt16("Int16", 3)
	msg.SetInt16Array("Int16Array", []int16{3, 4})
	msg.SetInt32("Int32", 5)
	msg.SetInt32Array("Int32Array", []int32{5, 6})
	msg.SetInt64("Int64", 7)
	msg.SetInt64Array("Int64Array", []int64{7, 8})
	msg.SetUInt8("UInt8", 11)
	msg.SetUInt8Array("UInt8Array", []uint8{11, 12})
	msg.SetUInt16("UInt16", 13)
	msg.SetUInt16Array("UInt16Array", []uint16{13, 14})
	msg.SetUInt32("UInt32", 15)
	msg.SetUInt32Array("UInt32Array", []uint32{15, 16})
	msg.SetUInt64("UInt64", 17)
	msg.SetUInt64Array("UInt64Array", []uint64{17, 18})
	msg.SetFloat32("Float32", 0.75)
	msg.SetFloat32Array("Float32Array", []float32{0.75, 0.751})
	msg.SetFloat64("Float64", 0.95)
	msg.SetFloat64Array("Float64Array", []float64{0.95, 0.951})

	var inner RvMessage

	err = inner.Create()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer inner.Destroy()

	inner.SetString("Dove sono", "sono dentro")
	msg.SetRvMessage("InnerMessage", inner)

	got, err := msg.JSON()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	expected := "{\"UInt8Array\": [11, 12], \"UInt16\": 13, \"UInt64\": 17, \"Float32Array\": [0.750000, 0.751000], \"Int16\": 3, \"Int64Array\": [7, 8], \"UInt32\": 15, \"UInt32Array\": [15, 16], \"UInt64Array\": [17, 18], \"InnerMessage\": {\"Dove sono\": \"sono dentro\"}, \"StringArray\": [\"text1.1\", \"text1.2\"], \"Int8Array\": [1, 2], \"Int16Array\": [3, 4], \"Int32Array\": [5, 6], \"UInt8\": 11, \"UInt16Array\": [13, 14], \"Float64Array\": [0.950000, 0.951000], \"String\": \"text1\", \"Int8\": 1, \"Int32\": 5, \"Int64\": 7, \"Float32\": 0.750000, \"Float64\": 0.950000}"

	var e, g AutoGenerated
	if err := json.Unmarshal([]byte(expected), &e); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := json.Unmarshal([]byte(got), &g); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if !reflect.DeepEqual(e, g) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestJSONRvMessage(t *testing.T) {
	var input string = "{\"nint8\": {\"inner\": \"abc\"}}"
	var output string

	var msg *RvMessage
	var err error

	if msg, err = JSON(input); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()

	if output, err = msg.JSON(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if input != output {
		t.Fatalf("Expected %s, got %s", input, output)
	}
}

func replier(subject string) {
	var queue RvQueue
	queue.Create()
	defer queue.Destroy()

	var transport RvNetTransport
	transport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
		Description("TestDescription"),
	)
	defer transport.Destroy()

	var callback RvCallback = func(t *RvNetTransport) func(msg *RvMessage) {
		return func(imsg *RvMessage) {
			var omsg RvMessage
			omsg.create(imsg.internal)
			t.SendReply(omsg, *imsg)
			omsg.Destroy()
		}
	}(&transport)

	var listener RvListener
	listener.Create(
		queue,
		callback,
		transport,
		subject,
	)
	defer listener.Destroy()

	queue.Dispatch()
	time.Sleep(time.Second)
}

func TestRvMessageDebugTheBug(t *testing.T) {
	sendSubject := "UNIT.TEST.DBG"
	replySubject := "DBG.TEST.UNIT"

	go replier(sendSubject)
	time.Sleep(time.Second * 2)

	var request, reply, inner RvMessage
	inner.Create()
	inner.SetFloat64("Pi", 3.145)
	request.Create()
	request.SetRvMessage("inner", inner)
	request.SetSendSubject(sendSubject)
	request.SetReplySubject(replySubject)
	reply.Create()

	var transport RvNetTransport
	err := transport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
		Description("TestDescription"),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	err = transport.SendRequest(request, &reply, WaitForEver)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	tmp, err := reply.GetRvMessage("inner")
	if err != nil {
		t.Fatalf("Expected %v, got %v", nil, err)
	}
	pi, err := tmp.GetFloat64("Pi")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if pi != 3.145 {
		t.Fatalf("Expected nil, got %v", err)
	}
}
