package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/*
	etcd github: https://github.com/etcd-io/etcd

	etcd offical doc: https://etcd.io/docs/v3.5/tutorials/reading-from-etcd/

	go lib source: https://pkg.go.dev/go.etcd.io/etcd/client/v3#section-readme

	practice: https://amyangfei.me/2020/12/19/best-practice-with-go-etcd/

	example: https://github.com/etcd-io/etcd/blob/main/tests/integration/clientv3/examples/example_test.go

	command line get all keys: https://github.com/etcd-io/etcd/issues/11792

	example of getting all keys: https://www.reddit.com/r/golang/comments/muz70g/get_all_the_keys_stored_in_the_etcd_database/
*/

func main() {
	fmt.Println("etcd begins")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		panic(err)
	}
	defer cli.Close()

	res, err := cli.Get(context.Background(), "key1")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res.Kvs[0].Value))
	fmt.Println("------------------------")

	// todoCtx := context.TODO()
	defaultOpts := []clientv3.OpOption{clientv3.WithFromKey()}
	// ctx, cancel := context.WithTimeout(todoCtx, 5*time.Second)
	// defer cancel()
	resp, err := cli.Get(context.Background(), "/", defaultOpts...)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to Get keys with prefix \\")
	}
	fmt.Println(resp.Kvs)
	for _, kv := range resp.Kvs {
		logrus.Infof("Key: %s", string(kv.Key))
	}
}
