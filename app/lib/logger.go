package lib

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func initLogger() {
	level := zap.NewAtomicLevel()
	if Config.Debug {
		level.SetLevel(zapcore.DebugLevel)
	}
	config := zap.Config{
		Level:             level,
		Encoding:          "console",
		Development:       Config.Debug,
		DisableStacktrace: true,
		DisableCaller:     true,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build zap config: %w", err))
	}
	Logger = logger.Sugar()
}
