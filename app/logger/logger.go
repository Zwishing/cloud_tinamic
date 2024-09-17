package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func init() {
	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
