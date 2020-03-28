package log

import (
	"context"
	"fmt"
)

/**
**
**
** there are ryanx's log addition
**
**
 */
// Debug logs a message at the debug log level.
func Debug(format string, args ...interface{}) {
	h.Log(context.Background(), _debugLevel, KVString(_log, fmt.Sprintf(format, args...)))
}

// Debugc logs a message at the debug log level.
func Debugc(ctx context.Context, format string, args ...interface{}) {
	h.Log(ctx, _debugLevel, KVString(_log, fmt.Sprintf(format, args...)))
}

// Debugv logs a message at the debug log level.
func Debugv(ctx context.Context, args ...D) {
	h.Log(ctx, _debugLevel, args...)
}

// Debugw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Debugw(ctx context.Context, args ...interface{}) {
	h.Log(ctx, _debugLevel, logw(args)...)
}

// Debugm logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Debugm(msg string, args ...interface{}) {
	args = append([]interface{}{"msg", msg}, args...)
	h.Log(context.Background(), _debugLevel, logw(args)...)
}

// Infow logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Infom(msg string, args ...interface{}) {
	args = append([]interface{}{"msg", msg}, args...)
	h.Log(context.Background(), _infoLevel, logw(args)...)
}

// Warnw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Warnm(msg string, args ...interface{}) {
	args = append([]interface{}{"msg", msg}, args...)
	h.Log(context.Background(), _warnLevel, logw(args)...)
}

// Errorw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Errorm(msg string, args ...interface{}) {
	args = append([]interface{}{"msg", msg}, args...)
	h.Log(context.Background(), _errorLevel, logw(args)...)
}
