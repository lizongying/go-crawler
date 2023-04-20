package logger

import (
	"fmt"
	"github.com/lizongying/go-crawler/internal/config"
	"github.com/lizongying/go-crawler/internal/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

type Logger struct {
	longFile      bool
	level         Level
	loggerDebug   *log.Logger
	loggerInfo    *log.Logger
	loggerWarning *log.Logger
	loggerError   *log.Logger
}

func (l *Logger) Debug(v ...any) {
	if l.level > LevelDebug {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	v = append([]any{strings.Join([]string{file, strconv.Itoa(line)}, ":")}, v...)
	l.loggerDebug.Println(v...)
}

func (l *Logger) DebugF(format string, v ...any) {
	if l.level > LevelDebug {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerDebug.Printf(format, v...)
}

func (l *Logger) Info(v ...any) {
	if l.level > LevelInfo {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	v = append([]any{strings.Join([]string{file, strconv.Itoa(line)}, ":")}, v...)
	l.loggerInfo.Println(v...)
}

func (l *Logger) InfoF(format string, v ...any) {
	if l.level > LevelInfo {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerInfo.Printf(format, v...)
}

func (l *Logger) Warning(v ...any) {
	if l.level > LevelWarn {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	v = append([]any{strings.Join([]string{file, strconv.Itoa(line)}, ":")}, v...)
	l.loggerWarning.Println(v...)
}

func (l *Logger) WarningF(format string, v ...any) {
	if l.level > LevelWarn {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerWarning.Printf(format, v...)
}

func (l *Logger) Error(v ...any) {
	if l.level > LevelError {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	v = append([]any{strings.Join([]string{file, strconv.Itoa(line)}, ":")}, v...)
	l.loggerError.Println(v...)
}

func (l *Logger) ErrorF(format string, v ...any) {
	if l.level > LevelError {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerError.Printf(format, v...)
}

func NewLogger(config *config.Config) (logger *Logger, err error) {
	levelStr := config.Log.Level
	level := LevelInfo
	if levelStr != "" {
		var LevelMap = map[string]Level{
			"DEBUG": LevelDebug,
			"INFO":  LevelInfo,
			"WARN":  LevelWarn,
			"ERROR": LevelError,
		}
		level = LevelMap[strings.ToUpper(levelStr)]
	}

	logger = &Logger{
		longFile: config.Log.LongFile,
		level:    level,
	}
	filename := config.Log.Filename
	if filename == "" {
		logger.loggerDebug = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime)
		logger.loggerInfo = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime)
		logger.loggerWarning = log.New(os.Stdout, "Warn:", log.Ldate|log.Ltime)
		logger.loggerError = log.New(os.Stdout, "Error:", log.Ldate|log.Ltime)
		return
	}

	if !utils.ExistsDir(filename) {
		err = os.MkdirAll(filepath.Dir(filename), 0744)
		if err != nil {
			log.Panicln(err)
			return
		}
	}
	if !utils.ExistsFile(filename) {
		file, errCreateFile := os.Create(filename)
		if errCreateFile != nil {
			log.Panicln(errCreateFile)
			return
		}
		err = file.Close()
		if err != nil {
			log.Panicln(err)
			return
		}
	}
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
		return
	}

	logger.loggerDebug = log.New(io.MultiWriter(os.Stderr, logFile), "Debug:", log.Ldate|log.Ltime)
	logger.loggerInfo = log.New(io.MultiWriter(os.Stderr, logFile), "Info:", log.Ldate|log.Ltime)
	logger.loggerWarning = log.New(io.MultiWriter(os.Stderr, logFile), "Warning:", log.Ldate|log.Ltime)
	logger.loggerError = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Ldate|log.Ltime)

	return
}
