// Code generated by protoc-gen-go. DO NOT EDIT.
// source: database.proto

/*
Package dataprotobuf is a generated protocol buffer package.

It is generated from these files:
	database.proto

It has these top-level messages:
	Title
	Titlelist
	Content
	Contentlist
*/
package dataprotobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Title struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Title) Reset()                    { *m = Title{} }
func (m *Title) String() string            { return proto.CompactTextString(m) }
func (*Title) ProtoMessage()               {}
func (*Title) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Title) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Title) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Titlelist struct {
	Rtncode int32    `protobuf:"varint,1,opt,name=rtncode" json:"rtncode,omitempty"`
	Titles  []*Title `protobuf:"bytes,2,rep,name=titles" json:"titles,omitempty"`
}

func (m *Titlelist) Reset()                    { *m = Titlelist{} }
func (m *Titlelist) String() string            { return proto.CompactTextString(m) }
func (*Titlelist) ProtoMessage()               {}
func (*Titlelist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Titlelist) GetRtncode() int32 {
	if m != nil {
		return m.Rtncode
	}
	return 0
}

func (m *Titlelist) GetTitles() []*Title {
	if m != nil {
		return m.Titles
	}
	return nil
}

type Content struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Category    string `protobuf:"bytes,2,opt,name=category" json:"category,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Results     string `protobuf:"bytes,4,opt,name=results" json:"results,omitempty"`
	Number      int32  `protobuf:"varint,5,opt,name=number" json:"number,omitempty"`
	Searchcount int32  `protobuf:"varint,6,opt,name=searchcount" json:"searchcount,omitempty"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Content) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Content) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Content) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Content) GetResults() string {
	if m != nil {
		return m.Results
	}
	return ""
}

func (m *Content) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Content) GetSearchcount() int32 {
	if m != nil {
		return m.Searchcount
	}
	return 0
}

type Contentlist struct {
	Id       int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Rtncode  int32      `protobuf:"varint,2,opt,name=rtncode" json:"rtncode,omitempty"`
	Contents []*Content `protobuf:"bytes,3,rep,name=contents" json:"contents,omitempty"`
}

func (m *Contentlist) Reset()                    { *m = Contentlist{} }
func (m *Contentlist) String() string            { return proto.CompactTextString(m) }
func (*Contentlist) ProtoMessage()               {}
func (*Contentlist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Contentlist) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contentlist) GetRtncode() int32 {
	if m != nil {
		return m.Rtncode
	}
	return 0
}

func (m *Contentlist) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*Title)(nil), "dataprotobuf.title")
	proto.RegisterType((*Titlelist)(nil), "dataprotobuf.titlelist")
	proto.RegisterType((*Content)(nil), "dataprotobuf.content")
	proto.RegisterType((*Contentlist)(nil), "dataprotobuf.contentlist")
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4e, 0x84, 0x30,
	0x10, 0x86, 0x43, 0x59, 0xd8, 0xdd, 0xc1, 0xec, 0xa1, 0x46, 0xd3, 0x78, 0x22, 0x9c, 0x48, 0x36,
	0x21, 0x51, 0xdf, 0xa4, 0x6f, 0x50, 0xca, 0xa8, 0x18, 0xb6, 0x35, 0xed, 0x70, 0xf0, 0x61, 0x7c,
	0x57, 0xe3, 0xa4, 0x8b, 0xa0, 0xb7, 0xf9, 0xe7, 0xff, 0x07, 0xbe, 0xfe, 0x70, 0x1a, 0x0c, 0x99,
	0xde, 0x44, 0xec, 0x3e, 0x82, 0x27, 0x2f, 0x6f, 0x7e, 0x34, 0x8f, 0xfd, 0xfc, 0xd2, 0x9c, 0xa1,
	0xa0, 0x91, 0x26, 0x94, 0x27, 0x10, 0xe3, 0xa0, 0xb2, 0x3a, 0x6b, 0x0b, 0x2d, 0xc6, 0x41, 0x4a,
	0xd8, 0x39, 0x73, 0x41, 0x25, 0xea, 0xac, 0x3d, 0x6a, 0x9e, 0x1b, 0x0d, 0x47, 0x0e, 0x4f, 0x63,
	0x24, 0xa9, 0x60, 0x1f, 0xc8, 0x59, 0x3f, 0x60, 0xba, 0xba, 0x4a, 0x79, 0x86, 0x92, 0x63, 0x51,
	0x89, 0x3a, 0x6f, 0xab, 0xa7, 0xdb, 0x6e, 0xfd, 0xcb, 0x8e, 0x3d, 0x9d, 0x22, 0xcd, 0x57, 0x06,
	0x7b, 0xeb, 0x1d, 0xa1, 0xa3, 0x7f, 0x0c, 0x0f, 0x70, 0xb0, 0x86, 0xf0, 0xd5, 0x87, 0xcf, 0xc4,
	0xb1, 0xe8, 0x85, 0x2f, 0xff, 0xe5, 0x63, 0x24, 0x8c, 0xf3, 0x44, 0x51, 0xed, 0x78, 0x7d, 0x95,
	0xf2, 0x1e, 0x4a, 0x37, 0x5f, 0x7a, 0x0c, 0xaa, 0xe0, 0xaf, 0x27, 0x25, 0x6b, 0xa8, 0x22, 0x9a,
	0x60, 0xdf, 0xac, 0x9f, 0x1d, 0xa9, 0x92, 0xcd, 0xf5, 0xaa, 0x79, 0x87, 0x2a, 0xe1, 0xf1, 0xab,
	0xff, 0x22, 0xae, 0x5a, 0x10, 0xdb, 0x16, 0x1e, 0xe1, 0x90, 0x0e, 0xa3, 0xca, 0xb9, 0x87, 0xbb,
	0x6d, 0x0f, 0xc9, 0xd5, 0x4b, 0xac, 0x2f, 0xd9, 0x7b, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0x99, 0x77, 0xe1, 0xb3, 0x01, 0x00, 0x00,
}
