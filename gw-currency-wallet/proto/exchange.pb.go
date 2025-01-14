// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: exchange.proto

package proto

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

type ConvertCurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	Amount        float32                `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConvertCurrencyRequest) Reset() {
	*x = ConvertCurrencyRequest{}
	mi := &file_exchange_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertCurrencyRequest) ProtoMessage() {}

func (x *ConvertCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertCurrencyRequest.ProtoReflect.Descriptor instead.
func (*ConvertCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertCurrencyRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ConvertCurrencyRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ConvertCurrencyRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Запрос для получения курса обмена для конкретной валюты
type CurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrencyRequest) Reset() {
	*x = CurrencyRequest{}
	mi := &file_exchange_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRequest) ProtoMessage() {}

func (x *CurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRequest.ProtoReflect.Descriptor instead.
func (*CurrencyRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *CurrencyRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

type ConvertCurrencyResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ConvertedAmount float32                `protobuf:"fixed32,1,opt,name=converted_amount,json=convertedAmount,proto3" json:"converted_amount,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ConvertCurrencyResponse) Reset() {
	*x = ConvertCurrencyResponse{}
	mi := &file_exchange_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertCurrencyResponse) ProtoMessage() {}

func (x *ConvertCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertCurrencyResponse.ProtoReflect.Descriptor instead.
func (*ConvertCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *ConvertCurrencyResponse) GetConvertedAmount() float32 {
	if x != nil {
		return x.ConvertedAmount
	}
	return 0
}

// Ответ с курсом обмена для конкретной валюты
type ExchangeRateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	Rate          float32                `protobuf:"fixed32,3,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeRateResponse) Reset() {
	*x = ExchangeRateResponse{}
	mi := &file_exchange_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateResponse) ProtoMessage() {}

func (x *ExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*ExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeRateResponse) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// Ответ с курсами обмена всех валют
type ExchangeRatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rates         map[string]float32     `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"` // ключ: валюта, значение: курс
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeRatesResponse) Reset() {
	*x = ExchangeRatesResponse{}
	mi := &file_exchange_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRatesResponse) ProtoMessage() {}

func (x *ExchangeRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRatesResponse.ProtoReflect.Descriptor instead.
func (*ExchangeRatesResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *ExchangeRatesResponse) GetRates() map[string]float32 {
	if x != nil {
		return x.Rates
	}
	return nil
}

// Пустое сообщение
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_exchange_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{5}
}

var File_exchange_proto protoreflect.FileDescriptor

var file_exchange_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x76, 0x0a, 0x16, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f,
	0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x44, 0x0a, 0x17, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x70, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x88, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a,
	0x18, 0x67, 0x77, 0x2d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2d, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_exchange_proto_rawDescOnce sync.Once
	file_exchange_proto_rawDescData = file_exchange_proto_rawDesc
)

func file_exchange_proto_rawDescGZIP() []byte {
	file_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_proto_rawDescData)
	})
	return file_exchange_proto_rawDescData
}

var file_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_exchange_proto_goTypes = []any{
	(*ConvertCurrencyRequest)(nil),  // 0: exchange.ConvertCurrencyRequest
	(*CurrencyRequest)(nil),         // 1: exchange.CurrencyRequest
	(*ConvertCurrencyResponse)(nil), // 2: exchange.ConvertCurrencyResponse
	(*ExchangeRateResponse)(nil),    // 3: exchange.ExchangeRateResponse
	(*ExchangeRatesResponse)(nil),   // 4: exchange.ExchangeRatesResponse
	(*Empty)(nil),                   // 5: exchange.Empty
	nil,                             // 6: exchange.ExchangeRatesResponse.RatesEntry
}
var file_exchange_proto_depIdxs = []int32{
	6, // 0: exchange.ExchangeRatesResponse.rates:type_name -> exchange.ExchangeRatesResponse.RatesEntry
	5, // 1: exchange.ExchangeService.GetExchangeRates:input_type -> exchange.Empty
	1, // 2: exchange.ExchangeService.GetExchangeRateForCurrency:input_type -> exchange.CurrencyRequest
	0, // 3: exchange.ExchangeService.ConvertCurrency:input_type -> exchange.ConvertCurrencyRequest
	4, // 4: exchange.ExchangeService.GetExchangeRates:output_type -> exchange.ExchangeRatesResponse
	3, // 5: exchange.ExchangeService.GetExchangeRateForCurrency:output_type -> exchange.ExchangeRateResponse
	2, // 6: exchange.ExchangeService.ConvertCurrency:output_type -> exchange.ConvertCurrencyResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_exchange_proto_init() }
func file_exchange_proto_init() {
	if File_exchange_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_proto_depIdxs,
		MessageInfos:      file_exchange_proto_msgTypes,
	}.Build()
	File_exchange_proto = out.File
	file_exchange_proto_rawDesc = nil
	file_exchange_proto_goTypes = nil
	file_exchange_proto_depIdxs = nil
}
