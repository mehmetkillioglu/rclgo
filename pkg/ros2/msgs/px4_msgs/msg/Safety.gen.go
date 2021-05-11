/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/safety.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/Safety", &Safety{})
}

// Do not create instances of this type directly. Always use NewSafety
// function instead.
type Safety struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	SafetySwitchAvailable bool `yaml:"safety_switch_available"`// Set to true if a safety switch is connected
	SafetyOff bool `yaml:"safety_off"`// Set to true if safety is off
	OverrideAvailable bool `yaml:"override_available"`// Set to true if external override system is connected
	OverrideEnabled bool `yaml:"override_enabled"`// Set to true if override is engaged
}

// NewSafety creates a new Safety with default values.
func NewSafety() *Safety {
	self := Safety{}
	self.SetDefaults(nil)
	return &self
}

func (t *Safety) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Safety) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Safety())
}
func (t *Safety) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Safety
	return (unsafe.Pointer)(C.px4_msgs__msg__Safety__create())
}
func (t *Safety) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Safety__destroy((*C.px4_msgs__msg__Safety)(pointer_to_free))
}
func (t *Safety) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__Safety)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.safety_switch_available = C.bool(t.SafetySwitchAvailable)
	mem.safety_off = C.bool(t.SafetyOff)
	mem.override_available = C.bool(t.OverrideAvailable)
	mem.override_enabled = C.bool(t.OverrideEnabled)
	return unsafe.Pointer(mem)
}
func (t *Safety) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__Safety)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.SafetySwitchAvailable = bool(mem.safety_switch_available)
	t.SafetyOff = bool(mem.safety_off)
	t.OverrideAvailable = bool(mem.override_available)
	t.OverrideEnabled = bool(mem.override_enabled)
}
func (t *Safety) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSafety = C.px4_msgs__msg__Safety
type CSafety__Sequence = C.px4_msgs__msg__Safety__Sequence

func Safety__Sequence_to_Go(goSlice *[]Safety, cSlice CSafety__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Safety, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Safety__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Safety * uintptr(i)),
		))
		(*goSlice)[i] = Safety{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Safety__Sequence_to_C(cSlice *CSafety__Sequence, goSlice []Safety) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Safety)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Safety * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Safety)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Safety * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__Safety)(v.AsCStruct())
	}
}
func Safety__Array_to_Go(goSlice []Safety, cSlice []CSafety) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Safety__Array_to_C(cSlice []CSafety, goSlice []Safety) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__Safety)(goSlice[i].AsCStruct())
	}
}

