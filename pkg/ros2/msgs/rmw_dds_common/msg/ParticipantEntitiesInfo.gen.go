/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rmw_dds_common
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrmw_dds_common__rosidl_typesupport_c -lrmw_dds_common__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rmw_dds_common/msg/participant_entities_info.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("rmw_dds_common/ParticipantEntitiesInfo", &ParticipantEntitiesInfo{})
}

// Do not create instances of this type directly. Always use NewParticipantEntitiesInfo
// function instead.
type ParticipantEntitiesInfo struct {
	Gid Gid `yaml:"gid"`
	NodeEntitiesInfoSeq []NodeEntitiesInfo `yaml:"node_entities_info_seq"`
}

// NewParticipantEntitiesInfo creates a new ParticipantEntitiesInfo with default values.
func NewParticipantEntitiesInfo() *ParticipantEntitiesInfo {
	self := ParticipantEntitiesInfo{}
	self.SetDefaults(nil)
	return &self
}

func (t *ParticipantEntitiesInfo) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Gid.SetDefaults(nil)
	
	return t
}

func (t *ParticipantEntitiesInfo) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rmw_dds_common__msg__ParticipantEntitiesInfo())
}
func (t *ParticipantEntitiesInfo) PrepareMemory() unsafe.Pointer { //returns *C.rmw_dds_common__msg__ParticipantEntitiesInfo
	return (unsafe.Pointer)(C.rmw_dds_common__msg__ParticipantEntitiesInfo__create())
}
func (t *ParticipantEntitiesInfo) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rmw_dds_common__msg__ParticipantEntitiesInfo__destroy((*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(pointer_to_free))
}
func (t *ParticipantEntitiesInfo) AsCStruct() unsafe.Pointer {
	mem := (*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(t.PrepareMemory())
	mem.gid = *(*C.rmw_dds_common__msg__Gid)(t.Gid.AsCStruct())
	NodeEntitiesInfo__Sequence_to_C(&mem.node_entities_info_seq, t.NodeEntitiesInfoSeq)
	return unsafe.Pointer(mem)
}
func (t *ParticipantEntitiesInfo) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(ros2_message_buffer)
	t.Gid.AsGoStruct(unsafe.Pointer(&mem.gid))
	NodeEntitiesInfo__Sequence_to_Go(&t.NodeEntitiesInfoSeq, mem.node_entities_info_seq)
}
func (t *ParticipantEntitiesInfo) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CParticipantEntitiesInfo = C.rmw_dds_common__msg__ParticipantEntitiesInfo
type CParticipantEntitiesInfo__Sequence = C.rmw_dds_common__msg__ParticipantEntitiesInfo__Sequence

func ParticipantEntitiesInfo__Sequence_to_Go(goSlice *[]ParticipantEntitiesInfo, cSlice CParticipantEntitiesInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ParticipantEntitiesInfo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rmw_dds_common__msg__ParticipantEntitiesInfo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rmw_dds_common__msg__ParticipantEntitiesInfo * uintptr(i)),
		))
		(*goSlice)[i] = ParticipantEntitiesInfo{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ParticipantEntitiesInfo__Sequence_to_C(cSlice *CParticipantEntitiesInfo__Sequence, goSlice []ParticipantEntitiesInfo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(C.malloc((C.size_t)(C.sizeof_struct_rmw_dds_common__msg__ParticipantEntitiesInfo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rmw_dds_common__msg__ParticipantEntitiesInfo * uintptr(i)),
		))
		*cIdx = *(*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(v.AsCStruct())
	}
}
func ParticipantEntitiesInfo__Array_to_Go(goSlice []ParticipantEntitiesInfo, cSlice []CParticipantEntitiesInfo) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ParticipantEntitiesInfo__Array_to_C(cSlice []CParticipantEntitiesInfo, goSlice []ParticipantEntitiesInfo) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.rmw_dds_common__msg__ParticipantEntitiesInfo)(goSlice[i].AsCStruct())
	}
}

