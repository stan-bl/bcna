// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bcna/genesis.proto

package types

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

// GenesisState defines the bcna module's genesis state.
type GenesisState struct {
	Params           Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BitcannaidList   []Bitcannaid  `protobuf:"bytes,2,rep,name=bitcannaidList,proto3" json:"bitcannaidList"`
	BitcannaidCount  uint64        `protobuf:"varint,3,opt,name=bitcannaidCount,proto3" json:"bitcannaidCount,omitempty"`
	SupplychainList  []Supplychain `protobuf:"bytes,4,rep,name=supplychainList,proto3" json:"supplychainList"`
	SupplychainCount uint64        `protobuf:"varint,5,opt,name=supplychainCount,proto3" json:"supplychainCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0474e07bc9bd608, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBitcannaidList() []Bitcannaid {
	if m != nil {
		return m.BitcannaidList
	}
	return nil
}

func (m *GenesisState) GetBitcannaidCount() uint64 {
	if m != nil {
		return m.BitcannaidCount
	}
	return 0
}

func (m *GenesisState) GetSupplychainList() []Supplychain {
	if m != nil {
		return m.SupplychainList
	}
	return nil
}

func (m *GenesisState) GetSupplychainCount() uint64 {
	if m != nil {
		return m.SupplychainCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bitcannaglobal.bcna.bcna.GenesisState")
}

func init() { proto.RegisterFile("bcna/genesis.proto", fileDescriptor_a0474e07bc9bd608) }

var fileDescriptor_a0474e07bc9bd608 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0xc6, 0x73, 0x6d, 0xed, 0x70, 0x15, 0xab, 0x87, 0x4a, 0xe8, 0x70, 0x06, 0x51, 0x08, 0x0a,
	0x09, 0xd4, 0xdd, 0x21, 0x45, 0xba, 0x38, 0x48, 0x8a, 0x8b, 0xdb, 0x25, 0x86, 0xf4, 0x20, 0xbd,
	0x0b, 0xcd, 0x05, 0xec, 0x5b, 0xf8, 0x58, 0xdd, 0xec, 0xe8, 0x24, 0x92, 0xbc, 0x88, 0xe4, 0x7f,
	0xd1, 0x94, 0x48, 0x96, 0x24, 0xfc, 0xf8, 0xee, 0xf7, 0x7d, 0xe1, 0x30, 0x09, 0x42, 0xc1, 0xdc,
	0x38, 0x12, 0x51, 0xc6, 0x33, 0x27, 0x5d, 0x4b, 0x25, 0x89, 0x19, 0x70, 0x15, 0x32, 0x21, 0x58,
	0x9c, 0xc8, 0x80, 0x25, 0x4e, 0x15, 0x81, 0xc7, 0xe4, 0x34, 0x96, 0xb1, 0x84, 0x90, 0x5b, 0x7d,
	0xe9, 0xfc, 0xe4, 0x04, 0x1c, 0x29, 0x5b, 0xb3, 0x55, 0xad, 0x98, 0x9c, 0x01, 0xfa, 0xf5, 0xf0,
	0xd7, 0x1a, 0x9f, 0x03, 0xce, 0xf2, 0x34, 0x4d, 0x36, 0xe1, 0x92, 0x71, 0xa1, 0xf9, 0xe5, 0x47,
	0x0f, 0x1f, 0xce, 0xf5, 0x86, 0x85, 0x62, 0x2a, 0x22, 0xf7, 0x78, 0xa8, 0x7d, 0x26, 0xb2, 0x90,
	0x3d, 0x9a, 0x5a, 0x4e, 0xd7, 0x26, 0xe7, 0x09, 0x72, 0xde, 0x60, 0xfb, 0x75, 0x61, 0xf8, 0xf5,
	0x29, 0xe2, 0xe3, 0xa3, 0xa6, 0xfc, 0x91, 0x67, 0xca, 0xec, 0x59, 0x7d, 0x7b, 0x34, 0xbd, 0xea,
	0xf6, 0x78, 0x7f, 0xf9, 0xda, 0xd5, 0x32, 0x10, 0x1b, 0x8f, 0x1b, 0x32, 0x93, 0xb9, 0x50, 0x66,
	0xdf, 0x42, 0xf6, 0xc0, 0x6f, 0x63, 0xf2, 0x8c, 0xc7, 0x7b, 0xff, 0x08, 0xf5, 0x03, 0xa8, 0xbf,
	0xee, 0xae, 0x5f, 0x34, 0x07, 0xea, 0xfe, 0xb6, 0x83, 0xdc, 0xe0, 0xe3, 0x3d, 0xa4, 0x17, 0x1c,
	0xc0, 0x82, 0x7f, 0xdc, 0x7b, 0xd8, 0x16, 0x14, 0xed, 0x0a, 0x8a, 0xbe, 0x0b, 0x8a, 0xde, 0x4b,
	0x6a, 0xec, 0x4a, 0x6a, 0x7c, 0x96, 0xd4, 0x78, 0xb9, 0x8d, 0xb9, 0x5a, 0xe6, 0x81, 0x13, 0xca,
	0x95, 0xeb, 0x71, 0x35, 0xab, 0xd6, 0xcc, 0x61, 0x8d, 0x0b, 0xb7, 0xf3, 0xa6, 0x5f, 0x6a, 0x93,
	0x46, 0x59, 0x30, 0x84, 0xfb, 0xb9, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x0e, 0x3d, 0x66,
	0x27, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SupplychainCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SupplychainCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SupplychainList) > 0 {
		for iNdEx := len(m.SupplychainList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SupplychainList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BitcannaidCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BitcannaidCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BitcannaidList) > 0 {
		for iNdEx := len(m.BitcannaidList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BitcannaidList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BitcannaidList) > 0 {
		for _, e := range m.BitcannaidList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BitcannaidCount != 0 {
		n += 1 + sovGenesis(uint64(m.BitcannaidCount))
	}
	if len(m.SupplychainList) > 0 {
		for _, e := range m.SupplychainList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SupplychainCount != 0 {
		n += 1 + sovGenesis(uint64(m.SupplychainCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitcannaidList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BitcannaidList = append(m.BitcannaidList, Bitcannaid{})
			if err := m.BitcannaidList[len(m.BitcannaidList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitcannaidCount", wireType)
			}
			m.BitcannaidCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BitcannaidCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplychainList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupplychainList = append(m.SupplychainList, Supplychain{})
			if err := m.SupplychainList[len(m.SupplychainList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplychainCount", wireType)
			}
			m.SupplychainCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SupplychainCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
