package sample

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func KeepAlive(ctx context.Context, cli *clientv3.Client, srv, addr string, ttl int) (err error) {
	for {
		err = keepAlive(ctx, cli, srv, addr, ttl)
		if err == context.DeadlineExceeded || err == context.Canceled {
			return
		}
		time.Sleep(time.Duration(rand.Int63n(10)) * time.Millisecond)
	}
}

func keepAlive(ctx context.Context, cli *clientv3.Client, srv, addr string, ttl int) error {
	sess, err := concurrency.NewSession(cli, concurrency.WithContext(ctx), concurrency.WithTTL(ttl))
	if err != nil {
		return err
	}
	defer sess.Close()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-sess.Done():
		return fmt.Errorf("%s/%s exit", srv, addr)
	}
}

//func TestGocloud() *runtimevar.Variable {
//return nil
//}
