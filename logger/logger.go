package logger

import (
	"fmt"
	"strings"

	env "github.com/esequielvirtuoso/go_utils_lib/envs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel         = "LOG_LEVEL"
	envLogLevelDefault  = "info"
	envLogOutput        = "LOG_OUTPUT"
	envLogOutputDefault = "stdout"
)

var (
	log logger
)

type loggerInterface interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

type logger struct {
	log *zap.Logger
}

// init starts logger
func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// GetLogger allow external modules to access log variable
func GetLogger() loggerInterface {
	return log
}

// Printf implements the Printf function required by elasticsearch library
func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

// Print implements the Print function required by mysql errors interface
func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

// Info register information logs
func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	if err := log.log.Sync(); err != nil {
		fmt.Println("error while logging")
	}
}

// Error register error logs
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	if err := log.log.Sync(); err != nil {
		fmt.Println("error while logging")
	}
}

// getLevel is responsible to define the logger level of information
func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(env.GetString(envLogLevel, envLogLevelDefault))) {
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	return strings.TrimSpace(env.GetString(envLogOutput, envLogOutputDefault))
}
