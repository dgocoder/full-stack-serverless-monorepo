package logger

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func Info(msg string, data interface{}) {
	l := log.Info().Caller(1)
	if data != nil {
		l.Any("data", data)
	}

	l.Msg(msg)
}

func Debug(msg string, data interface{}) {
	l := log.Debug().Caller(1)
	if data != nil {
		l.Any("data", data)
	}

	l.Msg(msg)
}

func Warn(msg string, data interface{}) {
	l := log.Warn().Caller(1)
	if data != nil {
		l.Any("data", data)
	}

	l.Msg(msg)
}

func Error(msg string, data interface{}) {
	l := log.Error().Caller(1)
	if data != nil {
		l.Any("data", data)
	}

	l.Msg(msg)
}

func Errorf(msg string, values string, data interface{}) {
	l := log.Error().Caller(1)
	if data != nil {
		l.Any("data", data)
	}

	l.Msg(fmt.Sprintf(msg, values))
}
