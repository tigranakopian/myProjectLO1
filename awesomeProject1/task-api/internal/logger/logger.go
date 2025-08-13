package logger

import (
	"context"
	"log"
)

type Logger struct {
	ch chan string
}

func NewLogger() *Logger {
	return &Logger{ch: make(chan string, 100)}
}

func (l *Logger) Log(msg string) {
	select {
	case l.ch <- msg:
	default:
		// drop if full
	}
}

func (l *Logger) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-l.ch:
			log.Println("[LOG]", msg)
		}
	}
}
