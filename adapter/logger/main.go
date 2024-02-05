package logger

import (
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"
)

func NewLogger(config util.Config) port.Logger {
	return NewZeroLogLogger(config)
}
