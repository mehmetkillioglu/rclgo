/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_srvs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_srvs__rosidl_typesupport_c -lstd_srvs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_srvs/srv/set_bool.h>
*/
import "C"

import (
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/typemap"
	"github.com/mehmetkillioglu/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("std_srvs/SetBool", SetBoolTypeSupport)
}

type _SetBoolTypeSupport struct {}

func (s _SetBoolTypeSupport) Request() types.MessageTypeSupport {
	return SetBool_RequestTypeSupport
}

func (s _SetBoolTypeSupport) Response() types.MessageTypeSupport {
	return SetBool_ResponseTypeSupport
}

func (s _SetBoolTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__std_srvs__srv__SetBool())
}

// Modifying this variable is undefined behavior.
var SetBoolTypeSupport types.ServiceTypeSupport = _SetBoolTypeSupport{}
