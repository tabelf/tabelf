package common

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	// nolint:staticcheck
	// ignore SA1019 Need to keep deprecated package for compatibility.
	"github.com/golang/protobuf/proto"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const (
	HTTP2Protocol = "HTTP/2"
	LocalHost     = "127.0.0.1"
)

type AccessLog struct {
	ClientIP   string
	Method     string
	Request    string
	Protocol   string
	Agent      string
	LogType    string
	GrpcStatus string
	Payload    []byte
	Response   []byte
	Duration   int64
	StatusCode int
}

// NewUnaryServerAccessLogInterceptor returns a new unary server interceptors tha log access log.
func NewUnaryServerAccessLogInterceptor(logger *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		startTime := time.Now()
		ip, _ := peer.FromContext(ctx)
		resp, err := handler(ctx, req)
		clientIP := strings.Split(ip.Addr.String(), ":")[0]
		// ignore probe requests
		if clientIP == LocalHost {
			return resp, err
		}
		code := grpc_logging.DefaultErrorToCode(err)
		l := AccessLog{
			ClientIP:   clientIP,
			Request:    info.FullMethod,
			Protocol:   HTTP2Protocol,
			Duration:   time.Since(startTime).Milliseconds(),
			LogType:    grpcLogType,
			GrpcStatus: code.String(),
		}
		if msg, ok := req.(proto.Message); ok {
			l.Payload, err = MarshalJSON(msg)
		}
		if msg, ok := resp.(proto.Message); ok {
			l.Response, err = MarshalJSON(msg)
		}
		l.Log(logger)
		return resp, err
	}
}

func (l AccessLog) Log(logger *zap.SugaredLogger) {
	logType := defaultLogType
	if l.LogType != "" {
		logType = l.LogType
	}
	logger.Infof(
		"%s %s %s $%q$ %s %d %d \"%s\" %s $%q$ %s %s",
		l.ClientIP, l.Method, l.Request, l.Payload, l.Protocol,
		l.StatusCode, l.Duration, l.Agent, serviceName, l.Response, logType, l.GrpcStatus,
	)
}

func MarshalJSON(msg interface{}) ([]byte, error) {
	if pb, ok := msg.(proto.Message); ok {
		b := &bytes.Buffer{}
		if err := grpc_zap.JsonPbMarshaller.Marshal(b, pb); err != nil {
			return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
		}
		return b.Bytes(), nil
	}
	return nil, fmt.Errorf("msg not valid: %v", msg)
}
