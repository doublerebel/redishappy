package logger

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Trace      *log.Logger
	Info       *log.Logger
	Warning    *log.Logger
	Error      *log.Logger
	NoteWorthy *log.Logger
)

func init() {
	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	NoteWorthy = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitLogging(logPath string) {}

func newLogFileWriter(logPath string) io.Writer {

	return &lumberjack.Logger{
		Filename:   logPath + "/redis-happy.log",
		MaxSize:    5, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}
}
