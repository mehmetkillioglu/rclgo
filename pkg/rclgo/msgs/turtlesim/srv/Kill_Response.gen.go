/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package turtlesim_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/kill.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/Kill_Response", Kill_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewKill_Response
// function instead.
type Kill_Response struct {
}

// NewKill_Response creates a new Kill_Response with default values.
func NewKill_Response() *Kill_Response {
	self := Kill_Response{}
	self.SetDefaults()
	return &self
}

func (t *Kill_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Kill_Response) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var Kill_ResponseTypeSupport types.MessageTypeSupport = _Kill_ResponseTypeSupport{}

type _Kill_ResponseTypeSupport struct{}

func (t _Kill_ResponseTypeSupport) New() types.Message {
	return NewKill_Response()
}

func (t _Kill_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__Kill_Response
	return (unsafe.Pointer)(C.turtlesim__srv__Kill_Response__create())
}

func (t _Kill_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__Kill_Response__destroy((*C.turtlesim__srv__Kill_Response)(pointer_to_free))
}

func (t _Kill_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _Kill_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _Kill_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__Kill_Response())
}

type CKill_Response = C.turtlesim__srv__Kill_Response
type CKill_Response__Sequence = C.turtlesim__srv__Kill_Response__Sequence

func Kill_Response__Sequence_to_Go(goSlice *[]Kill_Response, cSlice CKill_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Kill_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__Kill_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__Kill_Response * uintptr(i)),
		))
		Kill_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Kill_Response__Sequence_to_C(cSlice *CKill_Response__Sequence, goSlice []Kill_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__Kill_Response)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__Kill_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__Kill_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__Kill_Response * uintptr(i)),
		))
		Kill_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Kill_Response__Array_to_Go(goSlice []Kill_Response, cSlice []CKill_Response) {
	for i := 0; i < len(cSlice); i++ {
		Kill_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Kill_Response__Array_to_C(cSlice []CKill_Response, goSlice []Kill_Response) {
	for i := 0; i < len(goSlice); i++ {
		Kill_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}