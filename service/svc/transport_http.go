package svc

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	svchttp "github.com/linger1216/go-front/service/svc/http"
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

	engine.Handle("POST", "/eg/echo", func(c *gin.Context) {
		svchttp.NewServer(
			endpoints.CreateEchoEndpoint,
			svchttp.WrapS(c, DecodeHTTPCreateEchoRequest),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("DELETE", "/eg/echo/:ids", func(c *gin.Context) {
		svchttp.NewServer(
			endpoints.DeleteEchoEndpoint,
			svchttp.WrapS(c, DecodeHTTPDeleteEchoRequestV2),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("PUT", "/eg/echo", func(c *gin.Context) {
		svchttp.NewServer(
			endpoints.UpdateEchoEndpoint,
			svchttp.WrapS(c, DecodeHTTPUpdateEchoRequest),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("GET", "/eg/echo", func(c *gin.Context) {
		svchttp.NewServer(
			endpoints.ListEchoEndpoint,
			svchttp.WrapS(c, DecodeHTTPListEchoRequest),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("HEAD", "/eg/echo", func(c *gin.Context) {
		svchttp.NewServer(
			endpoints.ListEchoEndpoint,
			svchttp.WrapS(c, DecodeHTTPListHeadEchoRequest),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("GET", "/eg/echo/:ids", func(c *gin.Context) {
		svchttp.NewServer(
			endpoints.GetEchoEndpoint,
			svchttp.WrapS(c, DecodeHTTPGetEchoRequest),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
}
