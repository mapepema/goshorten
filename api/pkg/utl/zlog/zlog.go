package zlog

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Log represents zerolog Logger
type Log struct {
	logger *zerolog.Logger
}

// New instanciates new zero logger
func New() *Log {
	z := zerolog.New(os.Stdout)

	return &Log{
		logger: &z,
	}
}

// Log logs using zerolog
func (z *Log) Log(ctx fiber.Ctx, source, msg string, err error, params map[string]interface{}) {

	if params == nil {
		params = make(map[string]interface{})
	}

	params["source"] = source

	if id := ctx.Get("id"); id != "" {
		params["id"] = id
	}

	if err != nil {
		params["error"] = err
		z.logger.Error().Fields(params).Msg(msg)
		return
	}

	z.logger.Info().Fields(params).Msg(msg)
}
