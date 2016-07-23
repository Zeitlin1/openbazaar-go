// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile struct {
	Handle           string                   `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name             string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location         string                   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About            string                   `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription string                   `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website          string                   `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email            string                   `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	Social           []*Profile_SocialAccount `protobuf:"bytes,8,rep,name=social" json:"social,omitempty"`
	Nsfw             bool                     `protobuf:"varint,9,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor           bool                     `protobuf:"varint,10,opt,name=vendor" json:"vendor,omitempty"`
	Moderator        bool                     `protobuf:"varint,11,opt,name=moderator" json:"moderator,omitempty"`
	PrimaryColor     string                   `protobuf:"bytes,12,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor   string                   `protobuf:"bytes,13,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor        string                   `protobuf:"bytes,14,opt,name=textColor" json:"textColor,omitempty"`
	FollowerCount    uint32                   `protobuf:"varint,15,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount   uint32                   `protobuf:"varint,16,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount     uint32                   `protobuf:"varint,17,opt,name=listingCount" json:"listingCount,omitempty"`
	LastModified     uint32                   `protobuf:"varint,18,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
}

var fileDescriptor3 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x45, 0x81, 0x00, 0x03, 0xa1, 0xd4, 0xaa, 0x90, 0x85, 0x7a, 0x40, 0xa8, 0xaa, 0x50,
	0x0f, 0x39, 0xb4, 0x4f, 0x50, 0xd1, 0xeb, 0x4a, 0xab, 0xac, 0xf6, 0x01, 0x4c, 0xe2, 0x2c, 0x96,
	0x9c, 0x4c, 0x64, 0x9b, 0x65, 0x79, 0x8f, 0x7d, 0xe0, 0x8d, 0x27, 0x21, 0x10, 0xf6, 0x84, 0xff,
	0xcf, 0x1f, 0x63, 0x8f, 0x33, 0x10, 0x96, 0x06, 0x33, 0xa5, 0x65, 0x54, 0xfd, 0x3a, 0xdc, 0xbc,
	0x0f, 0x61, 0xf4, 0x58, 0x13, 0xb6, 0x84, 0xe0, 0x20, 0x8a, 0x54, 0x4b, 0xde, 0x5b, 0xf7, 0xb6,
	0x93, 0xb8, 0x49, 0x8c, 0xc1, 0xa0, 0x10, 0xb9, 0xe4, 0x5f, 0x88, 0xd2, 0x9a, 0xad, 0x60, 0xac,
	0x31, 0x11, 0x4e, 0x61, 0xc1, 0xfb, 0xc4, 0xdb, 0xcc, 0xbe, 0xc3, 0x50, 0xec, 0xf1, 0xe8, 0xf8,
	0x80, 0x36, 0xea, 0xc0, 0x7e, 0xc3, 0xc2, 0x1e, 0xd0, 0xb8, 0xff, 0xd2, 0x26, 0x46, 0x95, 0xf4,
	0xcf, 0x21, 0x09, 0x9f, 0x38, 0xe3, 0x30, 0x3a, 0xc9, 0xbd, 0x55, 0x4e, 0xf2, 0x80, 0x94, 0x4b,
	0xf4, 0xb5, 0x65, 0x2e, 0x94, 0xe6, 0xa3, 0xba, 0x36, 0x05, 0x16, 0x41, 0x60, 0x31, 0x51, 0x42,
	0xf3, 0xf1, 0xba, 0xbf, 0x9d, 0xfe, 0x59, 0x46, 0x4d, 0x4f, 0xd1, 0x13, 0xe1, 0x7f, 0x49, 0x82,
	0xc7, 0xc2, 0xc5, 0x8d, 0x45, 0x1d, 0xd9, 0xec, 0xc4, 0x27, 0x55, 0x91, 0x71, 0x4c, 0x6b, 0xdf,
	0xfd, 0xab, 0x2c, 0x52, 0x34, 0x1c, 0x88, 0x36, 0x89, 0xfd, 0x80, 0x49, 0x8e, 0xa9, 0x34, 0xc2,
	0x55, 0x5b, 0x53, 0xda, 0xba, 0x02, 0xb6, 0x81, 0x59, 0x69, 0x54, 0x2e, 0xcc, 0x79, 0x87, 0xba,
	0x12, 0x66, 0x74, 0xad, 0x0e, 0x63, 0xbf, 0x60, 0x6e, 0x65, 0x82, 0x45, 0xda, 0x5a, 0x21, 0x59,
	0x77, 0xd4, 0x9f, 0xe4, 0xe4, 0x9b, 0xab, 0x95, 0x39, 0x29, 0x57, 0xc0, 0x7e, 0x42, 0x98, 0xa1,
	0xd6, 0x78, 0x92, 0x66, 0xe7, 0x9b, 0xe1, 0x5f, 0x2b, 0x23, 0x8c, 0xbb, 0xd0, 0x9f, 0x55, 0x03,
	0x55, 0xbc, 0xd4, 0xda, 0x82, 0xb4, 0x3b, 0xea, 0xef, 0xad, 0x95, 0x75, 0xad, 0xf5, 0x8d, 0xac,
	0x0e, 0x23, 0x47, 0x58, 0xf7, 0x80, 0xa9, 0xca, 0x94, 0x4c, 0x39, 0x6b, 0x9c, 0x1b, 0xb6, 0x7a,
	0x86, 0xb0, 0xf3, 0xc4, 0xfe, 0x69, 0xdd, 0xb9, 0xbc, 0x8c, 0x10, 0xad, 0xfd, 0xb0, 0x1c, 0xad,
	0x34, 0x37, 0x43, 0xd4, 0x66, 0xff, 0x41, 0xab, 0x49, 0xc4, 0xac, 0x99, 0xa2, 0x3a, 0xec, 0x03,
	0x9a, 0xce, 0xbf, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x8e, 0x76, 0x6b, 0xae, 0x02, 0x00,
	0x00,
}
