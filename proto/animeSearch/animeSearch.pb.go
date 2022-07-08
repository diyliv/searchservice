// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/animeSearch/animeSearch.proto

/*
Package searchpb is a generated protocol buffer package.

It is generated from these files:
	proto/animeSearch/animeSearch.proto

It has these top-level messages:
	Images
	YoutubeTrailerImages
	Trailer
	AnimeSearch
	SearchResponse
	FinSearchResp
*/
package searchpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

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

type Images struct {
	ImageUrl      string `protobuf:"bytes,1,opt,name=imageUrl" json:"imageUrl,omitempty"`
	SmallImageUrl string `protobuf:"bytes,2,opt,name=smallImageUrl" json:"smallImageUrl,omitempty"`
	LargeImageUrl string `protobuf:"bytes,3,opt,name=largeImageUrl" json:"largeImageUrl,omitempty"`
}

func (m *Images) Reset()                    { *m = Images{} }
func (m *Images) String() string            { return proto.CompactTextString(m) }
func (*Images) ProtoMessage()               {}
func (*Images) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Images) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Images) GetSmallImageUrl() string {
	if m != nil {
		return m.SmallImageUrl
	}
	return ""
}

func (m *Images) GetLargeImageUrl() string {
	if m != nil {
		return m.LargeImageUrl
	}
	return ""
}

type YoutubeTrailerImages struct {
	ImageUrl        string `protobuf:"bytes,1,opt,name=imageUrl" json:"imageUrl,omitempty"`
	SmallImageUrl   string `protobuf:"bytes,2,opt,name=smallImageUrl" json:"smallImageUrl,omitempty"`
	MediumImageUrl  string `protobuf:"bytes,3,opt,name=mediumImageUrl" json:"mediumImageUrl,omitempty"`
	LargeImageUrl   string `protobuf:"bytes,4,opt,name=largeImageUrl" json:"largeImageUrl,omitempty"`
	MaximumImageUrl string `protobuf:"bytes,5,opt,name=maximumImageUrl" json:"maximumImageUrl,omitempty"`
}

func (m *YoutubeTrailerImages) Reset()                    { *m = YoutubeTrailerImages{} }
func (m *YoutubeTrailerImages) String() string            { return proto.CompactTextString(m) }
func (*YoutubeTrailerImages) ProtoMessage()               {}
func (*YoutubeTrailerImages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *YoutubeTrailerImages) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *YoutubeTrailerImages) GetSmallImageUrl() string {
	if m != nil {
		return m.SmallImageUrl
	}
	return ""
}

func (m *YoutubeTrailerImages) GetMediumImageUrl() string {
	if m != nil {
		return m.MediumImageUrl
	}
	return ""
}

func (m *YoutubeTrailerImages) GetLargeImageUrl() string {
	if m != nil {
		return m.LargeImageUrl
	}
	return ""
}

func (m *YoutubeTrailerImages) GetMaximumImageUrl() string {
	if m != nil {
		return m.MaximumImageUrl
	}
	return ""
}

type Trailer struct {
	YoutubeID            string                `protobuf:"bytes,1,opt,name=youtubeID" json:"youtubeID,omitempty"`
	YoutubeTrailerURL    string                `protobuf:"bytes,2,opt,name=youtubeTrailerURL" json:"youtubeTrailerURL,omitempty"`
	YoutubeTrailerImages *YoutubeTrailerImages `protobuf:"bytes,3,opt,name=youtubeTrailerImages" json:"youtubeTrailerImages,omitempty"`
}

func (m *Trailer) Reset()                    { *m = Trailer{} }
func (m *Trailer) String() string            { return proto.CompactTextString(m) }
func (*Trailer) ProtoMessage()               {}
func (*Trailer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Trailer) GetYoutubeID() string {
	if m != nil {
		return m.YoutubeID
	}
	return ""
}

func (m *Trailer) GetYoutubeTrailerURL() string {
	if m != nil {
		return m.YoutubeTrailerURL
	}
	return ""
}

