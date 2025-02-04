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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/illuminance.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Illuminance", IlluminanceTypeSupport)
}

// Do not create instances of this type directly. Always use NewIlluminance
// function instead.
type Illuminance struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp is the time the illuminance was measured
	Illuminance float64 `yaml:"illuminance"`// Measurement of the Photometric Illuminance in Lux.
	Variance float64 `yaml:"variance"`// 0 is interpreted as variance unknown
}

// NewIlluminance creates a new Illuminance with default values.
func NewIlluminance() *Illuminance {
	self := Illuminance{}
	self.SetDefaults()
	return &self
}

func (t *Illuminance) Clone() *Illuminance {
	c := &Illuminance{}
	c.Header = *t.Header.Clone()
	c.Illuminance = t.Illuminance
	c.Variance = t.Variance
	return c
}

func (t *Illuminance) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Illuminance) SetDefaults() {
	t.Header.SetDefaults()
	t.Illuminance = 0
	t.Variance = 0
}

// CloneIlluminanceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneIlluminanceSlice(dst, src []Illuminance) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var IlluminanceTypeSupport types.MessageTypeSupport = _IlluminanceTypeSupport{}

type _IlluminanceTypeSupport struct{}

func (t _IlluminanceTypeSupport) New() types.Message {
	return NewIlluminance()
}

func (t _IlluminanceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Illuminance
	return (unsafe.Pointer)(C.sensor_msgs__msg__Illuminance__create())
}

func (t _IlluminanceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Illuminance__destroy((*C.sensor_msgs__msg__Illuminance)(pointer_to_free))
}

func (t _IlluminanceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Illuminance)
	mem := (*C.sensor_msgs__msg__Illuminance)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.illuminance = C.double(m.Illuminance)
	mem.variance = C.double(m.Variance)
}

func (t _IlluminanceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Illuminance)
	mem := (*C.sensor_msgs__msg__Illuminance)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Illuminance = float64(mem.illuminance)
	m.Variance = float64(mem.variance)
}

func (t _IlluminanceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Illuminance())
}

type CIlluminance = C.sensor_msgs__msg__Illuminance
type CIlluminance__Sequence = C.sensor_msgs__msg__Illuminance__Sequence

func Illuminance__Sequence_to_Go(goSlice *[]Illuminance, cSlice CIlluminance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Illuminance, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__Illuminance__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Illuminance * uintptr(i)),
		))
		IlluminanceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Illuminance__Sequence_to_C(cSlice *CIlluminance__Sequence, goSlice []Illuminance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Illuminance)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__Illuminance * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__Illuminance)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Illuminance * uintptr(i)),
		))
		IlluminanceTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Illuminance__Array_to_Go(goSlice []Illuminance, cSlice []CIlluminance) {
	for i := 0; i < len(cSlice); i++ {
		IlluminanceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Illuminance__Array_to_C(cSlice []CIlluminance, goSlice []Illuminance) {
	for i := 0; i < len(goSlice); i++ {
		IlluminanceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
