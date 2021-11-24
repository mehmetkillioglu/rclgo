/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/defaults.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Defaults", DefaultsTypeSupport)
}

// Do not create instances of this type directly. Always use NewDefaults
// function instead.
type Defaults struct {
	BoolValue bool `yaml:"bool_value"`
	ByteValue byte `yaml:"byte_value"`
	CharValue byte `yaml:"char_value"`
	Float32Value float32 `yaml:"float32_value"`
	Float64Value float64 `yaml:"float64_value"`
	Int8Value int8 `yaml:"int8_value"`
	Uint8Value uint8 `yaml:"uint8_value"`
	Int16Value int16 `yaml:"int16_value"`
	Uint16Value uint16 `yaml:"uint16_value"`
	Int32Value int32 `yaml:"int32_value"`
	Uint32Value uint32 `yaml:"uint32_value"`
	Int64Value int64 `yaml:"int64_value"`
	Uint64Value uint64 `yaml:"uint64_value"`
}

// NewDefaults creates a new Defaults with default values.
func NewDefaults() *Defaults {
	self := Defaults{}
	self.SetDefaults()
	return &self
}

func (t *Defaults) Clone() *Defaults {
	c := &Defaults{}
	c.BoolValue = t.BoolValue
	c.ByteValue = t.ByteValue
	c.CharValue = t.CharValue
	c.Float32Value = t.Float32Value
	c.Float64Value = t.Float64Value
	c.Int8Value = t.Int8Value
	c.Uint8Value = t.Uint8Value
	c.Int16Value = t.Int16Value
	c.Uint16Value = t.Uint16Value
	c.Int32Value = t.Int32Value
	c.Uint32Value = t.Uint32Value
	c.Int64Value = t.Int64Value
	c.Uint64Value = t.Uint64Value
	return c
}

func (t *Defaults) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Defaults) SetDefaults() {
	t.BoolValue = true
	t.ByteValue = 50
	t.CharValue = 100
	t.Float32Value = 1.125
	t.Float64Value = 1.125
	t.Int8Value = -50
	t.Uint8Value = 200
	t.Int16Value = -1000
	t.Uint16Value = 2000
	t.Int32Value = -30000
	t.Uint32Value = 60000
	t.Int64Value = -40000000
	t.Uint64Value = 50000000
}

// CloneDefaultsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDefaultsSlice(dst, src []Defaults) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DefaultsTypeSupport types.MessageTypeSupport = _DefaultsTypeSupport{}

type _DefaultsTypeSupport struct{}

func (t _DefaultsTypeSupport) New() types.Message {
	return NewDefaults()
}

func (t _DefaultsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Defaults
	return (unsafe.Pointer)(C.test_msgs__msg__Defaults__create())
}

func (t _DefaultsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Defaults__destroy((*C.test_msgs__msg__Defaults)(pointer_to_free))
}

func (t _DefaultsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Defaults)
	mem := (*C.test_msgs__msg__Defaults)(dst)
	mem.bool_value = C.bool(m.BoolValue)
	mem.byte_value = C.uint8_t(m.ByteValue)
	mem.char_value = C.uchar(m.CharValue)
	mem.float32_value = C.float(m.Float32Value)
	mem.float64_value = C.double(m.Float64Value)
	mem.int8_value = C.int8_t(m.Int8Value)
	mem.uint8_value = C.uint8_t(m.Uint8Value)
	mem.int16_value = C.int16_t(m.Int16Value)
	mem.uint16_value = C.uint16_t(m.Uint16Value)
	mem.int32_value = C.int32_t(m.Int32Value)
	mem.uint32_value = C.uint32_t(m.Uint32Value)
	mem.int64_value = C.int64_t(m.Int64Value)
	mem.uint64_value = C.uint64_t(m.Uint64Value)
}

func (t _DefaultsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Defaults)
	mem := (*C.test_msgs__msg__Defaults)(ros2_message_buffer)
	m.BoolValue = bool(mem.bool_value)
	m.ByteValue = byte(mem.byte_value)
	m.CharValue = byte(mem.char_value)
	m.Float32Value = float32(mem.float32_value)
	m.Float64Value = float64(mem.float64_value)
	m.Int8Value = int8(mem.int8_value)
	m.Uint8Value = uint8(mem.uint8_value)
	m.Int16Value = int16(mem.int16_value)
	m.Uint16Value = uint16(mem.uint16_value)
	m.Int32Value = int32(mem.int32_value)
	m.Uint32Value = uint32(mem.uint32_value)
	m.Int64Value = int64(mem.int64_value)
	m.Uint64Value = uint64(mem.uint64_value)
}

func (t _DefaultsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Defaults())
}

type CDefaults = C.test_msgs__msg__Defaults
type CDefaults__Sequence = C.test_msgs__msg__Defaults__Sequence

func Defaults__Sequence_to_Go(goSlice *[]Defaults, cSlice CDefaults__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Defaults, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__Defaults__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Defaults * uintptr(i)),
		))
		DefaultsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Defaults__Sequence_to_C(cSlice *CDefaults__Sequence, goSlice []Defaults) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Defaults)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__Defaults * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__Defaults)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Defaults * uintptr(i)),
		))
		DefaultsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Defaults__Array_to_Go(goSlice []Defaults, cSlice []CDefaults) {
	for i := 0; i < len(cSlice); i++ {
		DefaultsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Defaults__Array_to_C(cSlice []CDefaults, goSlice []Defaults) {
	for i := 0; i < len(goSlice); i++ {
		DefaultsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
