package pkg

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var LevelMap = map[string]Level{
	"DEBUG": LevelDebug,
	"INFO":  LevelInfo,
	"WARN":  LevelWarn,
	"ERROR": LevelError,
}

type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}
