// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/nft_transfer/v1/mock_extension.proto

package mock

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// ClassMetadata defines a struct for the class metadata
type ClassMetadata struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Schema           string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	MintRestricted   bool   `protobuf:"varint,3,opt,name=mint_restricted,json=mintRestricted,proto3" json:"mint_restricted,omitempty"`
	UpdateRestricted bool   `protobuf:"varint,4,opt,name=update_restricted,json=updateRestricted,proto3" json:"update_restricted,omitempty"`
	Data             string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ClassMetadata) Reset()         { *m = ClassMetadata{} }
func (m *ClassMetadata) String() string { return proto.CompactTextString(m) }
func (*ClassMetadata) ProtoMessage()    {}
func (*ClassMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe79d10c731b049f, []int{0}
}
func (m *ClassMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassMetadata.Merge(m, src)
}
func (m *ClassMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ClassMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ClassMetadata proto.InternalMessageInfo

func (m *ClassMetadata) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ClassMetadata) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *ClassMetadata) GetMintRestricted() bool {
	if m != nil {
		return m.MintRestricted
	}
	return false
}

func (m *ClassMetadata) GetUpdateRestricted() bool {
	if m != nil {
		return m.UpdateRestricted
	}
	return false
}

func (m *ClassMetadata) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// NFTMetadata defines a struct for the nft metadata
type NFTMetadata struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *NFTMetadata) Reset()         { *m = NFTMetadata{} }
func (m *NFTMetadata) String() string { return proto.CompactTextString(m) }
func (*NFTMetadata) ProtoMessage()    {}
func (*NFTMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe79d10c731b049f, []int{1}
}
func (m *NFTMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTMetadata.Merge(m, src)
}
func (m *NFTMetadata) XXX_Size() int {
	return m.Size()
}
func (m *NFTMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NFTMetadata proto.InternalMessageInfo

func (m *NFTMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NFTMetadata) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Extension defines a data structure for storing data types that the system
// cannot recognize
type Extension struct {
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe79d10c731b049f, []int{2}
}
func (m *Extension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(m, src)
}
func (m *Extension) XXX_Size() int {
	return m.Size()
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

func (m *Extension) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*ClassMetadata)(nil), "ibc.applications.nft_transfer.v1.ClassMetadata")
	proto.RegisterType((*NFTMetadata)(nil), "ibc.applications.nft_transfer.v1.NFTMetadata")
	proto.RegisterType((*Extension)(nil), "ibc.applications.nft_transfer.v1.Extension")
}

func init() {
	proto.RegisterFile("ibc/applications/nft_transfer/v1/mock_extension.proto", fileDescriptor_fe79d10c731b049f)
}

