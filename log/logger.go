package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func init() {

	logFile := "/var/log/device_plugins/cxl.log"
	// level := zap.DebugLevel
	// dev := true

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	ll := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1024, //MB
		MaxBackups: 30,
		MaxAge:     90, //days
		Compress:   true,
	}

	writeSyncer := zapcore.AddSync(ll)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddStacktrace(zap.DebugLevel))
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	// ll.Rotate()

	// zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
	// 	return lumberjackSink{
	// 		Logger: ll,
	// 	}, nil
	// })

	// loggerConfig := zap.Config{
	// 	Level:         zap.NewAtomicLevelAt(level),
	// 	Development:   dev,
	// 	Encoding:      "console",
	// 	EncoderConfig: encoderConfig,
	// 	OutputPaths:   []string{fmt.Sprintf("lumberjack:%s", logFile)},
	// }
	// var err error
	// logger, err = loggerConfig.Build()

	// if err != nil {
	// 	panic(fmt.Sprintf("build zap logger from config error: %v", err))
	// }

	// zap.ReplaceGlobals(logger)

}

func Init() *log.Logger {

	return zap.NewStdLog(logger)

}
