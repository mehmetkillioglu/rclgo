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

#include <sensor_msgs/msg/nav_sat_fix.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/NavSatFix", NavSatFixTypeSupport)
}
const (
	NavSatFix_COVARIANCE_TYPE_UNKNOWN uint8 = 0
	NavSatFix_COVARIANCE_TYPE_APPROXIMATED uint8 = 1
	NavSatFix_COVARIANCE_TYPE_DIAGONAL_KNOWN uint8 = 2
	NavSatFix_COVARIANCE_TYPE_KNOWN uint8 = 3
)

// Do not create instances of this type directly. Always use NewNavSatFix
// function instead.
type NavSatFix struct {
	Header std_msgs_msg.Header `yaml:"header"`// header.stamp specifies the ROS time for this measurement (thecorresponding satellite time may be reported using thesensor_msgs/TimeReference message).header.frame_id is the frame of reference reported by the satellitereceiver, usually the location of the antenna.  This is aEuclidean frame relative to the vehicle, not a referenceellipsoid.
	Status NavSatStatus `yaml:"status"`// Satellite fix status information.
	Latitude float64 `yaml:"latitude"`// Latitude [degrees]. Positive is north of equator; negative is south.
	Longitude float64 `yaml:"longitude"`// Longitude [degrees]. Positive is east of prime meridian; negative is west.
	Altitude float64 `yaml:"altitude"`// Altitude [m]. Positive is above the WGS 84 ellipsoid(quiet NaN if no altitude is available).
	PositionCovariance [9]float64 `yaml:"position_covariance"`// Position covariance [m^2] defined relative to a tangential planethrough the reported position. The components are East, North, andUp (ENU), in row-major order.Beware: this coordinate system exhibits singularities at the poles.
	PositionCovarianceType uint8 `yaml:"position_covariance_type"`
}

// NewNavSatFix creates a new NavSatFix with default values.
func NewNavSatFix() *NavSatFix {
	self := NavSatFix{}
	self.SetDefaults()
	return &self
}

func (t *NavSatFix) Clone() *NavSatFix {
	c := &NavSatFix{}
	c.Header = *t.Header.Clone()
	c.Status = *t.Status.Clone()
	c.Latitude = t.Latitude
	c.Longitude = t.Longitude
	c.Altitude = t.Altitude
	c.PositionCovariance = t.PositionCovariance
	c.PositionCovarianceType = t.PositionCovarianceType
	return c
}

func (t *NavSatFix) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NavSatFix) SetDefaults() {
	t.Header.SetDefaults()
	t.Status.SetDefaults()
	t.Latitude = 0
	t.Longitude = 0
	t.Altitude = 0
	t.PositionCovariance = [9]float64{}
	t.PositionCovarianceType = 0
}

// CloneNavSatFixSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNavSatFixSlice(dst, src []NavSatFix) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NavSatFixTypeSupport types.MessageTypeSupport = _NavSatFixTypeSupport{}

type _NavSatFixTypeSupport struct{}

func (t _NavSatFixTypeSupport) New() types.Message {
	return NewNavSatFix()
}

func (t _NavSatFixTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__NavSatFix
	return (unsafe.Pointer)(C.sensor_msgs__msg__NavSatFix__create())
}

func (t _NavSatFixTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__NavSatFix__destroy((*C.sensor_msgs__msg__NavSatFix)(pointer_to_free))
}

func (t _NavSatFixTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NavSatFix)
	mem := (*C.sensor_msgs__msg__NavSatFix)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	NavSatStatusTypeSupport.AsCStruct(unsafe.Pointer(&mem.status), &m.Status)
	mem.latitude = C.double(m.Latitude)
	mem.longitude = C.double(m.Longitude)
	mem.altitude = C.double(m.Altitude)
	cSlice_position_covariance := mem.position_covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_position_covariance)), m.PositionCovariance[:])
	mem.position_covariance_type = C.uint8_t(m.PositionCovarianceType)
}

func (t _NavSatFixTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NavSatFix)
	mem := (*C.sensor_msgs__msg__NavSatFix)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	NavSatStatusTypeSupport.AsGoStruct(&m.Status, unsafe.Pointer(&mem.status))
	m.Latitude = float64(mem.latitude)
	m.Longitude = float64(mem.longitude)
	m.Altitude = float64(mem.altitude)
	cSlice_position_covariance := mem.position_covariance[:]
	primitives.Float64__Array_to_Go(m.PositionCovariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_position_covariance)))
	m.PositionCovarianceType = uint8(mem.position_covariance_type)
}

func (t _NavSatFixTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__NavSatFix())
}

type CNavSatFix = C.sensor_msgs__msg__NavSatFix
type CNavSatFix__Sequence = C.sensor_msgs__msg__NavSatFix__Sequence

func NavSatFix__Sequence_to_Go(goSlice *[]NavSatFix, cSlice CNavSatFix__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NavSatFix, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__NavSatFix__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__NavSatFix * uintptr(i)),
		))
		NavSatFixTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func NavSatFix__Sequence_to_C(cSlice *CNavSatFix__Sequence, goSlice []NavSatFix) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__NavSatFix)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__NavSatFix * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__NavSatFix)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__NavSatFix * uintptr(i)),
		))
		NavSatFixTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func NavSatFix__Array_to_Go(goSlice []NavSatFix, cSlice []CNavSatFix) {
	for i := 0; i < len(cSlice); i++ {
		NavSatFixTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NavSatFix__Array_to_C(cSlice []CNavSatFix, goSlice []NavSatFix) {
	for i := 0; i < len(goSlice); i++ {
		NavSatFixTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
