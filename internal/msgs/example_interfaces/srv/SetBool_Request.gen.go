/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/set_bool.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/SetBool_Request", SetBool_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetBool_Request
// function instead.
type SetBool_Request struct {
	Data bool `yaml:"data"`// e.g. for hardware enabling / disabling
}

// NewSetBool_Request creates a new SetBool_Request with default values.
func NewSetBool_Request() *SetBool_Request {
	self := SetBool_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetBool_Request) Clone() *SetBool_Request {
	c := &SetBool_Request{}
	c.Data = t.Data
	return c
}

func (t *SetBool_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetBool_Request) SetDefaults() {
	t.Data = false
}

// CloneSetBool_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetBool_RequestSlice(dst, src []SetBool_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetBool_RequestTypeSupport types.MessageTypeSupport = _SetBool_RequestTypeSupport{}

type _SetBool_RequestTypeSupport struct{}

func (t _SetBool_RequestTypeSupport) New() types.Message {
	return NewSetBool_Request()
}

func (t _SetBool_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__SetBool_Request
	return (unsafe.Pointer)(C.example_interfaces__srv__SetBool_Request__create())
}

func (t _SetBool_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__SetBool_Request__destroy((*C.example_interfaces__srv__SetBool_Request)(pointer_to_free))
}

func (t _SetBool_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetBool_Request)
	mem := (*C.example_interfaces__srv__SetBool_Request)(dst)
	mem.data = C.bool(m.Data)
}

func (t _SetBool_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetBool_Request)
	mem := (*C.example_interfaces__srv__SetBool_Request)(ros2_message_buffer)
	m.Data = bool(mem.data)
}

func (t _SetBool_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__SetBool_Request())
}

type CSetBool_Request = C.example_interfaces__srv__SetBool_Request
type CSetBool_Request__Sequence = C.example_interfaces__srv__SetBool_Request__Sequence

func SetBool_Request__Sequence_to_Go(goSlice *[]SetBool_Request, cSlice CSetBool_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetBool_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__srv__SetBool_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__SetBool_Request * uintptr(i)),
		))
		SetBool_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetBool_Request__Sequence_to_C(cSlice *CSetBool_Request__Sequence, goSlice []SetBool_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__srv__SetBool_Request)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__srv__SetBool_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__srv__SetBool_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__SetBool_Request * uintptr(i)),
		))
		SetBool_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetBool_Request__Array_to_Go(goSlice []SetBool_Request, cSlice []CSetBool_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetBool_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetBool_Request__Array_to_C(cSlice []CSetBool_Request, goSlice []SetBool_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetBool_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
