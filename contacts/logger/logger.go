package logger

import (
	"log"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTimestampFormat = "2006-01-02T15:04:05.000"
)

// Logger defines contract for implementing logs
type Logger interface {
	Error(string, error, ...Fields)
	Warn(string, error, ...Fields)
	Info(string, ...Fields)
}

type logger struct {
	env        string
	serverRoot string
	lg         *zap.Logger
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(logTimestampFormat))
}

var (
	onceNewLogger    sync.Once
	onceNewLoggerRes Logger
)

// NewLogger creates zap logger with default prod configs
// This function can be used to change logger configs
func NewLogger(env string) Logger {
	onceNewLogger.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
		zapLogger, err := cfg.Build()
		if err != nil {
			log.Fatalf("err creating logger: %v\n", err.Error())
		}

		onceNewLoggerRes = &logger{
			lg:         zapLogger,
			env:        env,
			serverRoot: "/",
		}
	})

	return onceNewLoggerRes
}

// Fields is a wrapper above map to be used
// for structured logging
type Fields map[string]interface{}

func (l logger) Error(msg string, err error, fields ...Fields) {
	l.lg.Error(msg, append(withFields(fields), zap.Error(err))...)
}

func (l logger) Warn(msg string, err error, fields ...Fields) {
	l.lg.Warn(msg, append(withFields(fields), zap.Error(err))...)
}

func (l logger) Info(msg string, fields ...Fields) {
	l.lg.Info(msg, withFields(fields)...)
}

func withFields(fields []Fields) []zapcore.Field {
	withFields := []zapcore.Field{}
	for _, fieldMap := range fields {
		for key, ele := range fieldMap {
			withFields = append(withFields, zap.Any(key, ele))
		}
	}

	return withFields
}

func Error(msg string, err error, fields ...Fields) {
	NewLogger("").Error(msg, err, fields...)
}

func Warn(msg string, err error, fields ...Fields) {
	NewLogger("").Warn(msg, err, fields...)
}

func Info(msg string, fields ...Fields) {
	NewLogger("").Info(msg, fields...)
}
