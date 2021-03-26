package resolver

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/resolver"
	"shorturl/wangjian-zero/core/lang"
	"shorturl/wangjian-zero/core/mathx"
)

func TestDirectBuilder_Build(t *testing.T) {
	tests := []int{
		0,
		1,
		2,
		subsetSize / 2,
		subsetSize,
		subsetSize * 2,
	}

	for _, test := range tests {
		test := test
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			var servers []string
			for i := 0; i < test; i++ {
				servers = append(servers, fmt.Sprintf("localhost:%d", i))
			}
			var b directBuilder
			cc := new(mockedClientConn)
			_, err := b.Build(resolver.Target{
				Scheme:   DirectScheme,
				Endpoint: strings.Join(servers, ","),
			}, cc, resolver.BuildOptions{})
			assert.Nil(t, err)
			size := mathx.MinInt(test, subsetSize)
			assert.Equal(t, size, len(cc.state.Addresses))
			m := make(map[string]lang.PlaceholderType)
			for _, each := range cc.state.Addresses {
				m[each.Addr] = lang.Placeholder
			}
			assert.Equal(t, size, len(m))
		})
	}
}

func TestDirectBuilder_Scheme(t *testing.T) {
	var b directBuilder
	assert.Equal(t, DirectScheme, b.Scheme())
}
