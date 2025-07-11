// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: vlan.proto

package pb

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

type CreateVlanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVlanRequest) Reset() {
	*x = CreateVlanRequest{}
	mi := &file_vlan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVlanRequest) ProtoMessage() {}

func (x *CreateVlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVlanRequest.ProtoReflect.Descriptor instead.
func (*CreateVlanRequest) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateVlanResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVlanResponse) Reset() {
	*x = CreateVlanResponse{}
	mi := &file_vlan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVlanResponse) ProtoMessage() {}

func (x *CreateVlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVlanResponse.ProtoReflect.Descriptor instead.
func (*CreateVlanResponse) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVlanResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetVlansRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVlansRequest) Reset() {
	*x = GetVlansRequest{}
	mi := &file_vlan_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVlansRequest) ProtoMessage() {}

func (x *GetVlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVlansRequest.ProtoReflect.Descriptor instead.
func (*GetVlansRequest) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{2}
}

func (x *GetVlansRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetVlansRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetVlansResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*VlanItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVlansResponse) Reset() {
	*x = GetVlansResponse{}
	mi := &file_vlan_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVlansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVlansResponse) ProtoMessage() {}

func (x *GetVlansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVlansResponse.ProtoReflect.Descriptor instead.
func (*GetVlansResponse) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{3}
}

func (x *GetVlansResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetVlansResponse) GetItems() []*VlanItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type VlanItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VlanItem) Reset() {
	*x = VlanItem{}
	mi := &file_vlan_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VlanItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VlanItem) ProtoMessage() {}

func (x *VlanItem) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VlanItem.ProtoReflect.Descriptor instead.
func (*VlanItem) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{4}
}

func (x *VlanItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VlanItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetVlanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVlanRequest) Reset() {
	*x = GetVlanRequest{}
	mi := &file_vlan_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVlanRequest) ProtoMessage() {}

func (x *GetVlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVlanRequest.ProtoReflect.Descriptor instead.
func (*GetVlanRequest) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{5}
}

func (x *GetVlanRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetVlanResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVlanResponse) Reset() {
	*x = GetVlanResponse{}
	mi := &file_vlan_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVlanResponse) ProtoMessage() {}

func (x *GetVlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vlan_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVlanResponse.ProtoReflect.Descriptor instead.
func (*GetVlanResponse) Descriptor() ([]byte, []int) {
	return file_vlan_proto_rawDescGZIP(), []int{6}
}

func (x *GetVlanResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetVlanResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_vlan_proto protoreflect.FileDescriptor

const file_vlan_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"vlan.proto\x12\x04vlan\"'\n" +
	"\x11CreateVlanRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"$\n" +
	"\x12CreateVlanResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"A\n" +
	"\x0fGetVlansRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x03R\x04page\x12\x1a\n" +
	"\bpageSize\x18\x02 \x01(\x03R\bpageSize\"N\n" +
	"\x10GetVlansResponse\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x03R\x05total\x12$\n" +
	"\x05items\x18\x02 \x03(\v2\x0e.vlan.VlanItemR\x05items\".\n" +
	"\bVlanItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\" \n" +
	"\x0eGetVlanRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"5\n" +
	"\x0fGetVlanResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name2\xc5\x01\n" +
	"\vVlanService\x12?\n" +
	"\n" +
	"CreateVlan\x12\x17.vlan.CreateVlanRequest\x1a\x18.vlan.CreateVlanResponse\x129\n" +
	"\bGetVlans\x12\x15.vlan.GetVlansRequest\x1a\x16.vlan.GetVlansResponse\x12:\n" +
	"\vGetVlanById\x12\x14.vlan.GetVlanRequest\x1a\x15.vlan.GetVlanResponseB*Z(github.com/magomedcoder/ipmanager/api/pbb\x06proto3"

var (
	file_vlan_proto_rawDescOnce sync.Once
	file_vlan_proto_rawDescData []byte
)

func file_vlan_proto_rawDescGZIP() []byte {
	file_vlan_proto_rawDescOnce.Do(func() {
		file_vlan_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vlan_proto_rawDesc), len(file_vlan_proto_rawDesc)))
	})
	return file_vlan_proto_rawDescData
}

var file_vlan_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vlan_proto_goTypes = []any{
	(*CreateVlanRequest)(nil),  // 0: vlan.CreateVlanRequest
	(*CreateVlanResponse)(nil), // 1: vlan.CreateVlanResponse
	(*GetVlansRequest)(nil),    // 2: vlan.GetVlansRequest
	(*GetVlansResponse)(nil),   // 3: vlan.GetVlansResponse
	(*VlanItem)(nil),           // 4: vlan.VlanItem
	(*GetVlanRequest)(nil),     // 5: vlan.GetVlanRequest
	(*GetVlanResponse)(nil),    // 6: vlan.GetVlanResponse
}
var file_vlan_proto_depIdxs = []int32{
	4, // 0: vlan.GetVlansResponse.items:type_name -> vlan.VlanItem
	0, // 1: vlan.VlanService.CreateVlan:input_type -> vlan.CreateVlanRequest
	2, // 2: vlan.VlanService.GetVlans:input_type -> vlan.GetVlansRequest
	5, // 3: vlan.VlanService.GetVlanById:input_type -> vlan.GetVlanRequest
	1, // 4: vlan.VlanService.CreateVlan:output_type -> vlan.CreateVlanResponse
	3, // 5: vlan.VlanService.GetVlans:output_type -> vlan.GetVlansResponse
	6, // 6: vlan.VlanService.GetVlanById:output_type -> vlan.GetVlanResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vlan_proto_init() }
func file_vlan_proto_init() {
	if File_vlan_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vlan_proto_rawDesc), len(file_vlan_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vlan_proto_goTypes,
		DependencyIndexes: file_vlan_proto_depIdxs,
		MessageInfos:      file_vlan_proto_msgTypes,
	}.Build()
	File_vlan_proto = out.File
	file_vlan_proto_goTypes = nil
	file_vlan_proto_depIdxs = nil
}
