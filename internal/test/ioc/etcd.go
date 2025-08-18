package ioc

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// etcdConnectTimeout 是连接 etcd 的超时时间
	etcdConnectTimeout = 5 * time.Second
)

// 初始化 etcd 客户端
func InitEtcdClient() *clientv3.Client {
	// 生产环境建议配置多个节点地址
	endpoints := []string{"http://localhost:23790"}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: etcdConnectTimeout,
	})
	if err != nil {
		panic(fmt.Errorf("连接失败: %w", err))
	}
	// 验证连接
	ctx, cancel := context.WithTimeout(context.Background(), etcdConnectTimeout)
	defer cancel()
	if _, err := cli.Status(ctx, endpoints[0]); err != nil {
		panic(fmt.Errorf("连接状态异常: %w", err))
	}
	return cli
}
