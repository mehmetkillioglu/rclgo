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

#include <px4_msgs/msg/vehicle_imu.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleImu", VehicleImuTypeSupport)
}
const (
	VehicleImu_CLIPPING_X uint8 = 1
	VehicleImu_CLIPPING_Y uint8 = 2
	VehicleImu_CLIPPING_Z uint8 = 4
)

// Do not create instances of this type directly. Always use NewVehicleImu
// function instead.
type VehicleImu struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`
	AccelDeviceId uint32 `yaml:"accel_device_id"`// Accelerometer unique device ID for the sensor that does not change between power cycles
	GyroDeviceId uint32 `yaml:"gyro_device_id"`// Gyroscope unique device ID for the sensor that does not change between power cycles
	DeltaAngle [3]float32 `yaml:"delta_angle"`// delta angle about the FRD body frame XYZ-axis in rad over the integration time frame (delta_angle_dt)
	DeltaVelocity [3]float32 `yaml:"delta_velocity"`// delta velocity in the FRD body frame XYZ-axis in m/s over the integration time frame (delta_velocity_dt)
	DeltaAngleDt uint16 `yaml:"delta_angle_dt"`// integration period in microseconds
	DeltaVelocityDt uint16 `yaml:"delta_velocity_dt"`// integration period in microseconds
	DeltaVelocityClipping uint8 `yaml:"delta_velocity_clipping"`// bitfield indicating if there was any accelerometer clipping (per axis) during the integration time frame
	CalibrationCount uint8 `yaml:"calibration_count"`// Calibration changed counter. Monotonically increases whenever calibration changes.
}

// NewVehicleImu creates a new VehicleImu with default values.
func NewVehicleImu() *VehicleImu {
	self := VehicleImu{}
	self.SetDefaults()
	return &self
}

func (t *VehicleImu) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *VehicleImu) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var VehicleImuTypeSupport types.MessageTypeSupport = _VehicleImuTypeSupport{}

type _VehicleImuTypeSupport struct{}

func (t _VehicleImuTypeSupport) New() types.Message {
	return NewVehicleImu()
}

func (t _VehicleImuTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleImu
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleImu__create())
}

func (t _VehicleImuTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleImu__destroy((*C.px4_msgs__msg__VehicleImu)(pointer_to_free))
}

func (t _VehicleImuTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleImu)
	mem := (*C.px4_msgs__msg__VehicleImu)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.accel_device_id = C.uint32_t(m.AccelDeviceId)
	mem.gyro_device_id = C.uint32_t(m.GyroDeviceId)
	cSlice_delta_angle := mem.delta_angle[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_angle)), m.DeltaAngle[:])
	cSlice_delta_velocity := mem.delta_velocity[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_velocity)), m.DeltaVelocity[:])
	mem.delta_angle_dt = C.uint16_t(m.DeltaAngleDt)
	mem.delta_velocity_dt = C.uint16_t(m.DeltaVelocityDt)
	mem.delta_velocity_clipping = C.uint8_t(m.DeltaVelocityClipping)
	mem.calibration_count = C.uint8_t(m.CalibrationCount)
}

func (t _VehicleImuTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleImu)
	mem := (*C.px4_msgs__msg__VehicleImu)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.AccelDeviceId = uint32(mem.accel_device_id)
	m.GyroDeviceId = uint32(mem.gyro_device_id)
	cSlice_delta_angle := mem.delta_angle[:]
	primitives.Float32__Array_to_Go(m.DeltaAngle[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_angle)))
	cSlice_delta_velocity := mem.delta_velocity[:]
	primitives.Float32__Array_to_Go(m.DeltaVelocity[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_velocity)))
	m.DeltaAngleDt = uint16(mem.delta_angle_dt)
	m.DeltaVelocityDt = uint16(mem.delta_velocity_dt)
	m.DeltaVelocityClipping = uint8(mem.delta_velocity_clipping)
	m.CalibrationCount = uint8(mem.calibration_count)
}

func (t _VehicleImuTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleImu())
}

type CVehicleImu = C.px4_msgs__msg__VehicleImu
type CVehicleImu__Sequence = C.px4_msgs__msg__VehicleImu__Sequence

func VehicleImu__Sequence_to_Go(goSlice *[]VehicleImu, cSlice CVehicleImu__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleImu, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleImu__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleImu * uintptr(i)),
		))
		VehicleImuTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleImu__Sequence_to_C(cSlice *CVehicleImu__Sequence, goSlice []VehicleImu) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleImu)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleImu * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleImu)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleImu * uintptr(i)),
		))
		VehicleImuTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleImu__Array_to_Go(goSlice []VehicleImu, cSlice []CVehicleImu) {
	for i := 0; i < len(cSlice); i++ {
		VehicleImuTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleImu__Array_to_C(cSlice []CVehicleImu, goSlice []VehicleImu) {
	for i := 0; i < len(goSlice); i++ {
		VehicleImuTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}