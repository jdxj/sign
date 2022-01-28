package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	etcd "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient(endpoints []string, ca, cert, key string) *etcd.Client {
	kp, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		panic(err)
	}
	d, err := os.ReadFile(ca)
	if err != nil {
		panic(err)
	}
	cp := x509.NewCertPool()
	cp.AppendCertsFromPEM(d)
	tc := &tls.Config{
		Certificates: []tls.Certificate{kp},
		RootCAs:      cp,
	}
	c, err := etcd.New(etcd.Config{
		Endpoints: endpoints,
		TLS:       tc,
	})
	if err != nil {
		panic(err)
	}
	return c
}
