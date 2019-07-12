// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/ext_authz/v2/ext_authz.proto

package v2

import (
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// External Authorization filter calls out to an external service over the
// gRPC Authorization API defined by
// :ref:`CheckRequest <envoy_api_msg_service.auth.v2.CheckRequest>`.
// A failed check will cause this filter to close the TCP connection.
type ExtAuthz struct {
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The external authorization gRPC service configuration.
	// The default timeout is set to 200ms by this filter.
	GrpcService *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. When it is set to true, Envoy will also allow traffic in case of
	// communication failure between authorization service and the proxy.
	// Defaults to false.
	FailureModeAllow     bool     `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec2615c2696024a, []int{0}
}
func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return m.Size()
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

func (m *ExtAuthz) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.network.ext_authz.v2.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/ext_authz/v2/ext_authz.proto", fileDescriptor_3ec2615c2696024a)
}

var fileDescriptor_3ec2615c2696024a = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x26, 0xb2, 0x65, 0x0a, 0xd2, 0x8b, 0x63, 0x87, 0x52, 0xc4, 0x43, 0x11, 0x49,
	0xa0, 0x82, 0x78, 0xdd, 0x40, 0x3c, 0x09, 0xa5, 0x9e, 0xf4, 0x52, 0x62, 0xfb, 0x5a, 0x83, 0xb5,
	0x09, 0x59, 0x96, 0x75, 0x7e, 0x21, 0xbf, 0x83, 0x27, 0x8f, 0x1e, 0xfd, 0x08, 0xd2, 0x9b, 0xdf,
	0x42, 0xd2, 0x6c, 0xd4, 0xa3, 0xb7, 0xbc, 0xfc, 0xdf, 0xef, 0xff, 0x7f, 0xef, 0xe1, 0x2b, 0xa8,
	0x8d, 0xd8, 0xd0, 0x4c, 0xd4, 0x05, 0x2f, 0x69, 0xc1, 0x2b, 0x0d, 0x8a, 0xd6, 0xa0, 0xd7, 0x42,
	0x3d, 0x53, 0x68, 0x74, 0xca, 0x56, 0xfa, 0xe9, 0x95, 0x9a, 0xa8, 0x2f, 0x88, 0x54, 0x42, 0x0b,
	0x2f, 0xec, 0x48, 0xe2, 0x48, 0xe2, 0x48, 0xb2, 0x25, 0x49, 0xdf, 0x6c, 0xa2, 0xd9, 0xa9, 0xcb,
	0x60, 0x92, 0x5b, 0x9f, 0x4c, 0x28, 0xa0, 0xa5, 0x92, 0x59, 0xba, 0x04, 0x65, 0x78, 0x06, 0xce,
	0x6f, 0x76, 0x6c, 0x58, 0xc5, 0x73, 0xa6, 0x81, 0xee, 0x1e, 0x4e, 0x38, 0x79, 0x43, 0x78, 0x74,
	0xdd, 0xe8, 0xb9, 0xb5, 0xf3, 0xce, 0xf0, 0x64, 0xa9, 0x99, 0x4e, 0xa5, 0x82, 0x82, 0x37, 0x53,
	0x14, 0xa0, 0x70, 0xbc, 0x18, 0xbf, 0xff, 0x7c, 0x0c, 0xf7, 0xd4, 0x20, 0x40, 0x09, 0xb6, 0x6a,
	0xdc, 0x89, 0xde, 0x1c, 0x1f, 0xfc, 0xcd, 0x99, 0x0e, 0x02, 0x14, 0x4e, 0x22, 0x9f, 0xb8, 0xc1,
	0x99, 0xe4, 0xc4, 0x44, 0xc4, 0x8e, 0x43, 0x6e, 0x94, 0xcc, 0xee, 0x5c, 0x57, 0x32, 0x29, 0xfb,
	0xc2, 0x3b, 0xc7, 0x5e, 0xc1, 0x78, 0xb5, 0x52, 0x90, 0xbe, 0x88, 0x1c, 0x52, 0x56, 0x55, 0x62,
	0x3d, 0x1d, 0x06, 0x28, 0x1c, 0x25, 0x47, 0x5b, 0xe5, 0x56, 0xe4, 0x30, 0xb7, 0xff, 0x8b, 0xfb,
	0xcf, 0xd6, 0x47, 0x5f, 0xad, 0x8f, 0xbe, 0x5b, 0x1f, 0xe1, 0x4b, 0x2e, 0x5c, 0x94, 0x54, 0xa2,
	0xd9, 0x90, 0xff, 0x9e, 0x6b, 0x71, 0xb8, 0x5b, 0x36, 0xb6, 0xeb, 0xc7, 0xe8, 0x61, 0x60, 0xa2,
	0xc7, 0xfd, 0xee, 0x16, 0x17, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x17, 0xc6, 0x56, 0xb0,
	0x01, 0x00, 0x00,
}

func (m *ExtAuthz) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtAuthz) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExtAuthz(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if m.GrpcService != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExtAuthz(dAtA, i, uint64(m.GrpcService.Size()))
		n1, err := m.GrpcService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FailureModeAllow {
		dAtA[i] = 0x18
		i++
		if m.FailureModeAllow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintExtAuthz(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExtAuthz) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovExtAuthz(uint64(l))
	}
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovExtAuthz(uint64(l))
	}
	if m.FailureModeAllow {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExtAuthz(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExtAuthz(x uint64) (n int) {
	return sovExtAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtAuthz) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtAuthz
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
			return fmt.Errorf("proto: ExtAuthz: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtAuthz: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
				return ErrInvalidLengthExtAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
				return ErrInvalidLengthExtAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &core.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeAllow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
			m.FailureModeAllow = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipExtAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExtAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtAuthz
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
					return 0, ErrIntOverflowExtAuthz
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
					return 0, ErrIntOverflowExtAuthz
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
				return 0, ErrInvalidLengthExtAuthz
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthExtAuthz
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExtAuthz
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
				next, err := skipExtAuthz(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthExtAuthz
				}
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
	ErrInvalidLengthExtAuthz = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtAuthz   = fmt.Errorf("proto: integer overflow")
)