func (m *Trailer) GetYoutubeTrailerImages() *YoutubeTrailerImages {
	if m != nil {
		return m.YoutubeTrailerImages
	}
	return nil
}

type AnimeSearch struct {
	MalId          int64    `protobuf:"varint,1,opt,name=malId" json:"malId,omitempty"`
	Url            string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	JpgImages      *Images  `protobuf:"bytes,3,opt,name=jpgImages" json:"jpgImages,omitempty"`
	YoutubeTrailer *Trailer `protobuf:"bytes,4,opt,name=youtubeTrailer" json:"youtubeTrailer,omitempty"`
	Title          string   `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	TitleEng       string   `protobuf:"bytes,6,opt,name=titleEng" json:"titleEng,omitempty"`
	Type           string   `protobuf:"bytes,7,opt,name=type" json:"type,omitempty"`
	Source         string   `protobuf:"bytes,8,opt,name=source" json:"source,omitempty"`
	Episodes       int64    `protobuf:"varint,9,opt,name=episodes" json:"episodes,omitempty"`
	Status         string   `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
	Rating         string   `protobuf:"bytes,11,opt,name=rating" json:"rating,omitempty"`
	Score          float32  `protobuf:"fixed32,12,opt,name=score" json:"score,omitempty"`
	Year           int64    `protobuf:"varint,13,opt,name=year" json:"year,omitempty"`
}

func (m *AnimeSearch) Reset()                    { *m = AnimeSearch{} }
func (m *AnimeSearch) String() string            { return proto.CompactTextString(m) }
func (*AnimeSearch) ProtoMessage()               {}
func (*AnimeSearch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AnimeSearch) GetMalId() int64 {
	if m != nil {
		return m.MalId
	}
	return 0
}

func (m *AnimeSearch) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AnimeSearch) GetJpgImages() *Images {
	if m != nil {
		return m.JpgImages
	}
	return nil
}

func (m *AnimeSearch) GetYoutubeTrailer() *Trailer {
	if m != nil {
		return m.YoutubeTrailer
	}
	return nil
}

func (m *AnimeSearch) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AnimeSearch) GetTitleEng() string {
	if m != nil {
		return m.TitleEng
	}
	return ""
}

func (m *AnimeSearch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnimeSearch) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AnimeSearch) GetEpisodes() int64 {
	if m != nil {
		return m.Episodes
	}
	return 0
}

func (m *AnimeSearch) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AnimeSearch) GetRating() string {
	if m != nil {
		return m.Rating
	}
	return ""
}

func (m *AnimeSearch) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *AnimeSearch) GetYear() int64 {
	if m != nil {
		return m.Year
	}
	return 0
}

