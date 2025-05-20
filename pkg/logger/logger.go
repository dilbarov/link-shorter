package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func SetupLogger(env string) zerolog.Logger {
	switch env {
	case "development":
		writer := zerolog.ConsoleWriter{
			Out:           os.Stdout,
			NoColor:       false,
			TimeFormat:    "15:04:05",
			FormatLevel:   func(_ interface{}) string { return "" }, // убираем level
			FormatMessage: func(i interface{}) string { return fmt.Sprintf("%s", i) },
			FormatTimestamp: func(i interface{}) string {
				return fmt.Sprintf("[%s]", time.Now().Format("15:04:05"))
			},
		}
		return zerolog.New(writer).With().Timestamp().Logger()

	case "production":
		fallthrough
	default:
		return zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
