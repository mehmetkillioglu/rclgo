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

#include <std_msgs/msg/multi_array_dimension.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/MultiArrayDimension", MultiArrayDimensionTypeSupport)
}

// Do not create instances of this type directly. Always use NewMultiArrayDimension
// function instead.
type MultiArrayDimension struct {
	Label string `yaml:"label"`// label of given dimension
	Size uint32 `yaml:"size"`// size of given dimension (in type units)
	Stride uint32 `yaml:"stride"`// stride of given dimension
}

// NewMultiArrayDimension creates a new MultiArrayDimension with default values.
func NewMultiArrayDimension() *MultiArrayDimension {
	self := MultiArrayDimension{}
	self.SetDefaults()
	return &self
}

func (t *MultiArrayDimension) Clone() *MultiArrayDimension {
	c := &MultiArrayDimension{}
	c.Label = t.Label
	c.Size = t.Size
	c.Stride = t.Stride
	return c
}

func (t *MultiArrayDimension) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultiArrayDimension) SetDefaults() {
	t.Label = ""
	t.Size = 0
	t.Stride = 0
}

// CloneMultiArrayDimensionSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultiArrayDimensionSlice(dst, src []MultiArrayDimension) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultiArrayDimensionTypeSupport types.MessageTypeSupport = _MultiArrayDimensionTypeSupport{}

type _MultiArrayDimensionTypeSupport struct{}

func (t _MultiArrayDimensionTypeSupport) New() types.Message {
	return NewMultiArrayDimension()
}

func (t _MultiArrayDimensionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__MultiArrayDimension
	return (unsafe.Pointer)(C.std_msgs__msg__MultiArrayDimension__create())
}

func (t _MultiArrayDimensionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__MultiArrayDimension__destroy((*C.std_msgs__msg__MultiArrayDimension)(pointer_to_free))
}

func (t _MultiArrayDimensionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiArrayDimension)
	mem := (*C.std_msgs__msg__MultiArrayDimension)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.label), m.Label)
	mem.size = C.uint32_t(m.Size)
	mem.stride = C.uint32_t(m.Stride)
}

func (t _MultiArrayDimensionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiArrayDimension)
	mem := (*C.std_msgs__msg__MultiArrayDimension)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Label, unsafe.Pointer(&mem.label))
	m.Size = uint32(mem.size)
	m.Stride = uint32(mem.stride)
}

func (t _MultiArrayDimensionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__MultiArrayDimension())
}

type CMultiArrayDimension = C.std_msgs__msg__MultiArrayDimension
type CMultiArrayDimension__Sequence = C.std_msgs__msg__MultiArrayDimension__Sequence

func MultiArrayDimension__Sequence_to_Go(goSlice *[]MultiArrayDimension, cSlice CMultiArrayDimension__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiArrayDimension, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__MultiArrayDimension__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__MultiArrayDimension * uintptr(i)),
		))
		MultiArrayDimensionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MultiArrayDimension__Sequence_to_C(cSlice *CMultiArrayDimension__Sequence, goSlice []MultiArrayDimension) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__MultiArrayDimension)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__MultiArrayDimension * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__MultiArrayDimension)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__MultiArrayDimension * uintptr(i)),
		))
		MultiArrayDimensionTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MultiArrayDimension__Array_to_Go(goSlice []MultiArrayDimension, cSlice []CMultiArrayDimension) {
	for i := 0; i < len(cSlice); i++ {
		MultiArrayDimensionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiArrayDimension__Array_to_C(cSlice []CMultiArrayDimension, goSlice []MultiArrayDimension) {
	for i := 0; i < len(goSlice); i++ {
		MultiArrayDimensionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
