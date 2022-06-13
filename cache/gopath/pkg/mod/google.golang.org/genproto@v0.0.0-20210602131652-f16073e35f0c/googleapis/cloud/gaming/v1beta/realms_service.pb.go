// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: google/cloud/gaming/v1beta/realms_service.proto

package gaming

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_google_cloud_gaming_v1beta_realms_service_proto protoreflect.FileDescriptor

var file_google_cloud_gaming_v1beta_realms_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xf7, 0x09, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x73, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x99, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x6d,
	0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d,
	0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0xd0, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x22, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x3a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0xda, 0x41, 0x15,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x2c, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x5f, 0x69, 0x64, 0xca, 0x41, 0x1a, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0xc8, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x6a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x2a, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xca, 0x41, 0x2a, 0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x11, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0xd2, 0x01,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3d, 0x32, 0x34, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a,
	0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0xda, 0x41, 0x11, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0xca, 0x41, 0x1a, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0xd6, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x61, 0x6c, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x61, 0x6c, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b,
	0x32, 0x42, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x1a, 0x4f, 0xca, 0x41, 0x1b,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x81, 0x01, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x50,
	0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x47, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_gaming_v1beta_realms_service_proto_goTypes = []interface{}{
	(*ListRealmsRequest)(nil),          // 0: google.cloud.gaming.v1beta.ListRealmsRequest
	(*GetRealmRequest)(nil),            // 1: google.cloud.gaming.v1beta.GetRealmRequest
	(*CreateRealmRequest)(nil),         // 2: google.cloud.gaming.v1beta.CreateRealmRequest
	(*DeleteRealmRequest)(nil),         // 3: google.cloud.gaming.v1beta.DeleteRealmRequest
	(*UpdateRealmRequest)(nil),         // 4: google.cloud.gaming.v1beta.UpdateRealmRequest
	(*PreviewRealmUpdateRequest)(nil),  // 5: google.cloud.gaming.v1beta.PreviewRealmUpdateRequest
	(*ListRealmsResponse)(nil),         // 6: google.cloud.gaming.v1beta.ListRealmsResponse
	(*Realm)(nil),                      // 7: google.cloud.gaming.v1beta.Realm
	(*longrunning.Operation)(nil),      // 8: google.longrunning.Operation
	(*PreviewRealmUpdateResponse)(nil), // 9: google.cloud.gaming.v1beta.PreviewRealmUpdateResponse
}
var file_google_cloud_gaming_v1beta_realms_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.gaming.v1beta.RealmsService.ListRealms:input_type -> google.cloud.gaming.v1beta.ListRealmsRequest
	1, // 1: google.cloud.gaming.v1beta.RealmsService.GetRealm:input_type -> google.cloud.gaming.v1beta.GetRealmRequest
	2, // 2: google.cloud.gaming.v1beta.RealmsService.CreateRealm:input_type -> google.cloud.gaming.v1beta.CreateRealmRequest
	3, // 3: google.cloud.gaming.v1beta.RealmsService.DeleteRealm:input_type -> google.cloud.gaming.v1beta.DeleteRealmRequest
	4, // 4: google.cloud.gaming.v1beta.RealmsService.UpdateRealm:input_type -> google.cloud.gaming.v1beta.UpdateRealmRequest
	5, // 5: google.cloud.gaming.v1beta.RealmsService.PreviewRealmUpdate:input_type -> google.cloud.gaming.v1beta.PreviewRealmUpdateRequest
	6, // 6: google.cloud.gaming.v1beta.RealmsService.ListRealms:output_type -> google.cloud.gaming.v1beta.ListRealmsResponse
	7, // 7: google.cloud.gaming.v1beta.RealmsService.GetRealm:output_type -> google.cloud.gaming.v1beta.Realm
	8, // 8: google.cloud.gaming.v1beta.RealmsService.CreateRealm:output_type -> google.longrunning.Operation
	8, // 9: google.cloud.gaming.v1beta.RealmsService.DeleteRealm:output_type -> google.longrunning.Operation
	8, // 10: google.cloud.gaming.v1beta.RealmsService.UpdateRealm:output_type -> google.longrunning.Operation
	9, // 11: google.cloud.gaming.v1beta.RealmsService.PreviewRealmUpdate:output_type -> google.cloud.gaming.v1beta.PreviewRealmUpdateResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_gaming_v1beta_realms_service_proto_init() }
