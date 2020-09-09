package http

import (
	"context"
	"net/http"
)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)
type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
type DecodeResponseFunc func(context.Context, *http.Response) (response interface{}, err error)
