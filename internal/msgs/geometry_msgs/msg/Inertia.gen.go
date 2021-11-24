/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/inertia.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Inertia", InertiaTypeSupport)
}

// Do not create instances of this type directly. Always use NewInertia
// function instead.
type Inertia struct {
	M float64 `yaml:"m"`// Mass [kg]
	Com Vector3 `yaml:"com"`// Center of mass [m]
	Ixx float64 `yaml:"ixx"`// Inertia Tensor [kg-m^2]| ixx ixy ixz |I = | ixy iyy iyz || ixz iyz izz |
	Ixy float64 `yaml:"ixy"`// Inertia Tensor [kg-m^2]| ixx ixy ixz |I = | ixy iyy iyz || ixz iyz izz |
	Ixz float64 `yaml:"ixz"`// Inertia Tensor [kg-m^2]| ixx ixy ixz |I = | ixy iyy iyz || ixz iyz izz |
	Iyy float64 `yaml:"iyy"`// Inertia Tensor [kg-m^2]| ixx ixy ixz |I = | ixy iyy iyz || ixz iyz izz |
	Iyz float64 `yaml:"iyz"`// Inertia Tensor [kg-m^2]| ixx ixy ixz |I = | ixy iyy iyz || ixz iyz izz |
	Izz float64 `yaml:"izz"`// Inertia Tensor [kg-m^2]| ixx ixy ixz |I = | ixy iyy iyz || ixz iyz izz |
}

// NewInertia creates a new Inertia with default values.
func NewInertia() *Inertia {
	self := Inertia{}
	self.SetDefaults()
	return &self
}

func (t *Inertia) Clone() *Inertia {
	c := &Inertia{}
	c.M = t.M
	c.Com = *t.Com.Clone()
	c.Ixx = t.Ixx
	c.Ixy = t.Ixy
	c.Ixz = t.Ixz
	c.Iyy = t.Iyy
	c.Iyz = t.Iyz
	c.Izz = t.Izz
	return c
}

func (t *Inertia) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Inertia) SetDefaults() {
	t.M = 0
	t.Com.SetDefaults()
	t.Ixx = 0
	t.Ixy = 0
	t.Ixz = 0
	t.Iyy = 0
	t.Iyz = 0
	t.Izz = 0
}

// CloneInertiaSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInertiaSlice(dst, src []Inertia) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InertiaTypeSupport types.MessageTypeSupport = _InertiaTypeSupport{}

type _InertiaTypeSupport struct{}

func (t _InertiaTypeSupport) New() types.Message {
	return NewInertia()
}

func (t _InertiaTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Inertia
	return (unsafe.Pointer)(C.geometry_msgs__msg__Inertia__create())
}

func (t _InertiaTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Inertia__destroy((*C.geometry_msgs__msg__Inertia)(pointer_to_free))
}

func (t _InertiaTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Inertia)
	mem := (*C.geometry_msgs__msg__Inertia)(dst)
	mem.m = C.double(m.M)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.com), &m.Com)
	mem.ixx = C.double(m.Ixx)
	mem.ixy = C.double(m.Ixy)
	mem.ixz = C.double(m.Ixz)
	mem.iyy = C.double(m.Iyy)
	mem.iyz = C.double(m.Iyz)
	mem.izz = C.double(m.Izz)
}

func (t _InertiaTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Inertia)
	mem := (*C.geometry_msgs__msg__Inertia)(ros2_message_buffer)
	m.M = float64(mem.m)
	Vector3TypeSupport.AsGoStruct(&m.Com, unsafe.Pointer(&mem.com))
	m.Ixx = float64(mem.ixx)
	m.Ixy = float64(mem.ixy)
	m.Ixz = float64(mem.ixz)
	m.Iyy = float64(mem.iyy)
	m.Iyz = float64(mem.iyz)
	m.Izz = float64(mem.izz)
}

func (t _InertiaTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Inertia())
}

type CInertia = C.geometry_msgs__msg__Inertia
type CInertia__Sequence = C.geometry_msgs__msg__Inertia__Sequence

func Inertia__Sequence_to_Go(goSlice *[]Inertia, cSlice CInertia__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Inertia, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Inertia__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Inertia * uintptr(i)),
		))
		InertiaTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Inertia__Sequence_to_C(cSlice *CInertia__Sequence, goSlice []Inertia) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Inertia)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Inertia * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Inertia)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Inertia * uintptr(i)),
		))
		InertiaTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Inertia__Array_to_Go(goSlice []Inertia, cSlice []CInertia) {
	for i := 0; i < len(cSlice); i++ {
		InertiaTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Inertia__Array_to_C(cSlice []CInertia, goSlice []Inertia) {
	for i := 0; i < len(goSlice); i++ {
		InertiaTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
