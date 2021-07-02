package micro_logger

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"testing"
)

var l logger.Logger
var l1 logger.Logger

func TestMain(m *testing.M) {
	cli := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})

	cli2 := &kafka.Writer{
		Addr:         kafka.TCP("192.168.26.182:9092", "192.168.26.183:9092", "192.168.26.184:9092"),
		Balancer:     &kafka.LeastBytes{},
	}

	opt := Zap(
		WithJsonEncoder(),
		WithCaller(),
		WithRedisWriter(cli, "test", "micro"),
		WithStdout(),
		WithLevel(logger.InfoLevel),
		WithStackTrace(logger.ErrorLevel),
	)

	opt1 := Zap(
		WithJsonEncoder(),
		WithCaller(),
		WithKafkaWriter(cli2, "test", "micro"),
		WithStdout(),
		WithLevel(logger.InfoLevel),
		WithStackTrace(logger.ErrorLevel),
		)
	l = NewZapLogger(opt)
	l1 = NewZapLogger(opt1)
	m.Run()
}

func TestZapLogger_Log(t *testing.T) {
    //l.Log(logger.InfoLevel, "hello", "world")
    //l.Log(logger.ErrorLevel, "TEST")
    l1.Log(logger.InfoLevel, "hello", "world")
    l1.Log(logger.ErrorLevel, "TEST")
}
