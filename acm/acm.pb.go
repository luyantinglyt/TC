// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acm.proto

/*
	Package acm is a generated protocol buffer package.

	It is generated from these files:
		acm.proto

	It has these top-level messages:
		ConcreteAccount
*/
package acm

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import permission "github.com/hyperledger/burrow/permission"
import crypto "github.com/hyperledger/burrow/crypto"

import github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConcreteAccount struct {
	Address     github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address"`
	PublicKey   crypto.PublicKey                             `protobuf:"bytes,2,opt,name=PublicKey" json:"PublicKey"`
	Sequence    uint64                                       `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Balance     uint64                                       `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`
	Code        Bytecode                                     `protobuf:"bytes,5,opt,name=Code,proto3,customtype=Bytecode" json:"Code"`
	Permissions permission.AccountPermissions                `protobuf:"bytes,6,opt,name=Permissions" json:"Permissions"`
}

func (m *ConcreteAccount) Reset()                    { *m = ConcreteAccount{} }
func (*ConcreteAccount) ProtoMessage()               {}
func (*ConcreteAccount) Descriptor() ([]byte, []int) { return fileDescriptorAcm, []int{0} }

func (m *ConcreteAccount) GetPublicKey() crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return crypto.PublicKey{}
}

func (m *ConcreteAccount) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ConcreteAccount) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *ConcreteAccount) GetPermissions() permission.AccountPermissions {
	if m != nil {
		return m.Permissions
	}
	return permission.AccountPermissions{}
}

func (*ConcreteAccount) XXX_MessageName() string {
	return "acm.ConcreteAccount"
}
func init() {
	proto.RegisterType((*ConcreteAccount)(nil), "acm.ConcreteAccount")
	golang_proto.RegisterType((*ConcreteAccount)(nil), "acm.ConcreteAccount")
}
func (m *ConcreteAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConcreteAccount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Address.Size()))
	n1, err := m.Address.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.PublicKey.Size()))
	n2, err := m.PublicKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Sequence != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Sequence))
	}
	if m.Balance != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Balance))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Code.Size()))
	n3, err := m.Code.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Permissions.Size()))
	n4, err := m.Permissions.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeVarintAcm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConcreteAccount) Size() (n int) {
	var l int
	_ = l
	l = m.Address.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.PublicKey.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.Sequence != 0 {
		n += 1 + sovAcm(uint64(m.Sequence))
	}
	if m.Balance != 0 {
		n += 1 + sovAcm(uint64(m.Balance))
	}
	l = m.Code.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.Permissions.Size()
	n += 1 + l + sovAcm(uint64(l))
	return n
}

func sovAcm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAcm(x uint64) (n int) {
	return sovAcm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConcreteAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcm
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
			return fmt.Errorf("proto: ConcreteAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConcreteAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Code.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAcm
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
func skipAcm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
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
				return 0, ErrInvalidLengthAcm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAcm
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
				next, err := skipAcm(dAtA[start:])
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
	ErrInvalidLengthAcm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcm   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("acm.proto", fileDescriptorAcm) }
func init() { golang_proto.RegisterFile("acm.proto", fileDescriptorAcm) }

var fileDescriptorAcm = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xeb, 0x36, 0xff, 0x7e, 0xb8, 0x95, 0xfe, 0xc5, 0x53, 0xd4, 0xc1, 0x2d, 0x88, 0xa1,
	0x03, 0x24, 0x12, 0x1f, 0x42, 0x62, 0x6b, 0x2a, 0xb1, 0x20, 0xa1, 0x2a, 0x6c, 0x6c, 0x89, 0x73,
	0x49, 0x23, 0x35, 0x71, 0x70, 0x6c, 0xa1, 0xbc, 0x09, 0x23, 0x8f, 0x82, 0x98, 0x3a, 0x32, 0x33,
	0x54, 0xa8, 0x7d, 0x11, 0x54, 0xe3, 0x96, 0x4c, 0x6c, 0x39, 0xf9, 0xdd, 0x73, 0xef, 0xd1, 0x31,
	0xee, 0x04, 0x2c, 0x75, 0x72, 0xc1, 0x25, 0x27, 0x8d, 0x80, 0xa5, 0x83, 0xd3, 0x38, 0x91, 0x73,
	0x15, 0x3a, 0x8c, 0xa7, 0x6e, 0xcc, 0x63, 0xee, 0x6a, 0x16, 0xaa, 0x47, 0xad, 0xb4, 0xd0, 0x5f,
	0x3f, 0x9e, 0x41, 0x3f, 0x07, 0x91, 0x26, 0x45, 0x91, 0xf0, 0xcc, 0xfc, 0xe9, 0x31, 0x51, 0xe6,
	0xd2, 0xf0, 0xa3, 0xf7, 0x3a, 0xfe, 0x3f, 0xe5, 0x19, 0x13, 0x20, 0x61, 0xc2, 0x18, 0x57, 0x99,
	0x24, 0x77, 0xb8, 0x35, 0x89, 0x22, 0x01, 0x45, 0x61, 0xa3, 0x11, 0x1a, 0xf7, 0xbc, 0x8b, 0xe5,
	0x6a, 0x58, 0xfb, 0x5c, 0x0d, 0x4f, 0x2a, 0xb7, 0xe7, 0x65, 0x0e, 0x62, 0x01, 0x51, 0x0c, 0xc2,
	0x0d, 0x95, 0x10, 0xfc, 0xd9, 0x35, 0x8b, 0x8d, 0xd7, 0xdf, 0x2d, 0x21, 0x97, 0xb8, 0x33, 0x53,
	0xe1, 0x22, 0x61, 0xb7, 0x50, 0xda, 0xf5, 0x11, 0x1a, 0x77, 0xcf, 0x0e, 0x1c, 0x33, 0xbc, 0x07,
	0x9e, 0xb5, 0x3d, 0xe2, 0xff, 0x4e, 0x92, 0x01, 0x6e, 0xdf, 0xc3, 0x93, 0x82, 0x8c, 0x81, 0xdd,
	0x18, 0xa1, 0xb1, 0xe5, 0xef, 0x35, 0xb1, 0x71, 0xcb, 0x0b, 0x16, 0xc1, 0x16, 0x59, 0x1a, 0xed,
	0x24, 0x39, 0xc6, 0xd6, 0x94, 0x47, 0x60, 0xff, 0xd3, 0xc9, 0xfb, 0x26, 0x79, 0xdb, 0x2b, 0x25,
	0x30, 0x1e, 0x81, 0xaf, 0x29, 0xb9, 0xc1, 0xdd, 0xd9, 0xbe, 0x98, 0xc2, 0x6e, 0xea, 0x50, 0xd4,
	0xa9, 0x94, 0x65, 0xca, 0xa8, 0x4c, 0x99, 0x84, 0x55, 0xe3, 0xb5, 0xf5, 0xf2, 0x3a, 0xac, 0x79,
	0x57, 0xcb, 0x35, 0x45, 0x1f, 0x6b, 0x8a, 0xbe, 0xd6, 0x14, 0xbd, 0x6d, 0x28, 0x5a, 0x6e, 0x28,
	0x7a, 0x38, 0xfc, 0xbb, 0xad, 0x80, 0xa5, 0x61, 0x53, 0x3f, 0xc2, 0xf9, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0x02, 0x22, 0xf4, 0xe5, 0x01, 0x00, 0x00,
}
