/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <test_msgs/srv/basic_types.h>
*/
import "C"

import (
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("test_msgs/BasicTypes", BasicTypesTypeSupport)
}

type _BasicTypesTypeSupport struct {}

func (s _BasicTypesTypeSupport) Request() types.MessageTypeSupport {
	return BasicTypes_RequestTypeSupport
}

func (s _BasicTypesTypeSupport) Response() types.MessageTypeSupport {
	return BasicTypes_ResponseTypeSupport
}

func (s _BasicTypesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__test_msgs__srv__BasicTypes())
}

// Modifying this variable is undefined behavior.
var BasicTypesTypeSupport types.ServiceTypeSupport = _BasicTypesTypeSupport{}
