/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs_msg
import (
	"unsafe"

	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	primitives "github.com/mehmetkillioglu/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int64_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/UInt64MultiArray", UInt64MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt64MultiArray
// function instead.
type UInt64MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint64 `yaml:"data"`// array of data
}

// NewUInt64MultiArray creates a new UInt64MultiArray with default values.
func NewUInt64MultiArray() *UInt64MultiArray {
	self := UInt64MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *UInt64MultiArray) Clone() *UInt64MultiArray {
	c := &UInt64MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]uint64, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *UInt64MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt64MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// CloneUInt64MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt64MultiArraySlice(dst, src []UInt64MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt64MultiArrayTypeSupport types.MessageTypeSupport = _UInt64MultiArrayTypeSupport{}

type _UInt64MultiArrayTypeSupport struct{}

func (t _UInt64MultiArrayTypeSupport) New() types.Message {
	return NewUInt64MultiArray()
}

func (t _UInt64MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt64MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__UInt64MultiArray__create())
}

func (t _UInt64MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt64MultiArray__destroy((*C.std_msgs__msg__UInt64MultiArray)(pointer_to_free))
}

func (t _UInt64MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt64MultiArray)
	mem := (*C.std_msgs__msg__UInt64MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Uint64__Sequence_to_C((*primitives.CUint64__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _UInt64MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt64MultiArray)
	mem := (*C.std_msgs__msg__UInt64MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Uint64__Sequence_to_Go(&m.Data, *(*primitives.CUint64__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _UInt64MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt64MultiArray())
}

type CUInt64MultiArray = C.std_msgs__msg__UInt64MultiArray
type CUInt64MultiArray__Sequence = C.std_msgs__msg__UInt64MultiArray__Sequence

func UInt64MultiArray__Sequence_to_Go(goSlice *[]UInt64MultiArray, cSlice CUInt64MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt64MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt64MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt64MultiArray * uintptr(i)),
		))
		UInt64MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt64MultiArray__Sequence_to_C(cSlice *CUInt64MultiArray__Sequence, goSlice []UInt64MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt64MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt64MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt64MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt64MultiArray * uintptr(i)),
		))
		UInt64MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt64MultiArray__Array_to_Go(goSlice []UInt64MultiArray, cSlice []CUInt64MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		UInt64MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt64MultiArray__Array_to_C(cSlice []CUInt64MultiArray, goSlice []UInt64MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		UInt64MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
