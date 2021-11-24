/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <sensor_msgs/srv/set_camera_info.h>
*/
import "C"

import (
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("sensor_msgs/SetCameraInfo", SetCameraInfoTypeSupport)
}

type _SetCameraInfoTypeSupport struct {}

func (s _SetCameraInfoTypeSupport) Request() types.MessageTypeSupport {
	return SetCameraInfo_RequestTypeSupport
}

func (s _SetCameraInfoTypeSupport) Response() types.MessageTypeSupport {
	return SetCameraInfo_ResponseTypeSupport
}

func (s _SetCameraInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__sensor_msgs__srv__SetCameraInfo())
}

// Modifying this variable is undefined behavior.
var SetCameraInfoTypeSupport types.ServiceTypeSupport = _SetCameraInfoTypeSupport{}
