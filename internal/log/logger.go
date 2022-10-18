package logging

import (
	"github.com/stackitcloud/stackit-argus-cli/internal/settings"
	nativeLogger "log"
	"sync"

	"go.uber.org/zap"
)

var (
	appKey  = "app"
	appName = "argus-default-dashboard"
)

// logger to init zap logger and a lock for singleton.
type logger struct {
	sugar *zap.SugaredLogger
	lock  sync.Mutex
	l     *zap.Logger
}

var log = logger{lock: sync.Mutex{}}

type Field = zap.Field

var (
	Skip       = zap.Skip
	Binary     = zap.Binary
	Bool       = zap.Bool
	Boolp      = zap.Boolp
	ByteString = zap.ByteString
	Float64    = zap.Float64
	Float64p   = zap.Float64p
	Float32    = zap.Float32
	Float32p   = zap.Float32p
	Durationp  = zap.Durationp
	Any        = zap.Any
	String     = zap.String
)

// Logger should have a info and error log level.
type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

var _ Logger = &logger{}

// Info logs on info level.
func (l *logger) Info(msg string, fields ...Field) {
	if l.sugar != nil {
		fields = append(fields, String(appKey, appName))
		l.l.Info(msg, fields...)
	} else {
		nativeLogger.Println(msg)
	}
}

// Error logs on error level.
func (l *logger) Error(msg string, fields ...Field) {
	if l.sugar != nil {
		fields = append(fields, String(appKey, appName))
		l.l.Error(msg, fields...)
	} else {
		nativeLogger.Println(msg)
	}
}

// Fatal logs on fatal level.
func (l *logger) Fatal(msg string, fields ...Field) {
	if l.sugar != nil {
		fields = append(fields, String(appKey, appName))
		l.l.Fatal(msg, fields...)
	} else {
		nativeLogger.Fatalf(msg)
	}
}

// Debug logs on debug level if it is enabled.
func (l *logger) Debug(msg string, fields ...Field) {
	if settings.DebugLog {
		if l.sugar != nil {
			fields = append(fields, String(appKey, appName))
			l.l.Info(msg, fields...)
		} else {
			nativeLogger.Println(msg)
		}
	}
}

// New initializes a new logger.
func New() Logger {
	log.lock.Lock()
	defer log.lock.Unlock()

	if log.sugar == nil {
		l, err := zap.NewProduction()
		if err != nil {
			return nil
		}

		log.sugar = l.Sugar()
		log.l = l
	}

	return &log
}
