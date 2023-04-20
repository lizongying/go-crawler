package internal

type Logger interface {
	Debug(v ...any)
	DebugF(format string, v ...any)
	Info(v ...any)
	InfoF(format string, v ...any)
	Warning(v ...any)
	WarningF(format string, v ...any)
	Error(v ...any)
	ErrorF(format string, v ...any)
}
