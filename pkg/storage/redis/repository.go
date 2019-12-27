package redis

import (
	"time"

	"github.com/RainJoe/go-web-project-template/pkg/config"
	"github.com/gomodule/redigo/redis"
)

// Storage stores  data in redis
type Storage struct {
	redisPool *redis.Pool
}

//NewStorage returns a new redis storage
func NewStorage(conf config.TOMLConfig) (*Storage, error) {
	var err error

	s := new(Storage)

	pool := &redis.Pool{
		MaxIdle:     100,
		MaxActive:   4000,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.RedisConfig.Addr, redis.DialPassword(conf.RedisConfig.PassWord), redis.DialDatabase(conf.RedisConfig.DB))
			if nil != err {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	s.redisPool = pool

	return s, err
}
