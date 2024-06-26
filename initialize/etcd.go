package initialize

import (
	"github.com/SliverFlow/core/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Etcd(c *config.Etcd) (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		Username:  c.Username,
		Password:  c.Password,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
