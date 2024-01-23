package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
	Expired time.Duration
}

func (r *RedisClient) Connect(addr string, passwd string, db int) error {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	pong, err := r.Client.Ping().Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return err
	}
	return nil
}

func (r *RedisClient) Set(key string, value string) error {
	err := r.Client.Set(key, value, r.Expired).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.Client.Get(key).Result()
	if err == redis.Nil {

	} else if err != nil {
		return "", err
	}
	return val, nil
}
