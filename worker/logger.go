package worker

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Print(level zerolog.Level, args ...interface{}) {
	log.WithLevel(level).Msg(fmt.Sprintf("%v", args...))
}

func (l *Logger) Debug(args ...interface{}) {
	log.Debug().Msg(fmt.Sprintf("%v", args...))
}

func (l *Logger) Info(args ...interface{}) {
	log.Info().Msg(fmt.Sprintf("%v", args...))
}

func (l *Logger) Warn(args ...interface{}) {
	log.Warn().Msg(fmt.Sprintf("%v", args...))
}

func (l *Logger) Error(args ...interface{}) {
	log.Error().Msg(fmt.Sprintf("%v", args...))
}

func (l *Logger) Fatal(args ...interface{}) {
	log.Fatal().Msg(fmt.Sprintf("%v", args...))
}
