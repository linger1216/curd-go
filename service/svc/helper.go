package svc

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strings"
)

const contentType = "application/json; charset=utf-8"

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k, _ := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// add the access key to context
	accessKey := r.URL.Query().Get("access_key")
	if len(accessKey) > 0 {
		ctx = context.WithValue(ctx, "access_key", accessKey)
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "request-url", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}

func writeHeader(w http.ResponseWriter, kvs []*KVResponse) error {
	for i := range kvs {
		w.Header().Set(kvs[i].Key, kvs[i].Value)
	}
	return nil
}

func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	encoder := jsoniter.ConfigFastest.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	switch x := response.(type) {
	case *CreateEchoResponse:
		return encoder.Encode(x.Ids)
	case *DeleteEchoResponse:
	case *UpdateEchoResponse:
	case *ListEchoResponse:
		if len(x.Headers) > 0 {
			return writeHeader(w, x.Headers)
		}
		return encoder.Encode(x.Echos)
	case *GetEchoResponse:
		return encoder.Encode(x.Echos)
	}
	return encoder.Encode(response)
}
