// Copyright 2021 Google LLC
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
// 	protoc        v3.13.0
// source: google/apps/drive/activity/v2/drive_activity_service.proto

package activity

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_google_apps_drive_activity_v2_drive_activity_service_proto protoreflect.FileDescriptor

var file_google_apps_drive_activity_v2_drive_activity_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc, 0x02, 0x0a, 0x14, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa8, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x3a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x1a, 0x88, 0x01, 0xca, 0x41,
	0x1c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x66,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2c, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x65,
	0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x42, 0xce, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x42, 0x19, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0xa2, 0x02, 0x04, 0x47, 0x41, 0x44, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x5c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_apps_drive_activity_v2_drive_activity_service_proto_goTypes = []interface{}{
	(*QueryDriveActivityRequest)(nil),  // 0: google.apps.drive.activity.v2.QueryDriveActivityRequest
	(*QueryDriveActivityResponse)(nil), // 1: google.apps.drive.activity.v2.QueryDriveActivityResponse
}
var file_google_apps_drive_activity_v2_drive_activity_service_proto_depIdxs = []int32{
	0, // 0: google.apps.drive.activity.v2.DriveActivityService.QueryDriveActivity:input_type -> google.apps.drive.activity.v2.QueryDriveActivityRequest
	1, // 1: google.apps.drive.activity.v2.DriveActivityService.QueryDriveActivity:output_type -> google.apps.drive.activity.v2.QueryDriveActivityResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_apps_drive_activity_v2_drive_activity_service_proto_init() }
func file_google_apps_drive_activity_v2_drive_activity_service_proto_init() {
	if File_google_apps_drive_activity_v2_drive_activity_service_proto != nil {
		return
	}
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_init()
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_drive_activity_v2_drive_activity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_apps_drive_activity_v2_drive_activity_service_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_activity_v2_drive_activity_service_proto_depIdxs,
	}.Build()
	File_google_apps_drive_activity_v2_drive_activity_service_proto = out.File
	file_google_apps_drive_activity_v2_drive_activity_service_proto_rawDesc = nil
	file_google_apps_drive_activity_v2_drive_activity_service_proto_goTypes = nil
	file_google_apps_drive_activity_v2_drive_activity_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DriveActivityServiceClient is the client API for DriveActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriveActivityServiceClient interface {
	// Query past activity in Google Drive.
	QueryDriveActivity(ctx context.Context, in *QueryDriveActivityRequest, opts ...grpc.CallOption) (*QueryDriveActivityResponse, error)
}

type driveActivityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriveActivityServiceClient(cc grpc.ClientConnInterface) DriveActivityServiceClient {
	return &driveActivityServiceClient{cc}
}

func (c *driveActivityServiceClient) QueryDriveActivity(ctx context.Context, in *QueryDriveActivityRequest, opts ...grpc.CallOption) (*QueryDriveActivityResponse, error) {
	out := new(QueryDriveActivityResponse)
	err := c.cc.Invoke(ctx, "/google.apps.drive.activity.v2.DriveActivityService/QueryDriveActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriveActivityServiceServer is the server API for DriveActivityService service.
type DriveActivityServiceServer interface {
	// Query past activity in Google Drive.
	QueryDriveActivity(context.Context, *QueryDriveActivityRequest) (*QueryDriveActivityResponse, error)
}

// UnimplementedDriveActivityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDriveActivityServiceServer struct {
}

func (*UnimplementedDriveActivityServiceServer) QueryDriveActivity(context.Context, *QueryDriveActivityRequest) (*QueryDriveActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDriveActivity not implemented")
}

func RegisterDriveActivityServiceServer(s *grpc.Server, srv DriveActivityServiceServer) {
	s.RegisterService(&_DriveActivityService_serviceDesc, srv)
}

func _DriveActivityService_QueryDriveActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDriveActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriveActivityServiceServer).QueryDriveActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.apps.drive.activity.v2.DriveActivityService/QueryDriveActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriveActivityServiceServer).QueryDriveActivity(ctx, req.(*QueryDriveActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DriveActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.apps.drive.activity.v2.DriveActivityService",
	HandlerType: (*DriveActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryDriveActivity",
			Handler:    _DriveActivityService_QueryDriveActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/apps/drive/activity/v2/drive_activity_service.proto",
}
