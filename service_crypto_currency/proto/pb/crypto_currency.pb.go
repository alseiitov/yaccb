// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: crypto_currency.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CryptoCurrencies []*CryptoCurrency `protobuf:"bytes,1,rep,name=CryptoCurrencies,proto3" json:"CryptoCurrencies,omitempty"`
}

func (x *GetCurrenciesResponse) Reset() {
	*x = GetCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesResponse) ProtoMessage() {}

func (x *GetCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_crypto_currency_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrenciesResponse) GetCryptoCurrencies() []*CryptoCurrency {
	if x != nil {
		return x.CryptoCurrencies
	}
	return nil
}

type CryptoCurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CryptoCurrency) Reset() {
	*x = CryptoCurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoCurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoCurrency) ProtoMessage() {}

func (x *CryptoCurrency) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoCurrency.ProtoReflect.Descriptor instead.
func (*CryptoCurrency) Descriptor() ([]byte, []int) {
	return file_crypto_currency_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoCurrency) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CryptoCurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *CryptoCurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_crypto_currency_proto protoreflect.FileDescriptor

var file_crypto_currency_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x10, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x0e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x80, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_currency_proto_rawDescOnce sync.Once
	file_crypto_currency_proto_rawDescData = file_crypto_currency_proto_rawDesc
)

func file_crypto_currency_proto_rawDescGZIP() []byte {
	file_crypto_currency_proto_rawDescOnce.Do(func() {
		file_crypto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_currency_proto_rawDescData)
	})
	return file_crypto_currency_proto_rawDescData
}

var file_crypto_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crypto_currency_proto_goTypes = []interface{}{
	(*GetCurrenciesResponse)(nil), // 0: proto.GetCurrenciesResponse
	(*CryptoCurrency)(nil),        // 1: proto.CryptoCurrency
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_crypto_currency_proto_depIdxs = []int32{
	1, // 0: proto.GetCurrenciesResponse.CryptoCurrencies:type_name -> proto.CryptoCurrency
	2, // 1: proto.CryptoCurrencyService.GetCryptoCurrencies:input_type -> google.protobuf.Empty
	0, // 2: proto.CryptoCurrencyService.GetCryptoCurrencies:output_type -> proto.GetCurrenciesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_crypto_currency_proto_init() }
func file_crypto_currency_proto_init() {
	if File_crypto_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesResponse); i {
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
		file_crypto_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoCurrency); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crypto_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_currency_proto_goTypes,
		DependencyIndexes: file_crypto_currency_proto_depIdxs,
		MessageInfos:      file_crypto_currency_proto_msgTypes,
	}.Build()
	File_crypto_currency_proto = out.File
	file_crypto_currency_proto_rawDesc = nil
	file_crypto_currency_proto_goTypes = nil
	file_crypto_currency_proto_depIdxs = nil
}
