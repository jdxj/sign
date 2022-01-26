package rpc

import (
	"context"
	"fmt"
	"sync"

	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"

	"github.com/jdxj/sign/internal/pkg/logger"
)

const (
	SignScheme      = "sign"
	SignSchemeLocal = "sign.local"
)

func Init(etcdAddr string) {
	client, err := clientV3.New(clientV3.Config{
		Endpoints: []string{etcdAddr},
	})
	if err != nil {
		panic(err)
	}
	resolver.Register(&SignBuilder{
		etcdClient: client,
	})
}

type SignBuilder struct {
	etcdClient *clientV3.Client
}

func (sb *SignBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx, cancel := context.WithCancel(context.Background())
	sr := &SignResolver{
		cc:             cc,
		ctx:            ctx,
		cancel:         cancel,
		wg:             &sync.WaitGroup{},
		etcdClient:     sb.etcdClient,
		etcdKeyService: fmt.Sprintf("%s/%s", registry, target.Endpoint),
		addresses:      make(map[string]struct{}),
	}
	sr.watch()
	logger.Debugf("Build")
	return sr, nil
}

func (sb *SignBuilder) Scheme() string {
	return SignScheme
}

type LocalBuilder struct {
}

func (lb *LocalBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (
	resolver.Resolver, error) {
	lr := &LocalResolver{
		cc:      cc,
		service: target.Endpoint,
	}
	// 立即更新
	lr.update()
	return lr, nil
}

func (lb *LocalBuilder) Scheme() string {
	return SignSchemeLocal
}
