/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	geometry_msgs_msg "github.com/tiiuae/rclgo/pkg/ros2/msgs/geometry_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/srv/get_plan.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("nav_msgs/GetPlan_Request", &GetPlan_Request{})
}

// Do not create instances of this type directly. Always use NewGetPlan_Request
// function instead.
type GetPlan_Request struct {
	Start geometry_msgs_msg.PoseStamped `yaml:"start"`// The start pose for the plan
	Goal geometry_msgs_msg.PoseStamped `yaml:"goal"`// The final pose of the goal position
	Tolerance float32 `yaml:"tolerance"`// If the goal is obstructed, how many meters the planner canrelax the constraint in x and y before failing.
}

// NewGetPlan_Request creates a new GetPlan_Request with default values.
func NewGetPlan_Request() *GetPlan_Request {
	self := GetPlan_Request{}
	self.SetDefaults(nil)
	return &self
}

func (t *GetPlan_Request) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Start.SetDefaults(nil)
	t.Goal.SetDefaults(nil)
	
	return t
}

func (t *GetPlan_Request) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__srv__GetPlan_Request())
}
func (t *GetPlan_Request) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__srv__GetPlan_Request
	return (unsafe.Pointer)(C.nav_msgs__srv__GetPlan_Request__create())
}
func (t *GetPlan_Request) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__srv__GetPlan_Request__destroy((*C.nav_msgs__srv__GetPlan_Request)(pointer_to_free))
}
func (t *GetPlan_Request) AsCStruct() unsafe.Pointer {
	mem := (*C.nav_msgs__srv__GetPlan_Request)(t.PrepareMemory())
	mem.start = *(*C.geometry_msgs__msg__PoseStamped)(t.Start.AsCStruct())
	mem.goal = *(*C.geometry_msgs__msg__PoseStamped)(t.Goal.AsCStruct())
	mem.tolerance = C.float(t.Tolerance)
	return unsafe.Pointer(mem)
}
func (t *GetPlan_Request) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.nav_msgs__srv__GetPlan_Request)(ros2_message_buffer)
	t.Start.AsGoStruct(unsafe.Pointer(&mem.start))
	t.Goal.AsGoStruct(unsafe.Pointer(&mem.goal))
	t.Tolerance = float32(mem.tolerance)
}
func (t *GetPlan_Request) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CGetPlan_Request = C.nav_msgs__srv__GetPlan_Request
type CGetPlan_Request__Sequence = C.nav_msgs__srv__GetPlan_Request__Sequence

func GetPlan_Request__Sequence_to_Go(goSlice *[]GetPlan_Request, cSlice CGetPlan_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetPlan_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__srv__GetPlan_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetPlan_Request * uintptr(i)),
		))
		(*goSlice)[i] = GetPlan_Request{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func GetPlan_Request__Sequence_to_C(cSlice *CGetPlan_Request__Sequence, goSlice []GetPlan_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__srv__GetPlan_Request)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__srv__GetPlan_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__srv__GetPlan_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetPlan_Request * uintptr(i)),
		))
		*cIdx = *(*C.nav_msgs__srv__GetPlan_Request)(v.AsCStruct())
	}
}
func GetPlan_Request__Array_to_Go(goSlice []GetPlan_Request, cSlice []CGetPlan_Request) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func GetPlan_Request__Array_to_C(cSlice []CGetPlan_Request, goSlice []GetPlan_Request) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.nav_msgs__srv__GetPlan_Request)(goSlice[i].AsCStruct())
	}
}

