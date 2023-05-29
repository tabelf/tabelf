package common

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	timeFormatter  = "2006-01-02 15:04:05"
	defaultLogType = "http"
	grpcLogType    = "grpc"
)

var serviceName = GetEnv("SERVICE_NAME", "default")

func SetServiceName(name string) {
	serviceName = name
}

type LogType string

const (
	ACCESS  LogType = "access"
	REQUEST LogType = "request"
	TRACK   LogType = "track"
	ERROR   LogType = "error"
	UNION   LogType = "union"
)

type Logger struct {
	Type    LogType
	LogPath string
}

type LoggerOpt struct {
	CustomEncoderConfig  *zapcore.EncoderConfig
	CustomEncoderConfigs map[zapcore.Level]*zapcore.EncoderConfig
	EnableStdout         bool
	EnableFile           bool
	IsUnion              bool
	IsJSONEncoder        bool
}

func (l *Logger) NewZapLogger(opt LoggerOpt) (logger *zap.Logger) {
	var writers []zapcore.WriteSyncer
	if opt.EnableStdout {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}
	if opt.EnableFile {
		writers = append(writers, zapcore.AddSync(l.fileRotateWriter()))
	}
	var core zapcore.Core
	if opt.IsUnion {
		encoderMap := make(map[zapcore.Level]zapcore.EncoderConfig)
		encoderMap[zapcore.ErrorLevel] = DefaultEncoderConfig(true)
		encoderMap[zapcore.InfoLevel] = DefaultEncoderConfig(false)
		encoderMap[zapcore.DebugLevel] = DefaultEncoderConfig(false)
		for level, encoder := range opt.CustomEncoderConfigs {
			encoderMap[level] = *encoder
		}
		cores := make([]zapcore.Core, len(encoderMap))
		i := 0
		for key, encoderConfig := range encoderMap {
			level := key
			encoder := zapcore.NewConsoleEncoder(encoderConfig)
			if opt.IsJSONEncoder {
				encoder = zapcore.NewJSONEncoder(encoderConfig)
			}
			cores[i] = zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(writers...),
				zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == level
				}),
			)
			i++
		}
		core = zapcore.NewTee(cores...)
	} else {
		var enc zapcore.EncoderConfig
		switch l.Type {
		case ERROR:
			enc = DefaultEncoderConfig(true)
		case TRACK:
			enc = TrackEncoderConfig()
		default:
			enc = DefaultEncoderConfig(false)
		}
		if opt.CustomEncoderConfig != nil {
			enc = *opt.CustomEncoderConfig
		}
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(enc),
			zapcore.NewMultiWriteSyncer(writers...),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		)
	}
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Named(serviceName)
}

func DefaultEncoderConfig(withStack bool) zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		NameKey:          "name",
		LevelKey:         "level",
		TimeKey:          "time",
		CallerKey:        "path",
		FunctionKey:      "func",
		MessageKey:       "msg",
		ConsoleSeparator: " ",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(timeFormatter),
	}
	if withStack {
		config.StacktraceKey = "stacktrace"
		config.LineEnding = "$" + zapcore.DefaultLineEnding
	}
	return config
}

func (l *Logger) fileRotateWriter() io.Writer {
	path := l.filePath()
	hook, err := rotateLogs.New(
		path+".%Y-%m-%d",
		rotateLogs.WithLinkName(path),
		rotateLogs.WithMaxAge(time.Hour*24*30),
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Panic(err)
	}
	return hook
}

func (l *Logger) filePath() string {
	return fmt.Sprintf("%s/%s.log", l.LogPath, l.Type)
}

type UnionLog struct {
	ClientIP   string
	Protocol   string
	Agent      string
	Method     string
	Request    string
	LogType    string
	GrpcStatus string
	Payload    []byte
	Response   []byte
	Duration   int64
	StatusCode int
}

func (l UnionLog) Log(ctx context.Context, logger *zap.SugaredLogger) {
	logType := defaultLogType
	if l.LogType != "" {
		logType = l.LogType
	}
	logger.Infow("",
		zap.String("client_ip", l.ClientIP),
		zap.String("protocol", l.Protocol),
		zap.String("agent", l.Agent),
		zap.String("method", l.Method),
		zap.String("request", l.Request),
		zap.ByteString("payload", l.Payload),
		zap.ByteString("response", l.Response),
		zap.Int64("duration", l.Duration),
		zap.Int("status_code", l.StatusCode),
		zap.String("log_type", logType),
		zap.String("grpc_status", l.GrpcStatus),
		zap.String("trace_id", TraceIDFromContext(ctx)),
		zap.String("span_id", SpanIDFromContext(ctx)),
	)
}

func TrackEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:          "time",
		FunctionKey:      "func",
		MessageKey:       "msg",
		ConsoleSeparator: " ",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(serviceName)
			enc.AppendString(zapcore.InfoLevel.CapitalString())
			enc.AppendString(t.Format(timeFormatter))
		},
	}
}

func (l UnionLog) Error(ctx context.Context, logger *zap.SugaredLogger, err error) {
	logger.Errorw(
		err.Error(),
		zap.String("trace_id", TraceIDFromContext(ctx)),
		zap.String("span_id", SpanIDFromContext(ctx)),
	)
}

func (l UnionLog) Track(ctx context.Context, logger *zap.SugaredLogger, msg interface{}) {
	logger.Debugw(fmt.Sprintf("%+v", msg),
		zap.String("trace_id", TraceIDFromContext(ctx)),
		zap.String("span_id", SpanIDFromContext(ctx)),
	)
}

func Error(logger *zap.SugaredLogger, err error) {
	logger.Errorf("%v", err)
}

func Track(logger *zap.SugaredLogger, message interface{}) {
	logger.Infof("%s", message)
}

func SpanIDFromContext(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}
	return ""
}

func TraceIDFromContext(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}
