/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "github.com/tiiuae/rclgo/pkg/rclgo/msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/tiiuae/rclgo/pkg/rclgo/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/interactive_marker_pose.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/InteractiveMarkerPose", InteractiveMarkerPoseTypeSupport)
}

// Do not create instances of this type directly. Always use NewInteractiveMarkerPose
// function instead.
type InteractiveMarkerPose struct {
	Header std_msgs_msg.Header `yaml:"header"`// Time/frame info.
	Pose geometry_msgs_msg.Pose `yaml:"pose"`// Initial pose. Also, defines the pivot point for rotations.
	Name string `yaml:"name"`// Identifying string. Must be globally unique inthe topic that this message is sent through.
}

// NewInteractiveMarkerPose creates a new InteractiveMarkerPose with default values.
func NewInteractiveMarkerPose() *InteractiveMarkerPose {
	self := InteractiveMarkerPose{}
	self.SetDefaults()
	return &self
}

func (t *InteractiveMarkerPose) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *InteractiveMarkerPose) SetDefaults() {
	t.Header.SetDefaults()
	t.Pose.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var InteractiveMarkerPoseTypeSupport types.MessageTypeSupport = _InteractiveMarkerPoseTypeSupport{}

type _InteractiveMarkerPoseTypeSupport struct{}

func (t _InteractiveMarkerPoseTypeSupport) New() types.Message {
	return NewInteractiveMarkerPose()
}

func (t _InteractiveMarkerPoseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarkerPose
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarkerPose__create())
}

func (t _InteractiveMarkerPoseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarkerPose__destroy((*C.visualization_msgs__msg__InteractiveMarkerPose)(pointer_to_free))
}

func (t _InteractiveMarkerPoseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InteractiveMarkerPose)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerPose)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	geometry_msgs_msg.PoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
}

func (t _InteractiveMarkerPoseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InteractiveMarkerPose)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerPose)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	geometry_msgs_msg.PoseTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
}

func (t _InteractiveMarkerPoseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarkerPose())
}

type CInteractiveMarkerPose = C.visualization_msgs__msg__InteractiveMarkerPose
type CInteractiveMarkerPose__Sequence = C.visualization_msgs__msg__InteractiveMarkerPose__Sequence

func InteractiveMarkerPose__Sequence_to_Go(goSlice *[]InteractiveMarkerPose, cSlice CInteractiveMarkerPose__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarkerPose, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerPose__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerPose * uintptr(i)),
		))
		InteractiveMarkerPoseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InteractiveMarkerPose__Sequence_to_C(cSlice *CInteractiveMarkerPose__Sequence, goSlice []InteractiveMarkerPose) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarkerPose)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerPose * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerPose)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerPose * uintptr(i)),
		))
		InteractiveMarkerPoseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InteractiveMarkerPose__Array_to_Go(goSlice []InteractiveMarkerPose, cSlice []CInteractiveMarkerPose) {
	for i := 0; i < len(cSlice); i++ {
		InteractiveMarkerPoseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarkerPose__Array_to_C(cSlice []CInteractiveMarkerPose, goSlice []InteractiveMarkerPose) {
	for i := 0; i < len(goSlice); i++ {
		InteractiveMarkerPoseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}