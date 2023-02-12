package logger

import "github.com/rs/zerolog/log"

func Info(msg string, data interface{}) {
	l := log.Info()
	if data != nil {
		l.Any("data", data)
	}

	l.Msg(msg)
}
