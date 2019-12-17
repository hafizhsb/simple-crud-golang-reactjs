package common

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func SetUUIDContext(ctx context.Context, r *http.Request) context.Context {
	uid, _ := uuid.NewV4()
	ctx = context.WithValue(ctx, ContextKeyCorrelationId, uid.String())
	return ctx
}

func SetRequestContext(ctx context.Context, r *http.Request) context.Context {
	for k, v := range map[contextKey]string{
		ContextKeyRequestMethod:     r.Method,
		ContextKeyRequestURI:        r.RequestURI,
		ContextKeyRequestPath:       r.URL.Path,
		ContextKeyRequestProto:      r.Proto,
		ContextKeyRequestHost:       r.Host,
		ContextKeyRequestRemoteAddr: r.RemoteAddr,
	} {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}

type contextKey int

const (
	// ContextKeyRequestMethod is populated in the context by
	// PopulateRequestContext. Its value is r.Method.
	ContextKeyRequestMethod contextKey = iota

	// ContextKeyRequestURI is populated in the context by
	// PopulateRequestContext. Its value is r.RequestURI.
	ContextKeyRequestURI

	// ContextKeyRequestPath is populated in the context by
	// PopulateRequestContext. Its value is r.URL.Path.
	ContextKeyRequestPath

	// ContextKeyRequestProto is populated in the context by
	// PopulateRequestContext. Its value is r.Proto.
	ContextKeyRequestProto

	// ContextKeyRequestHost is populated in the context by
	// PopulateRequestContext. Its value is r.Host.
	ContextKeyRequestHost

	// ContextKeyRequestRemoteAddr is populated in the context by
	// PopulateRequestContext. Its value is r.RemoteAddr.
	ContextKeyRequestRemoteAddr

	// ContextKeyRequestUid is populated in the context by
	// PopulateRequestContext. Its value is r.Header.Get("Uid").
	ContextKeyCorrelationId
)
