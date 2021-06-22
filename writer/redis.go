package writer

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	cli   redis.UniversalClient
	topic string
	key   string
}

func (r *Redis) Write(b []byte) (int, error) {
	n, err := r.cli.LPush(context.Background(), r.topic+":"+r.key, b).Result()
	return int(n), err
}

func NewRedisWriter(cli redis.UniversalClient, topic, key string) *Redis {
	return &Redis{
		cli:   cli,
		topic: topic,
		key:   key,
	}
}
