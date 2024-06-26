// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/incentive/v1beta1/apy.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Apy contains the calculated APY for a given collateral type at a specific
// instant in time.
type Apy struct {
	CollateralType string                                 `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	Apy            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=apy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"apy"`
}

func (m *Apy) Reset()         { *m = Apy{} }
func (m *Apy) String() string { return proto.CompactTextString(m) }
func (*Apy) ProtoMessage()    {}
func (*Apy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c1ad571f25cae9, []int{0}
}
func (m *Apy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Apy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Apy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Apy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Apy.Merge(m, src)
}
func (m *Apy) XXX_Size() int {
	return m.Size()
}
func (m *Apy) XXX_DiscardUnknown() {
	xxx_messageInfo_Apy.DiscardUnknown(m)
}

var xxx_messageInfo_Apy proto.InternalMessageInfo

func (m *Apy) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func init() {
	proto.RegisterType((*Apy)(nil), "kava.incentive.v1beta1.Apy")
}

func init() { proto.RegisterFile("kava/incentive/v1beta1/apy.proto", fileDescriptor_b2c1ad571f25cae9) }

var fileDescriptor_b2c1ad571f25cae9 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x4e, 0x2c, 0x4b,
	0xd4, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x4f, 0x2c, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x03, 0xa9,
	0xd0, 0x83, 0xab, 0xd0, 0x83, 0xaa, 0x90, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07,
	0xab, 0xd2, 0x87, 0x70, 0x20, 0x5a, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0x21, 0xe2, 0x20, 0x16,
	0x44, 0x54, 0xa9, 0x8e, 0x8b, 0xd9, 0xb1, 0xa0, 0x52, 0x48, 0x9d, 0x8b, 0x3f, 0x39, 0x3f, 0x27,
	0x27, 0xb1, 0x24, 0xb5, 0x28, 0x31, 0x27, 0xbe, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x88, 0x0f, 0x21, 0x1c, 0x52, 0x59, 0x90, 0x2a, 0xe4, 0xc7, 0xc5, 0x9c, 0x58, 0x50,
	0x29, 0xc1, 0x04, 0x92, 0x74, 0xb2, 0x39, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0xb5, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xa8, 0x9d, 0x50, 0x4a, 0xb7, 0x38, 0x25,
	0x5b, 0x1f, 0x64, 0x5a, 0xb1, 0x9e, 0x4b, 0x6a, 0xf2, 0xa5, 0x2d, 0xba, 0x5c, 0x50, 0x27, 0xb9,
	0xa4, 0x26, 0x07, 0x81, 0x0c, 0x72, 0x72, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x6d, 0x24, 0x43, 0x41, 0xbe, 0xd5, 0xcd, 0x49, 0x4c, 0x2a, 0x06, 0xb3, 0xf4, 0x2b,
	0x90, 0xc2, 0x06, 0x6c, 0x7a, 0x12, 0x1b, 0xd8, 0x37, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x59, 0xa8, 0x77, 0x3a, 0x01, 0x00, 0x00,
}

func (m *Apy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Apy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Apy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Apy.Size()
		i -= size
		if _, err := m.Apy.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintApy(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintApy(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApy(dAtA []byte, offset int, v uint64) int {
	offset -= sovApy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Apy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovApy(uint64(l))
	}
	l = m.Apy.Size()
	n += 1 + l + sovApy(uint64(l))
	return n
}

func sovApy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApy(x uint64) (n int) {
	return sovApy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Apy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApy
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
			return fmt.Errorf("proto: Apy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Apy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApy
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
				return ErrInvalidLengthApy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApy
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
				return ErrInvalidLengthApy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Apy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApy
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
func skipApy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApy
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
					return 0, ErrIntOverflowApy
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
					return 0, ErrIntOverflowApy
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
				return 0, ErrInvalidLengthApy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApy = fmt.Errorf("proto: unexpected end of group")
)
