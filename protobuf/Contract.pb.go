// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: Contract.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Contract struct {
	state                        protoimpl.MessageState `protogen:"open.v1"`
	ConId                        *int32                 `protobuf:"varint,1,opt,name=conId,proto3,oneof" json:"conId,omitempty"`
	Symbol                       *string                `protobuf:"bytes,2,opt,name=symbol,proto3,oneof" json:"symbol,omitempty"`
	SecType                      *string                `protobuf:"bytes,3,opt,name=secType,proto3,oneof" json:"secType,omitempty"`
	LastTradeDateOrContractMonth *string                `protobuf:"bytes,4,opt,name=lastTradeDateOrContractMonth,proto3,oneof" json:"lastTradeDateOrContractMonth,omitempty"`
	Strike                       *float64               `protobuf:"fixed64,5,opt,name=strike,proto3,oneof" json:"strike,omitempty"`
	Right                        *string                `protobuf:"bytes,6,opt,name=right,proto3,oneof" json:"right,omitempty"`
	Multiplier                   *float64               `protobuf:"fixed64,7,opt,name=multiplier,proto3,oneof" json:"multiplier,omitempty"`
	Exchange                     *string                `protobuf:"bytes,8,opt,name=exchange,proto3,oneof" json:"exchange,omitempty"`
	PrimaryExch                  *string                `protobuf:"bytes,9,opt,name=primaryExch,proto3,oneof" json:"primaryExch,omitempty"`
	Currency                     *string                `protobuf:"bytes,10,opt,name=currency,proto3,oneof" json:"currency,omitempty"`
	LocalSymbol                  *string                `protobuf:"bytes,11,opt,name=localSymbol,proto3,oneof" json:"localSymbol,omitempty"`
	TradingClass                 *string                `protobuf:"bytes,12,opt,name=tradingClass,proto3,oneof" json:"tradingClass,omitempty"`
	SecIdType                    *string                `protobuf:"bytes,13,opt,name=secIdType,proto3,oneof" json:"secIdType,omitempty"`
	SecId                        *string                `protobuf:"bytes,14,opt,name=secId,proto3,oneof" json:"secId,omitempty"`
	Description                  *string                `protobuf:"bytes,15,opt,name=description,proto3,oneof" json:"description,omitempty"`
	IssuerId                     *string                `protobuf:"bytes,16,opt,name=issuerId,proto3,oneof" json:"issuerId,omitempty"`
	DeltaNeutralContract         *DeltaNeutralContract  `protobuf:"bytes,17,opt,name=deltaNeutralContract,proto3,oneof" json:"deltaNeutralContract,omitempty"`
	IncludeExpired               *bool                  `protobuf:"varint,18,opt,name=includeExpired,proto3,oneof" json:"includeExpired,omitempty"`
	ComboLegsDescrip             *string                `protobuf:"bytes,19,opt,name=comboLegsDescrip,proto3,oneof" json:"comboLegsDescrip,omitempty"`
	ComboLegs                    []*ComboLeg            `protobuf:"bytes,20,rep,name=comboLegs,proto3" json:"comboLegs,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *Contract) Reset() {
	*x = Contract{}
	mi := &file_Contract_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_Contract_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_Contract_proto_rawDescGZIP(), []int{0}
}

func (x *Contract) GetConId() int32 {
	if x != nil && x.ConId != nil {
		return *x.ConId
	}
	return 0
}

func (x *Contract) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *Contract) GetSecType() string {
	if x != nil && x.SecType != nil {
		return *x.SecType
	}
	return ""
}

func (x *Contract) GetLastTradeDateOrContractMonth() string {
	if x != nil && x.LastTradeDateOrContractMonth != nil {
		return *x.LastTradeDateOrContractMonth
	}
	return ""
}

func (x *Contract) GetStrike() float64 {
	if x != nil && x.Strike != nil {
		return *x.Strike
	}
	return 0
}

func (x *Contract) GetRight() string {
	if x != nil && x.Right != nil {
		return *x.Right
	}
	return ""
}

func (x *Contract) GetMultiplier() float64 {
	if x != nil && x.Multiplier != nil {
		return *x.Multiplier
	}
	return 0
}

func (x *Contract) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *Contract) GetPrimaryExch() string {
	if x != nil && x.PrimaryExch != nil {
		return *x.PrimaryExch
	}
	return ""
}

func (x *Contract) GetCurrency() string {
	if x != nil && x.Currency != nil {
		return *x.Currency
	}
	return ""
}

func (x *Contract) GetLocalSymbol() string {
	if x != nil && x.LocalSymbol != nil {
		return *x.LocalSymbol
	}
	return ""
}

func (x *Contract) GetTradingClass() string {
	if x != nil && x.TradingClass != nil {
		return *x.TradingClass
	}
	return ""
}

func (x *Contract) GetSecIdType() string {
	if x != nil && x.SecIdType != nil {
		return *x.SecIdType
	}
	return ""
}

func (x *Contract) GetSecId() string {
	if x != nil && x.SecId != nil {
		return *x.SecId
	}
	return ""
}

func (x *Contract) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Contract) GetIssuerId() string {
	if x != nil && x.IssuerId != nil {
		return *x.IssuerId
	}
	return ""
}

func (x *Contract) GetDeltaNeutralContract() *DeltaNeutralContract {
	if x != nil {
		return x.DeltaNeutralContract
	}
	return nil
}

func (x *Contract) GetIncludeExpired() bool {
	if x != nil && x.IncludeExpired != nil {
		return *x.IncludeExpired
	}
	return false
}

func (x *Contract) GetComboLegsDescrip() string {
	if x != nil && x.ComboLegsDescrip != nil {
		return *x.ComboLegsDescrip
	}
	return ""
}

func (x *Contract) GetComboLegs() []*ComboLeg {
	if x != nil {
		return x.ComboLegs
	}
	return nil
}

var File_Contract_proto protoreflect.FileDescriptor

var file_Contract_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x0e, 0x43, 0x6f, 0x6d, 0x62,
	0x6f, 0x4c, 0x65, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x08, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x1c, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x1c, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x04, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x05, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x06, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x07, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x45, 0x78, 0x63, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x45, 0x78, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a,
	0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x27, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x65, 0x63,
	0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0c, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x73, 0x65, 0x63, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0d, 0x52, 0x05, 0x73,
	0x65, 0x63, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0e, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x0f, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x57, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4e, 0x65,
	0x75, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x10, 0x52,
	0x14, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x11, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x4c, 0x65,
	0x67, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x12, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x4c, 0x65, 0x67, 0x73, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x4c,
	0x65, 0x67, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x6f, 0x4c, 0x65, 0x67, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x62, 0x6f, 0x4c, 0x65, 0x67, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x6e,
	0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x73, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6b, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x45, 0x78, 0x63, 0x68, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65,
	0x63, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x65, 0x63, 0x49,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x42, 0x17,
	0x0a, 0x15, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63,
	0x6f, 0x6d, 0x62, 0x6f, 0x4c, 0x65, 0x67, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_Contract_proto_rawDescOnce sync.Once
	file_Contract_proto_rawDescData []byte
)

func file_Contract_proto_rawDescGZIP() []byte {
	file_Contract_proto_rawDescOnce.Do(func() {
		file_Contract_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_Contract_proto_rawDesc), len(file_Contract_proto_rawDesc)))
	})
	return file_Contract_proto_rawDescData
}

var file_Contract_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Contract_proto_goTypes = []any{
	(*Contract)(nil),             // 0: protobuf.Contract
	(*DeltaNeutralContract)(nil), // 1: protobuf.DeltaNeutralContract
	(*ComboLeg)(nil),             // 2: protobuf.ComboLeg
}
var file_Contract_proto_depIdxs = []int32{
	1, // 0: protobuf.Contract.deltaNeutralContract:type_name -> protobuf.DeltaNeutralContract
	2, // 1: protobuf.Contract.comboLegs:type_name -> protobuf.ComboLeg
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Contract_proto_init() }
func file_Contract_proto_init() {
	if File_Contract_proto != nil {
		return
	}
	file_ComboLeg_proto_init()
	file_DeltaNeutralContract_proto_init()
	file_Contract_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_Contract_proto_rawDesc), len(file_Contract_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Contract_proto_goTypes,
		DependencyIndexes: file_Contract_proto_depIdxs,
		MessageInfos:      file_Contract_proto_msgTypes,
	}.Build()
	File_Contract_proto = out.File
	file_Contract_proto_goTypes = nil
	file_Contract_proto_depIdxs = nil
}
