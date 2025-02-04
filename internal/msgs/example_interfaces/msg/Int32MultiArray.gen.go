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

#include <example_interfaces/msg/int32_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Int32MultiArray", Int32MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewInt32MultiArray
// function instead.
type Int32MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []int32 `yaml:"data"`// array of data
}

// NewInt32MultiArray creates a new Int32MultiArray with default values.
func NewInt32MultiArray() *Int32MultiArray {
	self := Int32MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Int32MultiArray) Clone() *Int32MultiArray {
	c := &Int32MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]int32, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Int32MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Int32MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// CloneInt32MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInt32MultiArraySlice(dst, src []Int32MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Int32MultiArrayTypeSupport types.MessageTypeSupport = _Int32MultiArrayTypeSupport{}

type _Int32MultiArrayTypeSupport struct{}

func (t _Int32MultiArrayTypeSupport) New() types.Message {
	return NewInt32MultiArray()
}

func (t _Int32MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Int32MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__Int32MultiArray__create())
}

func (t _Int32MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Int32MultiArray__destroy((*C.example_interfaces__msg__Int32MultiArray)(pointer_to_free))
}

func (t _Int32MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int32MultiArray)
	mem := (*C.example_interfaces__msg__Int32MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Int32MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int32MultiArray)
	mem := (*C.example_interfaces__msg__Int32MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Int32__Sequence_to_Go(&m.Data, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Int32MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Int32MultiArray())
}

type CInt32MultiArray = C.example_interfaces__msg__Int32MultiArray
type CInt32MultiArray__Sequence = C.example_interfaces__msg__Int32MultiArray__Sequence

func Int32MultiArray__Sequence_to_Go(goSlice *[]Int32MultiArray, cSlice CInt32MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int32MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Int32MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int32MultiArray * uintptr(i)),
		))
		Int32MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int32MultiArray__Sequence_to_C(cSlice *CInt32MultiArray__Sequence, goSlice []Int32MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Int32MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Int32MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Int32MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int32MultiArray * uintptr(i)),
		))
		Int32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int32MultiArray__Array_to_Go(goSlice []Int32MultiArray, cSlice []CInt32MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Int32MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int32MultiArray__Array_to_C(cSlice []CInt32MultiArray, goSlice []Int32MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Int32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
