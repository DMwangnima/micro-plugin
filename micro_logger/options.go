package micro_logger

import (
	"context"
	"github.com/DMwangnima/micro-plugin/writer"
	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapKey struct{}

type ZapOptions struct {
	encoder             zapcore.Encoder
	levelFunc           zap.LevelEnablerFunc
	syncer              zapcore.WriteSyncer
	stdoutFlag          bool
	addCaller           bool
	addStackTrace       bool
	stackTraceLevelFunc zap.LevelEnablerFunc
}

type ZapOption func(options *ZapOptions)

func WithJsonEncoder() ZapOption {
	return func(options *ZapOptions) {
		options.encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
}

func WithRedisWriter(cli redis.UniversalClient, topic, key string) ZapOption {
	return func(options *ZapOptions) {
		options.syncer = zapcore.AddSync(writer.NewRedisWriter(cli, topic, key))
	}
}

func WithKafkaWriter(w *kafka.Writer, topic, key string) ZapOption {
	return func(options *ZapOptions) {
		options.syncer = zapcore.AddSync(writer.NewKafkaWriter(w, topic, key))
	}
}

func WithStdout() ZapOption {
	return func(options *ZapOptions) {
		options.stdoutFlag = true
	}
}

func WithCaller() ZapOption {
	return func(options *ZapOptions) {
		options.addCaller = true
	}
}

func WithLevel(l logger.Level) ZapOption {
	return func(options *ZapOptions) {
		lvl := loggerToZapLevel(l)
		options.levelFunc = func(level zapcore.Level) bool {
			return level >= lvl
		}
	}
}

func WithStackTrace(l logger.Level) ZapOption {
	return func(options *ZapOptions) {
		lvl := loggerToZapLevel(l)
		options.addStackTrace = true
		options.stackTraceLevelFunc = func(level zapcore.Level) bool {
			return level >= lvl
		}
	}
}

func Zap(opts ...ZapOption) logger.Option {
	return func(options *logger.Options) {
		if options.Context == nil {
			options.Context = context.Background()
		}
		options.Context = context.WithValue(options.Context, zapKey{}, opts)
	}
}
