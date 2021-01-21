package logs

import (
	"context"
)

var (
	defaultLogger *Logger
	psm           = ""
	cluster       = ""
	localIP       = ""
)

func init() {
	defaultLogger = NewConsoleLogger()
	defaultLogger.StartLogger()
	defaultLogger.SetCallDepth(3)
}

func InitLogger(logger *Logger) {
	defaultLogger.Stop()
	defaultLogger = logger
	defaultLogger.StartLogger()
}

func AddProvider(p LogProvider) {
	defaultLogger.AddProvider(p)
}

func SetLevel(l int) {
	defaultLogger.SetLevel(l)
}

// Level 当前logger的Level
func Level() int {
	return defaultLogger.level
}

func SetCallDepth(depth int) {
	defaultLogger.SetCallDepth(depth)
}

func Stop() {
	defaultLogger.Stop()
}

func DefaultLogger() *Logger {
	return defaultLogger
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatal(format, v...)
}

// Panicf 出现问题了，强制退出
func Panicf(format string, v ...interface{}) {
	defaultLogger.Panicf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

func Warnf(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

func Noticef(format string, v ...interface{}) {
	defaultLogger.Notice(format, v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

func Debugf(format string, v ...interface{}) {
	if defaultLogger.level > LevelDebug {
		return
	}
	defaultLogger.Debug(format, v...)
}

func Tracef(format string, v ...interface{}) {
	defaultLogger.Trace(format, v...)
}

func Fatal(format string, v ...interface{}) {
	defaultLogger.Fatal(format, v...)
}

func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

func Notice(format string, v ...interface{}) {
	defaultLogger.Notice(format, v...)
}

func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

// InfoV2 不再format处理
func InfoV2(msg string) {
	defaultLogger.InfoV2(msg)
}

func Debug(format string, v ...interface{}) {
	if defaultLogger.level > LevelDebug {
		return
	}
	defaultLogger.Debug(format, v...)
}

func Trace(format string, v ...interface{}) {
	defaultLogger.Trace(format, v...)
}

func Flush() {
	defaultLogger.Flush()
}

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxFatal(ctx, format, v...)
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxError(ctx, format, v...)
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxWarn(ctx, format, v...)
}

func CtxNotice(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxNotice(ctx, format, v...)
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxInfo(ctx, format, v...)
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxDebug(ctx, format, v...)
}

func CtxTrace(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxTrace(ctx, format, v...)
}

func CtxFatalKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxFatalKVs(ctx, kvs...)
}

func CtxErrorKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxErrorKVs(ctx, kvs...)
}

func CtxWarnKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxWarnKVs(ctx, kvs...)
}

func CtxNoticeKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxNoticeKVs(ctx, kvs...)
}

func CtxInfoKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxInfoKVs(ctx, kvs...)
}

func CtxDebugKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxDebugKVs(ctx, kvs...)
}

func CtxTraceKvs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxTraceKVs(ctx, kvs...)
}

func CtxPushNotice(ctx context.Context, k, v interface{}) {
	ntc := getNotice(ctx)
	if ntc == nil {
		return
	}
	ntc.PushNotice(k, v)
}

func CtxFlushNotice(ctx context.Context) {
	ntc := getNotice(ctx)
	if ntc == nil {
		return
	}
	kvs := ntc.KVs()
	if len(kvs) == 0 {
		return
	}
	logid := string(logIDBytes(ctx))
	defaultLogger.WriteKvs(LevelNotice, localIP, psm, logid, cluster, kvs...)
}

func NewNoticeCtx(ctx context.Context) context.Context {
	ntc := newNoticeKVs()
	return context.WithValue(ctx, noticeCtxKey, ntc)
}
