package handlers

import "log"

type Applications struct {
	l *log.Logger
}

func NewApplication(l *log.Logger) *Applications {
	return &Applications{l}
}
