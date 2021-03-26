package clientinterceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"shorturl/go-zero/core/trace"
)

// TracingInterceptor is an interceptor that handles tracing.
func TracingInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, span := trace.StartClientSpan(ctx, cc.Target(), method)
	defer span.Finish()

	var pairs []string
	span.Visit(func(key, val string) bool {
		pairs = append(pairs, key, val)
		return true
	})
	ctx = metadata.AppendToOutgoingContext(ctx, pairs...)

	return invoker(ctx, method, req, reply, cc, opts...)
}
