package micro_logger

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"testing"
)

var l logger.Logger

func TestMain(m *testing.M) {
	cli := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})

	opt := Zap(
		WithJsonEncoder(),
		WithCaller(),
		WithRedisWriter(cli, "test", "micro"),
		WithStdout(),
		WithLevel(logger.InfoLevel),
		WithStackTrace(logger.ErrorLevel),
	)
	l = NewZapLogger(opt)
	m.Run()
}

func TestZapLogger_Log(t *testing.T) {
    l.Log(logger.InfoLevel, "hello", "world")
    l.Log(logger.ErrorLevel, "TEST")
}
