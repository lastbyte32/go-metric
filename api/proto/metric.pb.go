// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/proto/metric.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Types int32

const (
	Types_UNKNOWN Types = 0
	Types_COUNTER Types = 1
	Types_GAUGE   Types = 2
)

// Enum value maps for Types.
var (
	Types_name = map[int32]string{
		0: "UNKNOWN",
		1: "COUNTER",
		2: "GAUGE",
	}
	Types_value = map[string]int32{
		"UNKNOWN": 0,
		"COUNTER": 1,
		"GAUGE":   2,
	}
)

func (x Types) Enum() *Types {
	p := new(Types)
	*p = x
	return p
}

func (x Types) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Types) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_metric_proto_enumTypes[0].Descriptor()
}

func (Types) Type() protoreflect.EnumType {
	return &file_api_proto_metric_proto_enumTypes[0]
}

func (x Types) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Types.Descriptor instead.
func (Types) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_metric_proto_rawDescGZIP(), []int{0}
}

type CounterMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Delta int64  `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *CounterMetric) Reset() {
	*x = CounterMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterMetric) ProtoMessage() {}

func (x *CounterMetric) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterMetric.ProtoReflect.Descriptor instead.
func (*CounterMetric) Descriptor() ([]byte, []int) {
	return file_api_proto_metric_proto_rawDescGZIP(), []int{0}
}

func (x *CounterMetric) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CounterMetric) GetDelta() int64 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type GaugeMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GaugeMetric) Reset() {
	*x = GaugeMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeMetric) ProtoMessage() {}

func (x *GaugeMetric) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeMetric.ProtoReflect.Descriptor instead.
func (*GaugeMetric) Descriptor() ([]byte, []int) {
	return file_api_proto_metric_proto_rawDescGZIP(), []int{1}
}

func (x *GaugeMetric) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GaugeMetric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Types `protobuf:"varint,1,opt,name=type,proto3,enum=metrics.Types" json:"type,omitempty"`
	// Types that are assignable to Metric:
	//
	//	*Metric_Counter
	//	*Metric_Gauge
	Metric isMetric_Metric `protobuf_oneof:"metric"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_api_proto_metric_proto_rawDescGZIP(), []int{2}
}

func (x *Metric) GetType() Types {
	if x != nil {
		return x.Type
	}
	return Types_UNKNOWN
}

func (m *Metric) GetMetric() isMetric_Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (x *Metric) GetCounter() *CounterMetric {
	if x, ok := x.GetMetric().(*Metric_Counter); ok {
		return x.Counter
	}
	return nil
}

func (x *Metric) GetGauge() *GaugeMetric {
	if x, ok := x.GetMetric().(*Metric_Gauge); ok {
		return x.Gauge
	}
	return nil
}

type isMetric_Metric interface {
	isMetric_Metric()
}

type Metric_Counter struct {
	Counter *CounterMetric `protobuf:"bytes,2,opt,name=counter,proto3,oneof"`
}

type Metric_Gauge struct {
	Gauge *GaugeMetric `protobuf:"bytes,3,opt,name=gauge,proto3,oneof"`
}

func (*Metric_Counter) isMetric_Metric() {}

func (*Metric_Gauge) isMetric_Metric() {}

var File_api_proto_metric_proto protoreflect.FileDescriptor

var file_api_proto_metric_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0b, 0x47, 0x61, 0x75, 0x67, 0x65, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x06, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x48, 0x00, 0x52, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2a, 0x2c, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41, 0x55, 0x47,
	0x45, 0x10, 0x02, 0x32, 0x3e, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x33,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x28, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x61, 0x73, 0x74, 0x62, 0x79, 0x74, 0x65, 0x33, 0x32, 0x2f, 0x67, 0x6f, 0x2d,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_metric_proto_rawDescOnce sync.Once
	file_api_proto_metric_proto_rawDescData = file_api_proto_metric_proto_rawDesc
)

func file_api_proto_metric_proto_rawDescGZIP() []byte {
	file_api_proto_metric_proto_rawDescOnce.Do(func() {
		file_api_proto_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_metric_proto_rawDescData)
	})
	return file_api_proto_metric_proto_rawDescData
}

var file_api_proto_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_metric_proto_goTypes = []interface{}{
	(Types)(0),            // 0: metrics.Types
	(*CounterMetric)(nil), // 1: metrics.CounterMetric
	(*GaugeMetric)(nil),   // 2: metrics.GaugeMetric
	(*Metric)(nil),        // 3: metrics.Metric
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_api_proto_metric_proto_depIdxs = []int32{
	0, // 0: metrics.Metric.type:type_name -> metrics.Types
	1, // 1: metrics.Metric.counter:type_name -> metrics.CounterMetric
	2, // 2: metrics.Metric.gauge:type_name -> metrics.GaugeMetric
	3, // 3: metrics.Metrics.Update:input_type -> metrics.Metric
	4, // 4: metrics.Metrics.Update:output_type -> google.protobuf.Empty
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_metric_proto_init() }
func file_api_proto_metric_proto_init() {
	if File_api_proto_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_proto_metric_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Metric_Counter)(nil),
		(*Metric_Gauge)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_metric_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_metric_proto_goTypes,
		DependencyIndexes: file_api_proto_metric_proto_depIdxs,
		EnumInfos:         file_api_proto_metric_proto_enumTypes,
		MessageInfos:      file_api_proto_metric_proto_msgTypes,
	}.Build()
	File_api_proto_metric_proto = out.File
	file_api_proto_metric_proto_rawDesc = nil
	file_api_proto_metric_proto_goTypes = nil
	file_api_proto_metric_proto_depIdxs = nil
}