package rpc

import (
	"context"
	"fmt"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"

	"github.com/jdxj/sign/internal/pkg/logger"
	test_grpc "github.com/jdxj/sign/internal/proto/test-grpc"
)

type SignResolver struct {
	cc         resolver.ClientConn
	ctx        context.Context
	cancel     context.CancelFunc
	wg         *sync.WaitGroup
	etcdClient *clientV3.Client

	etcdKeyService string
	addresses      map[string]struct{}
}

func (sr *SignResolver) ResolveNow(rno resolver.ResolveNowOptions) {
	logger.Debugf("ResolveNow-getLatest")
	sr.getLatest()
}

// Close closes the resolver.
func (sr *SignResolver) Close() {
	sr.cancel()
	sr.wg.Wait()
}

func (sr *SignResolver) watch() {
	sr.wg.Add(1)
	go func() {
		sr.wg.Done()

		for {
			select {
			case <-sr.ctx.Done():
				logger.Infof("stop watch")
				return
			default:
			}

			logger.Debugf("watch-getLatest")
			sr.getLatest()

			watchChan := sr.etcdClient.Watch(sr.ctx, sr.etcdKeyService, clientV3.WithPrefix())
			for watchRsp := range watchChan {
				err := watchRsp.Err()
				if err != nil {
					logger.Errorf("watch err: %s", err)
					// todo: 是否要 continue
				}
				for _, event := range watchRsp.Events {
					sr.updateAddress(event.Type, string(event.Kv.Value))
				}
			}
		}
	}()
}

func (sr *SignResolver) getLatest() {
	rsp, err := sr.etcdClient.Get(sr.ctx, sr.etcdKeyService, clientV3.WithPrefix())
	if err != nil {
		logger.Errorf("get service address err: %s", err)
		return
	}

	if len(rsp.Kvs) == 0 {
		return
	}

	addresses := make([]resolver.Address, 0, len(rsp.Kvs))
	for _, kv := range rsp.Kvs {
		addr := resolver.Address{
			Addr: string(kv.Value),
		}
		addresses = append(addresses, addr)
	}
	err = sr.cc.UpdateState(resolver.State{
		Addresses: addresses,
	})
	if err != nil {
		logger.Errorf("update address state err: %s", err)
	}
}

func (sr *SignResolver) updateAddress(tpy mvccpb.Event_EventType, value string) {
	switch tpy {
	case mvccpb.PUT:
		sr.addresses[value] = struct{}{}
	case mvccpb.DELETE:
		delete(sr.addresses, value)
	}
	logger.Debugf("type: %s, addr: %s", tpy, value)

	addresses := make([]resolver.Address, 0, len(sr.addresses))
	for addr := range sr.addresses {
		address := resolver.Address{
			Addr: addr,
		}
		addresses = append(addresses, address)
	}

	err := sr.cc.UpdateState(resolver.State{Addresses: addresses})
	if err != nil {
		logger.Errorf("update address state err: %s", err)
	}
}

type LocalResolver struct {
	cc      resolver.ClientConn
	service string
}

func (lr *LocalResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	lr.update()
}

func (lr *LocalResolver) Close() {

}

func (lr *LocalResolver) update() {
	// todo: 添加待测试 service
	var addr resolver.Address
	switch lr.service {
	case test_grpc.ServiceName:
		addr.Addr = fmt.Sprintf("127.0.0.1:%d", test_grpc.ServicePort)
	case test_grpc.MServiceName:
		addr.Addr = fmt.Sprintf("127.0.0.1:%d", test_grpc.MServicePort)
	}
	err := lr.cc.UpdateState(resolver.State{
		Addresses: []resolver.Address{addr},
	})
	if err != nil {
		logger.Errorf("update address state err: %s", err)
	}
}
