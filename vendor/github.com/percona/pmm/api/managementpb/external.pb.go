// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/external.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddExternalRequest struct {
	// Node identifier on which a external exporter is been running. Required.
	RunsOnNodeId string `protobuf:"bytes,1,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// HTTP basic auth username for collecting metrics.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// HTTP basic auth password for collecting metrics.
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `protobuf:"bytes,5,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `protobuf:"bytes,6,opt,name=metrics_path,json=metricsPath,proto3" json:"metrics_path,omitempty"`
	// Listen port for scraping metrics.
	ListenPort uint32 `protobuf:"varint,7,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	// Node identifier on which a external service is been running. Required.
	NodeId string `protobuf:"bytes,8,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,9,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,10,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,11,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels         map[string]string `protobuf:"bytes,15,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddExternalRequest) Reset()         { *m = AddExternalRequest{} }
func (m *AddExternalRequest) String() string { return proto.CompactTextString(m) }
func (*AddExternalRequest) ProtoMessage()    {}
func (*AddExternalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab8accc93fe1440, []int{0}
}

func (m *AddExternalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddExternalRequest.Unmarshal(m, b)
}
func (m *AddExternalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddExternalRequest.Marshal(b, m, deterministic)
}
func (m *AddExternalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddExternalRequest.Merge(m, src)
}
func (m *AddExternalRequest) XXX_Size() int {
	return xxx_messageInfo_AddExternalRequest.Size(m)
}
func (m *AddExternalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddExternalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddExternalRequest proto.InternalMessageInfo

func (m *AddExternalRequest) GetRunsOnNodeId() string {
	if m != nil {
		return m.RunsOnNodeId
	}
	return ""
}

func (m *AddExternalRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AddExternalRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddExternalRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddExternalRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AddExternalRequest) GetMetricsPath() string {
	if m != nil {
		return m.MetricsPath
	}
	return ""
}

func (m *AddExternalRequest) GetListenPort() uint32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *AddExternalRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddExternalRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AddExternalRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *AddExternalRequest) GetReplicationSet() string {
	if m != nil {
		return m.ReplicationSet
	}
	return ""
}

func (m *AddExternalRequest) GetCustomLabels() map[string]string {
	if m != nil {
		return m.CustomLabels
	}
	return nil
}

