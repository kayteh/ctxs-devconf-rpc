// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blog.proto

package blog

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PostList struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostList) Reset()         { *m = PostList{} }
func (m *PostList) String() string { return proto.CompactTextString(m) }
func (*PostList) ProtoMessage()    {}
func (*PostList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{0}
}

func (m *PostList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostList.Unmarshal(m, b)
}
func (m *PostList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostList.Marshal(b, m, deterministic)
}
func (m *PostList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostList.Merge(m, src)
}
func (m *PostList) XXX_Size() int {
	return xxx_messageInfo_PostList.Size(m)
}
func (m *PostList) XXX_DiscardUnknown() {
	xxx_messageInfo_PostList.DiscardUnknown(m)
}

var xxx_messageInfo_PostList proto.InternalMessageInfo

func (m *PostList) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{1}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type PostQuery struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostQuery) Reset()         { *m = PostQuery{} }
func (m *PostQuery) String() string { return proto.CompactTextString(m) }
func (*PostQuery) ProtoMessage()    {}
func (*PostQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{2}
}

func (m *PostQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostQuery.Unmarshal(m, b)
}
func (m *PostQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostQuery.Marshal(b, m, deterministic)
}
func (m *PostQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostQuery.Merge(m, src)
}
func (m *PostQuery) XXX_Size() int {
	return xxx_messageInfo_PostQuery.Size(m)
}
func (m *PostQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PostQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PostQuery proto.InternalMessageInfo

func (m *PostQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*PostList)(nil), "blog.PostList")
	proto.RegisterType((*Post)(nil), "blog.Post")
	proto.RegisterType((*PostQuery)(nil), "blog.PostQuery")
}

func init() { proto.RegisterFile("blog/blog.proto", fileDescriptor_e737bdbda5671783) }

var fileDescriptor_e737bdbda5671783 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x4c, 0xd2, 0x54, 0xed, 0x53, 0x5a, 0x58, 0x44, 0x42, 0x7a, 0x09, 0x7b, 0xd0, 0x08, 0x76,
	0x03, 0xf5, 0xa0, 0x47, 0xad, 0x8a, 0x17, 0x0f, 0xda, 0xa3, 0xb7, 0x26, 0xd9, 0xa6, 0xc1, 0xad,
	0x2f, 0x64, 0x5f, 0xc4, 0x7c, 0xa3, 0x3f, 0x25, 0xbb, 0x51, 0x5a, 0x2a, 0x5e, 0x96, 0xf7, 0x66,
	0x66, 0x67, 0x67, 0x58, 0x18, 0xa5, 0x0a, 0x8b, 0xc4, 0x1c, 0xa2, 0xaa, 0x91, 0x90, 0xf9, 0x66,
	0x0e, 0xc7, 0x05, 0x62, 0xa1, 0x64, 0x62, 0xb1, 0xb4, 0x59, 0x26, 0x72, 0x5d, 0x51, 0xdb, 0x49,
	0xf8, 0x05, 0x1c, 0x3c, 0xa3, 0xa6, 0xa7, 0x52, 0x13, 0x8b, 0xa0, 0x5f, 0xa1, 0x26, 0x1d, 0xb8,
	0x51, 0x2f, 0x3e, 0x9c, 0x82, 0xb0, 0x56, 0x86, 0x9e, 0x77, 0x04, 0xbf, 0x01, 0xdf, 0xac, 0x6c,
	0x08, 0x5e, 0x99, 0x07, 0x6e, 0xe4, 0xc6, 0x83, 0xb9, 0x57, 0xe6, 0xec, 0x18, 0xfa, 0x54, 0x92,
	0x92, 0x81, 0x67, 0xa1, 0x6e, 0x61, 0x0c, 0xfc, 0x14, 0xf3, 0x36, 0xe8, 0x59, 0xd0, 0xce, 0x7c,
	0x0c, 0x03, 0xe3, 0xf0, 0xd2, 0xc8, 0xba, 0xdd, 0xb5, 0x99, 0x7e, 0xb9, 0xe0, 0xcf, 0x14, 0x16,
	0xec, 0x1a, 0x8e, 0x4c, 0xa2, 0x5b, 0xa5, 0x8c, 0x58, 0xb3, 0x13, 0xd1, 0x75, 0x10, 0xbf, 0x1d,
	0xc4, 0x83, 0xe9, 0x10, 0x0e, 0x37, 0x11, 0x8d, 0x9e, 0x3b, 0x2c, 0x86, 0xfd, 0x47, 0x49, 0x36,
	0xe4, 0x68, 0x43, 0xda, 0xe7, 0xc2, 0xad, 0x42, 0xdc, 0x61, 0xa7, 0x00, 0x77, 0xb5, 0x5c, 0x90,
	0xb4, 0xe2, 0x2d, 0x6e, 0x47, 0x77, 0x05, 0x70, 0x2f, 0x95, 0xfc, 0xd1, 0xfd, 0x31, 0xfd, 0x27,
	0x1a, 0x77, 0x66, 0xe7, 0xaf, 0x67, 0x45, 0x49, 0xab, 0x26, 0x15, 0x19, 0xae, 0x93, 0xb7, 0x45,
	0x4b, 0x72, 0x95, 0x64, 0xf4, 0xa9, 0x27, 0xb9, 0xfc, 0xc8, 0xf0, 0x7d, 0x39, 0xa9, 0xab, 0xcc,
	0x7e, 0x57, 0xba, 0x67, 0x2f, 0x5f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xd1, 0xd5, 0x6d,
	0xc2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogClient interface {
	ListAllPosts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PostList, error)
	GetPost(ctx context.Context, in *PostQuery, opts ...grpc.CallOption) (*Post, error)
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	DeletePost(ctx context.Context, in *PostQuery, opts ...grpc.CallOption) (*empty.Empty, error)
}

type blogClient struct {
	cc *grpc.ClientConn
}

func NewBlogClient(cc *grpc.ClientConn) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) ListAllPosts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PostList, error) {
	out := new(PostList)
	err := c.cc.Invoke(ctx, "/blog.Blog/ListAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetPost(ctx context.Context, in *PostQuery, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) DeletePost(ctx context.Context, in *PostQuery, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/blog.Blog/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
type BlogServer interface {
	ListAllPosts(context.Context, *empty.Empty) (*PostList, error)
	GetPost(context.Context, *PostQuery) (*Post, error)
	CreatePost(context.Context, *Post) (*Post, error)
	DeletePost(context.Context, *PostQuery) (*empty.Empty, error)
}

// UnimplementedBlogServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (*UnimplementedBlogServer) ListAllPosts(ctx context.Context, req *empty.Empty) (*PostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllPosts not implemented")
}
func (*UnimplementedBlogServer) GetPost(ctx context.Context, req *PostQuery) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (*UnimplementedBlogServer) CreatePost(ctx context.Context, req *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (*UnimplementedBlogServer) DeletePost(ctx context.Context, req *PostQuery) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}

func RegisterBlogServer(s *grpc.Server, srv BlogServer) {
	s.RegisterService(&_Blog_serviceDesc, srv)
}

func _Blog_ListAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).ListAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/ListAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).ListAllPosts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetPost(ctx, req.(*PostQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).DeletePost(ctx, req.(*PostQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllPosts",
			Handler:    _Blog_ListAllPosts_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Blog_GetPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Blog_CreatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Blog_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/blog.proto",
}
