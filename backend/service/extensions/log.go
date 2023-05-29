package extensions

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.elastic.co/apm"
	"go.uber.org/zap"

	"tabelf/backend/common"
)

type LogConf struct {
	Path         string `json:"path"`
	Service      string `json:"service"`
	EnableStdOut bool   `json:"enable_std_out"`
	EnableFile   bool   `json:"enable_file"`
}

type LogExt struct {
	LogConf
	UnionLogger *zap.SugaredLogger
}

func NewLogExt(c LogConf) *LogExt {
	return &LogExt{
		LogConf:     c,
		UnionLogger: nil,
	}
}

func (e *LogExt) Init() error {
	common.SetServiceName(e.Service)
	opt := common.LoggerOpt{EnableStdout: e.EnableStdOut, EnableFile: e.EnableFile, IsUnion: true, IsJSONEncoder: true}
	union := common.Logger{Type: common.UNION, LogPath: e.Path}
	e.UnionLogger = union.NewZapLogger(opt).Sugar()
	return nil
}

func (e *LogExt) Close() error { return nil }

func (e *LogExt) OpenAPIMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tx := apm.TransactionFromContext(r.Context())
		opts := apm.SpanOptions{
			Start:  time.Now(),
			Parent: apm.SpanFromContext(ctx).TraceContext(),
		}
		span := tx.StartSpanOptions("AccessLoggerMiddleware", "access_log", opts)
		defer span.End()

		clientIP := r.RemoteAddr
		if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
			clientIP = clientIP[:colon]
		}
		userAgent := r.UserAgent()
		if userAgent == "" {
			userAgent = "-"
		}
		var payload []byte
		if r.Body != nil {
			var err error
			payload, err = ioutil.ReadAll(r.Body)
			if err != nil {
				httpx.Error(w, err)
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
		}
		startTime := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		union := common.UnionLog{
			ClientIP:   clientIP,
			Method:     r.Method,
			Request:    r.URL.RequestURI(),
			Payload:    payload,
			Response:   wrapped.body,
			StatusCode: wrapped.status,
			Protocol:   r.Proto,
			Duration:   time.Since(startTime).Milliseconds(),
			Agent:      fmt.Sprintf("%s, %s", userAgent, r.Header.Get("X-Zh-Crumbs")),
		}
		union.Log(ctx, e.UnionLogger)
	})
}

func (e *LogExt) Error(ctx context.Context, err error) {
	sentry.CaptureException(fmt.Errorf("[%s:%s] %w", common.TraceIDFromContext(ctx), common.SpanIDFromContext(ctx), err))
	common.UnionLog{}.Error(ctx, e.UnionLogger, err)
}

func (e *LogExt) Track(ctx context.Context, msg interface{}) {
	common.UnionLog{}.Track(ctx, e.UnionLogger, msg)
}

// Debug used by entschema migrations.
func (e *LogExt) Debug(v ...interface{}) {
	common.UnionLog{}.Track(context.Background(), e.UnionLogger, fmt.Sprint(v...))
}

func (e *LogExt) Fatal(ctx context.Context, err error) {
	e.Error(ctx, err)
	os.Exit(1)
}

func (e *LogExt) Panic(ctx context.Context, err error) {
	e.Error(ctx, err)
	panic(1)
}

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
	body        []byte
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	rw.body = append(rw.body, body...)
	return rw.ResponseWriter.Write(body)
}
