package accounts

import (
	"github.com/coreos/etcd/clientv3"
	"gopkg.in/coreos/etcd.v2/Godeps/_workspace/src/golang.org/x/net/context"
	"log"
	"time"
)

var Endpoints []string
var DialTimeout time.Duration

func Config(endpoints []string, dialTimeout time.Duration) {
	Endpoints = endpoints
	DialTimeout = dialTimeout
}

func connect() (cli *clientv3.Client, err error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   Endpoints,
		DialTimeout: DialTimeout * time.Second,
	})
}

func Put(key, value string) error {
	cli, err := connect()
	if err != nil {
		log.Println("Unable to connect to etcd server")
		return err
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)

	putResp, err := kv.Put(context.TODO(), key, value)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("clusterid:", putResp.Header.ClusterId)
	return err
}

func Get(key string) (string, error) {
	cli, err := connect()
	if err != nil {
		log.Println("Unable to connect to etcd server")
		return "", err
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)

	getResp, err := kv.Get(context.TODO(), key)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if len(getResp.Kvs) > 0 {
		return string(getResp.Kvs[0].Value), nil
	}
	return "*** NO KEY ***", nil
}