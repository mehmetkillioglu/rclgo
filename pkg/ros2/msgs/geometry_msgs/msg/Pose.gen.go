/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geometry_msgs/msg/pose.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Pose", &Pose{})
}

// Do not create instances of this type directly. Always use NewPose
// function instead.
type Pose struct {
	Position Point `yaml:"position"`
	Orientation Quaternion `yaml:"orientation"`
}

// NewPose creates a new Pose with default values.
func NewPose() *Pose {
	self := Pose{}
	self.SetDefaults(nil)
	return &self
}

func (t *Pose) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Position.SetDefaults(nil)
	t.Orientation.SetDefaults(nil)
	
	return t
}

func (t *Pose) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Pose())
}
func (t *Pose) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Pose
	return (unsafe.Pointer)(C.geometry_msgs__msg__Pose__create())
}
func (t *Pose) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Pose__destroy((*C.geometry_msgs__msg__Pose)(pointer_to_free))
}
func (t *Pose) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__Pose)(t.PrepareMemory())
	mem.position = *(*C.geometry_msgs__msg__Point)(t.Position.AsCStruct())
	mem.orientation = *(*C.geometry_msgs__msg__Quaternion)(t.Orientation.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Pose) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__Pose)(ros2_message_buffer)
	t.Position.AsGoStruct(unsafe.Pointer(&mem.position))
	t.Orientation.AsGoStruct(unsafe.Pointer(&mem.orientation))
}
func (t *Pose) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPose = C.geometry_msgs__msg__Pose
type CPose__Sequence = C.geometry_msgs__msg__Pose__Sequence

func Pose__Sequence_to_Go(goSlice *[]Pose, cSlice CPose__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Pose, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Pose__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Pose * uintptr(i)),
		))
		(*goSlice)[i] = Pose{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Pose__Sequence_to_C(cSlice *CPose__Sequence, goSlice []Pose) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Pose)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Pose * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Pose)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Pose * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__Pose)(v.AsCStruct())
	}
}
func Pose__Array_to_Go(goSlice []Pose, cSlice []CPose) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Pose__Array_to_C(cSlice []CPose, goSlice []Pose) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__Pose)(goSlice[i].AsCStruct())
	}
}

