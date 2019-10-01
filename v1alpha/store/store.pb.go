// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1alpha/store/store.proto

package store

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GraphItemEncoding int32

const (
	GraphItemEncoding_RAW  GraphItemEncoding = 0
	GraphItemEncoding_JSON GraphItemEncoding = 1
)

var GraphItemEncoding_name = map[int32]string{
	0: "RAW",
	1: "JSON",
}

var GraphItemEncoding_value = map[string]int32{
	"RAW":  0,
	"JSON": 1,
}

func (x GraphItemEncoding) String() string {
	return proto.EnumName(GraphItemEncoding_name, int32(x))
}

func (GraphItemEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{0}
}

type GraphItem struct {
	GraphItemType        string            `protobuf:"bytes,1,opt,name=graphItemType,proto3" json:"graphItemType,omitempty"`
	K1                   []byte            `protobuf:"bytes,2,opt,name=k1,proto3" json:"k1,omitempty"`
	K2                   []byte            `protobuf:"bytes,3,opt,name=k2,proto3" json:"k2,omitempty"`
	Encoding             GraphItemEncoding `protobuf:"varint,4,opt,name=encoding,proto3,enum=cloud.deps.api.v1alpha.store.GraphItemEncoding" json:"encoding,omitempty"`
	GraphItemData        []byte            `protobuf:"bytes,5,opt,name=graphItemData,proto3" json:"graphItemData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GraphItem) Reset()         { *m = GraphItem{} }
func (m *GraphItem) String() string { return proto.CompactTextString(m) }
func (*GraphItem) ProtoMessage()    {}
func (*GraphItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{0}
}
func (m *GraphItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphItem.Unmarshal(m, b)
}
func (m *GraphItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphItem.Marshal(b, m, deterministic)
}
func (m *GraphItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphItem.Merge(m, src)
}
func (m *GraphItem) XXX_Size() int {
	return xxx_messageInfo_GraphItem.Size(m)
}
func (m *GraphItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphItem.DiscardUnknown(m)
}

var xxx_messageInfo_GraphItem proto.InternalMessageInfo

func (m *GraphItem) GetGraphItemType() string {
	if m != nil {
		return m.GraphItemType
	}
	return ""
}

func (m *GraphItem) GetK1() []byte {
	if m != nil {
		return m.K1
	}
	return nil
}

func (m *GraphItem) GetK2() []byte {
	if m != nil {
		return m.K2
	}
	return nil
}

func (m *GraphItem) GetEncoding() GraphItemEncoding {
	if m != nil {
		return m.Encoding
	}
	return GraphItemEncoding_RAW
}

func (m *GraphItem) GetGraphItemData() []byte {
	if m != nil {
		return m.GraphItemData
	}
	return nil
}

type GraphItemPair struct {
	Edge                 *GraphItem `protobuf:"bytes,1,opt,name=edge,proto3" json:"edge,omitempty"`
	Node                 *GraphItem `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GraphItemPair) Reset()         { *m = GraphItemPair{} }
func (m *GraphItemPair) String() string { return proto.CompactTextString(m) }
func (*GraphItemPair) ProtoMessage()    {}
func (*GraphItemPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{1}
}
func (m *GraphItemPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphItemPair.Unmarshal(m, b)
}
func (m *GraphItemPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphItemPair.Marshal(b, m, deterministic)
}
func (m *GraphItemPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphItemPair.Merge(m, src)
}
func (m *GraphItemPair) XXX_Size() int {
	return xxx_messageInfo_GraphItemPair.Size(m)
}
func (m *GraphItemPair) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphItemPair.DiscardUnknown(m)
}

var xxx_messageInfo_GraphItemPair proto.InternalMessageInfo

func (m *GraphItemPair) GetEdge() *GraphItem {
	if m != nil {
		return m.Edge
	}
	return nil
}

func (m *GraphItemPair) GetNode() *GraphItem {
	if m != nil {
		return m.Node
	}
	return nil
}

