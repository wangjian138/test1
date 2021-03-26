package clientinterceptors

import (
	"context"
	"path"
	"time"

	"google.golang.org/grpc"
	"shorturl/go-zero/core/logx"
	"shorturl/go-zero/core/timex"
)

const slowThreshold = time.Millisecond * 500

// DurationInterceptor is an interceptor that logs the processing time.
func DurationInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	serverName := path.Join(cc.Target(), method)
	start := timex.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		logx.WithContext(ctx).WithDuration(timex.Since(start)).Infof("fail - %s - %v - %s",
			serverName, req, err.Error())
	} else {
		elapsed := timex.Since(start)
		if elapsed > slowThreshold {
			logx.WithContext(ctx).WithDuration(elapsed).Slowf("[RPC] ok - slowcall - %s - %v - %v",
				serverName, req, reply)
		}
	}

	return err
}
