package core

import (
	"log"
	"os"
)

type Logger interface {
	Info(format string, args ...any)
	Error(format string, args ...any)
	Debug(format string, args ...any)
}

type StdLogger struct {
	std *log.Logger
}

func NewLogger() *StdLogger {
	return &StdLogger{
		std: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *StdLogger) Info(format string, args ...any) {
	l.std.Printf("[INFO] "+format+"\n", args...)
}

func (l *StdLogger) Error(format string, args ...any) {
	l.std.Printf("[ERROR] "+format+"\n", args...)
}

func (l *StdLogger) Debug(format string, args ...any) {
	l.std.Printf("[DEBUG] "+format+"\n", args...)
}
