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
	DebugF(format string, v ...any)
	Info(v ...any)
	InfoF(format string, v ...any)
	Warn(v ...any)
	WarnF(format string, v ...any)
	Error(v ...any)
	ErrorF(format string, v ...any)
}
