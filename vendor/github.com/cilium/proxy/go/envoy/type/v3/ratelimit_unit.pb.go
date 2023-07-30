// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.12
// source: envoy/type/v3/ratelimit_unit.proto

package typev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

// Identifies the unit of of time for rate limit.
type RateLimitUnit int32

const (
	// The time unit is not known.
	RateLimitUnit_UNKNOWN RateLimitUnit = 0
	// The time unit representing a second.
	RateLimitUnit_SECOND RateLimitUnit = 1
	// The time unit representing a minute.
	RateLimitUnit_MINUTE RateLimitUnit = 2
	// The time unit representing an hour.
	RateLimitUnit_HOUR RateLimitUnit = 3
	// The time unit representing a day.
	RateLimitUnit_DAY RateLimitUnit = 4
	// The time unit representing a month.
	RateLimitUnit_MONTH RateLimitUnit = 5
	// The time unit representing a year.
	RateLimitUnit_YEAR RateLimitUnit = 6
)

// Enum value maps for RateLimitUnit.
var (
	RateLimitUnit_name = map[int32]string{
		0: "UNKNOWN",
		1: "SECOND",
		2: "MINUTE",
		3: "HOUR",
		4: "DAY",
		5: "MONTH",
		6: "YEAR",
	}
	RateLimitUnit_value = map[string]int32{
		"UNKNOWN": 0,
		"SECOND":  1,
		"MINUTE":  2,
		"HOUR":    3,
		"DAY":     4,
		"MONTH":   5,
		"YEAR":    6,
	}
)

func (x RateLimitUnit) Enum() *RateLimitUnit {
	p := new(RateLimitUnit)
	*p = x
	return p
}

func (x RateLimitUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_type_v3_ratelimit_unit_proto_enumTypes[0].Descriptor()
}

func (RateLimitUnit) Type() protoreflect.EnumType {
	return &file_envoy_type_v3_ratelimit_unit_proto_enumTypes[0]
}

func (x RateLimitUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitUnit.Descriptor instead.
func (RateLimitUnit) EnumDescriptor() ([]byte, []int) {
	return file_envoy_type_v3_ratelimit_unit_proto_rawDescGZIP(), []int{0}
}

var File_envoy_type_v3_ratelimit_unit_proto protoreflect.FileDescriptor

var file_envoy_type_v3_ratelimit_unit_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x5c, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x4f, 0x4e, 0x54, 0x48, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x45, 0x41, 0x52, 0x10, 0x06,
	0x42, 0x78, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x42,
	0x12, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x79, 0x70, 0x65,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_type_v3_ratelimit_unit_proto_rawDescOnce sync.Once
	file_envoy_type_v3_ratelimit_unit_proto_rawDescData = file_envoy_type_v3_ratelimit_unit_proto_rawDesc
)

func file_envoy_type_v3_ratelimit_unit_proto_rawDescGZIP() []byte {
	file_envoy_type_v3_ratelimit_unit_proto_rawDescOnce.Do(func() {
		file_envoy_type_v3_ratelimit_unit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_v3_ratelimit_unit_proto_rawDescData)
	})
	return file_envoy_type_v3_ratelimit_unit_proto_rawDescData
}

var file_envoy_type_v3_ratelimit_unit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_type_v3_ratelimit_unit_proto_goTypes = []interface{}{
	(RateLimitUnit)(0), // 0: envoy.type.v3.RateLimitUnit
}
var file_envoy_type_v3_ratelimit_unit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_type_v3_ratelimit_unit_proto_init() }
func file_envoy_type_v3_ratelimit_unit_proto_init() {
	if File_envoy_type_v3_ratelimit_unit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_v3_ratelimit_unit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_v3_ratelimit_unit_proto_goTypes,
		DependencyIndexes: file_envoy_type_v3_ratelimit_unit_proto_depIdxs,
		EnumInfos:         file_envoy_type_v3_ratelimit_unit_proto_enumTypes,
	}.Build()
	File_envoy_type_v3_ratelimit_unit_proto = out.File
	file_envoy_type_v3_ratelimit_unit_proto_rawDesc = nil
	file_envoy_type_v3_ratelimit_unit_proto_goTypes = nil
	file_envoy_type_v3_ratelimit_unit_proto_depIdxs = nil
}