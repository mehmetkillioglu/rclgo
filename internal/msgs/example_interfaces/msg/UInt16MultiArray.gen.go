/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	primitives "github.com/mehmetkillioglu/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/u_int16_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/UInt16MultiArray", UInt16MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt16MultiArray
// function instead.
type UInt16MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint16 `yaml:"data"`// array of data
}

// NewUInt16MultiArray creates a new UInt16MultiArray with default values.
func NewUInt16MultiArray() *UInt16MultiArray {
	self := UInt16MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *UInt16MultiArray) Clone() *UInt16MultiArray {
	c := &UInt16MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]uint16, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *UInt16MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt16MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// CloneUInt16MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt16MultiArraySlice(dst, src []UInt16MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt16MultiArrayTypeSupport types.MessageTypeSupport = _UInt16MultiArrayTypeSupport{}

type _UInt16MultiArrayTypeSupport struct{}

func (t _UInt16MultiArrayTypeSupport) New() types.Message {
	return NewUInt16MultiArray()
}

func (t _UInt16MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__UInt16MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__UInt16MultiArray__create())
}

func (t _UInt16MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__UInt16MultiArray__destroy((*C.example_interfaces__msg__UInt16MultiArray)(pointer_to_free))
}

func (t _UInt16MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt16MultiArray)
	mem := (*C.example_interfaces__msg__UInt16MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Uint16__Sequence_to_C((*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _UInt16MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt16MultiArray)
	mem := (*C.example_interfaces__msg__UInt16MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Uint16__Sequence_to_Go(&m.Data, *(*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _UInt16MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__UInt16MultiArray())
}

type CUInt16MultiArray = C.example_interfaces__msg__UInt16MultiArray
type CUInt16MultiArray__Sequence = C.example_interfaces__msg__UInt16MultiArray__Sequence

func UInt16MultiArray__Sequence_to_Go(goSlice *[]UInt16MultiArray, cSlice CUInt16MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt16MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__UInt16MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt16MultiArray * uintptr(i)),
		))
		UInt16MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt16MultiArray__Sequence_to_C(cSlice *CUInt16MultiArray__Sequence, goSlice []UInt16MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__UInt16MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__UInt16MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__UInt16MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt16MultiArray * uintptr(i)),
		))
		UInt16MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt16MultiArray__Array_to_Go(goSlice []UInt16MultiArray, cSlice []CUInt16MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		UInt16MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt16MultiArray__Array_to_C(cSlice []CUInt16MultiArray, goSlice []UInt16MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		UInt16MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
