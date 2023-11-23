package loggers

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

func (l *Logger) Debugf(format string, v ...any) {
	if l.level > pkg.LevelDebug {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s\n", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
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

func (l *Logger) Infof(format string, v ...any) {
	if l.level > pkg.LevelInfo {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s\n", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
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

func (l *Logger) Warnf(format string, v ...any) {
	if l.level > pkg.LevelWarn {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s\n", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
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

func (l *Logger) Errorf(format string, v ...any) {
	if l.level > pkg.LevelError {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	if !l.longFile {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	format = fmt.Sprintf("%s %s\n", strings.Join([]string{file, strconv.Itoa(line)}, ":"), format)
	l.loggerError.Printf(format, v...)
}

func NewLogger(config *config.Config, stream *Stream) (logger *Logger, err error) {
	logger = &Logger{
		longFile: config.GetLogLongFile(),
		level:    config.GetLogLevel(),
	}

	var multiWriter []io.Writer

	multiWriter = append(multiWriter, os.Stdout)

	filename := config.Log.Filename
	if filename != "" {
		filename = strings.ReplaceAll(filename, "{name}", name)

		if !utils.ExistsDir(filename) {
			if err = os.MkdirAll(filepath.Dir(filename), 0744); err != nil {
				log.Panicln(err)
				return
			}
		}

		var file *os.File
		if !utils.ExistsFile(filename) {
			file, err = os.Create(filename)
			if err != nil {
				log.Panicln(err)
				return
			}
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Panicln(err)
				return
			}
		}

		multiWriter = append(multiWriter, file)
	}

	if stream != nil {
		multiWriter = append(multiWriter, stream)
	}

	writer := io.MultiWriter(multiWriter...)
	logger.loggerDebug = log.New(writer, "Debug: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	logger.loggerInfo = log.New(writer, "Info: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	logger.loggerWarn = log.New(writer, "Warn: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	logger.loggerError = log.New(writer, "Error: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	return
}
