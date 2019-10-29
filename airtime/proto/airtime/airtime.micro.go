// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/airtime/airtime.proto

package airtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Airtime service

type AirtimeService interface {
	AddCountry(ctx context.Context, in *AddCountryRequest, opts ...client.CallOption) (*AddCountryResponse, error)
	Countries(ctx context.Context, in *GetCountriesRequest, opts ...client.CallOption) (*GetCountriesResponse, error)
	AddProvider(ctx context.Context, in *AddProviderRequest, opts ...client.CallOption) (*AddProviderResponse, error)
	Providers(ctx context.Context, in *ProvidersRequest, opts ...client.CallOption) (*ProvidersResponse, error)
	AddProviderToCountry(ctx context.Context, in *AddProviderToCountyRequest, opts ...client.CallOption) (*AddProviderToCountryResponse, error)
	SendAirtime(ctx context.Context, in *SendAirtimeRequest, opts ...client.CallOption) (*SendAirtimeResponse, error)
	History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error)
	HistoryCount(ctx context.Context, in *HistoryCountRequest, opts ...client.CallOption) (*HistoryCountResponse, error)
}

type airtimeService struct {
	c    client.Client
	name string
}

func NewAirtimeService(name string, c client.Client) AirtimeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "airtime"
	}
	return &airtimeService{
		c:    c,
		name: name,
	}
}

func (c *airtimeService) AddCountry(ctx context.Context, in *AddCountryRequest, opts ...client.CallOption) (*AddCountryResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.AddCountry", in)
	out := new(AddCountryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) Countries(ctx context.Context, in *GetCountriesRequest, opts ...client.CallOption) (*GetCountriesResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.Countries", in)
	out := new(GetCountriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) AddProvider(ctx context.Context, in *AddProviderRequest, opts ...client.CallOption) (*AddProviderResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.AddProvider", in)
	out := new(AddProviderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) Providers(ctx context.Context, in *ProvidersRequest, opts ...client.CallOption) (*ProvidersResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.Providers", in)
	out := new(ProvidersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) AddProviderToCountry(ctx context.Context, in *AddProviderToCountyRequest, opts ...client.CallOption) (*AddProviderToCountryResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.AddProviderToCountry", in)
	out := new(AddProviderToCountryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) SendAirtime(ctx context.Context, in *SendAirtimeRequest, opts ...client.CallOption) (*SendAirtimeResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.SendAirtime", in)
	out := new(SendAirtimeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.Histories", in)
	out := new(HistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimeService) HistoryCount(ctx context.Context, in *HistoryCountRequest, opts ...client.CallOption) (*HistoryCountResponse, error) {
	req := c.c.NewRequest(c.name, "Airtime.HistoryCount", in)
	out := new(HistoryCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Airtime service

type AirtimeHandler interface {
	AddCountry(context.Context, *AddCountryRequest, *AddCountryResponse) error
	Countries(context.Context, *GetCountriesRequest, *GetCountriesResponse) error
	AddProvider(context.Context, *AddProviderRequest, *AddProviderResponse) error
	Providers(context.Context, *ProvidersRequest, *ProvidersResponse) error
	AddProviderToCountry(context.Context, *AddProviderToCountyRequest, *AddProviderToCountryResponse) error
	SendAirtime(context.Context, *SendAirtimeRequest, *SendAirtimeResponse) error
	History(context.Context, *HistoryRequest, *HistoryResponse) error
	HistoryCount(context.Context, *HistoryCountRequest, *HistoryCountResponse) error
}

func RegisterAirtimeHandler(s server.Server, hdlr AirtimeHandler, opts ...server.HandlerOption) error {
	type airtime interface {
		AddCountry(ctx context.Context, in *AddCountryRequest, out *AddCountryResponse) error
		Countries(ctx context.Context, in *GetCountriesRequest, out *GetCountriesResponse) error
		AddProvider(ctx context.Context, in *AddProviderRequest, out *AddProviderResponse) error
		Providers(ctx context.Context, in *ProvidersRequest, out *ProvidersResponse) error
		AddProviderToCountry(ctx context.Context, in *AddProviderToCountyRequest, out *AddProviderToCountryResponse) error
		SendAirtime(ctx context.Context, in *SendAirtimeRequest, out *SendAirtimeResponse) error
		History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error
		HistoryCount(ctx context.Context, in *HistoryCountRequest, out *HistoryCountResponse) error
	}
	type Airtime struct {
		airtime
	}
	h := &airtimeHandler{hdlr}
	return s.Handle(s.NewHandler(&Airtime{h}, opts...))
}

type airtimeHandler struct {
	AirtimeHandler
}

func (h *airtimeHandler) AddCountry(ctx context.Context, in *AddCountryRequest, out *AddCountryResponse) error {
	return h.AirtimeHandler.AddCountry(ctx, in, out)
}

func (h *airtimeHandler) Countries(ctx context.Context, in *GetCountriesRequest, out *GetCountriesResponse) error {
	return h.AirtimeHandler.Countries(ctx, in, out)
}

func (h *airtimeHandler) AddProvider(ctx context.Context, in *AddProviderRequest, out *AddProviderResponse) error {
	return h.AirtimeHandler.AddProvider(ctx, in, out)
}

func (h *airtimeHandler) Providers(ctx context.Context, in *ProvidersRequest, out *ProvidersResponse) error {
	return h.AirtimeHandler.Providers(ctx, in, out)
}

func (h *airtimeHandler) AddProviderToCountry(ctx context.Context, in *AddProviderToCountyRequest, out *AddProviderToCountryResponse) error {
	return h.AirtimeHandler.AddProviderToCountry(ctx, in, out)
}

func (h *airtimeHandler) SendAirtime(ctx context.Context, in *SendAirtimeRequest, out *SendAirtimeResponse) error {
	return h.AirtimeHandler.SendAirtime(ctx, in, out)
}

func (h *airtimeHandler) History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error {
	return h.AirtimeHandler.History(ctx, in, out)
}

func (h *airtimeHandler) HistoryCount(ctx context.Context, in *HistoryCountRequest, out *HistoryCountResponse) error {
	return h.AirtimeHandler.HistoryCount(ctx, in, out)
}