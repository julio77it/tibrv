package main

import (
	"fmt"

	"github.com/julio77it/tibrv"
)

func main() {
	var msg tibrv.RvMessage

	if err := msg.Create(); err != nil {
		panic(err)
	}
	defer msg.Destroy()

	if err := msg.SetString("String", "string"); err != nil {
		panic(err)
	}
	if err := msg.SetInt8("Int8", 1); err != nil {
		panic(err)
	}
	if err := msg.SetUInt8("UInt8", 1); err != nil {
		panic(err)
	}
	if err := msg.SetInt16("Int16", 1); err != nil {
		panic(err)
	}
	if err := msg.SetUInt16("UInt16", 1); err != nil {
		panic(err)
	}
	if err := msg.SetInt32("Int32", 1); err != nil {
		panic(err)
	}
	if err := msg.SetUInt32("UInt32", 1); err != nil {
		panic(err)
	}
	if err := msg.SetInt64("Int64", 1); err != nil {
		panic(err)
	}
	if err := msg.SetUInt64("UInt64", 1); err != nil {
		panic(err)
	}
	if err := msg.SetFloat32("Float32", 1.0); err != nil {
		panic(err)
	}
	if err := msg.SetFloat64("Float64", 1.0); err != nil {
		panic(err)
	}
	if err := msg.SetStringArray("StringArray", []string{"string1", "string2"}); err != nil {
		panic(err)
	}
	if err := msg.SetInt8Array("Int8Array", []int8{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetUInt8Array("UInt8Array", []uint8{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetInt16Array("Int16Array", []int16{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetUInt16Array("UInt16Array", []uint16{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetInt32Array("Int32Array", []int32{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetUInt32Array("UInt32Array", []uint32{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetInt64Array("Int64Array", []int64{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetUInt64Array("UInt64Array", []uint64{1, 2}); err != nil {
		panic(err)
	}
	if err := msg.SetFloat32Array("Float32Array", []float32{1.2, 2.3}); err != nil {
		panic(err)
	}
	if err := msg.SetFloat64Array("Float64Array", []float64{1.2, 2.3}); err != nil {
		panic(err)
	}

	var innerMsg tibrv.RvMessage
	if err := innerMsg.Create(); err != nil {
		panic(err)
	}
	if err := innerMsg.SetString("Dentro1", "textinside"); err != nil {
		panic(err)
	}
	if err := innerMsg.SetFloat64("Dentro2", 6.02214); err != nil {
		panic(err)
	}
	defer innerMsg.Destroy()
	if err := msg.SetRvMessage("RvMessage", innerMsg); err != nil {
		panic(err)
	}

	output, err := msg.JSON()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)

	var rmsg *tibrv.RvMessage

	rmsg, err = tibrv.JSON(output)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*rmsg)
}
