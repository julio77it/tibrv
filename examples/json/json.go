package main

import (
	"fmt"
	"github.com/julio77it/tibrv"
)

func main() {
	var msg tibrv.RvMessage

	msg.Create()
	defer msg.Destroy()

	msg.SetString("String", "string")
	msg.SetInt8("Int8", 1)
	msg.SetUInt8("UInt8", 1)
	msg.SetInt16("Int16", 1)
	msg.SetUInt16("UInt16", 1)
	msg.SetInt32("Int32", 1)
	msg.SetUInt32("UInt32", 1)
	msg.SetInt64("Int64", 1)
	msg.SetUInt64("UInt64", 1)
	msg.SetFloat32("Float32", 1.0)
	msg.SetFloat64("Float64", 1.0)
	msg.SetStringArray("StringArray", []string{"string1", "string2"})
	msg.SetInt8Array("Int8Array", []int8{1, 2})
	msg.SetUInt8Array("UInt8Array", []uint8{1, 2})
	msg.SetInt16Array("Int16Array", []int16{1, 2})
	msg.SetUInt16Array("UInt16Array", []uint16{1, 2})
	msg.SetInt32Array("Int32Array", []int32{1, 2})
	msg.SetUInt32Array("UInt32Array", []uint32{1, 2})
	msg.SetInt64Array("Int64Array", []int64{1, 2})
	msg.SetUInt64Array("UInt64Array", []uint64{1, 2})
	msg.SetFloat32Array("Float32Array", []float32{1.2, 2.3})
	msg.SetFloat64Array("Float64Array", []float64{1.2, 2.3})

	var innerMsg tibrv.RvMessage
	innerMsg.Create()
	innerMsg.SetString("Dentro1", "textinside")
	innerMsg.SetFloat64("Dentro2", 6.02214)
	defer innerMsg.Destroy()
	msg.SetRvMessage("RvMessage", innerMsg)

	fmt.Println(msg.JSON())
}
