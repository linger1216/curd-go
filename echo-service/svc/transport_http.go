package svc

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	svchttp "github.com/linger1216/go-front/echo-service/svc/http"
	http "net/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshal, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshal.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if head, ok := err.(svchttp.Headerer); ok {
		for k := range head.Headers() {
			w.Header().Set(k, head.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(svchttp.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func MakeHTTPHandler(engine *gin.Engine, endpoints Endpoints) {

	serverOptions := []svchttp.ServerOption{
		svchttp.ServerBefore(headersToContext),
		svchttp.ServerErrorEncoder(errorEncoder),
		svchttp.ServerErrorHandler(svchttp.NewNopErrorHandler()),
		svchttp.ServerAfter(svchttp.SetContentType(contentType)),
	}

	engine.Handle("POST", "/eg/echo", gin.WrapH(svchttp.NewServer(
		endpoints.CreateEchoEndpoint,
		DecodeHTTPCreateEchoRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	)))

	engine.Handle("DELETE", "/eg/echo/:ids", gin.WrapH(svchttp.NewServer(
		endpoints.DeleteEchoEndpoint,
		DecodeHTTPDeleteEchoRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	)))

	engine.Handle("PUT", "/eg/echo", gin.WrapH(svchttp.NewServer(
		endpoints.UpdateEchoEndpoint,
		DecodeHTTPUpdateEchoRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	)))

	engine.Handle("GET", "/eg/echo", gin.WrapH(svchttp.NewServer(
		endpoints.ListEchoEndpoint,
		DecodeHTTPListEchoRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	)))

	engine.Handle("HEAD", "/eg/echo", gin.WrapH(svchttp.NewServer(
		endpoints.ListEchoEndpoint,
		DecodeHTTPListHeadEchoRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	)))

	engine.Handle("GET", "/eg/echo/:ids", gin.WrapH(svchttp.NewServer(
		endpoints.GetEchoEndpoint,
		DecodeHTTPGetEchoRequest,
		EncodeHTTPGenericResponse,
		serverOptions...,
	)))
}
