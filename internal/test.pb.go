// Code generated by protoc-gen-gogo.
// source: test.proto
// DO NOT EDIT!

/*
	Package internal is a generated protocol buffer package.

	It is generated from these files:
		test.proto

	It has these top-level messages:
		Struct
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Struct struct {
	Method string   `protobuf:"bytes,1,opt,name=Method,proto3" json:"Method,omitempty"`
	Seq    uint64   `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Bucket [][]byte `protobuf:"bytes,3,rep,name=bucket" json:"bucket,omitempty"`
	Data   []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Struct) Reset()                    { *m = Struct{} }
func (m *Struct) String() string            { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()               {}
func (*Struct) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{0} }

func (m *Struct) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Struct) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Struct) GetBucket() [][]byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *Struct) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Struct)(nil), "internal.Struct")
}
func (m *Struct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Struct) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTest(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if m.Seq != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTest(dAtA, i, uint64(m.Seq))
	}
	if len(m.Bucket) > 0 {
		for _, b := range m.Bucket {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintTest(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTest(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func encodeFixed64Test(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Test(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Struct) Size() (n int) {
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.Seq != 0 {
		n += 1 + sovTest(uint64(m.Seq))
	}
	if len(m.Bucket) > 0 {
		for _, b := range m.Bucket {
			l = len(b)
			n += 1 + l + sovTest(uint64(l))
		}
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}

func sovTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Struct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Struct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Struct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bucket", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bucket = append(m.Bucket, make([]byte, postIndex-iNdEx))
			copy(m.Bucket[len(m.Bucket)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
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
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTest
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTest
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("test.proto", fileDescriptorTest) }

var fileDescriptorTest = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc,
	0x51, 0x8a, 0xe3, 0x62, 0x0b, 0x2e, 0x29, 0x2a, 0x4d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0xf3, 0x4d,
	0x2d, 0xc9, 0xc8, 0x4f, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x04, 0xb8,
	0x98, 0x83, 0x53, 0x0b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x40, 0x4c, 0x90, 0xca, 0xa4,
	0xd2, 0xe4, 0xec, 0xd4, 0x12, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x9e, 0x20, 0x28, 0x4f, 0x48, 0x88,
	0x8b, 0x25, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x45, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x76, 0x12,
	0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0x48, 0x62, 0x03, 0x3b, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xde, 0xe3, 0xea, 0x14,
	0x90, 0x00, 0x00, 0x00,
}
