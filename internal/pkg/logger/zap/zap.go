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
	Logger = NewLogger(Config{}, "global")
}

type zapLogger struct {
	logger *zap.SugaredLogger
}

const logDir = "./logs"

// Config is struct to configure logger
type Config struct {
	Targets []string `mapstructure:"targets"`
}

// NewLogger creates sugared zap logger with common config.
// It logs into writer from params.
func NewLogger(cfg Config, serviceName string) logger.Logger {
	// making log file
	os.MkdirAll(logDir, os.ModePerm)

	logFile, err := os.OpenFile(filepath.Join(logDir, "main.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("error opening log file(%s): %s", logFile.Name(), err)
	}

	// making encoder config
	var zcfg zapcore.EncoderConfig
	if os.Getenv("APP_ENV") == "development" {
		zcfg = zap.NewDevelopmentEncoderConfig()
	} else {
		zcfg = zap.NewProductionEncoderConfig()
	}
	// time layout 2006-01-02T15:04:05.000Z0700
	zcfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(zcfg)

	// colorized output
	zcfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(zcfg)

	// cores array
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	}

	// walk for tcp targets
	for _, target := range cfg.Targets {
		conn, err := net.Dial("tcp", target)
		if err != nil {
			log.Fatalf("error connecting to target(%s): %s", target, err)
		}
		cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(conn), zapcore.DebugLevel))
	}

	core := zapcore.NewTee(cores...)

	// make new sugared logger
	sugaredLogger := zap.New(core).Sugar()
	sugaredLogger = sugaredLogger.Named(serviceName)
	if Logger != nil {
		sugaredLogger.Infof("logger initialized with targets: %s", cfg.Targets)
	}

	return &zapLogger{
		logger: sugaredLogger,
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
	return &zapLogger{
		logger: l.logger.With(args...),
	}
}