type SearchResponse struct {
	SearchResponse *AnimeSearch `protobuf:"bytes,1,opt,name=searchResponse" json:"searchResponse,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SearchResponse) GetSearchResponse() *AnimeSearch {
	if m != nil {
		return m.SearchResponse
	}
	return nil
}

type FinSearchResp struct {
	Finresp []*SearchResponse `protobuf:"bytes,1,rep,name=finresp" json:"finresp,omitempty"`
}

func (m *FinSearchResp) Reset()                    { *m = FinSearchResp{} }
func (m *FinSearchResp) String() string            { return proto.CompactTextString(m) }
func (*FinSearchResp) ProtoMessage()               {}
func (*FinSearchResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FinSearchResp) GetFinresp() []*SearchResponse {
	if m != nil {
		return m.Finresp
	}
	return nil
}

func init() {
	proto.RegisterType((*Images)(nil), "search.Images")
	proto.RegisterType((*YoutubeTrailerImages)(nil), "search.YoutubeTrailerImages")
	proto.RegisterType((*Trailer)(nil), "search.Trailer")
	proto.RegisterType((*AnimeSearch)(nil), "search.AnimeSearch")
	proto.RegisterType((*SearchResponse)(nil), "search.SearchResponse")
	proto.RegisterType((*FinSearchResp)(nil), "search.FinSearchResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SearchService service

type SearchServiceClient interface {
	Search(ctx context.Context, in *google_protobuf.StringValue, opts ...grpc.CallOption) (*FinSearchResp, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *google_protobuf.StringValue, opts ...grpc.CallOption) (*FinSearchResp, error) {
	out := new(FinSearchResp)
	err := grpc.Invoke(ctx, "/search.SearchService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchService service

type SearchServiceServer interface {
	Search(context.Context, *google_protobuf.StringValue) (*FinSearchResp, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*google_protobuf.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/animeSearch/animeSearch.proto",
}

func init() { proto.RegisterFile("proto/animeSearch/animeSearch.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xeb, 0x24, 0x63, 0x92, 0xc2, 0x12, 0xaa, 0x55, 0x14, 0xa1, 0xc8, 0x20, 0x94,
	0x43, 0xe5, 0x20, 0x73, 0xe0, 0xc0, 0x01, 0x15, 0x01, 0x52, 0x24, 0x90, 0x2a, 0x87, 0x22, 0xc1,
	0x6d, 0x93, 0x4c, 0xcd, 0x22, 0xff, 0xac, 0x76, 0x6d, 0xc0, 0x4f, 0xc4, 0xeb, 0xf0, 0x28, 0x3c,
	0x02, 0xf2, 0xae, 0x9d, 0xc4, 0x6e, 0x8e, 0xbd, 0xcd, 0xf7, 0xed, 0x37, 0x33, 0xdf, 0x8c, 0x47,
	0x86, 0xa7, 0x42, 0xa6, 0x59, 0xba, 0x60, 0x09, 0x8f, 0x71, 0x85, 0x4c, 0x6e, 0xbe, 0x1f, 0xc6,
	0x9e, 0x7e, 0x25, 0xb6, 0xd2, 0x68, 0xf2, 0x24, 0x4c, 0xd3, 0x30, 0xc2, 0x85, 0x66, 0xd7, 0xf9,
	0xcd, 0xe2, 0x97, 0x64, 0x42, 0xa0, 0x54, 0x46, 0xe7, 0x0a, 0xb0, 0x97, 0x31, 0x0b, 0x51, 0x91,
	0x09, 0xf4, 0x79, 0x19, 0x5d, 0xcb, 0x88, 0x5a, 0x33, 0x6b, 0x3e, 0x08, 0x76, 0x98, 0x3c, 0x83,
	0xa1, 0x8a, 0x59, 0x14, 0x2d, 0x6b, 0x41, 0x47, 0x0b, 0x9a, 0x64, 0xa9, 0x8a, 0x98, 0x0c, 0x71,
	0xa7, 0xea, 0x1a, 0x55, 0x83, 0x74, 0xff, 0x5a, 0x30, 0xfe, 0x9a, 0xe6, 0x59, 0xbe, 0xc6, 0xcf,
	0x92, 0xf1, 0x08, 0xe5, 0x9d, 0x19, 0x78, 0x0e, 0xa3, 0x18, 0xb7, 0x3c, 0x8f, 0x5b, 0x0e, 0x5a,
	0xec, 0x6d, 0xa3, 0x27, 0x47, 0x8c, 0x92, 0x39, 0x9c, 0xc5, 0xec, 0x37, 0x8f, 0x0f, 0xca, 0x9d,
	0x6a, 0x5d, 0x9b, 0x76, 0xff, 0x58, 0xd0, 0xab, 0x66, 0x21, 0x53, 0x18, 0x14, 0x66, 0xba, 0xe5,
	0xbb, 0x6a, 0x8c, 0x3d, 0x41, 0x2e, 0xe0, 0x61, 0xd1, 0x98, 0xfd, 0x3a, 0xf8, 0x58, 0xcd, 0x72,
	0xfb, 0x81, 0x5c, 0xc1, 0xb8, 0x38, 0xb2, 0x29, 0x3d, 0x95, 0xe3, 0x4f, 0x3d, 0xf3, 0x8d, 0xbd,
	0x63, 0xdb, 0x0c, 0x8e, 0x66, 0xba, 0xff, 0x3a, 0xe0, 0x5c, 0xee, 0x8f, 0x85, 0x8c, 0xe1, 0x34,
	0x66, 0xd1, 0x72, 0xab, 0x9d, 0x76, 0x03, 0x03, 0xc8, 0x03, 0xe8, 0xe6, 0xbb, 0x1d, 0x97, 0x21,
	0xb9, 0x80, 0xc1, 0x0f, 0x11, 0x36, 0xda, 0x8f, 0xea, 0xf6, 0x55, 0xc3, 0xbd, 0x80, 0xbc, 0x82,
	0x51, 0xb3, 0xbb, 0x5e, 0xb0, 0xe3, 0x9f, 0xd5, 0x29, 0x15, 0x1d, 0xb4, 0x64, 0xa5, 0x9d, 0x8c,
	0x67, 0x11, 0x56, 0x8b, 0x36, 0xa0, 0x3c, 0x0c, 0x1d, 0xbc, 0x4f, 0x42, 0x6a, 0x9b, 0xc3, 0xa8,
	0x31, 0x21, 0x70, 0x92, 0x15, 0x02, 0x69, 0x4f, 0xf3, 0x3a, 0x26, 0xe7, 0x60, 0xab, 0x34, 0x97,
	0x1b, 0xa4, 0x7d, 0xcd, 0x56, 0xa8, 0xac, 0x83, 0x82, 0xab, 0x74, 0x8b, 0x8a, 0x0e, 0xf4, 0xbc,
	0x3b, 0xac, 0x73, 0x32, 0x96, 0xe5, 0x8a, 0x42, 0x95, 0xa3, 0x51, 0xc9, 0x4b, 0x96, 0xf1, 0x24,
	0xa4, 0x8e, 0xe1, 0x0d, 0x2a, 0x9d, 0xaa, 0x4d, 0x2a, 0x91, 0xde, 0x9f, 0x59, 0xf3, 0x4e, 0x60,
	0x40, 0xe9, 0xa6, 0x40, 0x26, 0xe9, 0x50, 0x57, 0xd7, 0xb1, 0xfb, 0x09, 0x46, 0x66, 0xd9, 0x01,
	0x2a, 0x91, 0x26, 0x0a, 0xc9, 0x6b, 0x18, 0xa9, 0x06, 0xa3, 0xb7, 0xef, 0xf8, 0x8f, 0xea, 0xf5,
	0x1c, 0x7c, 0xa1, 0xa0, 0x25, 0x75, 0x2f, 0x61, 0xf8, 0x81, 0x27, 0xfb, 0x8a, 0xe4, 0x05, 0xf4,
	0x6e, 0x78, 0x22, 0x51, 0x09, 0x6a, 0xcd, 0xba, 0x73, 0xc7, 0x3f, 0xaf, 0xcb, 0x34, 0xdb, 0x06,
	0xb5, 0xcc, 0xbf, 0x82, 0xa1, 0x79, 0x5a, 0xa1, 0xfc, 0xc9, 0x37, 0x48, 0xde, 0x80, 0x5d, 0xdd,
	0xc3, 0xd4, 0x33, 0xff, 0x0b, 0xaf, 0xfe, 0x5f, 0x78, 0xab, 0x4c, 0xf2, 0x24, 0xfc, 0xc2, 0xa2,
	0x1c, 0x27, 0x8f, 0xeb, 0xca, 0x0d, 0x07, 0xee, 0xbd, 0xb7, 0xf0, 0xad, 0x6f, 0x5e, 0xc4, 0x7a,
	0x6d, 0xeb, 0xd4, 0x97, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa9, 0x35, 0xab, 0xa7, 0x04,
	0x00, 0x00,
}
