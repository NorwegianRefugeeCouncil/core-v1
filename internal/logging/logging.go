package logging

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var logLevel zap.AtomicLevel

type Level string

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
	FatalLevel Level = "fatal"
)

type Config struct {
	Level Level
}

func init() {
	logLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = ""
	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		logLevel,
	))
}

func SetLogLevel(level string) error {
	lvl, err := parseLogLevel(level)
	if err != nil {
		return err
	}
	logLevel.SetLevel(lvl)
	return nil
}

func parseLogLevel(level string) (zapcore.Level, error) {
	switch level {
	case string(DebugLevel):
		return zapcore.DebugLevel, nil
	case string(InfoLevel):
		return zapcore.InfoLevel, nil
	case string(WarnLevel):
		return zapcore.WarnLevel, nil
	case string(ErrorLevel):
		return zapcore.ErrorLevel, nil
	case string(FatalLevel):
		return zapcore.FatalLevel, nil
	default:
		return zapcore.DebugLevel, errors.New("unknown log level: " + string(level))
	}
}
