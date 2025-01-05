package zap

import (
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger logger.Logger

func init() {
	Logger = NewLogger()
}

type zapLogger struct {
	logger *zap.SugaredLogger
}

const logDir = "./logs"

// NewLogger creates sugared zap logger with common config.
// It logs into writer from params.
func NewLogger(tcpTargets ...string) logger.Logger {
	// making log file
	os.MkdirAll(logDir, os.ModePerm)

	logFile, err := os.OpenFile(filepath.Join(logDir, "main.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("error opening log file(%s): %s", logFile.Name(), err)
	}

	// making encoder config
	var cfg zapcore.EncoderConfig
	if os.Getenv("APP_ENV") == "development" {
		cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg = zap.NewProductionEncoderConfig()
	}
	// time layout 2006-01-02T15:04:05.000Z0700
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(cfg)

	// colorized output
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(cfg)

	// cores array
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	}

	// walk for tcp targets
	for _, target := range tcpTargets {
		conn, err := net.Dial("tcp", target)
		if err != nil {
			log.Fatalf("error connecting to target(%s): %s", target, err)
		}
		cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(conn), zapcore.DebugLevel))
	}

	core := zapcore.NewTee(cores...)

	return &zapLogger{
		logger: zap.New(core).Sugar(),
	}
}

func (l *zapLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *zapLogger) WithFields(args ...interface{}) logger.Logger {
	zl := zapLogger{
		logger: l.logger.With(args...),
	}
	return &zl
}
