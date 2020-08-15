package core

import (
	"io"
	"log"

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

// Logger ログレベルによって出力を変更するロガー
type Logger struct {
	l     *log.Logger
	level LogLevel
}

// NewLogger コンストラクタ
func NewLogger(level LogLevel) *Logger {
	l := log.New(colorable.NewColorableStderr(), "", log.Ldate|log.Lmicroseconds|log.Lmsgprefix)

	return &Logger{
		l:     l,
		level: level,
	}
}

// SetOutput 出力先を変更
func (l *Logger) SetOutput(w io.Writer) {
	l.l.SetOutput(colorable.NewNonColorable(w))
}

// SetLogger Loggerを設定
func (l *Logger) SetLogger(logger *log.Logger) {
	l.l = logger
}

func (l *Logger) printf(level LogLevel, format string, v ...interface{}) {
	if l.level <= level {
		l.l.SetPrefix(ConvertLogLevelToMessage(level))
		l.l.Printf(format, v...)
	}
}

// Debug Debugログ
func (l *Logger) Debug(format string, v ...interface{}) {
	l.printf(Debug, format, v...)
}

// Info Infoログ
func (l *Logger) Info(format string, v ...interface{}) {
	l.printf(Info, format, v...)
}

// Warn Warnログ
func (l *Logger) Warn(format string, v ...interface{}) {
	l.printf(Warn, format, v...)
}

// Error Errorログ
func (l *Logger) Error(format string, v ...interface{}) {
	l.printf(Error, format, v...)
}
