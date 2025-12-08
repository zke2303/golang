package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局Logger变量
var Log *zap.Logger

// InitLogger 初始化日志
// filepath: 日志文件路径
// level: 日志级别
func InitLogger(filePath string, level string) {
	// 1.配置日志切割
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}

	// 2.设置日志级别
	var atomicLevel zap.AtomicLevel
	switch level {
	case "debug":
		atomicLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "error":
		atomicLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	// 3.配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}

	// 4.创建core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)

	// 5.构造Logger
	Log = zap.New(core, zap.AddCaller())
}
