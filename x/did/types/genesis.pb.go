// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/did/v2/genesis.proto

package types

import (
	fmt "fmt"
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

// DidDocVersionSet contains all versions of DID Documents and their metadata for a given DID.
// The latest version of the DID Document set is stored in the latest_version field.
type DidDocVersionSet struct {
	// Latest version of the DID Document set
	LatestVersion string `protobuf:"bytes,1,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	// All versions of the DID Document set
	DidDocs []*DidDocWithMetadata `protobuf:"bytes,2,rep,name=did_docs,json=didDocs,proto3" json:"did_docs,omitempty"`
}

func (m *DidDocVersionSet) Reset()         { *m = DidDocVersionSet{} }
func (m *DidDocVersionSet) String() string { return proto.CompactTextString(m) }
func (*DidDocVersionSet) ProtoMessage()    {}
func (*DidDocVersionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_83613517e395af68, []int{0}
}
func (m *DidDocVersionSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidDocVersionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidDocVersionSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidDocVersionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidDocVersionSet.Merge(m, src)
}
func (m *DidDocVersionSet) XXX_Size() int {
	return m.Size()
}
func (m *DidDocVersionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DidDocVersionSet.DiscardUnknown(m)
}

var xxx_messageInfo_DidDocVersionSet proto.InternalMessageInfo

func (m *DidDocVersionSet) GetLatestVersion() string {
	if m != nil {
		return m.LatestVersion
	}
	return ""
}

func (m *DidDocVersionSet) GetDidDocs() []*DidDocWithMetadata {
	if m != nil {
		return m.DidDocs
	}
	return nil
}

// GenesisState defines the cheqd DID module's genesis state.
type GenesisState struct {
	// Namespace for the DID module
	// Example: mainnet, testnet, local
	DidNamespace string `protobuf:"bytes,1,opt,name=did_namespace,json=didNamespace,proto3" json:"did_namespace,omitempty"`
	// All DID Document version sets (contains all versions of all DID Documents)
	VersionSets []*DidDocVersionSet `protobuf:"bytes,2,rep,name=version_sets,json=versionSets,proto3" json:"version_sets,omitempty"`
	// Fee parameters for the DID module
	// Defines fixed fees and burn percentage for each DID operation type (create, update, delete)
	FeeParams *FeeParams `protobuf:"bytes,3,opt,name=fee_params,json=feeParams,proto3" json:"fee_params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_83613517e395af68, []int{1}
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

func (m *GenesisState) GetDidNamespace() string {
	if m != nil {
		return m.DidNamespace
	}
	return ""
}

func (m *GenesisState) GetVersionSets() []*DidDocVersionSet {
	if m != nil {
		return m.VersionSets
	}
	return nil
}

func (m *GenesisState) GetFeeParams() *FeeParams {
	if m != nil {
		return m.FeeParams
	}
	return nil
}

func init() {
	proto.RegisterType((*DidDocVersionSet)(nil), "cheqd.did.v2.DidDocVersionSet")
	proto.RegisterType((*GenesisState)(nil), "cheqd.did.v2.GenesisState")
}

func init() { proto.RegisterFile("cheqd/did/v2/genesis.proto", fileDescriptor_83613517e395af68) }

var fileDescriptor_83613517e395af68 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xed, 0x7e, 0x85, 0x4f, 0xbb, 0x4d, 0x45, 0x72, 0xd0, 0xda, 0xc3, 0x12, 0x2a, 0x62, 0x2f,
	0x26, 0x10, 0xc1, 0x8b, 0xa7, 0x4a, 0xd1, 0x93, 0x22, 0x29, 0x28, 0x78, 0x09, 0xdb, 0xcc, 0xb4,
	0x5d, 0xb0, 0xd9, 0xd8, 0x1d, 0x83, 0xfe, 0x0b, 0x7f, 0x89, 0xbf, 0xc3, 0x63, 0x8f, 0x1e, 0xa5,
	0xfd, 0x23, 0xe2, 0x6e, 0xab, 0x16, 0xbc, 0x0c, 0xc3, 0x7b, 0x8f, 0xf7, 0x66, 0x1e, 0x6f, 0x65,
	0x63, 0x7c, 0x80, 0x08, 0x14, 0x44, 0x65, 0x1c, 0x8d, 0x30, 0x47, 0xa3, 0x4c, 0x58, 0x4c, 0x35,
	0x69, 0xdf, 0xb3, 0x5c, 0x08, 0x0a, 0xc2, 0x32, 0x6e, 0xed, 0xad, 0x29, 0x41, 0x01, 0xe8, 0xcc,
	0x09, 0x5b, 0x3b, 0x6b, 0xd4, 0x10, 0xd1, 0xe1, 0xed, 0x92, 0x6f, 0xf7, 0x14, 0xf4, 0x74, 0x76,
	0x83, 0x53, 0xa3, 0x74, 0xde, 0x47, 0xf2, 0x0f, 0xf8, 0xd6, 0xbd, 0x24, 0x34, 0x94, 0x96, 0x0e,
	0x6c, 0xb2, 0x80, 0x75, 0x6a, 0x49, 0xc3, 0xa1, 0x4b, 0xa5, 0x7f, 0xca, 0x37, 0x41, 0x41, 0x0a,
	0x3a, 0x33, 0xcd, 0x7f, 0x41, 0xb5, 0x53, 0x8f, 0x83, 0xf0, 0xf7, 0x39, 0xa1, 0x33, 0xbe, 0x55,
	0x34, 0xbe, 0x44, 0x92, 0x20, 0x49, 0x26, 0x1b, 0x60, 0x31, 0xd3, 0x7e, 0x65, 0xdc, 0xbb, 0x70,
	0xaf, 0xf4, 0x49, 0x12, 0xfa, 0xfb, 0xbc, 0xf1, 0xe5, 0x96, 0xcb, 0x09, 0x9a, 0x42, 0x66, 0xb8,
	0xcc, 0xf4, 0x40, 0xc1, 0xd5, 0x0a, 0xf3, 0xbb, 0xdc, 0x5b, 0x9e, 0x94, 0x1a, 0xa4, 0x55, 0xac,
	0xf8, 0x2b, 0xf6, 0xe7, 0x9f, 0xa4, 0x5e, 0x7e, 0xef, 0xc6, 0x3f, 0xe1, 0x7c, 0x88, 0x98, 0x16,
	0x72, 0x2a, 0x27, 0xa6, 0x59, 0x0d, 0x58, 0xa7, 0x1e, 0xef, 0xae, 0x1b, 0x9c, 0x23, 0x5e, 0x5b,
	0x3a, 0xa9, 0x0d, 0x57, 0xeb, 0x59, 0xf7, 0x6d, 0x2e, 0xd8, 0x6c, 0x2e, 0xd8, 0xc7, 0x5c, 0xb0,
	0x97, 0x85, 0xa8, 0xcc, 0x16, 0xa2, 0xf2, 0xbe, 0x10, 0x95, 0xbb, 0xc3, 0x91, 0xa2, 0xf1, 0xe3,
	0x20, 0xcc, 0xf4, 0x24, 0x72, 0x2d, 0xdb, 0x79, 0x94, 0x6b, 0xc0, 0xe8, 0xc9, 0x56, 0x4e, 0xcf,
	0x05, 0x9a, 0xc1, 0x7f, 0x5b, 0xf9, 0xf1, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xeb, 0xc7,
	0x91, 0xd1, 0x01, 0x00, 0x00,
}

func (m *DidDocVersionSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidDocVersionSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidDocVersionSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DidDocs) > 0 {
		for iNdEx := len(m.DidDocs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidDocs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.LatestVersion) > 0 {
		i -= len(m.LatestVersion)
		copy(dAtA[i:], m.LatestVersion)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.LatestVersion)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if m.FeeParams != nil {
		{
			size, err := m.FeeParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VersionSets) > 0 {
		for iNdEx := len(m.VersionSets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VersionSets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DidNamespace) > 0 {
		i -= len(m.DidNamespace)
		copy(dAtA[i:], m.DidNamespace)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidNamespace)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *DidDocVersionSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LatestVersion)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.DidDocs) > 0 {
		for _, e := range m.DidDocs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DidNamespace)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.VersionSets) > 0 {
		for _, e := range m.VersionSets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.FeeParams != nil {
		l = m.FeeParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidDocVersionSet) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DidDocVersionSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidDocVersionSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocs", wireType)
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
			m.DidDocs = append(m.DidDocs, &DidDocWithMetadata{})
			if err := m.DidDocs[len(m.DidDocs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
				return fmt.Errorf("proto: wrong wireType = %d for field DidNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionSets", wireType)
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
			m.VersionSets = append(m.VersionSets, &DidDocVersionSet{})
			if err := m.VersionSets[len(m.VersionSets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeParams", wireType)
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
			if m.FeeParams == nil {
				m.FeeParams = &FeeParams{}
			}
			if err := m.FeeParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
