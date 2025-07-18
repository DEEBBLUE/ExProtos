// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: DataBase.proto

package Database

import (
	Req "github.com/DEEBBLUE/ExProtos/api/Req"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_DataBase_proto protoreflect.FileDescriptor

const file_DataBase_proto_rawDesc = "" +
	"\n" +
	"\x0eDataBase.proto\x12\bdatabase\x1a\tReq.proto2\xf9\x04\n" +
	"\bDatabase\x121\n" +
	"\n" +
	"CreateUser\x12\x12.req.CreateUserReq\x1a\x0f.req.DefaultRes\x124\n" +
	"\n" +
	"RepeatUser\x12\x12.req.RepeatUserReq\x1a\x12.req.RepeatUserRes\x129\n" +
	"\x0eChangeRoleUser\x12\x16.req.ChangeRoleUserReq\x1a\x0f.req.DefaultRes\x12I\n" +
	"\x16ChangeVerifeStatusUser\x12\x1e.req.ChangeVerifeStatusUserReq\x1a\x0f.req.DefaultRes\x12?\n" +
	"\x11ChangeBalanceUser\x12\x19.req.ChangeBalanceUserReq\x1a\x0f.req.DefaultRes\x12@\n" +
	"\x0eCreateExchange\x12\x16.req.CreateExchangeReq\x1a\x16.req.CreateExchangeRes\x12@\n" +
	"\x0eRepeatExchange\x12\x16.req.RepeatExchangeReq\x1a\x16.req.RepeatExchangeRes\x12I\n" +
	"\x16InitBankDetailExchange\x12\x1e.req.InitBankDetailExchangeReq\x1a\x0f.req.DefaultRes\x12+\n" +
	"\aConfirm\x12\x0f.req.ConfirmReq\x1a\x0f.req.DefaultRes\x12A\n" +
	"\x11RepeatUserHistory\x12\x16.req.RepeatUserListReq\x1a\x14.req.RepeatListExResB+Z)github.com/DEEBBLUE/ExProtos/api/Databaseb\x06proto3"

var file_DataBase_proto_goTypes = []any{
	(*Req.CreateUserReq)(nil),             // 0: req.CreateUserReq
	(*Req.RepeatUserReq)(nil),             // 1: req.RepeatUserReq
	(*Req.ChangeRoleUserReq)(nil),         // 2: req.ChangeRoleUserReq
	(*Req.ChangeVerifeStatusUserReq)(nil), // 3: req.ChangeVerifeStatusUserReq
	(*Req.ChangeBalanceUserReq)(nil),      // 4: req.ChangeBalanceUserReq
	(*Req.CreateExchangeReq)(nil),         // 5: req.CreateExchangeReq
	(*Req.RepeatExchangeReq)(nil),         // 6: req.RepeatExchangeReq
	(*Req.InitBankDetailExchangeReq)(nil), // 7: req.InitBankDetailExchangeReq
	(*Req.ConfirmReq)(nil),                // 8: req.ConfirmReq
	(*Req.RepeatUserListReq)(nil),         // 9: req.RepeatUserListReq
	(*Req.DefaultRes)(nil),                // 10: req.DefaultRes
	(*Req.RepeatUserRes)(nil),             // 11: req.RepeatUserRes
	(*Req.CreateExchangeRes)(nil),         // 12: req.CreateExchangeRes
	(*Req.RepeatExchangeRes)(nil),         // 13: req.RepeatExchangeRes
	(*Req.RepeatListExRes)(nil),           // 14: req.RepeatListExRes
}
var file_DataBase_proto_depIdxs = []int32{
	0,  // 0: database.Database.CreateUser:input_type -> req.CreateUserReq
	1,  // 1: database.Database.RepeatUser:input_type -> req.RepeatUserReq
	2,  // 2: database.Database.ChangeRoleUser:input_type -> req.ChangeRoleUserReq
	3,  // 3: database.Database.ChangeVerifeStatusUser:input_type -> req.ChangeVerifeStatusUserReq
	4,  // 4: database.Database.ChangeBalanceUser:input_type -> req.ChangeBalanceUserReq
	5,  // 5: database.Database.CreateExchange:input_type -> req.CreateExchangeReq
	6,  // 6: database.Database.RepeatExchange:input_type -> req.RepeatExchangeReq
	7,  // 7: database.Database.InitBankDetailExchange:input_type -> req.InitBankDetailExchangeReq
	8,  // 8: database.Database.Confirm:input_type -> req.ConfirmReq
	9,  // 9: database.Database.RepeatUserHistory:input_type -> req.RepeatUserListReq
	10, // 10: database.Database.CreateUser:output_type -> req.DefaultRes
	11, // 11: database.Database.RepeatUser:output_type -> req.RepeatUserRes
	10, // 12: database.Database.ChangeRoleUser:output_type -> req.DefaultRes
	10, // 13: database.Database.ChangeVerifeStatusUser:output_type -> req.DefaultRes
	10, // 14: database.Database.ChangeBalanceUser:output_type -> req.DefaultRes
	12, // 15: database.Database.CreateExchange:output_type -> req.CreateExchangeRes
	13, // 16: database.Database.RepeatExchange:output_type -> req.RepeatExchangeRes
	10, // 17: database.Database.InitBankDetailExchange:output_type -> req.DefaultRes
	10, // 18: database.Database.Confirm:output_type -> req.DefaultRes
	14, // 19: database.Database.RepeatUserHistory:output_type -> req.RepeatListExRes
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_DataBase_proto_init() }
func file_DataBase_proto_init() {
	if File_DataBase_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_DataBase_proto_rawDesc), len(file_DataBase_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_DataBase_proto_goTypes,
		DependencyIndexes: file_DataBase_proto_depIdxs,
	}.Build()
	File_DataBase_proto = out.File
	file_DataBase_proto_goTypes = nil
	file_DataBase_proto_depIdxs = nil
}
