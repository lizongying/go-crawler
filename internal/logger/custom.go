package logger

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
)

type Logger struct {
	level pkg.Level
}

func (l *Logger) Debug(v ...any) {
	if l.level > pkg.LevelDebug {
		return
	}
}

func (l *Logger) Debugf(format string, v ...any) {
	if l.level > pkg.LevelDebug {
		return
	}
}

func (l *Logger) Info(v ...any) {
	if l.level > pkg.LevelInfo {
		return
	}
}

func (l *Logger) Infof(format string, v ...any) {
	if l.level > pkg.LevelInfo {
		return
	}
}

func (l *Logger) Warn(v ...any) {
	if l.level > pkg.LevelWarn {
		return
	}
}

func (l *Logger) Warnf(format string, v ...any) {
	if l.level > pkg.LevelWarn {
		return
	}
}

func (l *Logger) Error(v ...any) {
	if l.level > pkg.LevelError {
		return
	}
}

func (l *Logger) Errorf(format string, v ...any) {
	if l.level > pkg.LevelError {
		return
	}
}

func NewLogger(config *config.Config) (logger *Logger, err error) {
	logger = &Logger{
		level: config.GetLogLevel(),
	}

	return
}
