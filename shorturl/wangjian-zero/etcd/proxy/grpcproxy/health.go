// Copyright 2017 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcproxy

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
	"shorturl/wangjian-zero/etcd/clientv3"
	"shorturl/wangjian-zero/etcd/etcdserver/api/etcdhttp"
	"shorturl/wangjian-zero/etcd/etcdserver/api/v3rpc/rpctypes"
)

// HandleHealth registers health handler on '/health'.
func HandleHealth(lg *zap.Logger, mux *http.ServeMux, c *clientv3.Client) {
	if lg == nil {
		lg = zap.NewNop()
	}
	mux.Handle(etcdhttp.PathHealth, etcdhttp.NewHealthHandler(lg, func() etcdhttp.Health { return checkHealth(c) }))
}

func checkHealth(c *clientv3.Client) etcdhttp.Health {
	h := etcdhttp.Health{Health: "false"}
	ctx, cancel := context.WithTimeout(c.Ctx(), time.Second)
	_, err := c.Get(ctx, "a")
	cancel()
	if err == nil || err == rpctypes.ErrPermissionDenied {
		h.Health = "true"
	}
	return h
}
