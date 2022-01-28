package registry

import (
	etcd "go.etcd.io/etcd/client/v3"
)

type EtcdRegistry struct {
}

func (er *EtcdRegistry) d() {
	etcd.Client{}
}

func NewEtcdRegistry() {

}

type Registry interface {
	Register() error
	Deregister() error
	GetService() error
	ListServices()
}

type Service struct {
}
