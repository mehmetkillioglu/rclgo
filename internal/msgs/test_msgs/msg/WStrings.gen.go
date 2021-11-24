/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	primitives "github.com/mehmetkillioglu/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/w_strings.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/WStrings", WStringsTypeSupport)
}

// Do not create instances of this type directly. Always use NewWStrings
// function instead.
type WStrings struct {
	WstringValue string `yaml:"wstring_value"`
	WstringValueDefault1 string `yaml:"wstring_value_default1"`
	WstringValueDefault2 string `yaml:"wstring_value_default2"`
	WstringValueDefault3 string `yaml:"wstring_value_default3"`
	ArrayOfWstrings [3]string `yaml:"array_of_wstrings"`// wstring WSTRING_CONST="Hello world!"wstring<=22 bounded_wstring_valuewstring<=22 bounded_wstring_value_default1 "Hello world!"
	BoundedSequenceOfWstrings []string `yaml:"bounded_sequence_of_wstrings"`// wstring WSTRING_CONST="Hello world!"wstring<=22 bounded_wstring_valuewstring<=22 bounded_wstring_value_default1 "Hello world!"
	UnboundedSequenceOfWstrings []string `yaml:"unbounded_sequence_of_wstrings"`// wstring WSTRING_CONST="Hello world!"wstring<=22 bounded_wstring_valuewstring<=22 bounded_wstring_value_default1 "Hello world!"
}

// NewWStrings creates a new WStrings with default values.
func NewWStrings() *WStrings {
	self := WStrings{}
	self.SetDefaults()
	return &self
}

func (t *WStrings) Clone() *WStrings {
	c := &WStrings{}
	c.WstringValue = t.WstringValue
	c.WstringValueDefault1 = t.WstringValueDefault1
	c.WstringValueDefault2 = t.WstringValueDefault2
	c.WstringValueDefault3 = t.WstringValueDefault3
	c.ArrayOfWstrings = t.ArrayOfWstrings
	if t.BoundedSequenceOfWstrings != nil {
		c.BoundedSequenceOfWstrings = make([]string, len(t.BoundedSequenceOfWstrings))
		copy(c.BoundedSequenceOfWstrings, t.BoundedSequenceOfWstrings)
	}
	if t.UnboundedSequenceOfWstrings != nil {
		c.UnboundedSequenceOfWstrings = make([]string, len(t.UnboundedSequenceOfWstrings))
		copy(c.UnboundedSequenceOfWstrings, t.UnboundedSequenceOfWstrings)
	}
	return c
}

func (t *WStrings) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WStrings) SetDefaults() {
	t.WstringValue = ""
	t.WstringValueDefault1 = "Hello world!"
	t.WstringValueDefault2 = "Hellö wörld!"
	t.WstringValueDefault3 = "ハローワールド"
	t.ArrayOfWstrings = [3]string{}
	t.BoundedSequenceOfWstrings = nil
	t.UnboundedSequenceOfWstrings = nil
}

// CloneWStringsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWStringsSlice(dst, src []WStrings) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WStringsTypeSupport types.MessageTypeSupport = _WStringsTypeSupport{}

type _WStringsTypeSupport struct{}

func (t _WStringsTypeSupport) New() types.Message {
	return NewWStrings()
}

func (t _WStringsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__WStrings
	return (unsafe.Pointer)(C.test_msgs__msg__WStrings__create())
}

func (t _WStringsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__WStrings__destroy((*C.test_msgs__msg__WStrings)(pointer_to_free))
}

func (t _WStringsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WStrings)
	mem := (*C.test_msgs__msg__WStrings)(dst)
	primitives.U16StringAsCStruct(unsafe.Pointer(&mem.wstring_value), m.WstringValue)
	primitives.U16StringAsCStruct(unsafe.Pointer(&mem.wstring_value_default1), m.WstringValueDefault1)
	primitives.U16StringAsCStruct(unsafe.Pointer(&mem.wstring_value_default2), m.WstringValueDefault2)
	primitives.U16StringAsCStruct(unsafe.Pointer(&mem.wstring_value_default3), m.WstringValueDefault3)
	cSlice_array_of_wstrings := mem.array_of_wstrings[:]
	primitives.U16String__Array_to_C(*(*[]primitives.CU16String)(unsafe.Pointer(&cSlice_array_of_wstrings)), m.ArrayOfWstrings[:])
	primitives.U16String__Sequence_to_C((*primitives.CU16String__Sequence)(unsafe.Pointer(&mem.bounded_sequence_of_wstrings)), m.BoundedSequenceOfWstrings)
	primitives.U16String__Sequence_to_C((*primitives.CU16String__Sequence)(unsafe.Pointer(&mem.unbounded_sequence_of_wstrings)), m.UnboundedSequenceOfWstrings)
}

func (t _WStringsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WStrings)
	mem := (*C.test_msgs__msg__WStrings)(ros2_message_buffer)
	primitives.U16StringAsGoStruct(&m.WstringValue, unsafe.Pointer(&mem.wstring_value))
	primitives.U16StringAsGoStruct(&m.WstringValueDefault1, unsafe.Pointer(&mem.wstring_value_default1))
	primitives.U16StringAsGoStruct(&m.WstringValueDefault2, unsafe.Pointer(&mem.wstring_value_default2))
	primitives.U16StringAsGoStruct(&m.WstringValueDefault3, unsafe.Pointer(&mem.wstring_value_default3))
	cSlice_array_of_wstrings := mem.array_of_wstrings[:]
	primitives.U16String__Array_to_Go(m.ArrayOfWstrings[:], *(*[]primitives.CU16String)(unsafe.Pointer(&cSlice_array_of_wstrings)))
	primitives.U16String__Sequence_to_Go(&m.BoundedSequenceOfWstrings, *(*primitives.CU16String__Sequence)(unsafe.Pointer(&mem.bounded_sequence_of_wstrings)))
	primitives.U16String__Sequence_to_Go(&m.UnboundedSequenceOfWstrings, *(*primitives.CU16String__Sequence)(unsafe.Pointer(&mem.unbounded_sequence_of_wstrings)))
}

func (t _WStringsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__WStrings())
}

type CWStrings = C.test_msgs__msg__WStrings
type CWStrings__Sequence = C.test_msgs__msg__WStrings__Sequence

func WStrings__Sequence_to_Go(goSlice *[]WStrings, cSlice CWStrings__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WStrings, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__WStrings__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__WStrings * uintptr(i)),
		))
		WStringsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func WStrings__Sequence_to_C(cSlice *CWStrings__Sequence, goSlice []WStrings) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__WStrings)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__WStrings * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__WStrings)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__WStrings * uintptr(i)),
		))
		WStringsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func WStrings__Array_to_Go(goSlice []WStrings, cSlice []CWStrings) {
	for i := 0; i < len(cSlice); i++ {
		WStringsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WStrings__Array_to_C(cSlice []CWStrings, goSlice []WStrings) {
	for i := 0; i < len(goSlice); i++ {
		WStringsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
