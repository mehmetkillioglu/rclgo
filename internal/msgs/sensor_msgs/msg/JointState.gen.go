/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/mehmetkillioglu/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/mehmetkillioglu/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/joint_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/JointState", JointStateTypeSupport)
}

// Do not create instances of this type directly. Always use NewJointState
// function instead.
type JointState struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Name []string `yaml:"name"`
	Position []float64 `yaml:"position"`
	Velocity []float64 `yaml:"velocity"`
	Effort []float64 `yaml:"effort"`
}

// NewJointState creates a new JointState with default values.
func NewJointState() *JointState {
	self := JointState{}
	self.SetDefaults()
	return &self
}

func (t *JointState) Clone() *JointState {
	c := &JointState{}
	c.Header = *t.Header.Clone()
	if t.Name != nil {
		c.Name = make([]string, len(t.Name))
		copy(c.Name, t.Name)
	}
	if t.Position != nil {
		c.Position = make([]float64, len(t.Position))
		copy(c.Position, t.Position)
	}
	if t.Velocity != nil {
		c.Velocity = make([]float64, len(t.Velocity))
		copy(c.Velocity, t.Velocity)
	}
	if t.Effort != nil {
		c.Effort = make([]float64, len(t.Effort))
		copy(c.Effort, t.Effort)
	}
	return c
}

func (t *JointState) CloneMsg() types.Message {
	return t.Clone()
}

func (t *JointState) SetDefaults() {
	t.Header.SetDefaults()
	t.Name = nil
	t.Position = nil
	t.Velocity = nil
	t.Effort = nil
}

// CloneJointStateSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJointStateSlice(dst, src []JointState) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JointStateTypeSupport types.MessageTypeSupport = _JointStateTypeSupport{}

type _JointStateTypeSupport struct{}

func (t _JointStateTypeSupport) New() types.Message {
	return NewJointState()
}

func (t _JointStateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__JointState
	return (unsafe.Pointer)(C.sensor_msgs__msg__JointState__create())
}

func (t _JointStateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__JointState__destroy((*C.sensor_msgs__msg__JointState)(pointer_to_free))
}

func (t _JointStateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*JointState)
	mem := (*C.sensor_msgs__msg__JointState)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.name)), m.Name)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.position)), m.Position)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.velocity)), m.Velocity)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.effort)), m.Effort)
}

func (t _JointStateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*JointState)
	mem := (*C.sensor_msgs__msg__JointState)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.String__Sequence_to_Go(&m.Name, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.name)))
	primitives.Float64__Sequence_to_Go(&m.Position, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.position)))
	primitives.Float64__Sequence_to_Go(&m.Velocity, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.velocity)))
	primitives.Float64__Sequence_to_Go(&m.Effort, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.effort)))
}

func (t _JointStateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__JointState())
}

type CJointState = C.sensor_msgs__msg__JointState
type CJointState__Sequence = C.sensor_msgs__msg__JointState__Sequence

func JointState__Sequence_to_Go(goSlice *[]JointState, cSlice CJointState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JointState, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__JointState__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JointState * uintptr(i)),
		))
		JointStateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func JointState__Sequence_to_C(cSlice *CJointState__Sequence, goSlice []JointState) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__JointState)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__JointState * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__JointState)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JointState * uintptr(i)),
		))
		JointStateTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func JointState__Array_to_Go(goSlice []JointState, cSlice []CJointState) {
	for i := 0; i < len(cSlice); i++ {
		JointStateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func JointState__Array_to_C(cSlice []CJointState, goSlice []JointState) {
	for i := 0; i < len(goSlice); i++ {
		JointStateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
