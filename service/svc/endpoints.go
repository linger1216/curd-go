package svc

import (
	"context"
	"github.com/linger1216/go-front/service/svc/endpoint"
)

type Endpoints struct {
	CreateEchoEndpoint endpoint.Endpoint
	DeleteEchoEndpoint endpoint.Endpoint
	UpdateEchoEndpoint endpoint.Endpoint
	ListEchoEndpoint   endpoint.Endpoint
	GetEchoEndpoint    endpoint.Endpoint
}

// Endpoints

func (e Endpoints) CreateEcho(ctx context.Context, in *CreateEchoRequest) (*CreateEchoResponse, error) {
	response, err := e.CreateEchoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*CreateEchoResponse), nil
}

func (e Endpoints) DeleteEcho(ctx context.Context, in *DeleteEchoRequest) (*DeleteEchoResponse, error) {
	response, err := e.DeleteEchoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*DeleteEchoResponse), nil
}

func (e Endpoints) UpdateEcho(ctx context.Context, in *UpdateEchoRequest) (*UpdateEchoResponse, error) {
	response, err := e.UpdateEchoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*UpdateEchoResponse), nil
}

func (e Endpoints) ListEcho(ctx context.Context, in *ListEchoRequest) (*ListEchoResponse, error) {
	response, err := e.ListEchoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*ListEchoResponse), nil
}

func (e Endpoints) GetEcho(ctx context.Context, in *GetEchoRequest) (*GetEchoResponse, error) {
	response, err := e.GetEchoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*GetEchoResponse), nil
}

// Make Endpoints

func MakeCreateEchoEndpoint(s EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateEchoRequest)
		v, err := s.CreateEcho(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetEchoEndpoint(s EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetEchoRequest)
		v, err := s.GetEcho(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeListEchoEndpoint(s EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*ListEchoRequest)
		v, err := s.ListEcho(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeUpdateEchoEndpoint(s EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdateEchoRequest)
		v, err := s.UpdateEcho(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDeleteEchoEndpoint(s EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*DeleteEchoRequest)
		v, err := s.DeleteEcho(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
