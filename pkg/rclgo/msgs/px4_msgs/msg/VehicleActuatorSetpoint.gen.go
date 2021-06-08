/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_actuator_setpoint.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleActuatorSetpoint", VehicleActuatorSetpointTypeSupport)
}
const (
	VehicleActuatorSetpoint_NUM_ACTUATOR_SETPOINT uint8 = 16
)

// Do not create instances of this type directly. Always use NewVehicleActuatorSetpoint
// function instead.
type VehicleActuatorSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp the data this control response is based on was sampled
	Actuator [16]float32 `yaml:"actuator"`
}

// NewVehicleActuatorSetpoint creates a new VehicleActuatorSetpoint with default values.
func NewVehicleActuatorSetpoint() *VehicleActuatorSetpoint {
	self := VehicleActuatorSetpoint{}
	self.SetDefaults()
	return &self
}

func (t *VehicleActuatorSetpoint) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleActuatorSetpoint) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleActuatorSetpointTypeSupport types.MessageTypeSupport = _VehicleActuatorSetpointTypeSupport{}

type _VehicleActuatorSetpointTypeSupport struct{}

func (t _VehicleActuatorSetpointTypeSupport) New() types.Message {
	return NewVehicleActuatorSetpoint()
}

func (t _VehicleActuatorSetpointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleActuatorSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleActuatorSetpoint__create())
}

func (t _VehicleActuatorSetpointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleActuatorSetpoint__destroy((*C.px4_msgs__msg__VehicleActuatorSetpoint)(pointer_to_free))
}

func (t _VehicleActuatorSetpointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleActuatorSetpoint)
	mem := (*C.px4_msgs__msg__VehicleActuatorSetpoint)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_actuator := mem.actuator[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_actuator)), m.Actuator[:])
}

func (t _VehicleActuatorSetpointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleActuatorSetpoint)
	mem := (*C.px4_msgs__msg__VehicleActuatorSetpoint)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_actuator := mem.actuator[:]
	primitives.Float32__Array_to_Go(m.Actuator[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_actuator)))
}

func (t _VehicleActuatorSetpointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleActuatorSetpoint())
}

type CVehicleActuatorSetpoint = C.px4_msgs__msg__VehicleActuatorSetpoint
type CVehicleActuatorSetpoint__Sequence = C.px4_msgs__msg__VehicleActuatorSetpoint__Sequence

func VehicleActuatorSetpoint__Sequence_to_Go(goSlice *[]VehicleActuatorSetpoint, cSlice CVehicleActuatorSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleActuatorSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleActuatorSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleActuatorSetpoint * uintptr(i)),
		))
		VehicleActuatorSetpointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleActuatorSetpoint__Sequence_to_C(cSlice *CVehicleActuatorSetpoint__Sequence, goSlice []VehicleActuatorSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleActuatorSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleActuatorSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleActuatorSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleActuatorSetpoint * uintptr(i)),
		))
		VehicleActuatorSetpointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleActuatorSetpoint__Array_to_Go(goSlice []VehicleActuatorSetpoint, cSlice []CVehicleActuatorSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		VehicleActuatorSetpointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleActuatorSetpoint__Array_to_C(cSlice []CVehicleActuatorSetpoint, goSlice []VehicleActuatorSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		VehicleActuatorSetpointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}