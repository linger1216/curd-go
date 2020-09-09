package svc

import (
	"context"
	"net/http"
)

type CreateEchoRequest struct {
}

type CreateEchoResponse struct {
}

func DecodeHTTPCreateEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type DeleteEchoRequest struct {
}

type DeleteEchoResponse struct {
}

func DecodeHTTPDeleteEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type UpdateEchoRequest struct {
}

type UpdateEchoResponse struct {
}

func DecodeHTTPUpdateEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type ListEchoRequest struct {
}

type ListEchoResponse struct {
}

func DecodeHTTPListEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeHTTPListHeadEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type GetEchoRequest struct {
}

type GetEchoResponse struct {
}

func DecodeHTTPGetEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
