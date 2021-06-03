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

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/set_bool.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/SetBool_Request", &SetBool_Request{})
}

// Do not create instances of this type directly. Always use NewSetBool_Request
// function instead.
type SetBool_Request struct {
	Data bool `yaml:"data"`// e.g. for hardware enabling / disabling
}

// NewSetBool_Request creates a new SetBool_Request with default values.
func NewSetBool_Request() *SetBool_Request {
	self := SetBool_Request{}
	self.SetDefaults(nil)
	return &self
}

func (t *SetBool_Request) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *SetBool_Request) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__SetBool_Request())
}
func (t *SetBool_Request) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__SetBool_Request
	return (unsafe.Pointer)(C.example_interfaces__srv__SetBool_Request__create())
}
func (t *SetBool_Request) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__SetBool_Request__destroy((*C.example_interfaces__srv__SetBool_Request)(pointer_to_free))
}
func (t *SetBool_Request) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__srv__SetBool_Request)(t.PrepareMemory())
	mem.data = C.bool(t.Data)
	return unsafe.Pointer(mem)
}
func (t *SetBool_Request) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__srv__SetBool_Request)(ros2_message_buffer)
	t.Data = bool(mem.data)
}
func (t *SetBool_Request) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
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
		(*goSlice)[i] = SetBool_Request{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
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
		*cIdx = *(*C.example_interfaces__srv__SetBool_Request)(v.AsCStruct())
	}
}
func SetBool_Request__Array_to_Go(goSlice []SetBool_Request, cSlice []CSetBool_Request) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func SetBool_Request__Array_to_C(cSlice []CSetBool_Request, goSlice []SetBool_Request) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__srv__SetBool_Request)(goSlice[i].AsCStruct())
	}
}

