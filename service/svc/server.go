package svc

import (
	"context"
)

type EchoServer interface {
	CreateEcho(ctx context.Context, in *CreateEchoRequest) (*CreateEchoResponse, error)
	DeleteEcho(ctx context.Context, in *DeleteEchoRequest) (*DeleteEchoResponse, error)
	UpdateEcho(ctx context.Context, in *UpdateEchoRequest) (*UpdateEchoResponse, error)
	ListEcho(ctx context.Context, in *ListEchoRequest) (*ListEchoResponse, error)
	GetEcho(ctx context.Context, in *GetEchoRequest) (*GetEchoResponse, error)
	Close() error
}

type EchoCacheServer interface {
	UpdateEcho(ctx context.Context, in *UpdateEchoRequest) error
	ListEcho(ctx context.Context, in *ListEchoRequest) (*ListEchoResponse, error)
	UpdateListEcho(ctx context.Context, in *ListEchoRequest, resp *ListEchoResponse) error
	GetEcho(ctx context.Context, in *GetEchoRequest) (*GetEchoResponse, error)
	UpdateGetEcho(ctx context.Context, in *GetEchoRequest, resp *GetEchoResponse) error
	Close() error
}
