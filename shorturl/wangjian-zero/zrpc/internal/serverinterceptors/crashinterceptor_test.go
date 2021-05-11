package serverinterceptors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"shorturl/wangjian-zero/core/logx"
	"shorturl/wangjian-zero/grpc"
)

func init() {
	logx.Disable()
}

func TestStreamCrashInterceptor(t *testing.T) {
	err := StreamCrashInterceptor(nil, nil, nil, func(
		srv interface{}, stream grpc.ServerStream) error {
		panic("mock panic")
	})
	assert.NotNil(t, err)
}

func TestUnaryCrashInterceptor(t *testing.T) {
	interceptor := UnaryCrashInterceptor()
	_, err := interceptor(context.Background(), nil, nil,
		func(ctx context.Context, req interface{}) (interface{}, error) {
			panic("mock panic")
		})
	assert.NotNil(t, err)
}
