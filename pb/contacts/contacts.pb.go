// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/infobloxopen/atlas-contacts-app/proto/contacts.proto

/*
Package contacts is a generated protocol buffer package.

It is generated from these files:
	github.com/infobloxopen/atlas-contacts-app/proto/contacts.proto

It has these top-level messages:
	Contact
	CreateContactRequest
	CreateContactResponse
	ReadContactRequest
	ReadContactResponse
	UpdateContactRequest
	UpdateContactResponse
	DeleteContactRequest
	ListContactsResponse
	SMSRequest
*/
package contacts

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Contact struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	MiddleName   string `protobuf:"bytes,3,opt,name=middle_name,json=middleName" json:"middle_name,omitempty"`
	LastName     string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	EmailAddress string `protobuf:"bytes,5,opt,name=email_address,json=emailAddress" json:"email_address,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Contact) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contact) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Contact) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *Contact) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Contact) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

type CreateContactRequest struct {
	Payload *Contact `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *CreateContactRequest) Reset()                    { *m = CreateContactRequest{} }
func (m *CreateContactRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateContactRequest) ProtoMessage()               {}
func (*CreateContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateContactRequest) GetPayload() *Contact {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateContactResponse struct {
	Result *Contact `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *CreateContactResponse) Reset()                    { *m = CreateContactResponse{} }
func (m *CreateContactResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateContactResponse) ProtoMessage()               {}
func (*CreateContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateContactResponse) GetResult() *Contact {
	if m != nil {
		return m.Result
	}
	return nil
}

type ReadContactRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadContactRequest) Reset()                    { *m = ReadContactRequest{} }
func (m *ReadContactRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadContactRequest) ProtoMessage()               {}
func (*ReadContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadContactRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadContactResponse struct {
	Result *Contact `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *ReadContactResponse) Reset()                    { *m = ReadContactResponse{} }
func (m *ReadContactResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadContactResponse) ProtoMessage()               {}
func (*ReadContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadContactResponse) GetResult() *Contact {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateContactRequest struct {
	Payload *Contact `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *UpdateContactRequest) Reset()                    { *m = UpdateContactRequest{} }
func (m *UpdateContactRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactRequest) ProtoMessage()               {}
func (*UpdateContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateContactRequest) GetPayload() *Contact {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UpdateContactResponse struct {
	Result *Contact `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *UpdateContactResponse) Reset()                    { *m = UpdateContactResponse{} }
func (m *UpdateContactResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactResponse) ProtoMessage()               {}
func (*UpdateContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateContactResponse) GetResult() *Contact {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteContactRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteContactRequest) Reset()                    { *m = DeleteContactRequest{} }
func (m *DeleteContactRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactRequest) ProtoMessage()               {}
func (*DeleteContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteContactRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListContactsResponse struct {
	Results []*Contact `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ListContactsResponse) Reset()                    { *m = ListContactsResponse{} }
func (m *ListContactsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListContactsResponse) ProtoMessage()               {}
func (*ListContactsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListContactsResponse) GetResults() []*Contact {
	if m != nil {
		return m.Results
	}
	return nil
}

type SMSRequest struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *SMSRequest) Reset()                    { *m = SMSRequest{} }
func (m *SMSRequest) String() string            { return proto.CompactTextString(m) }
func (*SMSRequest) ProtoMessage()               {}
func (*SMSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SMSRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SMSRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Contact)(nil), "api.contacts.Contact")
	proto.RegisterType((*CreateContactRequest)(nil), "api.contacts.CreateContactRequest")
	proto.RegisterType((*CreateContactResponse)(nil), "api.contacts.CreateContactResponse")
	proto.RegisterType((*ReadContactRequest)(nil), "api.contacts.ReadContactRequest")
	proto.RegisterType((*ReadContactResponse)(nil), "api.contacts.ReadContactResponse")
	proto.RegisterType((*UpdateContactRequest)(nil), "api.contacts.UpdateContactRequest")
	proto.RegisterType((*UpdateContactResponse)(nil), "api.contacts.UpdateContactResponse")
	proto.RegisterType((*DeleteContactRequest)(nil), "api.contacts.DeleteContactRequest")
	proto.RegisterType((*ListContactsResponse)(nil), "api.contacts.ListContactsResponse")
	proto.RegisterType((*SMSRequest)(nil), "api.contacts.SMSRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Contacts service

type ContactsClient interface {
	Create(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error)
	Read(ctx context.Context, in *ReadContactRequest, opts ...grpc.CallOption) (*ReadContactResponse, error)
	Update(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error)
	Delete(ctx context.Context, in *DeleteContactRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	List(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListContactsResponse, error)
	SendSMS(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type contactsClient struct {
	cc *grpc.ClientConn
}

func NewContactsClient(cc *grpc.ClientConn) ContactsClient {
	return &contactsClient{cc}
}

func (c *contactsClient) Create(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error) {
	out := new(CreateContactResponse)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Read(ctx context.Context, in *ReadContactRequest, opts ...grpc.CallOption) (*ReadContactResponse, error) {
	out := new(ReadContactResponse)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Update(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error) {
	out := new(UpdateContactResponse)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Delete(ctx context.Context, in *DeleteContactRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) List(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListContactsResponse, error) {
	out := new(ListContactsResponse)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SendSMS(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/api.contacts.Contacts/SendSMS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contacts service

type ContactsServer interface {
	Create(context.Context, *CreateContactRequest) (*CreateContactResponse, error)
	Read(context.Context, *ReadContactRequest) (*ReadContactResponse, error)
	Update(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error)
	Delete(context.Context, *DeleteContactRequest) (*google_protobuf.Empty, error)
	List(context.Context, *google_protobuf.Empty) (*ListContactsResponse, error)
	SendSMS(context.Context, *SMSRequest) (*google_protobuf.Empty, error)
}

func RegisterContactsServer(s *grpc.Server, srv ContactsServer) {
	s.RegisterService(&_Contacts_serviceDesc, srv)
}

func _Contacts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Create(ctx, req.(*CreateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Read(ctx, req.(*ReadContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Update(ctx, req.(*UpdateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Delete(ctx, req.(*DeleteContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).List(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.contacts.Contacts/SendSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SendSMS(ctx, req.(*SMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.contacts.Contacts",
	HandlerType: (*ContactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Contacts_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Contacts_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Contacts_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Contacts_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Contacts_List_Handler,
		},
		{
			MethodName: "SendSMS",
			Handler:    _Contacts_SendSMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/infobloxopen/atlas-contacts-app/proto/contacts.proto",
}

func init() {
	proto.RegisterFile("github.com/infobloxopen/atlas-contacts-app/proto/contacts.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xd3, 0x48,
	0x18, 0x96, 0x9d, 0x34, 0x69, 0xde, 0x76, 0xab, 0xee, 0x6c, 0xda, 0xa6, 0xee, 0x56, 0x9b, 0xba,
	0xab, 0xdd, 0xaa, 0xda, 0x78, 0x96, 0x20, 0x71, 0x48, 0x0f, 0xd0, 0x0f, 0xa8, 0x84, 0x28, 0x87,
	0x44, 0xbd, 0x54, 0x42, 0x65, 0x12, 0x4f, 0xdc, 0xa9, 0x6c, 0x8f, 0xf1, 0x4c, 0x80, 0x82, 0x2a,
	0x21, 0x7e, 0x02, 0x5c, 0xf8, 0x15, 0xdc, 0xdb, 0x13, 0xff, 0x81, 0x33, 0x37, 0x2e, 0xfc, 0x0b,
	0x14, 0x8f, 0x9d, 0x0f, 0xe7, 0x43, 0x50, 0x71, 0xb3, 0xdf, 0xf7, 0x99, 0xe7, 0x79, 0x1f, 0xcf,
	0xa3, 0xd7, 0x70, 0xd7, 0x61, 0xf2, 0xac, 0xd3, 0xb4, 0x5a, 0xdc, 0xc3, 0xcc, 0x6f, 0xf3, 0xa6,
	0xcb, 0x5f, 0xf2, 0x80, 0xfa, 0x98, 0x48, 0x97, 0x88, 0x4a, 0x8b, 0xfb, 0x92, 0xb4, 0xa4, 0xa8,
	0x90, 0x20, 0xc0, 0x41, 0xc8, 0x25, 0xc7, 0x49, 0xc9, 0x8a, 0x5e, 0xd1, 0x3c, 0x09, 0x98, 0x95,
	0xd4, 0x8c, 0x35, 0x87, 0x73, 0xc7, 0xa5, 0x0a, 0xda, 0xec, 0xb4, 0x31, 0xf5, 0x02, 0x79, 0xa1,
	0xa0, 0xc6, 0x9f, 0x71, 0x93, 0x04, 0x0c, 0x13, 0xdf, 0xe7, 0x92, 0x48, 0xc6, 0xfd, 0x98, 0xc8,
	0xd8, 0x19, 0x98, 0xc4, 0xbd, 0x68, 0x4b, 0xc5, 0xd1, 0xaa, 0x38, 0xd4, 0xaf, 0x3c, 0x27, 0x2e,
	0xb3, 0x89, 0xa4, 0x78, 0xe4, 0x21, 0x3e, 0xfc, 0xdf, 0x00, 0x58, 0xbc, 0x20, 0x8e, 0x43, 0x43,
	0xcc, 0x83, 0x88, 0x7e, 0x8c, 0x54, 0x6d, 0x92, 0xe9, 0x01, 0x16, 0x87, 0x87, 0x5e, 0x8f, 0xa2,
	0xfb, 0xa2, 0xce, 0x9a, 0x1f, 0x35, 0xc8, 0xef, 0x2b, 0xbb, 0x68, 0x01, 0x74, 0x66, 0x97, 0xb4,
	0xb2, 0xb6, 0x95, 0xad, 0xeb, 0xcc, 0x46, 0xeb, 0x00, 0x6d, 0x16, 0x0a, 0x79, 0xea, 0x13, 0x8f,
	0x96, 0xf4, 0xb2, 0xb6, 0x55, 0xa8, 0x17, 0xa2, 0xca, 0x63, 0xe2, 0x51, 0xf4, 0x17, 0xcc, 0x79,
	0xcc, 0xb6, 0x5d, 0xaa, 0xfa, 0x99, 0xa8, 0x0f, 0xaa, 0x14, 0x01, 0xd6, 0xa0, 0xe0, 0x92, 0xe4,
	0x78, 0x36, 0x6a, 0xcf, 0x76, 0x0b, 0x51, 0xd3, 0x82, 0xdf, 0xa8, 0x47, 0x98, 0x7b, 0x4a, 0x6c,
	0x3b, 0xa4, 0x42, 0x94, 0x66, 0xba, 0x80, 0xbd, 0xc2, 0xf5, 0xb7, 0x4f, 0x99, 0x6c, 0xa8, 0x3f,
	0xd5, 0xea, 0xf3, 0x51, 0x7f, 0x57, 0xb5, 0x6b, 0xb9, 0xeb, 0xab, 0x55, 0x7d, 0x56, 0x33, 0x0f,
	0xa1, 0xb8, 0x1f, 0x52, 0x22, 0x69, 0x3c, 0x75, 0x9d, 0x3e, 0xeb, 0x50, 0x21, 0x11, 0x86, 0x7c,
	0x40, 0x2e, 0x5c, 0x4e, 0x94, 0x83, 0xb9, 0xea, 0x92, 0x35, 0x78, 0x95, 0x56, 0x02, 0x4f, 0x50,
	0xe6, 0x03, 0x58, 0x4a, 0x11, 0x89, 0x80, 0xfb, 0x82, 0xa2, 0x0a, 0xe4, 0x42, 0x2a, 0x3a, 0xae,
	0x9c, 0x4e, 0x14, 0x83, 0xcc, 0xbf, 0x01, 0xd5, 0x29, 0xb1, 0x53, 0xe3, 0xa4, 0xbe, 0xa5, 0x79,
	0x00, 0x7f, 0x0c, 0xa1, 0x6e, 0xa6, 0x75, 0x08, 0xc5, 0xe3, 0xc0, 0xfe, 0x35, 0xe6, 0x53, 0x44,
	0x37, 0x1b, 0xe8, 0x1f, 0x28, 0x1e, 0x50, 0x97, 0x8e, 0x0c, 0x94, 0xb6, 0x7f, 0x08, 0xc5, 0x47,
	0x4c, 0xc8, 0x18, 0x25, 0x7a, 0x72, 0x18, 0xf2, 0x8a, 0x49, 0x94, 0xb4, 0x72, 0x66, 0xca, 0xe0,
	0x31, 0xca, 0xbc, 0x03, 0xd0, 0x38, 0x6a, 0x4c, 0x90, 0x41, 0x25, 0xc8, 0x7b, 0x54, 0x08, 0xe2,
	0x24, 0x71, 0x4d, 0x5e, 0xab, 0x6f, 0x66, 0x60, 0x36, 0x51, 0x47, 0x1e, 0xe4, 0xd4, 0xd5, 0x23,
	0x33, 0x25, 0x37, 0x26, 0x59, 0xc6, 0xe6, 0x54, 0x8c, 0x32, 0x62, 0x1a, 0x6f, 0x3f, 0x7f, 0x7d,
	0xaf, 0x17, 0xcd, 0x42, 0x6f, 0xa1, 0xd4, 0x92, 0x8f, 0x8d, 0x5a, 0x90, 0xed, 0xde, 0x3d, 0x2a,
	0x0f, 0x13, 0x8d, 0xa6, 0xc6, 0xd8, 0x98, 0x82, 0x88, 0x85, 0x96, 0x23, 0xa1, 0x45, 0xb4, 0xd0,
	0x13, 0xc2, 0xaf, 0x99, 0x7d, 0x89, 0x5e, 0x41, 0x4e, 0xdd, 0x68, 0xda, 0xd3, 0xb8, 0xc0, 0xa4,
	0x3d, 0x8d, 0xcd, 0x82, 0xf9, 0x6f, 0x24, 0xb5, 0x61, 0x2c, 0x0f, 0x48, 0xc5, 0x9e, 0x2c, 0x66,
	0x5f, 0xf6, 0x0d, 0x32, 0xc8, 0xa9, 0x14, 0xa4, 0xb5, 0xc7, 0x65, 0xc3, 0x58, 0xb6, 0xd4, 0xe2,
	0xb4, 0x92, 0xad, 0x6a, 0xdd, 0xef, 0x6e, 0x55, 0x73, 0xf3, 0xfa, 0x6a, 0xb5, 0xd0, 0xdb, 0x46,
	0xca, 0xe6, 0x76, 0xda, 0xe6, 0x31, 0x64, 0xbb, 0x41, 0x42, 0x13, 0x48, 0x8c, 0xd4, 0x00, 0xe3,
	0x42, 0x67, 0xfe, 0x1e, 0x71, 0xcf, 0xa1, 0xfe, 0x5d, 0xa1, 0x13, 0xc8, 0x37, 0xa8, 0x6f, 0x37,
	0x8e, 0x1a, 0xa8, 0x34, 0xcc, 0xd0, 0x4f, 0xdb, 0xc4, 0xc1, 0xd7, 0x23, 0xbe, 0x15, 0x13, 0x0d,
	0xcf, 0x8a, 0x85, 0x27, 0x6a, 0xda, 0xb6, 0x11, 0x6f, 0xae, 0xbd, 0x2f, 0xda, 0xbb, 0xdd, 0x0f,
	0x1a, 0xf2, 0xfb, 0x41, 0x34, 0x9f, 0xc0, 0xc2, 0x43, 0x7e, 0xe6, 0x97, 0xf7, 0xa8, 0x4b, 0x3c,
	0x12, 0xb2, 0x16, 0xaa, 0x9e, 0x49, 0x19, 0x88, 0x1a, 0xc6, 0x3f, 0xfe, 0x23, 0x33, 0x56, 0xce,
	0x9b, 0xc9, 0xf9, 0x7b, 0x09, 0xb6, 0x7b, 0xb0, 0x9a, 0xb9, 0x65, 0xfd, 0xbf, 0xad, 0x6b, 0x7a,
	0x75, 0x91, 0x04, 0x81, 0xcb, 0x5a, 0xd1, 0x5f, 0x03, 0x9f, 0x0b, 0xee, 0xd7, 0x46, 0x2a, 0x27,
	0x3f, 0xf5, 0xeb, 0x6c, 0xf6, 0xac, 0xee, 0x24, 0x0f, 0xcd, 0x5c, 0xf4, 0x59, 0x6e, 0x7f, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x45, 0x2a, 0xa6, 0x84, 0x07, 0x00, 0x00,
}
