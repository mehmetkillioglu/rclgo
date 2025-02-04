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

#include <sensor_msgs/msg/image.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Image", ImageTypeSupport)
}

// Do not create instances of this type directly. Always use NewImage
// function instead.
type Image struct {
	Header std_msgs_msg.Header `yaml:"header"`// Header timestamp should be acquisition time of image
	Height uint32 `yaml:"height"`// image height, that is, number of rows
	Width uint32 `yaml:"width"`// image width, that is, number of columns
	Encoding string `yaml:"encoding"`// Encoding of pixels -- channel meaning, ordering, size
	IsBigendian uint8 `yaml:"is_bigendian"`// is this data bigendian?
	Step uint32 `yaml:"step"`// Full row length in bytes
	Data []uint8 `yaml:"data"`// actual matrix data, size is (step * rows)
}

// NewImage creates a new Image with default values.
func NewImage() *Image {
	self := Image{}
	self.SetDefaults()
	return &self
}

func (t *Image) Clone() *Image {
	c := &Image{}
	c.Header = *t.Header.Clone()
	c.Height = t.Height
	c.Width = t.Width
	c.Encoding = t.Encoding
	c.IsBigendian = t.IsBigendian
	c.Step = t.Step
	if t.Data != nil {
		c.Data = make([]uint8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Image) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Image) SetDefaults() {
	t.Header.SetDefaults()
	t.Height = 0
	t.Width = 0
	t.Encoding = ""
	t.IsBigendian = 0
	t.Step = 0
	t.Data = nil
}

// CloneImageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneImageSlice(dst, src []Image) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ImageTypeSupport types.MessageTypeSupport = _ImageTypeSupport{}

type _ImageTypeSupport struct{}

func (t _ImageTypeSupport) New() types.Message {
	return NewImage()
}

func (t _ImageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Image
	return (unsafe.Pointer)(C.sensor_msgs__msg__Image__create())
}

func (t _ImageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Image__destroy((*C.sensor_msgs__msg__Image)(pointer_to_free))
}

func (t _ImageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Image)
	mem := (*C.sensor_msgs__msg__Image)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.height = C.uint32_t(m.Height)
	mem.width = C.uint32_t(m.Width)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.encoding), m.Encoding)
	mem.is_bigendian = C.uint8_t(m.IsBigendian)
	mem.step = C.uint32_t(m.Step)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _ImageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Image)
	mem := (*C.sensor_msgs__msg__Image)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Height = uint32(mem.height)
	m.Width = uint32(mem.width)
	primitives.StringAsGoStruct(&m.Encoding, unsafe.Pointer(&mem.encoding))
	m.IsBigendian = uint8(mem.is_bigendian)
	m.Step = uint32(mem.step)
	primitives.Uint8__Sequence_to_Go(&m.Data, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _ImageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Image())
}

type CImage = C.sensor_msgs__msg__Image
type CImage__Sequence = C.sensor_msgs__msg__Image__Sequence

func Image__Sequence_to_Go(goSlice *[]Image, cSlice CImage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Image, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__Image__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Image * uintptr(i)),
		))
		ImageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Image__Sequence_to_C(cSlice *CImage__Sequence, goSlice []Image) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Image)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__Image * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__Image)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Image * uintptr(i)),
		))
		ImageTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Image__Array_to_Go(goSlice []Image, cSlice []CImage) {
	for i := 0; i < len(cSlice); i++ {
		ImageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Image__Array_to_C(cSlice []CImage, goSlice []Image) {
	for i := 0; i < len(goSlice); i++ {
		ImageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
