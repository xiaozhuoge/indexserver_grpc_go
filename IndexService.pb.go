// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: IndexService.proto

package indexserver_grpc_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_IndexService_proto protoreflect.FileDescriptor

var file_IndexService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x73, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbd, 0x06, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a, 0x21,
	0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6e, 0x64, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x27, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x2a, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6e, 0x64,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x71, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x46,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_IndexService_proto_goTypes = []interface{}{
	(*AddEventLogRequest)(nil),                       // 0: indexserver_grpc.AddEventLogRequest
	(*SaveOrUpdateUserOnLineAndLocationRequest)(nil), // 1: indexserver_grpc.SaveOrUpdateUserOnLineAndLocationRequest
	(*DeleteUserRequest)(nil),                        // 2: indexserver_grpc.DeleteUserRequest
	(*AddUserFragmentRequest)(nil),                   // 3: indexserver_grpc.AddUserFragmentRequest
	(*DeleteFragmentRequest)(nil),                    // 4: indexserver_grpc.DeleteFragmentRequest
	(*QueryNearbyPersonRequest)(nil),                 // 5: indexserver_grpc.QueryNearbyPersonRequest
	(*QueryNearbyFragmentRequest)(nil),               // 6: indexserver_grpc.QueryNearbyFragmentRequest
	(*QueryEventLogRequest)(nil),                     // 7: indexserver_grpc.QueryEventLogRequest
	(*OperResponse)(nil),                             // 8: indexserver_grpc.OperResponse
	(*SearchUserOnLineAndLocationResponse)(nil),      // 9: indexserver_grpc.SearchUserOnLineAndLocationResponse
	(*SearchUserFragmentResponse)(nil),               // 10: indexserver_grpc.SearchUserFragmentResponse
	(*SearchEventLogResponse)(nil),                   // 11: indexserver_grpc.SearchEventLogResponse
}
var file_IndexService_proto_depIdxs = []int32{
	0,  // 0: indexserver_grpc.IndexService.AddEventLog:input_type -> indexserver_grpc.AddEventLogRequest
	1,  // 1: indexserver_grpc.IndexService.SaveOrUpdateUserOnLineAndLocation:input_type -> indexserver_grpc.SaveOrUpdateUserOnLineAndLocationRequest
	2,  // 2: indexserver_grpc.IndexService.DeleteUser:input_type -> indexserver_grpc.DeleteUserRequest
	3,  // 3: indexserver_grpc.IndexService.AddUserFragment:input_type -> indexserver_grpc.AddUserFragmentRequest
	4,  // 4: indexserver_grpc.IndexService.DeleteFragment:input_type -> indexserver_grpc.DeleteFragmentRequest
	5,  // 5: indexserver_grpc.IndexService.QueryNearbyPerson:input_type -> indexserver_grpc.QueryNearbyPersonRequest
	6,  // 6: indexserver_grpc.IndexService.QueryNearbyFragment:input_type -> indexserver_grpc.QueryNearbyFragmentRequest
	7,  // 7: indexserver_grpc.IndexService.QueryEventLog:input_type -> indexserver_grpc.QueryEventLogRequest
	8,  // 8: indexserver_grpc.IndexService.AddEventLog:output_type -> indexserver_grpc.OperResponse
	8,  // 9: indexserver_grpc.IndexService.SaveOrUpdateUserOnLineAndLocation:output_type -> indexserver_grpc.OperResponse
	8,  // 10: indexserver_grpc.IndexService.DeleteUser:output_type -> indexserver_grpc.OperResponse
	8,  // 11: indexserver_grpc.IndexService.AddUserFragment:output_type -> indexserver_grpc.OperResponse
	8,  // 12: indexserver_grpc.IndexService.DeleteFragment:output_type -> indexserver_grpc.OperResponse
	9,  // 13: indexserver_grpc.IndexService.QueryNearbyPerson:output_type -> indexserver_grpc.SearchUserOnLineAndLocationResponse
	10, // 14: indexserver_grpc.IndexService.QueryNearbyFragment:output_type -> indexserver_grpc.SearchUserFragmentResponse
	11, // 15: indexserver_grpc.IndexService.QueryEventLog:output_type -> indexserver_grpc.SearchEventLogResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_IndexService_proto_init() }
func file_IndexService_proto_init() {
	if File_IndexService_proto != nil {
		return
	}
	file_IndexMsg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_IndexService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_IndexService_proto_goTypes,
		DependencyIndexes: file_IndexService_proto_depIdxs,
	}.Build()
	File_IndexService_proto = out.File
	file_IndexService_proto_rawDesc = nil
	file_IndexService_proto_goTypes = nil
	file_IndexService_proto_depIdxs = nil
}
