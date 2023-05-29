package common

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"go.elastic.co/apm"
)

const (
	// DateLayout 日期序列化.
	DateLayout = "2006-01-02"
	// DateTimeLayout 时间序列化.
	DateTimeLayout = "2006-01-02T15:04:05"
)

// NewUUID 生成string类型的uuid.
func NewUUID() string {
	uid, err := uuid.New().MarshalBinary()
	if err != nil {
		return ""
	}
	return hex.EncodeToString(uid)
}

// GetEnv 获取环境变量，不存在则使用默认值.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}

// Date time.Date的快捷方法，省略sec，nsec，loc.
func Time(year int, month time.Month, day, hour, min int) time.Time {
	return time.Date(year, month, day, hour, min, 0, 0, time.Local)
}

type ZError struct {
	Code    string
	Message string
	TraceID string
	SpanID  string
}

func NewZError(ctx context.Context, code interface{}, message string) *ZError {
	return &ZError{
		Code:    fmt.Sprintf("%v", code),
		Message: message,
		TraceID: TraceIDFromContext(ctx),
		SpanID:  SpanIDFromContext(ctx),
	}
}

func (z ZError) Error() string {
	return fmt.Sprintf("%s: %s", z.Code, z.Message)
}

// NewApmSpan 根据context，在当前transaction中生成新的span记录.
func NewApmSpan(ctx context.Context, name, spanType string) *apm.Span {
	tx := apm.TransactionFromContext(ctx)
	opts := apm.SpanOptions{
		Start:  time.Now(),
		Parent: apm.SpanFromContext(ctx).TraceContext(),
	}
	return tx.StartSpanOptions(name, spanType, opts)
}
