package logger

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	name string
)

type Logger struct {
	longFile    bool
	level       pkg.Level
	loggerDebug *log.Logger
	loggerInfo  *log.Logger
	loggerWarn  *log.Logger
	loggerError *log.Logger
}

func (l *Logger) Debug(v ...any) {
	if l.level > pkg.LevelDebug {
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
	if l.level > pkg.LevelDebug {
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
	if l.level > pkg.LevelInfo {
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
	if l.level > pkg.LevelInfo {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerInfo.Printf(format, v...)
}

func (l *Logger) Warn(v ...any) {
	if l.level > pkg.LevelWarn {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	v = append([]any{strings.Join([]string{file, strconv.Itoa(line)}, ":")}, v...)
	l.loggerWarn.Println(v...)
}

func (l *Logger) WarnF(format string, v ...any) {
	if l.level > pkg.LevelWarn {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerWarn.Printf(format, v...)
}

func (l *Logger) Error(v ...any) {
	if l.level > pkg.LevelError {
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
	if l.level > pkg.LevelError {
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
	level := pkg.LevelInfo
	if levelStr != "" {
		l, ok := pkg.LevelMap[strings.ToUpper(levelStr)]
		if ok {
			level = l
		}
	}

	logger = &Logger{
		longFile: config.Log.LongFile,
		level:    level,
	}
	filename := config.Log.Filename
	if filename == "" {
		logger.loggerDebug = log.New(os.Stdout, "Debug:", log.Ldate|log.Ltime)
		logger.loggerInfo = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime)
		logger.loggerWarn = log.New(os.Stdout, "Warn:", log.Ldate|log.Ltime)
		logger.loggerError = log.New(os.Stdout, "Error:", log.Ldate|log.Ltime)
		return
	}

	if name != "" {
		filename = strings.ReplaceAll(filename, "{name}", name)
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
	logger.loggerWarn = log.New(io.MultiWriter(os.Stderr, logFile), "Warn:", log.Ldate|log.Ltime)
	logger.loggerError = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Ldate|log.Ltime)

	return
}
