package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap *zap.SugaredLogger

func init() {
	Zap = newLogger()
}

// New creates sugared zap logger with common config.
// It logs into writer from params.
func newLogger() *zap.SugaredLogger {
	// making log file
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// TEMP!!!
	logDir := filepath.Join(ex, "logs")
	if _, err := os.Stat(logDir); err != nil {
		os.MkdirAll(logDir, os.ModePerm)
	}

	logFile, err := os.OpenFile(filepath.Join(logDir, "main.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
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

	// TODO add kibana, logstash

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	return zap.New(core).Sugar()
}