func file_google_cloud_gaming_v1beta_realms_service_proto_init() {
	if File_google_cloud_gaming_v1beta_realms_service_proto != nil {
		return
	}
	file_google_cloud_gaming_v1beta_realms_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_gaming_v1beta_realms_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_gaming_v1beta_realms_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_gaming_v1beta_realms_service_proto_depIdxs,
	}.Build()
	File_google_cloud_gaming_v1beta_realms_service_proto = out.File
	file_google_cloud_gaming_v1beta_realms_service_proto_rawDesc = nil
	file_google_cloud_gaming_v1beta_realms_service_proto_goTypes = nil
	file_google_cloud_gaming_v1beta_realms_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RealmsServiceClient is the client API for RealmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RealmsServiceClient interface {
	// Lists realms in a given project and location.
	ListRealms(ctx context.Context, in *ListRealmsRequest, opts ...grpc.CallOption) (*ListRealmsResponse, error)
	// Gets details of a single realm.
	GetRealm(ctx context.Context, in *GetRealmRequest, opts ...grpc.CallOption) (*Realm, error)
	// Creates a new realm in a given project and location.
	CreateRealm(ctx context.Context, in *CreateRealmRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Deletes a single realm.
	DeleteRealm(ctx context.Context, in *DeleteRealmRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Patches a single realm.
	UpdateRealm(ctx context.Context, in *UpdateRealmRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Previews patches to a single realm.
	PreviewRealmUpdate(ctx context.Context, in *PreviewRealmUpdateRequest, opts ...grpc.CallOption) (*PreviewRealmUpdateResponse, error)
}

type realmsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRealmsServiceClient(cc grpc.ClientConnInterface) RealmsServiceClient {
	return &realmsServiceClient{cc}
}

func (c *realmsServiceClient) ListRealms(ctx context.Context, in *ListRealmsRequest, opts ...grpc.CallOption) (*ListRealmsResponse, error) {
	out := new(ListRealmsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.RealmsService/ListRealms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) GetRealm(ctx context.Context, in *GetRealmRequest, opts ...grpc.CallOption) (*Realm, error) {
	out := new(Realm)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.RealmsService/GetRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) CreateRealm(ctx context.Context, in *CreateRealmRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.RealmsService/CreateRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) DeleteRealm(ctx context.Context, in *DeleteRealmRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.RealmsService/DeleteRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) UpdateRealm(ctx context.Context, in *UpdateRealmRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.RealmsService/UpdateRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) PreviewRealmUpdate(ctx context.Context, in *PreviewRealmUpdateRequest, opts ...grpc.CallOption) (*PreviewRealmUpdateResponse, error) {
	out := new(PreviewRealmUpdateResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.RealmsService/PreviewRealmUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealmsServiceServer is the server API for RealmsService service.
type RealmsServiceServer interface {
	// Lists realms in a given project and location.
	ListRealms(context.Context, *ListRealmsRequest) (*ListRealmsResponse, error)
	// Gets details of a single realm.
	GetRealm(context.Context, *GetRealmRequest) (*Realm, error)
	// Creates a new realm in a given project and location.
	CreateRealm(context.Context, *CreateRealmRequest) (*longrunning.Operation, error)
	// Deletes a single realm.
	DeleteRealm(context.Context, *DeleteRealmRequest) (*longrunning.Operation, error)
	// Patches a single realm.
	UpdateRealm(context.Context, *UpdateRealmRequest) (*longrunning.Operation, error)
	// Previews patches to a single realm.
	PreviewRealmUpdate(context.Context, *PreviewRealmUpdateRequest) (*PreviewRealmUpdateResponse, error)
}

// UnimplementedRealmsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRealmsServiceServer struct {
}

func (*UnimplementedRealmsServiceServer) ListRealms(context.Context, *ListRealmsRequest) (*ListRealmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRealms not implemented")
}
func (*UnimplementedRealmsServiceServer) GetRealm(context.Context, *GetRealmRequest) (*Realm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealm not implemented")
}
func (*UnimplementedRealmsServiceServer) CreateRealm(context.Context, *CreateRealmRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRealm not implemented")
}
func (*UnimplementedRealmsServiceServer) DeleteRealm(context.Context, *DeleteRealmRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRealm not implemented")
}
func (*UnimplementedRealmsServiceServer) UpdateRealm(context.Context, *UpdateRealmRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRealm not implemented")
}
func (*UnimplementedRealmsServiceServer) PreviewRealmUpdate(context.Context, *PreviewRealmUpdateRequest) (*PreviewRealmUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRealmUpdate not implemented")
}

func RegisterRealmsServiceServer(s *grpc.Server, srv RealmsServiceServer) {
	s.RegisterService(&_RealmsService_serviceDesc, srv)
}

func _RealmsService_ListRealms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRealmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).ListRealms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.RealmsService/ListRealms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).ListRealms(ctx, req.(*ListRealmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_GetRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).GetRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.RealmsService/GetRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).GetRealm(ctx, req.(*GetRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_CreateRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).CreateRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.RealmsService/CreateRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).CreateRealm(ctx, req.(*CreateRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_DeleteRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).DeleteRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.RealmsService/DeleteRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).DeleteRealm(ctx, req.(*DeleteRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_UpdateRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).UpdateRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.RealmsService/UpdateRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).UpdateRealm(ctx, req.(*UpdateRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_PreviewRealmUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewRealmUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).PreviewRealmUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.RealmsService/PreviewRealmUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).PreviewRealmUpdate(ctx, req.(*PreviewRealmUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RealmsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.gaming.v1beta.RealmsService",
	HandlerType: (*RealmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRealms",
			Handler:    _RealmsService_ListRealms_Handler,
		},
		{
			MethodName: "GetRealm",
			Handler:    _RealmsService_GetRealm_Handler,
		},
		{
			MethodName: "CreateRealm",
			Handler:    _RealmsService_CreateRealm_Handler,
		},
		{
			MethodName: "DeleteRealm",
			Handler:    _RealmsService_DeleteRealm_Handler,
		},
		{
			MethodName: "UpdateRealm",
			Handler:    _RealmsService_UpdateRealm_Handler,
		},
		{
			MethodName: "PreviewRealmUpdate",
			Handler:    _RealmsService_PreviewRealmUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/gaming/v1beta/realms_service.proto",
}