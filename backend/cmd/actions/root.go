package actions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"tabelf/backend/common"
)

const Service = "binghuang"

var (
	ErrFatal    = errors.New("fatal error")
	ErrPanic    = errors.New("panic error")
	UnionLogger *zap.SugaredLogger
	rootCmd     = &cobra.Command{Use: Service}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("env", os.Getenv("ENV"), "environment. eg: development, production")
}

func InitDefaultLogger() {
	common.SetServiceName(Service)
	opt := common.LoggerOpt{EnableStdout: true, IsUnion: true, IsJSONEncoder: true}
	union := common.Logger{Type: common.UNION}
	UnionLogger = union.NewZapLogger(opt).Sugar()
}

func Fatal(msg string) {
	common.UnionLog{}.Error(context.TODO(), UnionLogger, fmt.Errorf("%w: %s", ErrFatal, msg))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v...))
}

func Panic(err error) {
	common.UnionLog{}.Error(context.TODO(), UnionLogger, err)
	panic(err)
}

func Panicf(format string, v ...interface{}) {
	Panic(fmt.Errorf("%w: %s", ErrPanic, fmt.Sprintf(format, v...)))
}

func Track(msg interface{}) {
	common.UnionLog{}.Track(context.TODO(), UnionLogger, msg)
}
