package utils

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"nonolive/nonoutils/config"
)

var _pool *pool.Pool

func init() {
	var err error
	url := config.GetGlobalStringValue("redis_url", "")
	_pool, err = pool.New("tcp", url, 20)
	if err != nil {
		fmt.Printf("connect to redis(%v) fail.\n", url)
		panic(err)
	}
	fmt.Printf("connect to redis(%v) ok.\n", url)
}

func WithinRedis(f func(*redis.Client) error) error {
	conn, err := _pool.Get()
	if err != nil {
		return err
	}
	defer func() {
		if conn != nil {
			conn.PipeClear()
		}
		_pool.Put(conn)
	}()
	return f(conn)
}
