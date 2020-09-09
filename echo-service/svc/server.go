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
