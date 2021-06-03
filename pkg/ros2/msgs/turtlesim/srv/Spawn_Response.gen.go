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

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/spawn.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("turtlesim/Spawn_Response", &Spawn_Response{})
}

// Do not create instances of this type directly. Always use NewSpawn_Response
// function instead.
type Spawn_Response struct {
	Name rosidl_runtime_c.String `yaml:"name"`
}

// NewSpawn_Response creates a new Spawn_Response with default values.
func NewSpawn_Response() *Spawn_Response {
	self := Spawn_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *Spawn_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Name.SetDefaults("")
	
	return t
}

func (t *Spawn_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__Spawn_Response())
}
func (t *Spawn_Response) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__Spawn_Response
	return (unsafe.Pointer)(C.turtlesim__srv__Spawn_Response__create())
}
func (t *Spawn_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__Spawn_Response__destroy((*C.turtlesim__srv__Spawn_Response)(pointer_to_free))
}
func (t *Spawn_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.turtlesim__srv__Spawn_Response)(t.PrepareMemory())
	mem.name = *(*C.rosidl_runtime_c__String)(t.Name.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Spawn_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.turtlesim__srv__Spawn_Response)(ros2_message_buffer)
	t.Name.AsGoStruct(unsafe.Pointer(&mem.name))
}
func (t *Spawn_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSpawn_Response = C.turtlesim__srv__Spawn_Response
type CSpawn_Response__Sequence = C.turtlesim__srv__Spawn_Response__Sequence

func Spawn_Response__Sequence_to_Go(goSlice *[]Spawn_Response, cSlice CSpawn_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Spawn_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__Spawn_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__Spawn_Response * uintptr(i)),
		))
		(*goSlice)[i] = Spawn_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Spawn_Response__Sequence_to_C(cSlice *CSpawn_Response__Sequence, goSlice []Spawn_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__Spawn_Response)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__Spawn_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__Spawn_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__Spawn_Response * uintptr(i)),
		))
		*cIdx = *(*C.turtlesim__srv__Spawn_Response)(v.AsCStruct())
	}
}
func Spawn_Response__Array_to_Go(goSlice []Spawn_Response, cSlice []CSpawn_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Spawn_Response__Array_to_C(cSlice []CSpawn_Response, goSlice []Spawn_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.turtlesim__srv__Spawn_Response)(goSlice[i].AsCStruct())
	}
}

