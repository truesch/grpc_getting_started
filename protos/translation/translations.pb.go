// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: protos/translations.proto

package translation

import (
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

type Languages int32

const (
	Languages_NONE Languages = 0
	Languages_DE   Languages = 1
	Languages_ES   Languages = 2
	Languages_CZ   Languages = 3
	Languages_FA   Languages = 4
	Languages_ZH   Languages = 5
	Languages_EN   Languages = 6
)

// Enum value maps for Languages.
var (
	Languages_name = map[int32]string{
		0: "NONE",
		1: "DE",
		2: "ES",
		3: "CZ",
		4: "FA",
		5: "ZH",
		6: "EN",
	}
	Languages_value = map[string]int32{
		"NONE": 0,
		"DE":   1,
		"ES":   2,
		"CZ":   3,
		"FA":   4,
		"ZH":   5,
		"EN":   6,
	}
)

func (x Languages) Enum() *Languages {
	p := new(Languages)
	*p = x
	return p
}

func (x Languages) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Languages) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_translations_proto_enumTypes[0].Descriptor()
}

func (Languages) Type() protoreflect.EnumType {
	return &file_protos_translations_proto_enumTypes[0]
}

func (x Languages) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Languages.Descriptor instead.
func (Languages) EnumDescriptor() ([]byte, []int) {
	return file_protos_translations_proto_rawDescGZIP(), []int{0}
}

type Vendors int32

const (
	Vendors_GoogleTranslate Vendors = 0
	Vendors_DeepL           Vendors = 1
)

// Enum value maps for Vendors.
var (
	Vendors_name = map[int32]string{
		0: "GoogleTranslate",
		1: "DeepL",
	}
	Vendors_value = map[string]int32{
		"GoogleTranslate": 0,
		"DeepL":           1,
	}
)

func (x Vendors) Enum() *Vendors {
	p := new(Vendors)
	*p = x
	return p
}

func (x Vendors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vendors) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_translations_proto_enumTypes[1].Descriptor()
}

func (Vendors) Type() protoreflect.EnumType {
	return &file_protos_translations_proto_enumTypes[1]
}

func (x Vendors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vendors.Descriptor instead.
func (Vendors) EnumDescriptor() ([]byte, []int) {
	return file_protos_translations_proto_rawDescGZIP(), []int{1}
}

type TranslationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string    `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	SourceLang Languages `protobuf:"varint,2,opt,name=SourceLang,proto3,enum=Languages" json:"SourceLang,omitempty"`
	TargetLang Languages `protobuf:"varint,3,opt,name=TargetLang,proto3,enum=Languages" json:"TargetLang,omitempty"`
	Vendor     *Vendors  `protobuf:"varint,4,opt,name=Vendor,proto3,enum=Vendors,oneof" json:"Vendor,omitempty"`
}

func (x *TranslationInput) Reset() {
	*x = TranslationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_translations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationInput) ProtoMessage() {}

func (x *TranslationInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_translations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationInput.ProtoReflect.Descriptor instead.
func (*TranslationInput) Descriptor() ([]byte, []int) {
	return file_protos_translations_proto_rawDescGZIP(), []int{0}
}

func (x *TranslationInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TranslationInput) GetSourceLang() Languages {
	if x != nil {
		return x.SourceLang
	}
	return Languages_NONE
}

func (x *TranslationInput) GetTargetLang() Languages {
	if x != nil {
		return x.TargetLang
	}
	return Languages_NONE
}

func (x *TranslationInput) GetVendor() Vendors {
	if x != nil && x.Vendor != nil {
		return *x.Vendor
	}
	return Vendors_GoogleTranslate
}

type TranslationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text        string    `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	SourceLang  Languages `protobuf:"varint,2,opt,name=SourceLang,proto3,enum=Languages" json:"SourceLang,omitempty"`
	TargetLang  Languages `protobuf:"varint,3,opt,name=TargetLang,proto3,enum=Languages" json:"TargetLang,omitempty"`
	BilledChars int32     `protobuf:"varint,4,opt,name=BilledChars,proto3" json:"BilledChars,omitempty"`
}

func (x *TranslationOutput) Reset() {
	*x = TranslationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_translations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationOutput) ProtoMessage() {}

func (x *TranslationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_translations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationOutput.ProtoReflect.Descriptor instead.
func (*TranslationOutput) Descriptor() ([]byte, []int) {
	return file_protos_translations_proto_rawDescGZIP(), []int{1}
}

func (x *TranslationOutput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TranslationOutput) GetSourceLang() Languages {
	if x != nil {
		return x.SourceLang
	}
	return Languages_NONE
}

func (x *TranslationOutput) GetTargetLang() Languages {
	if x != nil {
		return x.TargetLang
	}
	return Languages_NONE
}

func (x *TranslationOutput) GetBilledChars() int32 {
	if x != nil {
		return x.BilledChars
	}
	return 0
}

var File_protos_translations_proto protoreflect.FileDescriptor

var file_protos_translations_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67,
	0x12, 0x2a, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x06,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x22, 0xa1,
	0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67,
	0x12, 0x20, 0x0a, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x68, 0x61, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x68, 0x61,
	0x72, 0x73, 0x2a, 0x45, 0x0a, 0x09, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x45, 0x10,
	0x01, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x53, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x5a, 0x10,
	0x03, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x41, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x5a, 0x48, 0x10,
	0x05, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x4e, 0x10, 0x06, 0x2a, 0x29, 0x0a, 0x07, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x65, 0x65,
	0x70, 0x4c, 0x10, 0x01, 0x32, 0x41, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_translations_proto_rawDescOnce sync.Once
	file_protos_translations_proto_rawDescData = file_protos_translations_proto_rawDesc
)

func file_protos_translations_proto_rawDescGZIP() []byte {
	file_protos_translations_proto_rawDescOnce.Do(func() {
		file_protos_translations_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_translations_proto_rawDescData)
	})
	return file_protos_translations_proto_rawDescData
}

var file_protos_translations_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_translations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_translations_proto_goTypes = []interface{}{
	(Languages)(0),            // 0: Languages
	(Vendors)(0),              // 1: Vendors
	(*TranslationInput)(nil),  // 2: TranslationInput
	(*TranslationOutput)(nil), // 3: TranslationOutput
}
var file_protos_translations_proto_depIdxs = []int32{
	0, // 0: TranslationInput.SourceLang:type_name -> Languages
	0, // 1: TranslationInput.TargetLang:type_name -> Languages
	1, // 2: TranslationInput.Vendor:type_name -> Vendors
	0, // 3: TranslationOutput.SourceLang:type_name -> Languages
	0, // 4: TranslationOutput.TargetLang:type_name -> Languages
	2, // 5: Translation.Translate:input_type -> TranslationInput
	3, // 6: Translation.Translate:output_type -> TranslationOutput
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_translations_proto_init() }
func file_protos_translations_proto_init() {
	if File_protos_translations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_translations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationInput); i {
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
		file_protos_translations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationOutput); i {
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
	file_protos_translations_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_translations_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_translations_proto_goTypes,
		DependencyIndexes: file_protos_translations_proto_depIdxs,
		EnumInfos:         file_protos_translations_proto_enumTypes,
		MessageInfos:      file_protos_translations_proto_msgTypes,
	}.Build()
	File_protos_translations_proto = out.File
	file_protos_translations_proto_rawDesc = nil
	file_protos_translations_proto_goTypes = nil
	file_protos_translations_proto_depIdxs = nil
}
