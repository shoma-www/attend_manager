package core

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/mattn/go-colorable"
)

// LogLevel ログの出力レベル
type LogLevel int

// LogLevel
const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

// ConvertStringToLogLevel 文字列からログ出力レベルに変換する
func ConvertStringToLogLevel(level string) LogLevel {
	switch level {
	case "DEBUG":
		return Debug
	case "INFO":
		return Info
	case "WARN":
		return Warn
	case "ERROR":
		return Error
	default:
		return Debug
	}
}

// ConvertLogLevelToMessage LogLevelをメッセージ用表記に変更
// nolint
func ConvertLogLevelToMessage(level LogLevel) string {
	switch level {
	case Debug:
		return "\x1b[34m[DEBUG] "
	case Info:
		return "\x1b[32m[INFO ] "
	case Warn:
		return "\x1b[33m[WARN ] "
	case Error:
		return "\x1b[31m[ERROR] "
	default:
		return "\x1b[34m[DEBUG] "
	}
}

// ContextKey コンテキストキー
type ContextKey string

// UUIDContextKey UUIDのキー
const (
	UUIDContextKey   ContextKey = "uuid"
	LoggerContextKey ContextKey = "logger"
)

// Logger LoggerInterface
type Logger interface {
	SetOutput(w io.Writer)
	SetLogger(logger *log.Logger)
	SetUUID(uuid string) Logger
	WithUUID(ctx context.Context) Logger
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}

// SetLogger from context.Context
func SetLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, LoggerContextKey, l)
}

// GetLogger from context.Context
func GetLogger(ctx context.Context) Logger {
	if l, ok := ctx.Value(LoggerContextKey).(Logger); ok {
		return l
	}
	return NewLogger(Info)
}

// logger ログレベルによって出力を変更するロガー
type logger struct {
	l     *log.Logger
	level LogLevel
	uuid  string
}

// NewLogger コンストラクタ
func NewLogger(level LogLevel) Logger {
	l := log.New(colorable.NewColorableStderr(), "", log.Ldate|log.Lmicroseconds|log.Lmsgprefix)

	return &logger{
		l:     l,
		level: level,
	}
}

func (l *logger) clone() *logger {
	return &logger{
		l:     l.l,
		level: l.level,
	}
}

// WithUUID UUIDをセットする
func (l *logger) WithUUID(ctx context.Context) Logger {
	var ok bool
	cl := l.clone()
	if cl.uuid, ok = ctx.Value(UUIDContextKey).(string); !ok {
		cl.uuid = ""
	}
	return cl
}

// WithUUID UUIDをセットする
func (l *logger) SetUUID(uuid string) Logger {
	cl := l.clone()
	cl.uuid = uuid
	return cl
}

// SetOutput 出力先を変更
func (l *logger) SetOutput(w io.Writer) {
	l.l.SetOutput(colorable.NewNonColorable(w))
}

// Setlogger Loggerを設定
func (l *logger) SetLogger(logger *log.Logger) {
	l.l = logger
}

func (l *logger) printf(level LogLevel, format string, v ...interface{}) {
	if l.level <= level {
		ls := ConvertLogLevelToMessage(level)
		ss := []string{}
		if l.uuid != "" {
			ss = append(ss, l.uuid, " ")
		}
		ss = append(ss, ls, format)
		l.l.Printf(strings.Join(ss, ""), v...)
	}
}

// Debug Debugログ
// nolint
func (l *logger) Debug(format string, v ...interface{}) {
	l.printf(Debug, format, v...)
}

// Info Infoログ
// nolint
func (l *logger) Info(format string, v ...interface{}) {
	l.printf(Info, format, v...)
}

// Warn Warnログ
// nolint
func (l *logger) Warn(format string, v ...interface{}) {
	l.printf(Warn, format, v...)
}

// Error Errorログ
// nolint
func (l *logger) Error(format string, v ...interface{}) {
	l.printf(Error, format, v...)
}
