package micro_logger

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapLogger struct {
	logger      *zap.Logger
	options     logger.Options
	zapOptions  ZapOptions
}

func NewZapLogger(opts ...logger.Option) logger.Logger {
    z := new(zapLogger)
    for _, opt := range opts {
    	opt(&z.options)
	}
    if zapOpts, ok := z.options.Context.Value(zapKey{}).([]ZapOption); ok {
    	for _, zapOpt := range zapOpts {
            zapOpt(&z.zapOptions)
		}
	}
	var finalCore, outCore, stdoutCore zapcore.Core
	outCore = zapcore.NewCore(z.zapOptions.encoder, z.zapOptions.syncer, z.zapOptions.levelFunc)
	if z.zapOptions.stdoutFlag {
    	stdoutCore = zapcore.NewCore(z.zapOptions.encoder, zapcore.Lock(os.Stdout), z.zapOptions.levelFunc)
    	finalCore = zapcore.NewTee(outCore, stdoutCore)
	} else {
		finalCore = outCore
	}
	z.logger = zap.New(finalCore)
	if z.zapOptions.addCaller {
		z.logger = z.logger.WithOptions(zap.AddCaller())
	}
	return z
}

// don't use Init
func (l *zapLogger) Init(opts ...logger.Option) error {
	return nil
}

func (l *zapLogger) Options() logger.Options {
	return l.options
}

func (l *zapLogger) Fields(fields map[string]interface{}) logger.Logger {
	if len(fields) == 0 {
		return l
	}
	newFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		newFields = append(newFields, zap.Any(k, v))
	}
	newZapLogger := l.logger.With(newFields...)
	return &zapLogger{
		logger:     newZapLogger,
		options:    l.options,
		zapOptions: l.zapOptions,
	}
}

func (l *zapLogger) Log(level logger.Level, v ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(v...)
	switch lvl {
	case zap.DebugLevel:
		l.logger.Debug(msg)
	case zap.InfoLevel:
		l.logger.Info(msg)
	case zap.WarnLevel:
		l.logger.Warn(msg)
	case zap.ErrorLevel:
		l.logger.Error(msg)
	case zap.FatalLevel:
		l.logger.Fatal(msg)
	}
}

func (l *zapLogger) Logf(level logger.Level, format string, v ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, v...)
	switch lvl {
	case zap.DebugLevel:
		l.logger.Debug(msg)
	case zap.InfoLevel:
		l.logger.Info(msg)
	case zap.WarnLevel:
		l.logger.Warn(msg)
	case zap.ErrorLevel:
		l.logger.Error(msg)
	case zap.FatalLevel:
		l.logger.Fatal(msg)
	}
}

func (l *zapLogger) String() string {
	return "zap"
}

func loggerToZapLevel(level logger.Level) zapcore.Level {
	switch level {
	case logger.TraceLevel, logger.DebugLevel:
		return zap.DebugLevel
	case logger.InfoLevel:
		return zap.InfoLevel
	case logger.WarnLevel:
		return zap.WarnLevel
	case logger.ErrorLevel:
		return zap.ErrorLevel
	case logger.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}