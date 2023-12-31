// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eywacontract/eywacontract/handshake_type.proto

package types

import (
	fmt "fmt"
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

type HandshakeType struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Receiver      string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender        string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	ServerAddress string `protobuf:"bytes,4,opt,name=serverAddress,proto3" json:"serverAddress,omitempty"`
	RoomId        string `protobuf:"bytes,5,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Time          string `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
}

func (m *HandshakeType) Reset()         { *m = HandshakeType{} }
func (m *HandshakeType) String() string { return proto.CompactTextString(m) }
func (*HandshakeType) ProtoMessage()    {}
func (*HandshakeType) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9837696f113f05, []int{0}
}
func (m *HandshakeType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandshakeType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandshakeType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandshakeType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeType.Merge(m, src)
}
func (m *HandshakeType) XXX_Size() int {
	return m.Size()
}
func (m *HandshakeType) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeType.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeType proto.InternalMessageInfo

func (m *HandshakeType) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *HandshakeType) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *HandshakeType) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *HandshakeType) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *HandshakeType) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *HandshakeType) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*HandshakeType)(nil), "eywacontract.eywacontract.HandshakeType")
}

func init() {
	proto.RegisterFile("eywacontract/eywacontract/handshake_type.proto", fileDescriptor_2f9837696f113f05)
}

var fileDescriptor_2f9837696f113f05 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xad, 0x2c, 0x4f,
	0x4c, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x47, 0xe1, 0x64, 0x24, 0xe6, 0xa5, 0x14,
	0x67, 0x24, 0x66, 0xa7, 0xc6, 0x97, 0x54, 0x16, 0xa4, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0x49, 0x22, 0x2b, 0x41, 0xd1, 0xac, 0xb4, 0x9a, 0x91, 0x8b, 0xd7, 0x03, 0xa6, 0x27, 0xa4, 0xb2,
	0x20, 0x55, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc2, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b, 0x2d, 0x92,
	0x60, 0x02, 0x4b, 0xc0, 0xf9, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12,
	0xcc, 0x60, 0x19, 0x28, 0x4f, 0x48, 0x85, 0x8b, 0xb7, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0xc8, 0x31,
	0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x05, 0x2c, 0x8d, 0x2a, 0x08, 0xd2, 0x5d, 0x94, 0x9f,
	0x9f, 0xeb, 0x99, 0x22, 0xc1, 0x0a, 0xd1, 0x0d, 0xe1, 0x09, 0x09, 0x71, 0xb1, 0x94, 0x64, 0xe6,
	0xa6, 0x4a, 0xb0, 0x81, 0x45, 0xc1, 0x6c, 0x27, 0x9b, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0x52, 0x02, 0xf9, 0x4a, 0x17, 0x1e, 0x0c, 0x15, 0xa8, 0xa1, 0x02, 0x0a, 0x8b,
	0xe2, 0x24, 0x36, 0x70, 0x68, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x18, 0xd4, 0xd3,
	0x3f, 0x01, 0x00, 0x00,
}

func (m *HandshakeType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakeType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandshakeType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintHandshakeType(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RoomId) > 0 {
		i -= len(m.RoomId)
		copy(dAtA[i:], m.RoomId)
		i = encodeVarintHandshakeType(dAtA, i, uint64(len(m.RoomId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ServerAddress) > 0 {
		i -= len(m.ServerAddress)
		copy(dAtA[i:], m.ServerAddress)
		i = encodeVarintHandshakeType(dAtA, i, uint64(len(m.ServerAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintHandshakeType(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintHandshakeType(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintHandshakeType(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHandshakeType(dAtA []byte, offset int, v uint64) int {
	offset -= sovHandshakeType(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HandshakeType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovHandshakeType(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovHandshakeType(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovHandshakeType(uint64(l))
	}
	l = len(m.ServerAddress)
	if l > 0 {
		n += 1 + l + sovHandshakeType(uint64(l))
	}
	l = len(m.RoomId)
	if l > 0 {
		n += 1 + l + sovHandshakeType(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovHandshakeType(uint64(l))
	}
	return n
}

func sovHandshakeType(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHandshakeType(x uint64) (n int) {
	return sovHandshakeType(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HandshakeType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshakeType
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
			return fmt.Errorf("proto: HandshakeType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakeType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshakeType
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
				return ErrInvalidLengthHandshakeType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshakeType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshakeType
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
				return ErrInvalidLengthHandshakeType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshakeType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshakeType
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
				return ErrInvalidLengthHandshakeType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshakeType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshakeType
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
				return ErrInvalidLengthHandshakeType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshakeType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshakeType
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
				return ErrInvalidLengthHandshakeType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshakeType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshakeType
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
				return ErrInvalidLengthHandshakeType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshakeType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshakeType(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshakeType
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
func skipHandshakeType(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandshakeType
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
					return 0, ErrIntOverflowHandshakeType
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
					return 0, ErrIntOverflowHandshakeType
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
				return 0, ErrInvalidLengthHandshakeType
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHandshakeType
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHandshakeType
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHandshakeType        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandshakeType          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHandshakeType = fmt.Errorf("proto: unexpected end of group")
)