type PutRequest struct {
	Items                []*GraphItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{2}
}
func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetItems() []*GraphItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{3}
}
func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Items                []*GraphItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{4}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetItems() []*GraphItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{5}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{6}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ListResponse struct {
	Items                []*GraphItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{7}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*GraphItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type FindRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	EdgeTypes            []string `protobuf:"bytes,2,rep,name=edgeTypes,proto3" json:"edgeTypes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{8}
}
func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FindRequest) GetEdgeTypes() []string {
	if m != nil {
		return m.EdgeTypes
	}
	return nil
}

type FindResponse struct {
	Pairs                []*GraphItemPair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4a99b9d29aef3a, []int{9}
}
func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetPairs() []*GraphItemPair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func init() {
	proto.RegisterEnum("cloud.deps.api.v1alpha.store.GraphItemEncoding", GraphItemEncoding_name, GraphItemEncoding_value)
	proto.RegisterType((*GraphItem)(nil), "cloud.deps.api.v1alpha.store.GraphItem")
	proto.RegisterType((*GraphItemPair)(nil), "cloud.deps.api.v1alpha.store.GraphItemPair")
	proto.RegisterType((*PutRequest)(nil), "cloud.deps.api.v1alpha.store.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "cloud.deps.api.v1alpha.store.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "cloud.deps.api.v1alpha.store.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "cloud.deps.api.v1alpha.store.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "cloud.deps.api.v1alpha.store.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "cloud.deps.api.v1alpha.store.ListResponse")
	proto.RegisterType((*FindRequest)(nil), "cloud.deps.api.v1alpha.store.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "cloud.deps.api.v1alpha.store.FindResponse")
}

func init() { proto.RegisterFile("v1alpha/store/store.proto", fileDescriptor_6f4a99b9d29aef3a) }

var fileDescriptor_6f4a99b9d29aef3a = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0xdd, 0x36, 0x93, 0x1f, 0x85, 0x15, 0x07, 0x13, 0xf5, 0x60, 0x59, 0x08, 0x4c,
	0x41, 0xae, 0x62, 0x8e, 0xa8, 0x87, 0xa2, 0x00, 0x82, 0x40, 0x09, 0x5b, 0x10, 0x52, 0x25, 0x0e,
	0x4b, 0xbc, 0x4a, 0xad, 0x24, 0xde, 0xc5, 0xbb, 0x06, 0xf5, 0x11, 0x78, 0x2f, 0x1e, 0x8b, 0x03,
	0xda, 0x9f, 0xa4, 0xb5, 0x90, 0x1a, 0x57, 0xc0, 0x25, 0xda, 0x19, 0xcd, 0xf7, 0x33, 0x9e, 0x4f,
	0x81, 0xbb, 0xdf, 0x46, 0x64, 0xc9, 0xcf, 0xc9, 0xa1, 0x90, 0xac, 0xa4, 0xe6, 0x37, 0xe1, 0x25,
	0x93, 0x0c, 0xed, 0xcf, 0x96, 0xac, 0xca, 0x92, 0x8c, 0x72, 0x91, 0x10, 0x9e, 0x27, 0x76, 0x32,
	0xd1, 0x33, 0xd1, 0x4f, 0x07, 0xda, 0x2f, 0x4b, 0xc2, 0xcf, 0x5f, 0x49, 0xba, 0x42, 0xf7, 0xa0,
	0x37, 0x5f, 0x17, 0x1f, 0x2e, 0x38, 0x0d, 0x9c, 0xd0, 0x89, 0xdb, 0xb8, 0xde, 0x44, 0x7d, 0x68,
	0x2d, 0x46, 0x41, 0x2b, 0x74, 0xe2, 0x2e, 0x6e, 0x2d, 0x46, 0xba, 0x4e, 0x03, 0xd7, 0xd6, 0x29,
	0x9a, 0xc0, 0x1e, 0x2d, 0x66, 0x2c, 0xcb, 0x8b, 0x79, 0xe0, 0x85, 0x4e, 0xdc, 0x4f, 0x0f, 0x93,
	0xeb, 0x4c, 0x24, 0x1b, 0x03, 0xcf, 0x2d, 0x0c, 0x6f, 0x08, 0x6a, 0x96, 0xc6, 0x44, 0x92, 0xc0,
	0xd7, 0x3a, 0xf5, 0x66, 0xf4, 0xc3, 0x81, 0xde, 0x86, 0x65, 0x4a, 0xf2, 0x12, 0x3d, 0x05, 0x8f,
	0x66, 0x73, 0xb3, 0x41, 0x27, 0x7d, 0xd0, 0xd0, 0x00, 0xd6, 0x20, 0x05, 0x2e, 0x58, 0x46, 0xf5,
	0x8e, 0x37, 0x01, 0x2b, 0x50, 0x34, 0x01, 0x98, 0x56, 0x12, 0xd3, 0xaf, 0x15, 0x15, 0x12, 0x1d,
	0x81, 0x9f, 0x4b, 0xba, 0x12, 0x81, 0x13, 0xba, 0x37, 0xe1, 0x32, 0xa8, 0xa8, 0x07, 0x1d, 0x4d,
	0x26, 0x38, 0x2b, 0x04, 0x8d, 0x4e, 0xa0, 0x37, 0xa6, 0x4b, 0x2a, 0xe9, 0x3f, 0xa2, 0x1f, 0x40,
	0x7f, 0xcd, 0x67, 0x15, 0x26, 0xd0, 0x79, 0x93, 0x8b, 0x8d, 0x7d, 0x04, 0x1e, 0x27, 0xf6, 0x33,
	0xfa, 0x58, 0xbf, 0xd1, 0x1d, 0xf0, 0x67, 0xac, 0x2a, 0xa4, 0xfe, 0x3c, 0x3e, 0x36, 0x85, 0x9a,
	0x94, 0x2a, 0x32, 0xae, 0x8e, 0x8c, 0x7e, 0x47, 0x6f, 0xa1, 0x6b, 0xc8, 0x0c, 0xf9, 0xdf, 0xba,
	0x3d, 0x82, 0xce, 0x8b, 0xbc, 0xc8, 0xd6, 0xde, 0x06, 0xe0, 0x2e, 0xe8, 0x85, 0xb6, 0xd6, 0xc5,
	0xea, 0x89, 0xf6, 0xa1, 0xad, 0xee, 0xa7, 0x52, 0x2a, 0x82, 0x56, 0xe8, 0xc6, 0x6d, 0x7c, 0xd9,
	0x88, 0xde, 0x43, 0xd7, 0xc0, 0xad, 0x9b, 0x63, 0xf0, 0x39, 0xc9, 0xcb, 0xb5, 0x9b, 0x47, 0x0d,
	0xdd, 0xa8, 0x78, 0x61, 0x83, 0x3c, 0xb8, 0x0f, 0xb7, 0xff, 0x08, 0x2f, 0xda, 0x05, 0x17, 0x1f,
	0x7f, 0x1a, 0xdc, 0x42, 0x7b, 0xe0, 0xbd, 0x3e, 0x7d, 0x77, 0x32, 0x70, 0xd2, 0x5f, 0x2e, 0x80,
	0x1e, 0x3c, 0x55, 0x5c, 0xe8, 0x0c, 0xdc, 0x69, 0x25, 0x51, 0x7c, 0xbd, 0xe2, 0x65, 0x8a, 0x86,
	0x0f, 0x1b, 0x4c, 0xda, 0xad, 0x66, 0xb0, 0x63, 0x4e, 0x8a, 0xb6, 0x2c, 0x54, 0x0b, 0xd2, 0xf0,
	0x71, 0xb3, 0x61, 0x2b, 0xf2, 0x19, 0x3c, 0x75, 0x58, 0xb4, 0xc5, 0xd7, 0x95, 0x24, 0x0d, 0x0f,
	0x9a, 0x8c, 0x5a, 0x7a, 0x6a, 0x2e, 0xf5, 0x91, 0x0b, 0x59, 0x52, 0xb2, 0xda, 0x26, 0x73, 0x25,
	0x14, 0xdb, 0x64, 0x6a, 0x01, 0x98, 0x43, 0x5f, 0xd5, 0x63, 0xf6, 0xbd, 0xf8, 0xaf, 0x42, 0xcf,
	0x76, 0xcf, 0x7c, 0xdd, 0xfd, 0xb2, 0xa3, 0xff, 0x93, 0x9f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x29, 0x03, 0x3f, 0xf1, 0xb0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GraphStoreClient is the client API for GraphStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GraphStoreClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	FindUpstream(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	FindDownstream(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type graphStoreClient struct {
	cc *grpc.ClientConn
}

func NewGraphStoreClient(cc *grpc.ClientConn) GraphStoreClient {
	return &graphStoreClient{cc}
}

func (c *graphStoreClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.api.v1alpha.store.GraphStore/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.api.v1alpha.store.GraphStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphStoreClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.api.v1alpha.store.GraphStore/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphStoreClient) FindUpstream(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.api.v1alpha.store.GraphStore/FindUpstream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphStoreClient) FindDownstream(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.api.v1alpha.store.GraphStore/FindDownstream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraphStoreServer is the server API for GraphStore service.
type GraphStoreServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	FindUpstream(context.Context, *FindRequest) (*FindResponse, error)
	FindDownstream(context.Context, *FindRequest) (*FindResponse, error)
}

// UnimplementedGraphStoreServer can be embedded to have forward compatible implementations.
type UnimplementedGraphStoreServer struct {
}

func (*UnimplementedGraphStoreServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedGraphStoreServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedGraphStoreServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGraphStoreServer) FindUpstream(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUpstream not implemented")
}
func (*UnimplementedGraphStoreServer) FindDownstream(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDownstream not implemented")
}

func RegisterGraphStoreServer(s *grpc.Server, srv GraphStoreServer) {
	s.RegisterService(&_GraphStore_serviceDesc, srv)
}

func _GraphStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.api.v1alpha.store.GraphStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphStoreServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.api.v1alpha.store.GraphStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphStore_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphStoreServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.api.v1alpha.store.GraphStore/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphStoreServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphStore_FindUpstream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphStoreServer).FindUpstream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.api.v1alpha.store.GraphStore/FindUpstream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphStoreServer).FindUpstream(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphStore_FindDownstream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphStoreServer).FindDownstream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.api.v1alpha.store.GraphStore/FindDownstream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphStoreServer).FindDownstream(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GraphStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.deps.api.v1alpha.store.GraphStore",
	HandlerType: (*GraphStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _GraphStore_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GraphStore_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GraphStore_List_Handler,
		},
		{
			MethodName: "FindUpstream",
			Handler:    _GraphStore_FindUpstream_Handler,
		},
		{
			MethodName: "FindDownstream",
			Handler:    _GraphStore_FindDownstream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha/store/store.proto",
}
