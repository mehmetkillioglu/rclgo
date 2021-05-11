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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/estimator_states.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/EstimatorStates", &EstimatorStates{})
}

// Do not create instances of this type directly. Always use NewEstimatorStates
// function instead.
type EstimatorStates struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	States [24]float32 `yaml:"states"`// Internal filter states
	NStates uint8 `yaml:"n_states"`// Number of states effectively used
	Covariances [24]float32 `yaml:"covariances"`// Diagonal Elements of Covariance Matrix
}

// NewEstimatorStates creates a new EstimatorStates with default values.
func NewEstimatorStates() *EstimatorStates {
	self := EstimatorStates{}
	self.SetDefaults(nil)
	return &self
}

func (t *EstimatorStates) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *EstimatorStates) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorStates())
}
func (t *EstimatorStates) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorStates
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorStates__create())
}
func (t *EstimatorStates) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorStates__destroy((*C.px4_msgs__msg__EstimatorStates)(pointer_to_free))
}
func (t *EstimatorStates) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__EstimatorStates)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	cSlice_states := mem.states[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_states)), t.States[:])
	mem.n_states = C.uint8_t(t.NStates)
	cSlice_covariances := mem.covariances[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_covariances)), t.Covariances[:])
	return unsafe.Pointer(mem)
}
func (t *EstimatorStates) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__EstimatorStates)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_states := mem.states[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.States[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_states)))
	t.NStates = uint8(mem.n_states)
	cSlice_covariances := mem.covariances[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Covariances[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_covariances)))
}
func (t *EstimatorStates) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CEstimatorStates = C.px4_msgs__msg__EstimatorStates
type CEstimatorStates__Sequence = C.px4_msgs__msg__EstimatorStates__Sequence

func EstimatorStates__Sequence_to_Go(goSlice *[]EstimatorStates, cSlice CEstimatorStates__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorStates, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorStates__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorStates * uintptr(i)),
		))
		(*goSlice)[i] = EstimatorStates{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func EstimatorStates__Sequence_to_C(cSlice *CEstimatorStates__Sequence, goSlice []EstimatorStates) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorStates)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorStates * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorStates)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorStates * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__EstimatorStates)(v.AsCStruct())
	}
}
func EstimatorStates__Array_to_Go(goSlice []EstimatorStates, cSlice []CEstimatorStates) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorStates__Array_to_C(cSlice []CEstimatorStates, goSlice []EstimatorStates) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__EstimatorStates)(goSlice[i].AsCStruct())
	}
}

