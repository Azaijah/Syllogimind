package main

import (
	"os"

	"github.com/Azaijah/Syllogimind/syllogism"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	syllo := syllogism.Syllogism{}

	syllo.Generate()
	syllo.Show()

}