var fileDescriptor_fe79d10c731b049f = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0xc6, 0xb3, 0xf1, 0x8c, 0x66, 0xc5, 0x7f, 0x8b, 0xc8, 0x61, 0x71, 0x86, 0x80, 0x24, 0x20,
	0xde, 0x11, 0xc4, 0x46, 0x2b, 0x15, 0xed, 0xb4, 0x08, 0x56, 0x36, 0x61, 0x6e, 0xb3, 0x49, 0x56,
	0x73, 0xbb, 0xc7, 0xee, 0x24, 0xe8, 0x5b, 0xf8, 0x08, 0xbe, 0x82, 0x6f, 0x61, 0x99, 0xd2, 0x52,
	0x92, 0xc6, 0xc7, 0x90, 0xdb, 0x24, 0xc7, 0xa1, 0xdd, 0x37, 0x33, 0xbf, 0x8f, 0xf9, 0x98, 0xa1,
	0x67, 0x32, 0xe6, 0x11, 0xa4, 0xe9, 0x50, 0x72, 0x40, 0xa9, 0x95, 0x8d, 0x54, 0x0f, 0x3b, 0x68,
	0x40, 0xd9, 0x9e, 0x30, 0xd1, 0xb8, 0x15, 0x25, 0x9a, 0x3f, 0x77, 0xc4, 0x0b, 0x0a, 0x65, 0xa5,
	0x56, 0x61, 0x6a, 0x34, 0x6a, 0x56, 0x93, 0x31, 0x0f, 0x8b, 0xb6, 0xb0, 0x68, 0x0b, 0xc7, 0xad,
	0x83, 0xbd, 0xbe, 0xee, 0x6b, 0x07, 0x47, 0x99, 0x9a, 0xfb, 0xea, 0x1f, 0x84, 0x6e, 0x5e, 0x0f,
	0xc1, 0xda, 0x3b, 0x81, 0xd0, 0x05, 0x04, 0xe6, 0xd3, 0x35, 0x6e, 0x04, 0xa0, 0x36, 0x3e, 0xa9,
	0x91, 0x66, 0xb5, 0xbd, 0x2c, 0xd9, 0x3e, 0xad, 0x58, 0x3e, 0x10, 0x09, 0xf8, 0x65, 0x37, 0x58,
	0x54, 0xac, 0x41, 0xb7, 0x13, 0xa9, 0xb0, 0x63, 0x84, 0x45, 0x23, 0x39, 0x8a, 0xae, 0xbf, 0x52,
	0x23, 0xcd, 0xf5, 0xf6, 0x56, 0xd6, 0x6e, 0xe7, 0x5d, 0x76, 0x4c, 0x77, 0x47, 0x69, 0x17, 0x50,
	0x14, 0x51, 0xcf, 0xa1, 0x3b, 0xf3, 0x41, 0x01, 0x66, 0xd4, 0xcb, 0xf2, 0xf8, 0xab, 0x6e, 0x97,
	0xd3, 0xe7, 0xde, 0xcf, 0xfb, 0x21, 0xa9, 0x5f, 0xd0, 0x8d, 0xfb, 0xdb, 0x87, 0x3c, 0x30, 0xa3,
	0x9e, 0x82, 0x44, 0x2c, 0xd2, 0x3a, 0x9d, 0x9b, 0xcb, 0xff, 0xcc, 0x47, 0xb4, 0x7a, 0xb3, 0xbc,
	0x5d, 0x8e, 0x91, 0xbf, 0xd8, 0xd5, 0xe5, 0xe7, 0x34, 0x20, 0x93, 0x69, 0x40, 0xbe, 0xa7, 0x01,
	0x79, 0x9b, 0x05, 0xa5, 0xc9, 0x2c, 0x28, 0x7d, 0xcd, 0x82, 0xd2, 0x63, 0xa3, 0x2f, 0x71, 0x30,
	0x8a, 0x43, 0xae, 0x93, 0x28, 0x96, 0xa0, 0x9e, 0xa4, 0x00, 0x99, 0x3d, 0xe9, 0x24, 0x7f, 0x12,
	0xbe, 0xa6, 0xc2, 0xc6, 0x15, 0x77, 0xe1, 0xd3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x0a,
	0xbb, 0xf8, 0xd2, 0x01, 0x00, 0x00,
}

func (this *ClassMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClassMetadata)
	if !ok {
		that2, ok := that.(ClassMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Schema != that1.Schema {
		return false
	}
	if this.MintRestricted != that1.MintRestricted {
		return false
	}
	if this.UpdateRestricted != that1.UpdateRestricted {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	return true
}
func (this *NFTMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NFTMetadata)
	if !ok {
		that2, ok := that.(NFTMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	return true
}
func (this *Extension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extension)
	if !ok {
		that2, ok := that.(Extension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	return true
}
func (m *ClassMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMockExtension(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if m.UpdateRestricted {
		i--
		if m.UpdateRestricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.MintRestricted {
		i--
		if m.MintRestricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintMockExtension(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMockExtension(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFTMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMockExtension(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMockExtension(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Extension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Extension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Extension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMockExtension(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMockExtension(dAtA []byte, offset int, v uint64) int {
	offset -= sovMockExtension(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMockExtension(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovMockExtension(uint64(l))
	}
	if m.MintRestricted {
		n += 2
	}
	if m.UpdateRestricted {
		n += 2
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMockExtension(uint64(l))
	}
	return n
}

func (m *NFTMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMockExtension(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMockExtension(uint64(l))
	}
	return n
}

func (m *Extension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMockExtension(uint64(l))
	}
	return n
}

func sovMockExtension(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMockExtension(x uint64) (n int) {
	return sovMockExtension(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockExtension
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClassMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRestricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MintRestricted = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateRestricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UpdateRestricted = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMockExtension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockExtension
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NFTMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockExtension
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NFTMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMockExtension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockExtension
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Extension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockExtension
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Extension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Extension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockExtension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMockExtension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockExtension
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMockExtension(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMockExtension
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMockExtension
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMockExtension
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMockExtension
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMockExtension
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMockExtension        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMockExtension          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMockExtension = fmt.Errorf("proto: unexpected end of group")
)
