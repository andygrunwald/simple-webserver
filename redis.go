package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Storage is the basic interface to communicate
// with different storage endpoints.
type Storage interface {
	// Ping will "ping" the storage backend.
	// The function can use to check the connection from
	// the app to the storage backend.
	Ping() (string, error)
}

// RedisStorage reflects the Storage with
// a Redis NoSQL Server as backend
// See http://redis.io/
type RedisStorage struct {
	connectionPool *redis.Pool
}

// NewRedisStorage will return a new Storage type
// with Redis as backend
func NewRedisStorage(server string) Storage {
	return &RedisStorage{
		connectionPool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

// Ping will "ping" the storage backend.
// The function can use to check the connection from
// the app to the storage backend.
func (r *RedisStorage) Ping() (string, error) {
	conn := r.connectionPool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("PING"))
	return res, err
}