type AddExternalResponse struct {
	Service              *inventorypb.ExternalService  `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ExternalExporter     *inventorypb.ExternalExporter `protobuf:"bytes,2,opt,name=external_exporter,json=externalExporter,proto3" json:"external_exporter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AddExternalResponse) Reset()         { *m = AddExternalResponse{} }
func (m *AddExternalResponse) String() string { return proto.CompactTextString(m) }
func (*AddExternalResponse) ProtoMessage()    {}
func (*AddExternalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab8accc93fe1440, []int{1}
}

func (m *AddExternalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddExternalResponse.Unmarshal(m, b)
}
func (m *AddExternalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddExternalResponse.Marshal(b, m, deterministic)
}
func (m *AddExternalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddExternalResponse.Merge(m, src)
}
func (m *AddExternalResponse) XXX_Size() int {
	return xxx_messageInfo_AddExternalResponse.Size(m)
}
func (m *AddExternalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddExternalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddExternalResponse proto.InternalMessageInfo

func (m *AddExternalResponse) GetService() *inventorypb.ExternalService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *AddExternalResponse) GetExternalExporter() *inventorypb.ExternalExporter {
	if m != nil {
		return m.ExternalExporter
	}
	return nil
}

func init() {
	proto.RegisterType((*AddExternalRequest)(nil), "management.AddExternalRequest")
	proto.RegisterMapType((map[string]string)(nil), "management.AddExternalRequest.CustomLabelsEntry")
	proto.RegisterType((*AddExternalResponse)(nil), "management.AddExternalResponse")
}

func init() {
	proto.RegisterFile("managementpb/external.proto", fileDescriptor_dab8accc93fe1440)
}

var fileDescriptor_dab8accc93fe1440 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xcd, 0xb6, 0xdd, 0x6d, 0x4f, 0xb6, 0x5f, 0xa3, 0xc8, 0x90, 0x8a, 0xbb, 0xee, 0x85,
	0x56, 0xa5, 0x89, 0xae, 0x22, 0xd2, 0x1b, 0x69, 0xa5, 0xa0, 0x20, 0xb5, 0xa4, 0x08, 0xe2, 0x4d,
	0x98, 0x4d, 0x0e, 0xbb, 0x43, 0x93, 0x99, 0x38, 0x33, 0x49, 0xdb, 0xbb, 0xd2, 0x57, 0x10, 0xdf,
	0xc2, 0xb7, 0xf1, 0x01, 0x0a, 0xc5, 0x07, 0x91, 0x7c, 0xb5, 0xbb, 0x56, 0xbd, 0x9b, 0x73, 0x7e,
	0xff, 0x39, 0x67, 0x4e, 0xfe, 0x33, 0x81, 0x8d, 0x84, 0x09, 0x36, 0xc6, 0x04, 0x85, 0x49, 0x47,
	0x1e, 0x9e, 0x18, 0x54, 0x82, 0xc5, 0x6e, 0xaa, 0xa4, 0x91, 0x04, 0xae, 0xa1, 0xf3, 0x6a, 0xcc,
	0xcd, 0x24, 0x1b, 0xb9, 0xa1, 0x4c, 0xbc, 0xe4, 0x98, 0x9b, 0x23, 0x79, 0xec, 0x8d, 0xe5, 0x56,
	0x29, 0xdc, 0xca, 0x59, 0xcc, 0x23, 0x66, 0xa4, 0xd2, 0xde, 0xd5, 0xb2, 0xaa, 0xe1, 0xdc, 0x1b,
	0x4b, 0x39, 0x8e, 0xd1, 0x63, 0x29, 0xf7, 0x98, 0x10, 0xd2, 0x30, 0xc3, 0xa5, 0xd0, 0x35, 0xa5,
	0x5c, 0xe4, 0x28, 0x8c, 0x54, 0xa7, 0xe9, 0xc8, 0x63, 0x63, 0x14, 0xa6, 0x21, 0xce, 0x34, 0xd1,
	0xa8, 0x72, 0x1e, 0x62, 0xcd, 0x06, 0x3f, 0xe6, 0x81, 0xec, 0x44, 0xd1, 0x5e, 0x7d, 0x5a, 0x1f,
	0xbf, 0x66, 0xa8, 0x0d, 0xd9, 0x82, 0x55, 0x95, 0x09, 0x1d, 0x48, 0x11, 0x08, 0x19, 0x61, 0xc0,
	0x23, 0x6a, 0xf5, 0xad, 0xcd, 0xa5, 0xdd, 0xf6, 0xe5, 0x45, 0xaf, 0xf5, 0xd9, 0xf2, 0xbb, 0x05,
	0xfe, 0x28, 0xf6, 0x65, 0x84, 0xef, 0x23, 0xf2, 0x18, 0xba, 0x75, 0xdd, 0x40, 0xb0, 0x04, 0x69,
	0x6b, 0x46, 0x6b, 0xd7, 0x6c, 0x9f, 0x25, 0x48, 0x1c, 0x58, 0xcc, 0x74, 0xd1, 0x2b, 0x41, 0x3a,
	0x57, 0xc8, 0xfc, 0xab, 0xb8, 0x60, 0x29, 0xd3, 0xfa, 0x58, 0xaa, 0x88, 0xce, 0x57, 0xac, 0x89,
	0xc9, 0x5d, 0x68, 0xeb, 0x70, 0x82, 0x09, 0xd2, 0x85, 0x92, 0xd4, 0x11, 0x79, 0x00, 0xdd, 0x04,
	0x8d, 0xe2, 0xa1, 0x0e, 0x52, 0x66, 0x26, 0xb4, 0x5d, 0x52, 0xbb, 0xce, 0x1d, 0x30, 0x33, 0x21,
	0x4f, 0xc1, 0x8e, 0xb9, 0x36, 0x28, 0x82, 0x54, 0x2a, 0x43, 0x3b, 0x7d, 0x6b, 0x73, 0x79, 0x17,
	0x2e, 0x2f, 0x7a, 0xed, 0xb5, 0x5b, 0xf4, 0xec, 0x6c, 0xde, 0x87, 0x0a, 0x1f, 0x48, 0x65, 0x48,
	0x0f, 0x3a, 0xcd, 0xc4, 0x8b, 0x33, 0x53, 0xb4, 0x45, 0x35, 0x6b, 0x1f, 0x6c, 0x14, 0x39, 0x57,
	0x52, 0x14, 0x66, 0xd2, 0xa5, 0xaa, 0xdf, 0x54, 0x8a, 0x50, 0xe8, 0x84, 0x71, 0xa6, 0x0d, 0x2a,
	0x0a, 0x25, 0x6d, 0x42, 0xf2, 0x08, 0x56, 0x15, 0xa6, 0x31, 0x0f, 0x4b, 0xe7, 0x02, 0x8d, 0x86,
	0xda, 0xa5, 0x62, 0x65, 0x2a, 0x7d, 0x88, 0x86, 0x7c, 0x82, 0xe5, 0x30, 0xd3, 0x46, 0x26, 0x41,
	0xcc, 0x46, 0x18, 0x6b, 0xba, 0xda, 0x9f, 0xdb, 0xb4, 0x87, 0xcf, 0xdc, 0xeb, 0x6b, 0xe4, 0xde,
	0xb4, 0xcd, 0x7d, 0x5b, 0xee, 0xf9, 0x50, 0x6e, 0xd9, 0x13, 0x46, 0x9d, 0xfa, 0xdd, 0x70, 0x2a,
	0xe5, 0xbc, 0x81, 0xf5, 0x1b, 0x12, 0xb2, 0x06, 0x73, 0x47, 0x78, 0x5a, 0xf9, 0xeb, 0x17, 0x4b,
	0x72, 0x07, 0x16, 0x72, 0x16, 0x67, 0xb5, 0x8f, 0x7e, 0x15, 0x6c, 0xb7, 0x5e, 0x5b, 0x83, 0xef,
	0x16, 0xdc, 0x9e, 0xe9, 0xab, 0x53, 0x29, 0x34, 0x92, 0x97, 0xd0, 0xa9, 0x4d, 0x2e, 0xeb, 0xd8,
	0x43, 0xc7, 0xbd, 0xba, 0x74, 0x6e, 0xa3, 0x3e, 0xac, 0x14, 0x7e, 0x23, 0x25, 0xef, 0x60, 0xbd,
	0x79, 0x26, 0x01, 0x9e, 0x14, 0xe6, 0xa0, 0x2a, 0x7b, 0xda, 0xc3, 0x8d, 0xbf, 0xec, 0xdf, 0xab,
	0x25, 0xfe, 0x1a, 0xfe, 0x91, 0x19, 0x9e, 0x5b, 0xb0, 0xd8, 0xc8, 0x48, 0x0e, 0xf6, 0xd4, 0x19,
	0xc9, 0xfd, 0xff, 0x7f, 0x34, 0xa7, 0xf7, 0x4f, 0x5e, 0x0d, 0x37, 0x78, 0x78, 0xfe, 0xf3, 0xd7,
	0xb7, 0x56, 0x7f, 0xb0, 0xe1, 0xe5, 0xcf, 0xbd, 0x6b, 0xad, 0xd7, 0x08, 0xbd, 0x9d, 0x28, 0xda,
	0xb6, 0x9e, 0xec, 0xae, 0x7c, 0xe9, 0x4e, 0xff, 0x02, 0x46, 0xed, 0xf2, 0x89, 0xbd, 0xf8, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x0e, 0x60, 0x27, 0x81, 0x19, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExternalClient is the client API for External service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExternalClient interface {
	// AddExternal adds external service and adds external exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
	AddExternal(ctx context.Context, in *AddExternalRequest, opts ...grpc.CallOption) (*AddExternalResponse, error)
}

type externalClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalClient(cc grpc.ClientConnInterface) ExternalClient {
	return &externalClient{cc}
}

func (c *externalClient) AddExternal(ctx context.Context, in *AddExternalRequest, opts ...grpc.CallOption) (*AddExternalResponse, error) {
	out := new(AddExternalResponse)
	err := c.cc.Invoke(ctx, "/management.External/AddExternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServer is the server API for External service.
type ExternalServer interface {
	// AddExternal adds external service and adds external exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
	AddExternal(context.Context, *AddExternalRequest) (*AddExternalResponse, error)
}

// UnimplementedExternalServer can be embedded to have forward compatible implementations.
type UnimplementedExternalServer struct {
}

func (*UnimplementedExternalServer) AddExternal(ctx context.Context, req *AddExternalRequest) (*AddExternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExternal not implemented")
}

func RegisterExternalServer(s *grpc.Server, srv ExternalServer) {
	s.RegisterService(&_External_serviceDesc, srv)
}

func _External_AddExternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServer).AddExternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.External/AddExternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServer).AddExternal(ctx, req.(*AddExternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _External_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.External",
	HandlerType: (*ExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddExternal",
			Handler:    _External_AddExternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/external.proto",
}
