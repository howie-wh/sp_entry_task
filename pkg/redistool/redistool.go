package redistool

import (
	"entry_task/pkg/common"
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)


// RedisClient redis client
type RedisClient struct {
	pool    *redis.Pool
	connOpt common.RedisConnOpt
	// 数据接收
	chanRx chan common.RedisDataArray
	// 是否退出
	isExit bool
}

// NewRedis new redis client
func NewRedis(opt common.RedisConnOpt) *RedisClient {
	return &RedisClient{
		connOpt: opt,
		pool:    newPool(opt),
		chanRx:  make(chan common.RedisDataArray, 100),
	}
}

// newPool 线程池
func newPool(opt common.RedisConnOpt) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// MaxActive:   10,
		// Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", opt.Host, opt.Port))
			if err != nil {
				log.Fatalf("Redis.Dial: %v", err)
				return nil, err
			}
			//当redis不存在密码时不进行认证处理
			if gotool.StrUtils.HasNotEmpty(opt.Password) {
				if _, err := c.Do("AUTH", opt.Password); err != nil {
					c.Close()
					log.Fatalf("Redis.AUTH: %v", err)
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", opt.Index); err != nil {
				c.Close()
				log.Fatalf("Redis.SELECT: %v", err)
				return nil, err
			}
			return c, nil
		},
	}
}

// Start 启动接收任务协程
func (r *RedisClient) Start() {
	r.isExit = false
	// 开启协程用于循环接收数据
	go r.loopRead()
}

// Stop 停止接收任务
func (r *RedisClient) Stop() {
	r.isExit = true
	// 关闭数据接收通道
	close(r.chanRx)
	// 关闭redis线程池
	r.pool.Close()
}

// Write 向redis中写入多组数据
func (r *RedisClient) Write(data common.RedisDataArray) {
	r.chanRx <- data
}

// loopRead 循环接收数据
func (r *RedisClient) loopRead() {
	for !r.isExit {
		select {
		case rx := <-r.chanRx:
			for _, it := range rx {
				if len(it.Key) > 0 {
					if len(it.Field) > 0 {
						if _, err := r.HSET(it.Key, it.Field, it.Value); err != nil {
							log.Printf("[%s, %s, %s]: %s\n", it.Key, it.Field, it.Value, err.Error())
						}
					} else {
						if _, err := r.SET(it.Key, it.Value); err != nil {
							log.Printf("[%s, %s, %s]: %s\n", it.Key, it.Field, it.Value, err.Error())
						}
					}
					if it.Expire > 0 {
						r.EXPIRE(it.Key, it.Expire)
					}
				}
			}
		}
	}

}

func (r *RedisClient) Error() error {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Err()
}

// 常用Redis操作命令的封装

func (r *RedisClient) SCAN(patten string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	var out []string
	var cursor uint64 = 0xffffff
	isfirst := true
	for cursor != 0 {
		if isfirst {
			cursor = 0
			isfirst = false
		}
		arr, err := conn.Do("SCAN", cursor, "MATCH", patten, "COUNT", 100)
		if err != nil {
			return out, err
		}
		switch arr := arr.(type) {
		case []interface{}:
			cursor, _ = redis.Uint64(arr[0], nil)
			it, _ := redis.Strings(arr[1], nil)
			out = append(out, it...)
		}
	}
	out = gotool.StrArrayUtils.ArrayDuplication(out)
	return out, nil
}

func (r *RedisClient) DEL(key string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DEL", key))
}

func (r *RedisClient) DELALL(key []string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	arr := make([]interface{}, len(key))
	for i, v := range key {
		arr[i] = v
	}
	return redis.Int(conn.Do("DEL", arr...))
}

func (r *RedisClient) GET(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (r *RedisClient) SET(key string, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SET", key, value))
}

func (r *RedisClient) SETEX(key string, sec int, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SETEX", key, sec, value))
}

func (r *RedisClient) EXPIRE(key string, sec int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("EXPIRE", key, sec))
}

func (r *RedisClient) HGETALL(key string) (map[string]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

func (r *RedisClient) HGET(key string, field string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

func (r *RedisClient) HSET(key string, field string, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("HSET", key, field, value))
}
