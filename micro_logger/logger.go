package micro_logger

import (
	"github.com/asim/go-micro/v3/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	enabler     zapcore.LevelEnabler
	encoder     zapcore.Encoder
	writeSyncer zapcore.WriteSyncer
	logger      *zap.Logger
	options     logger.Options
}

func (l *zapLogger) Init(opts ...logger.Option) error {
	panic("implement me")
}

func (l *zapLogger) Options() logger.Options {
	return l.options
}

func (l *zapLogger) Fields(fields map[string]interface{}) logger.Logger {
	panic("implement me")
}

func (l *zapLogger) Log(level logger.Level, v ...interface{}) {
	panic("implement me")
}

func (l *zapLogger) Logf(level logger.Level, format string, v ...interface{}) {
	panic("implement me")
}

func (l *zapLogger) String() string {
	return "zap"
}



